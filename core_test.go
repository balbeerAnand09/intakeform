package servers_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/empty"

	"github.com/google/uuid"

	"github.com/golang/mock/gomock"
	"go.appointy.com/google/intakeform/internal/servers"
	"go.appointy.com/google/intakeform/internal/servers/mocks"
	"go.appointy.com/google/intakeform/internal/stores"
	"go.appointy.com/google/intakeform/internal/stores/mocks"
	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/google/pb/locations"
)

func getMockData() map[string]*intake_form.IntakeFormItem {
	v := &intake_form.IntakeFormItem{
		Key:   "abc",
		Value: &any.Any{},
	}
	m := make(map[string]*intake_form.IntakeFormItem, 0)
	m["k1"] = v

	return m
}

// getMockIntakeForm returns Intake form
func getMockIntakeForm() *intake_form.IntakeForm {
	return &intake_form.IntakeForm{
		Base: &locations.LocationRoot{
			LocationId: uuid.New().String(),
			ProgramId:  uuid.New().String(),
		},
		Id:   uuid.New().String(),
		Data: getMockData(),
	}
}

// getMockIntakeFormReq returns Mock IntakeForm Request
func getMockIntakeFormReq() *intake_form.IntakeFormIdentifier {
	return &intake_form.IntakeFormIdentifier{
		Base: &locations.LocationRoot{
			LocationId: uuid.New().String(),
			ProgramId:  uuid.New().String(),
		},
		Id: uuid.New().String(),
	}
}

func getMocks(t *testing.T) (*gomock.Controller, intake_form.IntakeFormsServer, *servers_mocks.MockLocationsClient, *store_mocks.MockIntakeFormStore) {
	ctrl := gomock.NewController(t)

	locCli := servers_mocks.NewMockLocationsClient(ctrl)
	store := store_mocks.NewMockIntakeFormStore(ctrl)

	srv := servers.NewCoreIntakeFormServer(store, locCli)

	return ctrl, srv, locCli, store
}

func TestCoreServer_AddIntakeForms(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *intake_form.IntakeForm
	}
	type wants struct {
		want    *intake_form.IntakeFormIdentifier
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want: &intake_form.IntakeFormIdentifier{
					Base: &locations.LocationRoot{},
				},
				wantErr: false,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				call := loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.in.Base.ProgramId,
							Id:        a.in.Base.LocationId,
						},
					).Return(nil, nil).Times(1)
				sm.EXPECT().
					AddIntakeForms(a.ctx, a.in).
					Return(w.want, nil).
					Times(1).
					After(call)
			},
		},
		{
			name: "location not found",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want: &intake_form.IntakeFormIdentifier{
					Base: &locations.LocationRoot{},
				},
				wantErr: true,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.in.Base.ProgramId,
							Id:        a.in.Base.LocationId,
						},
					).Times(1).
					Return(nil, errors.New("locatoin not found"))
			},
		},
		{
			name: "store error",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want: &intake_form.IntakeFormIdentifier{
					Base: &locations.LocationRoot{},
				},
				wantErr: true,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				call := loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.in.Base.ProgramId,
							Id:        a.in.Base.LocationId,
						},
					).Return(nil, nil).Times(1)
				sm.EXPECT().
					AddIntakeForms(a.ctx, a.in).
					Return(nil, errors.New("store error")).
					Times(1).
					After(call)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, srv, locCli, store := getMocks(t)
			defer ctrl.Finish()
			tt.setup(tt.a, locCli, store, tt.w)
			_, err := srv.AddIntakeForms(tt.a.ctx, tt.a.in)
			if (err != nil) != tt.w.wantErr {
				t.Errorf("CoreServer.AddIntakeForm() error = %v, wantErr %v", err, tt.w.wantErr)
				return
			}
		})
	}
}

func TestCoreServer_UpdateIntakeForm(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *intake_form.IntakeForm
	}
	type wants struct {
		emp     *empty.Empty
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				call := sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(a.in, nil).
					Times(1)
				sm.EXPECT().
					UpdateIntakeForm(a.ctx, a.in).
					Return(nil).
					Times(1).
					After(call)
			},
		},
		{
			name: "intake form not found",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(nil, errors.New("intake form not found")).
					Times(1)
			},
		},
		{
			name: "store error",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				call := sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(a.in, nil).
					Times(1)
				sm.EXPECT().
					UpdateIntakeForm(a.ctx, a.in).
					Return(errors.New("store error")).
					Times(1).
					After(call)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, srv, _, store := getMocks(t)
			defer ctrl.Finish()
			tt.setup(tt.a, store, tt.w)
			_, err := srv.UpdateIntakeForm(tt.a.ctx, tt.a.in)
			if (err != nil) != tt.w.wantErr {
				t.Errorf("CoreServer.UpdateIntakeForm error = %v, wantErr %v", err, tt.w.wantErr)
				return
			}
		})
	}
}

func TestCoreServer_DeleteIntakeForm(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *intake_form.IntakeFormIdentifier
	}
	type wants struct {
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants)
	}{
		{
			name: "Sucess",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				wantErr: false,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				getCall := sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.in.Id).
					Return(nil, nil).
					Times(1)

				sm.EXPECT().DeleteIntakeForm(gomock.Any(), a.in.Id).Return(nil).Times(1).After(getCall)
			},
		},
		{
			name: "Wrong Id",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				wantErr: true,
			},
			setup: func(a *args, _ *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.in.Id = ""

				sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.in.Id).
					Return(nil, errors.New("")).
					Times(1)
			},
		},
		{
			name: "Store Delete Error",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				wantErr: true,
			},
			setup: func(a *args, _ *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {

				getCall := sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.in.Id).
					Return(nil, nil).
					Times(1)

				sm.EXPECT().DeleteIntakeForm(gomock.Any(), a.in.Id).Return(errors.New("")).Times(1).After(getCall)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, s, cm, sm := getMocks(t)
			defer c.Finish()

			tt.setup(tt.a, cm, sm, tt.w)

			_, err := s.DeleteIntakeForm(tt.a.ctx, tt.a.in)
			if (err != nil) != tt.w.wantErr {
				t.Errorf("CoreServer.DeleteIntakeForm() error = %v, wantErr %v", err, tt.w.wantErr)
				return
			}
		})
	}
}

func TestCoreServer_GetIntakeForms(t *testing.T) {
	type args struct {
		ctx context.Context
		loc *locations.LocationRoot
	}
	type wants struct {
		want    *intake_form.IntakeFormList
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				loc: &locations.LocationRoot{},
			},
			w: &wants{
				want:    &intake_form.IntakeFormList{},
				wantErr: false,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.loc.LocationId = uuid.New().String()
				a.loc.ProgramId = uuid.New().String()
				for i := 0; i < 5; i++ {
					mp := getMockIntakeForm()
					mp.Base = a.loc
					w.want.Forms = append(w.want.Forms, mp)
				}
				call := loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.loc.ProgramId,
							Id:        a.loc.LocationId,
						},
					).Return(nil, nil).Times(1)
				sm.EXPECT().
					GetIntakeForms(a.ctx, a.loc.LocationId).
					Return(w.want, nil).
					Times(1).
					After(call)
			},
		},
		{
			name: "location not found",
			a: &args{
				ctx: context.Background(),
				loc: &locations.LocationRoot{},
			},
			w: &wants{
				want:    &intake_form.IntakeFormList{},
				wantErr: true,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.loc.LocationId = uuid.New().String()
				a.loc.ProgramId = uuid.New().String()
				loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.loc.ProgramId,
							Id:        a.loc.LocationId,
						},
					).Times(1).
					Return(nil, errors.New("location not found"))
			},
		},
		{
			a: &args{
				ctx: context.Background(),
				loc: &locations.LocationRoot{},
			},
			w: &wants{
				want:    &intake_form.IntakeFormList{},
				wantErr: true,
			},
			setup: func(a *args, loc *servers_mocks.MockLocationsClient, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.loc.LocationId = uuid.New().String()
				a.loc.ProgramId = uuid.New().String()
				call := loc.EXPECT().
					GetLocations(
						a.ctx,
						&locations.LocationIdentifier{
							ProgramId: a.loc.ProgramId,
							Id:        a.loc.LocationId,
						},
					).Return(nil, nil).Times(1)
				sm.EXPECT().
					GetIntakeForms(a.ctx, a.loc.LocationId).
					Return(nil, errors.New("store error")).
					Times(1).
					After(call)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, srv, locCli, store := getMocks(t)
			defer ctrl.Finish()
			tt.setup(tt.a, locCli, store, tt.w)
			mplst, err := srv.GetIntakeForms(tt.a.ctx, tt.a.loc)
			if (err != nil) != tt.w.wantErr {
				t.Errorf("CoreServer.GetIntakeForms() error = %v, wantErr %v", err, tt.w.wantErr)
				return
			}
			if !tt.w.wantErr && !reflect.DeepEqual(mplst, tt.w.want) {
				t.Errorf("CoreServer.GetIntakeForms() = %v, want %v", mplst, tt.w.want)
			}
		})
	}
}

func TestCoreServer_GetIntakeFormById(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *intake_form.IntakeFormIdentifier
	}
	type wants struct {
		want    *intake_form.IntakeForm
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				want:    getMockIntakeForm(),
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.in.Base.LocationId = w.want.Base.LocationId
				a.in.Base.ProgramId = w.want.Base.ProgramId
				a.in.Id = w.want.Id
				sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(w.want, nil).
					Times(1)
			},
		},
		{
			name: "intake form not found",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				want:    getMockIntakeForm(),
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(nil, stores.ErrNotFound).
					Times(1)
			},
		},
		{
			name: "wrong location",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq(),
			},
			w: &wants{
				want:    getMockIntakeForm(),
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, w *wants) {
				a.in.Base = &locations.LocationRoot{ProgramId: "", LocationId: ""}
				a.in.Id = w.want.Id
				sm.EXPECT().
					GetIntakeFormById(a.ctx, a.in.Id).
					Return(w.want, nil).
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl, srv, _, store := getMocks(t)
			defer ctrl.Finish()
			tt.setup(tt.a, store, tt.w)
			mp, err := srv.GetIntakeFormById(tt.a.ctx, tt.a.in)
			if (err != nil) != tt.w.wantErr {
				t.Errorf("CoreServer.GetIntakeFormById() error = %v, wantErr %v", err, tt.w.wantErr)
				return
			}
			if !tt.w.wantErr && !reflect.DeepEqual(mp, tt.w.want) {
				t.Errorf("CoreServer.GetIntakeFormById() = %v, want %v", mp, tt.w.want)
			}
		})
	}
}

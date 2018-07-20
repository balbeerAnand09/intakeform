package cache_test

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	goCache "github.com/srikrsna/go-cache"
	"github.com/srikrsna/go-cache/mocks"

	"go.appointy.com/google/intakeform/internal/cache"
	"go.appointy.com/google/intakeform/internal/stores"
	"go.appointy.com/google/intakeform/internal/stores/mocks"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/google/uuid"
	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/google/pb/locations"
)

// Test Helpers
type ofType struct{ t interface{} }

func OfType(t interface{}) gomock.Matcher {
	return &ofType{t}
}

func (o *ofType) Matches(x interface{}) bool {
	return reflect.DeepEqual(reflect.TypeOf(x), reflect.TypeOf(o.t))
}

func (o *ofType) String() string {
	return "is of type " + reflect.TypeOf(o.t).String()
}

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

func getMocks(t *testing.T) (stores.IntakeFormStore, *store_mocks.MockIntakeFormStore, *cache_mock.MockCache, *gomock.Controller) {
	t.Helper()
	ctrl := gomock.NewController(t)

	storeMock := store_mocks.NewMockIntakeFormStore(ctrl)
	cacheMock := cache_mock.NewMockCache(ctrl)

	c := cache.NewCachedIntakeFormStore(storeMock, cacheMock)

	return c, storeMock, cacheMock, ctrl
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
		setup func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want:    &intake_form.IntakeFormIdentifier{},
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				call := sm.EXPECT().
					AddIntakeForms(a.ctx, a.in).
					Return(w.want, nil).
					Times(1)
				cm.EXPECT().Set(gomock.Any(), a.in, gomock.Any()).Return(nil).Times(1).After(call)
			},
		},
		{
			name: "store add failed",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want:    &intake_form.IntakeFormIdentifier{},
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				sm.EXPECT().
					AddIntakeForms(a.ctx, a.in).
					Return(nil, errors.New("store add failed")).
					Times(1)
			},
		},
		{
			name: "cache failed",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				want:    &intake_form.IntakeFormIdentifier{},
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				call := sm.EXPECT().
					AddIntakeForms(a.ctx, a.in).
					Return(w.want, nil).
					Times(1)
				cm.EXPECT().
					Set(gomock.Any(), a.in, gomock.Any()).
					Return(errors.New("cache failed")).
					Times(1).
					After(call)
			},
		},
	}
	for _, tt := range tests {
		c, sm, cm, ctrl := getMocks(t)
		defer ctrl.Finish()
		tt.setup(tt.a, sm, cm, tt.w)
		if _, err := c.AddIntakeForms(tt.a.ctx, tt.a.in); (err != nil) != tt.w.wantErr {
			t.Errorf("AddIntakeForm, want err: %v, got: %v", tt.w.wantErr, err)
		}
	}
}

func TestCoreServer_UpdateIntakeForm(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *intake_form.IntakeForm
	}
	type wants struct {
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants)
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
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				call := sm.EXPECT().
					UpdateIntakeForm(a.ctx, a.in).
					Return(nil).
					Times(1)
				cm.EXPECT().Set(gomock.Any(), a.in, gomock.Any()).Return(nil).Times(1).After(call)
			},
		},
		{
			name: "store add failed",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				sm.EXPECT().
					UpdateIntakeForm(a.ctx, a.in).
					Return(errors.New("store add failed")).
					Times(1)
			},
		},
		{
			name: "cache failed",
			a: &args{
				ctx: context.Background(),
				in:  getMockIntakeForm(),
			},
			w: &wants{
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				call := sm.EXPECT().
					UpdateIntakeForm(a.ctx, a.in).
					Return(nil).
					Times(1)
				cm.EXPECT().
					Set(gomock.Any(), a.in, gomock.Any()).
					Return(errors.New("cache failed")).
					Times(1).
					After(call)
			},
		},
	}
	for _, tt := range tests {
		c, sm, cm, ctrl := getMocks(t)
		defer ctrl.Finish()
		tt.setup(tt.a, sm, cm, tt.w)
		if err := c.UpdateIntakeForm(tt.a.ctx, tt.a.in); (err != nil) != tt.w.wantErr {
			t.Errorf("UpdateIntakeForm, want err: %v, got: %v", tt.w.wantErr, err)
		}
	}
}

func TestCachedLocationStore_DeleteIntakeForm(t *testing.T) {
	type args struct {
		ctx context.Context
		in  string
	}
	type wants struct {
		err error
	}
	tests := []struct {
		name  string
		args  *args
		setup func(*args, *store_mocks.MockIntakeFormStore, *cache_mock.MockCache, *wants)
		wants *wants
	}{
		{
			name: "Success",
			args: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq().Id,
			},
			wants: &wants{
				err: nil,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				storeCall := sm.EXPECT().
					DeleteIntakeForm(gomock.Any(), a.in).
					Return(nil).
					Times(1)

				cm.EXPECT().
					Delete(gomock.Any()).
					Return(nil).
					After(storeCall).
					Times(1)
			},
		},
		{
			name: "Store Fail",
			args: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq().Id,
			},
			wants: &wants{
				err: errors.New(""),
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				sm.EXPECT().
					DeleteIntakeForm(gomock.Any(), a.in).
					Return(w.err).
					Times(1)
			},
		},
		{
			name: "Cache Fail",
			args: &args{
				ctx: context.Background(),
				in:  getMockIntakeFormReq().Id,
			},
			wants: &wants{
				err: nil,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				storeCall := sm.EXPECT().
					DeleteIntakeForm(gomock.Any(), a.in).
					Return(nil).
					Times(1)

				cm.EXPECT().
					Delete(gomock.Any()).
					Return(errors.New("")).
					After(storeCall).
					Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, storeMock, cacheMock, ctrl := getMocks(t)
			defer ctrl.Finish()

			tt.setup(tt.args, storeMock, cacheMock, tt.wants)

			err := c.DeleteIntakeForm(tt.args.ctx, tt.args.in)

			if err != tt.wants.err {
				t.Errorf("want err: %v, got: %v", tt.wants.err, err)
			}
		})
	}
}

func TestCoreServer_GetIntakeForms(t *testing.T) {
	c, sm, _, ctrl := getMocks(t)
	defer ctrl.Finish()
	sm.EXPECT().
		GetIntakeForms(gomock.Any(), gomock.Any()).
		Return(&intake_form.IntakeFormList{}, nil).
		Times(1)
	c.GetIntakeForms(context.Background(), "")
}

func TestCoreServer_GetIntakeFormById(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	type wants struct {
		mp      *intake_form.IntakeForm
		wantErr bool
	}
	tests := []struct {
		name  string
		a     *args
		w     *wants
		setup func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants)
	}{
		{
			name: "success",
			a: &args{
				ctx: context.Background(),
				id:  uuid.New().String(),
			},
			w: &wants{
				mp:      getMockIntakeForm(),
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				w.mp.Id = a.id
				cm.EXPECT().
					Get(gomock.Any(), OfType(w.mp), gomock.Any()).
					Return(nil).
					Do(func(key string, v interface{}, d time.Duration) {
						p := v.(*intake_form.IntakeForm)
						*p = *w.mp
					}).
					Times(1)
			},
		},
		{
			name: "cache miss",
			a: &args{
				ctx: context.Background(),
				id:  uuid.New().String(),
			},
			w: &wants{
				mp:      getMockIntakeForm(),
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				w.mp.Id = a.id
				checkCall := cm.EXPECT().
					Get(gomock.Any(), OfType(w.mp), gomock.Any()).
					Return(goCache.ErrCacheMiss).
					Times(1)
				storeCall := sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.id).
					Return(w.mp, nil).
					After(checkCall).
					Times(1)
				cm.EXPECT().
					Set(gomock.Any(), w.mp, gomock.Any()).
					Return(nil).
					After(storeCall).
					Times(1)
			},
		},
		{
			name: "Miss Save Error",
			a: &args{
				ctx: context.Background(),
				id:  uuid.New().String(),
			},
			w: &wants{
				mp:      getMockIntakeForm(),
				wantErr: false,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				w.mp.Id = a.id
				checkCall := cm.EXPECT().
					Get(gomock.Any(), OfType(w.mp), gomock.Any()).
					Return(goCache.ErrCacheMiss).
					Times(1)
				storeCall := sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.id).
					Return(w.mp, nil).
					After(checkCall).
					Times(1)
				cm.EXPECT().
					Set(gomock.Any(), w.mp, gomock.Any()).
					Return(errors.New("")).
					After(storeCall).
					Times(1)
			},
		},
		{
			name: "miss and store Error",
			a: &args{
				ctx: context.Background(),
				id:  uuid.New().String(),
			},
			w: &wants{
				mp:      getMockIntakeForm(),
				wantErr: true,
			},
			setup: func(a *args, sm *store_mocks.MockIntakeFormStore, cm *cache_mock.MockCache, w *wants) {
				w.mp.Id = a.id
				checkCall := cm.EXPECT().
					Get(gomock.Any(), OfType(w.mp), gomock.Any()).
					Return(errors.New("")).
					Times(1)
				sm.EXPECT().
					GetIntakeFormById(gomock.Any(), a.id).
					Return(w.mp, errors.New("store error")).
					After(checkCall).
					Times(1)
			},
		},
	}
	for _, tt := range tests {
		c, sm, cm, ctrl := getMocks(t)
		defer ctrl.Finish()
		tt.setup(tt.a, sm, cm, tt.w)
		mp, err := c.GetIntakeFormById(tt.a.ctx, tt.a.id)
		if (err != nil) != tt.w.wantErr {
			t.Errorf("GetIntakeFormById, want err: %v, got: %v", tt.w.wantErr, err)
			return
		}
		if !tt.w.wantErr && !reflect.DeepEqual(tt.w.mp, mp) {
			t.Errorf("want id: %v got: %v", tt.w.wantErr, err)
		}
	}
}

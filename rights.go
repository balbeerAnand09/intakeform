package rights

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/google/pb/locations"
	"go.appointy.com/google/pb/rights"
)

// RightsServer ...
type RightsServer struct {
	core      intake_form.IntakeFormsServer
	validator rights.RightValidatorClient
}

// IntakeFormRights ...
type IntakeFormRights = int32

const (
	//AddIntakeForms ...
	AddIntakeForms IntakeFormRights = iota
	//UpdateIntakeForm ...
	UpdateIntakeForm
	//DeleteIntakeForm ...
	DeleteIntakeForm
	//GetIntakeForms ...
	GetIntakeForms
	//GetIntakeFormsById ...
	GetIntakeFormsById
)

// NewRightsIntakeFormServer ...
func NewRightsIntakeFormServer(c rights.RightValidatorClient, s intake_form.IntakeFormsServer) intake_form.IntakeFormsServer {
	return &RightsServer{
		core:      s,
		validator: c,
	}
}

// AddIntakeForms adds a intake forms to the store
func (s *RightsServer) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {
	res, err := s.validator.IsValid(ctx, &rights.IsValidRequest{
		ResourcePath: fmt.Sprintf("/programs/%s/locations/%s", in.Base.GetProgramId(), in.Base.GetLocationId()),
		Value:        AddIntakeForms,
		UserId:       ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.core.AddIntakeForms(ctx, in)
}

// UpdateIntakeForm updates the intake form in the store
func (s *RightsServer) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm) (*empty.Empty, error) {
	//Trying to update intake form to store
	res, err := s.validator.IsValid(ctx, &rights.IsValidRequest{
		ResourcePath: fmt.Sprintf("/programs/%s/locations/%s", in.Base.GetProgramId(), in.Base.GetLocationId()),
		Value:        UpdateIntakeForm,
		UserId:       ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.core.UpdateIntakeForm(ctx, in)
}

// DeleteIntakeForm deletes intake form with the given ID
func (s *RightsServer) DeleteIntakeForm(ctx context.Context, in *intake_form.IntakeFormIdentifier) (*empty.Empty, error) {
	//Trying to delete intake form to store
	res, err := s.validator.IsValid(ctx, &rights.IsValidRequest{
		ResourcePath: fmt.Sprintf("/programs/%s/locations/%s", in.Base.GetProgramId(), in.Base.GetLocationId()),
		Value:        DeleteIntakeForm,
		UserId:       ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.core.DeleteIntakeForm(ctx, in)
}

// GetIntakeForms get intake form by location
func (s *RightsServer) GetIntakeForms(ctx context.Context, loc *locations.LocationRoot) (*intake_form.IntakeFormList, error) {
	//Trying to retrieve existing intake form of given ID from store
	res, err := s.validator.IsValid(ctx, &rights.IsValidRequest{
		ResourcePath: fmt.Sprintf("/programs/%s/location/%s", loc.ProgramId, loc.LocationId),
		Value:        GetIntakeForms,
		UserId:       ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.core.GetIntakeForms(ctx, loc)
}

// GetIntakeFormById get intake form by id
func (s *RightsServer) GetIntakeFormById(ctx context.Context, in *intake_form.IntakeFormIdentifier) (*intake_form.IntakeForm, error) {
	//Trying to retrieve existing intake form of given ID from store
	res, err := s.validator.IsValid(ctx, &rights.IsValidRequest{
		ResourcePath: fmt.Sprintf("/programs/%s/location/%s", in.Base.ProgramId, in.Base.LocationId),
		Value:        GetIntakeFormsById,
		UserId:       ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	if !res.IsValid {
		return nil, status.Errorf(codes.PermissionDenied, res.Reason)
	}

	return s.core.GetIntakeFormById(ctx, in)
}

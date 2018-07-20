package servers

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.appointy.com/google/intakeform/internal/stores"
	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/google/pb/locations"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Generic Error to be retured to client to hide possible sensitive information
var errInternal = status.Error(codes.Internal, "internal server error")

// CoreServer is the business logic layer of the IntakeformServer
type CoreServer struct {
	store  stores.IntakeFormStore
	locCli locations.LocationsClient
}

// NewCoreIntakeFormServer returns a IntakeformServer implementation with Core business logic
func NewCoreIntakeFormServer(s stores.IntakeFormStore, l locations.LocationsClient) intake_form.IntakeFormsServer {
	return &CoreServer{
		store:  s,
		locCli: l,
	}
}

// AddIntakeForms adds a intake form to the store
func (c *CoreServer) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {

	// Checking whether location exists or not.
	_, err := c.locCli.GetLocations(ctx, &locations.LocationIdentifier{ProgramId: in.Base.ProgramId, Id: in.Base.LocationId})
	if err != nil {
		return nil, err
	}
	// Add intake form to store
	id, err := c.store.AddIntakeForms(ctx, in)
	if err != nil {
		ctxzap.Extract(ctx).Error("unable to save intake forms to store", zap.Error(err))
		return nil, errInternal
	}
	return id, nil
}

// UpdateIntakeForm updates the intake form in the store
func (c *CoreServer) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm) (*empty.Empty, error) {

	// Try to get intake form
	if _, err := c.store.GetIntakeFormById(ctx, in.Id); err != nil {
		return nil, err
	}
	// Try to save intake form
	if err := c.store.UpdateIntakeForm(ctx, in); err != nil {
		ctxzap.Extract(ctx).Error("unable to save intake form to store", zap.Error(err))
		return nil, errInternal
	}
	return nil, nil
}

// DeleteIntakeForm deletes intake form with the given ID
func (c *CoreServer) DeleteIntakeForm(ctx context.Context, in *intake_form.IntakeFormIdentifier) (*empty.Empty, error) {

	// Try to get intake form
	if _, err := c.store.GetIntakeFormById(ctx, in.Id); err != nil {
		return nil, err
	}

	// Try to delete location
	if err := c.store.DeleteIntakeForm(ctx, in.Id); err != nil {
		ctxzap.Extract(ctx).Error("unable to delete intake form from store", zap.Error(err))
		return nil, errInternal
	}
	return nil, nil
}

// GetIntakeForms get intake form by location
func (c *CoreServer) GetIntakeForms(ctx context.Context, loc *locations.LocationRoot) (*intake_form.IntakeFormList, error) {

	// Checking whether location exists or not.
	_, err := c.locCli.GetLocations(ctx, &locations.LocationIdentifier{ProgramId: loc.ProgramId, Id: loc.LocationId})
	if err != nil {
		return nil, err
	}
	// Trying to get the intake form by location
	lst, err := c.store.GetIntakeForms(ctx, loc.LocationId)
	if err != nil {
		return nil, err
	}
	return lst, nil
}

// GetIntakeFormById get intake form by id
func (c *CoreServer) GetIntakeFormById(ctx context.Context, in *intake_form.IntakeFormIdentifier) (*intake_form.IntakeForm, error) {
	// Trying to get intake forms From store
	mp, err := c.store.GetIntakeFormById(ctx, in.Id)
	if err != nil {
		if err == stores.ErrNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		ctxzap.Extract(ctx).Error("unable to retrieve intake forms from store", zap.Error(err))
		return nil, errInternal
	}
	// Checking if intake forms belongs to the location
	if in.Base.LocationId != mp.Base.LocationId || in.Base.ProgramId != mp.Base.ProgramId {
		return nil, status.Error(codes.FailedPrecondition, "intake forms do not belong to location")
	}
	return mp, nil
}

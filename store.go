package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/srikrsna/go-cache"
	"go.appointy.com/google/intakeform/internal/stores"
	"go.appointy.com/google/pb/intake_form"
	"go.appointy.com/pb/location"
)

// CachedIntakeFormStore is a wrapper around any IntakeFormStore implementation to Cache IntakeForm
type CachedIntakeFormStore struct {
	cache.Cache

	prefix string
	store  stores.IntakeFormStore

	d time.Duration
}

// NewCachedIntakeFormStore wraps a IntakeFormStore in a caching layer
func NewCachedIntakeFormStore(s stores.IntakeFormStore, c cache.Cache) stores.IntakeFormStore {
	return &CachedIntakeFormStore{
		store:  s,
		prefix: proto.MessageName(&location.Location{}),
		d:      time.Second * 5,
		Cache:  c,
	}
}

// cacheKey returns a key for the cache using IntakeForm id
func (c *CachedIntakeFormStore) cacheKey(id string) string {
	return fmt.Sprintf("%s-%s", c.prefix, id)
}

// AddIntakeForms adds a intake form using underlying store and caches the result
func (c *CachedIntakeFormStore) AddIntakeForms(ctx context.Context, in *intake_form.IntakeForm) (*intake_form.IntakeFormIdentifier, error) {
	got, err := c.store.AddIntakeForms(ctx, in)
	if err != nil {
		return nil, err
	}
	if err := c.Set(c.cacheKey(in.Id), in, c.d); err != nil {
		ctxzap.Extract(ctx).Error("unable to set intake form to cache", zap.Error(err))
	}
	return got, nil
}

// UpdateIntakeForm updates the intake form using underlying store and caches the result
func (c *CachedIntakeFormStore) UpdateIntakeForm(ctx context.Context, in *intake_form.IntakeForm) error {
	if err := c.store.UpdateIntakeForm(ctx, in); err != nil {
		return err
	}
	if err := c.Set(c.cacheKey(in.Id), in, c.d); err != nil {
		ctxzap.Extract(ctx).Error("unable to set intake form to cache", zap.Error(err))
	}
	return nil
}

// DeleteIntakeForm bypass cache and forwards request to uinderlying store
func (c *CachedIntakeFormStore) DeleteIntakeForm(ctx context.Context, ID string) error {
	if err := c.store.DeleteIntakeForm(ctx, ID); err != nil {
		return err
	}

	if err := c.Delete(c.cacheKey(ID)); err != nil {
		ctxzap.Extract(ctx).Warn("unable to delete value from cache", zap.Error(err))
	}

	return nil
}

// GetIntakeForms bypass cache and forward to the store to retrieve all the intake forms by Location
func (c *CachedIntakeFormStore) GetIntakeForms(ctx context.Context, loc string) (*intake_form.IntakeFormList, error) {
	return c.store.GetIntakeForms(ctx, loc)
}

// GetIntakeFormById checks the intake form in cache and returns it, if exists, and if not it will revert to the store
func (c *CachedIntakeFormStore) GetIntakeFormById(ctx context.Context, id string) (*intake_form.IntakeForm, error) {
	key := c.cacheKey(id)
	mp := &intake_form.IntakeForm{}
	err := c.Get(key, mp, c.d)
	if err == nil {
		return mp, nil
	}
	if err != cache.ErrCacheMiss {
		ctxzap.Extract(ctx).Error("unable to get intake form from cache", zap.Error(err), zap.String("cache key", key))
	}
	// get intake form from store
	mp, err = c.store.GetIntakeFormById(ctx, id)
	if err != nil {
		return nil, err
	}
	if err = c.Set(key, mp, c.d); err != nil {
		ctxzap.Extract(ctx).Error("unable to set intake form to cache", zap.Error(err))
	}
	return mp, nil
}

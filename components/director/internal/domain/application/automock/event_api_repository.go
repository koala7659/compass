// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import (
	context "context"

	model "github.com/kyma-incubator/compass/components/director/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// EventAPIRepository is an autogenerated mock type for the EventAPIRepository type
type EventAPIRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, items
func (_m *EventAPIRepository) Create(ctx context.Context, items *model.EventDefinition) error {
	ret := _m.Called(ctx, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.EventDefinition) error); ok {
		r0 = rf(ctx, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByApplicationID provides a mock function with given fields: ctx, tenantID, appID
func (_m *EventAPIRepository) DeleteAllByApplicationID(ctx context.Context, tenantID string, appID string) error {
	ret := _m.Called(ctx, tenantID, appID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, appID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListForApplication provides a mock function with given fields: ctx, tenantID, applicationID, pageSize, cursor
func (_m *EventAPIRepository) ListForApplication(ctx context.Context, tenantID string, applicationID string, pageSize int, cursor string) (*model.EventDefinitionPage, error) {
	ret := _m.Called(ctx, tenantID, applicationID, pageSize, cursor)

	var r0 *model.EventDefinitionPage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) *model.EventDefinitionPage); ok {
		r0 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EventDefinitionPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, string) error); ok {
		r1 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

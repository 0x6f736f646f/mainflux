// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	things "github.com/absmach/magistrala/things"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// ChangeStatus provides a mock function with given fields: ctx, client
func (_m *Repository) ChangeStatus(ctx context.Context, client things.Client) (things.Client, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for ChangeStatus")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) (things.Client, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) things.Client); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RetrieveAll provides a mock function with given fields: ctx, pm
func (_m *Repository) RetrieveAll(ctx context.Context, pm things.Page) (things.ClientsPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 things.ClientsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) (things.ClientsPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) things.ClientsPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(things.ClientsPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAllByIDs provides a mock function with given fields: ctx, pm
func (_m *Repository) RetrieveAllByIDs(ctx context.Context, pm things.Page) (things.ClientsPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAllByIDs")
	}

	var r0 things.ClientsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) (things.ClientsPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) things.ClientsPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(things.ClientsPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveByID provides a mock function with given fields: ctx, id
func (_m *Repository) RetrieveByID(ctx context.Context, id string) (things.Client, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveByID")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (things.Client, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) things.Client); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBySecret provides a mock function with given fields: ctx, key
func (_m *Repository) RetrieveBySecret(ctx context.Context, key string) (things.Client, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBySecret")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (things.Client, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) things.Client); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, client
func (_m *Repository) Save(ctx context.Context, client ...things.Client) ([]things.Client, error) {
	_va := make([]interface{}, len(client))
	for _i := range client {
		_va[_i] = client[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 []things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...things.Client) ([]things.Client, error)); ok {
		return rf(ctx, client...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...things.Client) []things.Client); ok {
		r0 = rf(ctx, client...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]things.Client)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...things.Client) error); ok {
		r1 = rf(ctx, client...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchClients provides a mock function with given fields: ctx, pm
func (_m *Repository) SearchClients(ctx context.Context, pm things.Page) (things.ClientsPage, error) {
	ret := _m.Called(ctx, pm)

	if len(ret) == 0 {
		panic("no return value specified for SearchClients")
	}

	var r0 things.ClientsPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) (things.ClientsPage, error)); ok {
		return rf(ctx, pm)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Page) things.ClientsPage); ok {
		r0 = rf(ctx, pm)
	} else {
		r0 = ret.Get(0).(things.ClientsPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Page) error); ok {
		r1 = rf(ctx, pm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, client
func (_m *Repository) Update(ctx context.Context, client things.Client) (things.Client, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) (things.Client, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) things.Client); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateIdentity provides a mock function with given fields: ctx, client
func (_m *Repository) UpdateIdentity(ctx context.Context, client things.Client) (things.Client, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIdentity")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) (things.Client, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) things.Client); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateSecret provides a mock function with given fields: ctx, client
func (_m *Repository) UpdateSecret(ctx context.Context, client things.Client) (things.Client, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSecret")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) (things.Client, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) things.Client); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTags provides a mock function with given fields: ctx, client
func (_m *Repository) UpdateTags(ctx context.Context, client things.Client) (things.Client, error) {
	ret := _m.Called(ctx, client)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTags")
	}

	var r0 things.Client
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) (things.Client, error)); ok {
		return rf(ctx, client)
	}
	if rf, ok := ret.Get(0).(func(context.Context, things.Client) things.Client); ok {
		r0 = rf(ctx, client)
	} else {
		r0 = ret.Get(0).(things.Client)
	}

	if rf, ok := ret.Get(1).(func(context.Context, things.Client) error); ok {
		r1 = rf(ctx, client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

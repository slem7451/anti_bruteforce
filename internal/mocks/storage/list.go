// Code generated by mockery v2.53.3. DO NOT EDIT.

package appmock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// List is an autogenerated mock type for the list type
type List struct {
	mock.Mock
}

type List_Expecter struct {
	mock *mock.Mock
}

func (_m *List) EXPECT() *List_Expecter {
	return &List_Expecter{mock: &_m.Mock}
}

// AddToBlacklist provides a mock function with given fields: _a0, _a1
func (_m *List) AddToBlacklist(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddToBlacklist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List_AddToBlacklist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddToBlacklist'
type List_AddToBlacklist_Call struct {
	*mock.Call
}

// AddToBlacklist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) AddToBlacklist(_a0 interface{}, _a1 interface{}) *List_AddToBlacklist_Call {
	return &List_AddToBlacklist_Call{Call: _e.mock.On("AddToBlacklist", _a0, _a1)}
}

func (_c *List_AddToBlacklist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_AddToBlacklist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_AddToBlacklist_Call) Return(_a0 error) *List_AddToBlacklist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *List_AddToBlacklist_Call) RunAndReturn(run func(context.Context, string) error) *List_AddToBlacklist_Call {
	_c.Call.Return(run)
	return _c
}

// AddToWhitelist provides a mock function with given fields: _a0, _a1
func (_m *List) AddToWhitelist(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddToWhitelist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List_AddToWhitelist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddToWhitelist'
type List_AddToWhitelist_Call struct {
	*mock.Call
}

// AddToWhitelist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) AddToWhitelist(_a0 interface{}, _a1 interface{}) *List_AddToWhitelist_Call {
	return &List_AddToWhitelist_Call{Call: _e.mock.On("AddToWhitelist", _a0, _a1)}
}

func (_c *List_AddToWhitelist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_AddToWhitelist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_AddToWhitelist_Call) Return(_a0 error) *List_AddToWhitelist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *List_AddToWhitelist_Call) RunAndReturn(run func(context.Context, string) error) *List_AddToWhitelist_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFromBlacklist provides a mock function with given fields: _a0, _a1
func (_m *List) DeleteFromBlacklist(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFromBlacklist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List_DeleteFromBlacklist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFromBlacklist'
type List_DeleteFromBlacklist_Call struct {
	*mock.Call
}

// DeleteFromBlacklist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) DeleteFromBlacklist(_a0 interface{}, _a1 interface{}) *List_DeleteFromBlacklist_Call {
	return &List_DeleteFromBlacklist_Call{Call: _e.mock.On("DeleteFromBlacklist", _a0, _a1)}
}

func (_c *List_DeleteFromBlacklist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_DeleteFromBlacklist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_DeleteFromBlacklist_Call) Return(_a0 error) *List_DeleteFromBlacklist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *List_DeleteFromBlacklist_Call) RunAndReturn(run func(context.Context, string) error) *List_DeleteFromBlacklist_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteFromWhitelist provides a mock function with given fields: _a0, _a1
func (_m *List) DeleteFromWhitelist(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteFromWhitelist")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List_DeleteFromWhitelist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteFromWhitelist'
type List_DeleteFromWhitelist_Call struct {
	*mock.Call
}

// DeleteFromWhitelist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) DeleteFromWhitelist(_a0 interface{}, _a1 interface{}) *List_DeleteFromWhitelist_Call {
	return &List_DeleteFromWhitelist_Call{Call: _e.mock.On("DeleteFromWhitelist", _a0, _a1)}
}

func (_c *List_DeleteFromWhitelist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_DeleteFromWhitelist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_DeleteFromWhitelist_Call) Return(_a0 error) *List_DeleteFromWhitelist_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *List_DeleteFromWhitelist_Call) RunAndReturn(run func(context.Context, string) error) *List_DeleteFromWhitelist_Call {
	_c.Call.Return(run)
	return _c
}

// IsIPInBlacklist provides a mock function with given fields: _a0, _a1
func (_m *List) IsIPInBlacklist(_a0 context.Context, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsIPInBlacklist")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List_IsIPInBlacklist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsIPInBlacklist'
type List_IsIPInBlacklist_Call struct {
	*mock.Call
}

// IsIPInBlacklist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) IsIPInBlacklist(_a0 interface{}, _a1 interface{}) *List_IsIPInBlacklist_Call {
	return &List_IsIPInBlacklist_Call{Call: _e.mock.On("IsIPInBlacklist", _a0, _a1)}
}

func (_c *List_IsIPInBlacklist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_IsIPInBlacklist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_IsIPInBlacklist_Call) Return(_a0 bool, _a1 error) *List_IsIPInBlacklist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *List_IsIPInBlacklist_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *List_IsIPInBlacklist_Call {
	_c.Call.Return(run)
	return _c
}

// IsIPInWhitelist provides a mock function with given fields: _a0, _a1
func (_m *List) IsIPInWhitelist(_a0 context.Context, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsIPInWhitelist")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (bool, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List_IsIPInWhitelist_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsIPInWhitelist'
type List_IsIPInWhitelist_Call struct {
	*mock.Call
}

// IsIPInWhitelist is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *List_Expecter) IsIPInWhitelist(_a0 interface{}, _a1 interface{}) *List_IsIPInWhitelist_Call {
	return &List_IsIPInWhitelist_Call{Call: _e.mock.On("IsIPInWhitelist", _a0, _a1)}
}

func (_c *List_IsIPInWhitelist_Call) Run(run func(_a0 context.Context, _a1 string)) *List_IsIPInWhitelist_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *List_IsIPInWhitelist_Call) Return(_a0 bool, _a1 error) *List_IsIPInWhitelist_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *List_IsIPInWhitelist_Call) RunAndReturn(run func(context.Context, string) (bool, error)) *List_IsIPInWhitelist_Call {
	_c.Call.Return(run)
	return _c
}

// NewList creates a new instance of List. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewList(t interface {
	mock.TestingT
	Cleanup(func())
}) *List {
	mock := &List{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

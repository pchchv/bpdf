// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	entity "github.com/pchchv/bpdf/core/entity"
	mock "github.com/stretchr/testify/mock"
)

// Math is an autogenerated mock type for the Math type
type Math struct {
	mock.Mock
}

type Math_Expecter struct {
	mock *mock.Mock
}

func (_m *Math) EXPECT() *Math_Expecter {
	return &Math_Expecter{mock: &_m.Mock}
}

// GetInnerCenterCell provides a mock function with given fields: inner, outer
func (_m *Math) GetInnerCenterCell(inner *entity.Dimensions, outer *entity.Dimensions) *entity.Cell {
	ret := _m.Called(inner, outer)

	if len(ret) == 0 {
		panic("no return value specified for GetInnerCenterCell")
	}

	var r0 *entity.Cell
	if rf, ok := ret.Get(0).(func(*entity.Dimensions, *entity.Dimensions) *entity.Cell); ok {
		r0 = rf(inner, outer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Cell)
		}
	}

	return r0
}

// Math_GetInnerCenterCell_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInnerCenterCell'
type Math_GetInnerCenterCell_Call struct {
	*mock.Call
}

// GetInnerCenterCell is a helper method to define mock.On call
//   - inner *entity.Dimensions
//   - outer *entity.Dimensions
func (_e *Math_Expecter) GetInnerCenterCell(inner interface{}, outer interface{}) *Math_GetInnerCenterCell_Call {
	return &Math_GetInnerCenterCell_Call{Call: _e.mock.On("GetInnerCenterCell", inner, outer)}
}

func (_c *Math_GetInnerCenterCell_Call) Run(run func(inner *entity.Dimensions, outer *entity.Dimensions)) *Math_GetInnerCenterCell_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Dimensions), args[1].(*entity.Dimensions))
	})
	return _c
}

func (_c *Math_GetInnerCenterCell_Call) Return(_a0 *entity.Cell) *Math_GetInnerCenterCell_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Math_GetInnerCenterCell_Call) RunAndReturn(run func(*entity.Dimensions, *entity.Dimensions) *entity.Cell) *Math_GetInnerCenterCell_Call {
	_c.Call.Return(run)
	return _c
}

// Resize provides a mock function with given fields: inner, outer, percent, justReferenceWidth
func (_m *Math) Resize(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool) *entity.Dimensions {
	ret := _m.Called(inner, outer, percent, justReferenceWidth)

	if len(ret) == 0 {
		panic("no return value specified for Resize")
	}

	var r0 *entity.Dimensions
	if rf, ok := ret.Get(0).(func(*entity.Dimensions, *entity.Dimensions, float64, bool) *entity.Dimensions); ok {
		r0 = rf(inner, outer, percent, justReferenceWidth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Dimensions)
		}
	}

	return r0
}

// Math_Resize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Resize'
type Math_Resize_Call struct {
	*mock.Call
}

// Resize is a helper method to define mock.On call
//   - inner *entity.Dimensions
//   - outer *entity.Dimensions
//   - percent float64
//   - justReferenceWidth bool
func (_e *Math_Expecter) Resize(inner interface{}, outer interface{}, percent interface{}, justReferenceWidth interface{}) *Math_Resize_Call {
	return &Math_Resize_Call{Call: _e.mock.On("Resize", inner, outer, percent, justReferenceWidth)}
}

func (_c *Math_Resize_Call) Run(run func(inner *entity.Dimensions, outer *entity.Dimensions, percent float64, justReferenceWidth bool)) *Math_Resize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Dimensions), args[1].(*entity.Dimensions), args[2].(float64), args[3].(bool))
	})
	return _c
}

func (_c *Math_Resize_Call) Return(_a0 *entity.Dimensions) *Math_Resize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Math_Resize_Call) RunAndReturn(run func(*entity.Dimensions, *entity.Dimensions, float64, bool) *entity.Dimensions) *Math_Resize_Call {
	_c.Call.Return(run)
	return _c
}

// NewMath creates a new instance of Math. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMath(t interface {
	mock.TestingT
	Cleanup(func())
}) *Math {
	mock := &Math{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

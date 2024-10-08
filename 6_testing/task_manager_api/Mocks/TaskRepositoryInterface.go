// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "task_manager_api/Domain"

	mock "github.com/stretchr/testify/mock"
)

// TaskRepositoryInterface is an autogenerated mock type for the TaskRepositoryInterface type
type TaskRepositoryInterface struct {
	mock.Mock
}

// AddTask provides a mock function with given fields: c, newTask
func (_m *TaskRepositoryInterface) AddTask(c context.Context, newTask domain.Task) domain.CodedError {
	ret := _m.Called(c, newTask)

	if len(ret) == 0 {
		panic("no return value specified for AddTask")
	}

	var r0 domain.CodedError
	if rf, ok := ret.Get(0).(func(context.Context, domain.Task) domain.CodedError); ok {
		r0 = rf(c, newTask)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.CodedError)
		}
	}

	return r0
}

// DeleteTask provides a mock function with given fields: c, taskID
func (_m *TaskRepositoryInterface) DeleteTask(c context.Context, taskID string) domain.CodedError {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTask")
	}

	var r0 domain.CodedError
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.CodedError); ok {
		r0 = rf(c, taskID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.CodedError)
		}
	}

	return r0
}

// GetAllTasks provides a mock function with given fields: c
func (_m *TaskRepositoryInterface) GetAllTasks(c context.Context) ([]domain.Task, domain.CodedError) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTasks")
	}

	var r0 []domain.Task
	var r1 domain.CodedError
	if rf, ok := ret.Get(0).(func(context.Context) ([]domain.Task, domain.CodedError)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) domain.CodedError); ok {
		r1 = rf(c)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.CodedError)
		}
	}

	return r0, r1
}

// GetTaskByID provides a mock function with given fields: c, taskID
func (_m *TaskRepositoryInterface) GetTaskByID(c context.Context, taskID string) (domain.Task, domain.CodedError) {
	ret := _m.Called(c, taskID)

	if len(ret) == 0 {
		panic("no return value specified for GetTaskByID")
	}

	var r0 domain.Task
	var r1 domain.CodedError
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.Task, domain.CodedError)); ok {
		return rf(c, taskID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Task); ok {
		r0 = rf(c, taskID)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) domain.CodedError); ok {
		r1 = rf(c, taskID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.CodedError)
		}
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: c, taskID, updatedTask
func (_m *TaskRepositoryInterface) UpdateTask(c context.Context, taskID string, updatedTask domain.Task) (domain.Task, domain.CodedError) {
	ret := _m.Called(c, taskID, updatedTask)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTask")
	}

	var r0 domain.Task
	var r1 domain.CodedError
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Task) (domain.Task, domain.CodedError)); ok {
		return rf(c, taskID, updatedTask)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Task) domain.Task); ok {
		r0 = rf(c, taskID, updatedTask)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, domain.Task) domain.CodedError); ok {
		r1 = rf(c, taskID, updatedTask)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.CodedError)
		}
	}

	return r0, r1
}

// NewTaskRepositoryInterface creates a new instance of TaskRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTaskRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TaskRepositoryInterface {
	mock := &TaskRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

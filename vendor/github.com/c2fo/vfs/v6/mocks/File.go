// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import time "time"
import vfs "github.com/c2fo/vfs/v6"

// File is an autogenerated mock type for the File type
type File struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *File) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyToFile provides a mock function with given fields: file
func (_m *File) CopyToFile(file vfs.File) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(vfs.File) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyToLocation provides a mock function with given fields: location
func (_m *File) CopyToLocation(location vfs.Location) (vfs.File, error) {
	ret := _m.Called(location)

	var r0 vfs.File
	if rf, ok := ret.Get(0).(func(vfs.Location) vfs.File); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(vfs.Location) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields:
func (_m *File) Delete() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields:
func (_m *File) Exists() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LastModified provides a mock function with given fields:
func (_m *File) LastModified() (*time.Time, error) {
	ret := _m.Called()

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func() *time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Location provides a mock function with given fields:
func (_m *File) Location() vfs.Location {
	ret := _m.Called()

	var r0 vfs.Location
	if rf, ok := ret.Get(0).(func() vfs.Location); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.Location)
		}
	}

	return r0
}

// MoveToFile provides a mock function with given fields: file
func (_m *File) MoveToFile(file vfs.File) error {
	ret := _m.Called(file)

	var r0 error
	if rf, ok := ret.Get(0).(func(vfs.File) error); ok {
		r0 = rf(file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MoveToLocation provides a mock function with given fields: location
func (_m *File) MoveToLocation(location vfs.Location) (vfs.File, error) {
	ret := _m.Called(location)

	var r0 vfs.File
	if rf, ok := ret.Get(0).(func(vfs.Location) vfs.File); ok {
		r0 = rf(location)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(vfs.Location) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *File) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Path provides a mock function with given fields:
func (_m *File) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Read provides a mock function with given fields: p
func (_m *File) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Seek provides a mock function with given fields: offset, whence
func (_m *File) Seek(offset int64, whence int) (int64, error) {
	ret := _m.Called(offset, whence)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int64, int) int64); ok {
		r0 = rf(offset, whence)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int) error); ok {
		r1 = rf(offset, whence)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Size provides a mock function with given fields:
func (_m *File) Size() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// String provides a mock function with given fields:
func (_m *File) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Touch provides a mock function with given fields:
func (_m *File) Touch() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// URI provides a mock function with given fields:
func (_m *File) URI() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Write provides a mock function with given fields: p
func (_m *File) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

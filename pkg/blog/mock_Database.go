// Code generated by mockery v1.0.0. DO NOT EDIT.

package blog

import mock "github.com/stretchr/testify/mock"

// MockDatabase is an autogenerated mock type for the Database type
type MockDatabase struct {
	mock.Mock
}

// FindAllBlogPosts provides a mock function with given fields: _a0
func (_m *MockDatabase) FindAllBlogPosts(_a0 *RequestParams) ([]*PostData, error) {
	ret := _m.Called(_a0)

	var r0 []*PostData
	if rf, ok := ret.Get(0).(func(*RequestParams) []*PostData); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*PostData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*RequestParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindBlogPostByKey provides a mock function with given fields: _a0
func (_m *MockDatabase) FindBlogPostByKey(_a0 *RequestParams) (*PostData, error) {
	ret := _m.Called(_a0)

	var r0 *PostData
	if rf, ok := ret.Get(0).(func(*RequestParams) *PostData); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*PostData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*RequestParams) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveBlogPost provides a mock function with given fields: _a0
func (_m *MockDatabase) RemoveBlogPost(_a0 *RequestParams) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*RequestParams) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreBlogPost provides a mock function with given fields: _a0
func (_m *MockDatabase) StoreBlogPost(_a0 *PostData) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*PostData) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlogPost provides a mock function with given fields: _a0
func (_m *MockDatabase) UpdateBlogPost(_a0 *PostData) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*PostData) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
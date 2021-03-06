// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import dynalock "github.com/wolfeidau/dynalock"
import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AtomicDelete provides a mock function with given fields: key, previous
func (_m *Store) AtomicDelete(key string, previous *dynalock.KVPair) (bool, error) {
	ret := _m.Called(key, previous)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *dynalock.KVPair) bool); ok {
		r0 = rf(key, previous)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *dynalock.KVPair) error); ok {
		r1 = rf(key, previous)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtomicPut provides a mock function with given fields: key, options
func (_m *Store) AtomicPut(key string, options ...dynalock.WriteOption) (bool, *dynalock.KVPair, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...dynalock.WriteOption) bool); ok {
		r0 = rf(key, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *dynalock.KVPair
	if rf, ok := ret.Get(1).(func(string, ...dynalock.WriteOption) *dynalock.KVPair); ok {
		r1 = rf(key, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*dynalock.KVPair)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, ...dynalock.WriteOption) error); ok {
		r2 = rf(key, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Delete provides a mock function with given fields: key
func (_m *Store) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: key, options
func (_m *Store) Exists(key string, options ...dynalock.ReadOption) (bool, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, ...dynalock.ReadOption) bool); ok {
		r0 = rf(key, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...dynalock.ReadOption) error); ok {
		r1 = rf(key, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key, options
func (_m *Store) Get(key string, options ...dynalock.ReadOption) (*dynalock.KVPair, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynalock.KVPair
	if rf, ok := ret.Get(0).(func(string, ...dynalock.ReadOption) *dynalock.KVPair); ok {
		r0 = rf(key, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynalock.KVPair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...dynalock.ReadOption) error); ok {
		r1 = rf(key, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: prefix, options
func (_m *Store) List(prefix string, options ...dynalock.ReadOption) ([]*dynalock.KVPair, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, prefix)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*dynalock.KVPair
	if rf, ok := ret.Get(0).(func(string, ...dynalock.ReadOption) []*dynalock.KVPair); ok {
		r0 = rf(prefix, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dynalock.KVPair)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...dynalock.ReadOption) error); ok {
		r1 = rf(prefix, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLock provides a mock function with given fields: key, options
func (_m *Store) NewLock(key string, options ...dynalock.LockOption) (dynalock.Locker, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dynalock.Locker
	if rf, ok := ret.Get(0).(func(string, ...dynalock.LockOption) dynalock.Locker); ok {
		r0 = rf(key, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dynalock.Locker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...dynalock.LockOption) error); ok {
		r1 = rf(key, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: key, options
func (_m *Store) Put(key string, options ...dynalock.WriteOption) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, ...dynalock.WriteOption) error); ok {
		r0 = rf(key, options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

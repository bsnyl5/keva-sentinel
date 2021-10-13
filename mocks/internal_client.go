// Code generated by mockgen. DO NOT EDIT.
package sentinels

import (
	"github.com/keva-dev/go-sentinel"
	"github.com/stretchr/testify/mock"
)

type MockNoOpClient struct {
	mock.Mock
}

func (r *MockNoOpClient) Info() (string, error) {
	args := r.Called()
	return args.Get(0).(string), args.Error(1)
}

func (r *MockNoOpClient) Ping() (string, error) {
	args := r.Called()
	return args.Get(0).(string), args.Error(1)
}

func (r *MockNoOpClient) SlaveOf(arg1 string, arg2 string) error {
	args := r.Called(arg1, arg2)
	return args.Error(0)
}

func (r *MockNoOpClient) SlaveOfNoOne() error {
	args := r.Called()
	return args.Error(0)
}

func (r *MockNoOpClient) SubscribeHelloChan() sentinel.HelloChan {
	args := r.Called()
	return args.Get(0).(sentinel.HelloChan)
}
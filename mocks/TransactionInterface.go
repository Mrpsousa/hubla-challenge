// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/mrpsousa/api/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// TransactionInterface is an autogenerated mock type for the TransactionInterface type
type TransactionInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: transaction
func (_m *TransactionInterface) Create(transaction *entity.Transaction) error {
	ret := _m.Called(transaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Transaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/KraDM09/gophermart/internal/app/storage/models"
	mock "github.com/stretchr/testify/mock"

	pgx "github.com/jackc/pgx/v5"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// Begin provides a mock function with given fields: ctx
func (_m *Storage) Begin(ctx context.Context) (pgx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 pgx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (pgx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) pgx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pgx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrder provides a mock function with given fields: ctx, userID, number
func (_m *Storage) CreateOrder(ctx context.Context, userID int, number string) error {
	ret := _m.Called(ctx, userID, number)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) error); ok {
		r0 = rf(ctx, userID, number)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, login, password
func (_m *Storage) CreateUser(ctx context.Context, login string, password string) (*models.User, error) {
	ret := _m.Called(ctx, login, password)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.User, error)); ok {
		return rf(ctx, login, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, login, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWithdrawal provides a mock function with given fields: ctx, tx, userID, sum, order
func (_m *Storage) CreateWithdrawal(ctx context.Context, tx pgx.Tx, userID int, sum float32, order string) error {
	ret := _m.Called(ctx, tx, userID, sum, order)

	if len(ret) == 0 {
		panic("no return value specified for CreateWithdrawal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx, int, float32, string) error); ok {
		r0 = rf(ctx, tx, userID, sum, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrderByNumber provides a mock function with given fields: ctx, login
func (_m *Storage) GetOrderByNumber(ctx context.Context, login string) (*models.Order, error) {
	ret := _m.Called(ctx, login)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByNumber")
	}

	var r0 *models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.Order, error)); ok {
		return rf(ctx, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.Order); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrdersByUserID provides a mock function with given fields: ctx, userID
func (_m *Storage) GetOrdersByUserID(ctx context.Context, userID int) (*[]models.Order, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersByUserID")
	}

	var r0 *[]models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*[]models.Order, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *[]models.Order); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: ctx, userID
func (_m *Storage) GetUserByID(ctx context.Context, userID int) (*models.UserBalance, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *models.UserBalance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.UserBalance, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.UserBalance); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserBalance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByIDForUpdate provides a mock function with given fields: ctx, tx, userID
func (_m *Storage) GetUserByIDForUpdate(ctx context.Context, tx pgx.Tx, userID int) (*models.UserBalance, error) {
	ret := _m.Called(ctx, tx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByIDForUpdate")
	}

	var r0 *models.UserBalance
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx, int) (*models.UserBalance, error)); ok {
		return rf(ctx, tx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx, int) *models.UserBalance); ok {
		r0 = rf(ctx, tx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserBalance)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, pgx.Tx, int) error); ok {
		r1 = rf(ctx, tx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByLogin provides a mock function with given fields: ctx, login
func (_m *Storage) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	ret := _m.Called(ctx, login)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByLogin")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, login)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithdrawals provides a mock function with given fields: ctx, userID
func (_m *Storage) GetWithdrawals(ctx context.Context, userID int) (*[]models.Withdrawal, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetWithdrawals")
	}

	var r0 *[]models.Withdrawal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*[]models.Withdrawal, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *[]models.Withdrawal); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]models.Withdrawal)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: ctx, tx, sum, gamblerID
func (_m *Storage) UpdateBalance(ctx context.Context, tx pgx.Tx, sum float32, gamblerID int) error {
	ret := _m.Called(ctx, tx, sum, gamblerID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx, float32, int) error); ok {
		r0 = rf(ctx, tx, sum, gamblerID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOrder provides a mock function with given fields: ctx, tx, status, number, accrual
func (_m *Storage) UpdateOrder(ctx context.Context, tx pgx.Tx, status string, number string, accrual *float32) error {
	ret := _m.Called(ctx, tx, status, number, accrual)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, pgx.Tx, string, string, *float32) error); ok {
		r0 = rf(ctx, tx, status, number, accrual)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v1.0.0
package mocks

import (
	context "context"

	"github.com/models"

	mock "github.com/stretchr/testify/mock"
)

// Usecase is an autogenerated mock type for the Usecase type
type Usecase struct {
	mock.Mock
}
// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetUserByReferralCode(ctx context.Context,code string)(*models.UserDto,error) {
	ret := _m.Called(ctx,code)

	var r0 *models.UserDto
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserDto); ok {
		r0 = rf(ctx, code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) Update(ctx context.Context, ar *models.NewCommandUser, isAdmin bool ,token string) error {
	ret := _m.Called(ctx, ar, isAdmin,token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandUser, bool,string) error); ok {
		r0 = rf(ctx, ar, isAdmin,token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) Create(ctx context.Context, ar *models.NewCommandUser, isAdmin bool,token string) (*models.NewCommandUser,error) {
	ret := _m.Called(ctx, ar, isAdmin,token)

	var r0 *models.NewCommandUser
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandUser,bool,string) *models.NewCommandUser); ok {
		r0 = rf(ctx, ar,isAdmin,token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.NewCommandUser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandUser,bool,string) error); ok {
		r1 = rf(ctx, ar,isAdmin,token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) Login(ctx context.Context, ar *models.Login) (*models.GetToken, error) {
	ret := _m.Called(ctx, ar)

	var r0 *models.GetToken
	if rf, ok := ret.Get(0).(func(context.Context, *models.Login) *models.GetToken); ok {
		r0 = rf(ctx, ar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GetToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Login) error); ok {
		r1 = rf(ctx, ar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) ValidateTokenUser(ctx context.Context, token string) (*models.UserInfoDto, error) {
	ret := _m.Called(ctx, token)

	var r0 *models.UserInfoDto
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserInfoDto); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInfoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) GetUserInfo(ctx context.Context, token string,orderId string) (*models.UserInfoDto, error) {
	ret := _m.Called(ctx, token,orderId)

	var r0 *models.UserInfoDto
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.UserInfoDto); ok {
		r0 = rf(ctx, token,orderId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInfoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, token,orderId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) ChangePassword(ctx context.Context,token string,password string,email string,phoneNumber string)(*models.ResponseDelete,error) {
	ret := _m.Called(ctx, token, password,email,phoneNumber)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, string,string, string,string) *models.ResponseDelete); ok {
		r0 = rf(ctx, token,password,email,phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string, string,string) error); ok {
		r1 = rf(ctx, token,password,email,phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Usecase) CheckEmailORPhoneNumber(ctx context.Context,email string,phoneNumber string,otp string)(*models.ResponseDelete,error) {
	ret := _m.Called(ctx, email,phoneNumber,otp)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, string,string,string) *models.ResponseDelete); ok {
		r0 = rf(ctx,email,phoneNumber,otp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context,string,string,string) error); ok {
		r1 = rf(ctx, email,phoneNumber,otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetUserByEmail(ctx context.Context,email string, loginType string)(*models.UserDto,error) {
	ret := _m.Called(ctx,email,loginType)

	var r0 *models.UserDto
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.UserDto); ok {
		r0 = rf(ctx, email,loginType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, email,loginType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) LoginByGoogle(ctx context.Context,loginType string,email string,profilePicture string,name string)(*models.GetToken, error) {
	ret := _m.Called(ctx, loginType,email,profilePicture,name)

	var r0 *models.GetToken
	if rf, ok := ret.Get(0).(func(context.Context, string,string,string,string) *models.GetToken); ok {
		r0 = rf(ctx, loginType,email,profilePicture,name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.GetToken)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string,string,string) error); ok {
		r1 = rf(ctx, loginType,email,profilePicture,name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) Delete(ctx context.Context,userId string,token string)(*models.ResponseDelete, error) {
	ret := _m.Called(ctx, userId,token)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.ResponseDelete); ok {
		r0 = rf(ctx, userId,token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, userId,token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) RequestOTP(ctx context.Context,phoneNumber string)(*models.RequestOTP,error) {
	ret := _m.Called(ctx, phoneNumber)

	var r0 *models.RequestOTP
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.RequestOTP); ok {
		r0 = rf(ctx, phoneNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.RequestOTP)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phoneNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) VerifiedEmail(ctx context.Context, token string, codeOTP string) (*models.UserInfoDto, error) {
	ret := _m.Called(ctx, token, codeOTP)

	var r0 *models.UserInfoDto
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.UserInfoDto); ok {
		r0 = rf(ctx, token,codeOTP)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserInfoDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, token,codeOTP)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetch provides a mock function with given fields: ctx, cursor, num
func (_m *Usecase) GetCreditByID(ctx context.Context, id string) (*models.UserPoint, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.UserPoint
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.UserPoint); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserPoint)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Usecase) List(ctx context.Context, page, limit, offset int,search string) (*models.UserWithPagination, error) {
	ret := _m.Called(ctx, page,limit,offset,search)

	var r0 *models.UserWithPagination
	if rf, ok := ret.Get(0).(func(context.Context, int,int,int,string) *models.UserWithPagination); ok {
		r0 = rf(ctx, page,limit,offset,search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserWithPagination)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int,int,int,string) error); ok {
		r1 = rf(ctx, page,limit,offset,search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByTitle provides a mock function with given fields: ctx, title
func (_m *Usecase) GetUserDetailById(ctx context.Context,id string,token string)(*models.UserDto, error) {
	ret := _m.Called(ctx, id,token)

	var r0 *models.UserDto
	if rf, ok := ret.Get(0).(func(context.Context, string,string) *models.UserDto); ok {
		r0 = rf(ctx,id, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.UserDto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string,string) error); ok {
		r1 = rf(ctx, id,token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: _a0, _a1
func (_m *Usecase) Subscription(ctx context.Context, s *models.NewCommandSubscribe) (*models.ResponseDelete, error) {
	ret := _m.Called(ctx, s)

	var r0 *models.ResponseDelete
	if rf, ok := ret.Get(0).(func(context.Context, *models.NewCommandSubscribe) *models.ResponseDelete); ok {
		r0 = rf(ctx, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResponseDelete)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.NewCommandSubscribe) error); ok {
		r1 = rf(ctx, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	auth "github.com/smartcontractkit/chainlink/v2/core/auth"
	bridges "github.com/smartcontractkit/chainlink/v2/core/bridges"

	context "context"

	mock "github.com/stretchr/testify/mock"

	sessions "github.com/smartcontractkit/chainlink/v2/core/sessions"
)

// AuthenticationProvider is an autogenerated mock type for the AuthenticationProvider type
type AuthenticationProvider struct {
	mock.Mock
}

// AuthorizedUserWithSession provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) AuthorizedUserWithSession(ctx context.Context, sessionID string) (sessions.User, error) {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for AuthorizedUserWithSession")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, sessionID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClearNonCurrentSessions provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) ClearNonCurrentSessions(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for ClearNonCurrentSessions")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateAndSetAuthToken provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) CreateAndSetAuthToken(ctx context.Context, user *sessions.User) (*auth.Token, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateAndSetAuthToken")
	}

	var r0 *auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) (*auth.Token, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) *auth.Token); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auth.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *sessions.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateSession provides a mock function with given fields: ctx, sr
func (_m *AuthenticationProvider) CreateSession(ctx context.Context, sr sessions.SessionRequest) (string, error) {
	ret := _m.Called(ctx, sr)

	if len(ret) == 0 {
		panic("no return value specified for CreateSession")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, sessions.SessionRequest) (string, error)); ok {
		return rf(ctx, sr)
	}
	if rf, ok := ret.Get(0).(func(context.Context, sessions.SessionRequest) string); ok {
		r0 = rf(ctx, sr)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, sessions.SessionRequest) error); ok {
		r1 = rf(ctx, sr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) CreateUser(ctx context.Context, user *sessions.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAuthToken provides a mock function with given fields: ctx, user
func (_m *AuthenticationProvider) DeleteAuthToken(ctx context.Context, user *sessions.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) DeleteUser(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUserSession provides a mock function with given fields: ctx, sessionID
func (_m *AuthenticationProvider) DeleteUserSession(ctx context.Context, sessionID string) error {
	ret := _m.Called(ctx, sessionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserSession")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, sessionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindExternalInitiator provides a mock function with given fields: ctx, eia
func (_m *AuthenticationProvider) FindExternalInitiator(ctx context.Context, eia *auth.Token) (*bridges.ExternalInitiator, error) {
	ret := _m.Called(ctx, eia)

	if len(ret) == 0 {
		panic("no return value specified for FindExternalInitiator")
	}

	var r0 *bridges.ExternalInitiator
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) (*bridges.ExternalInitiator, error)); ok {
		return rf(ctx, eia)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *auth.Token) *bridges.ExternalInitiator); ok {
		r0 = rf(ctx, eia)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bridges.ExternalInitiator)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *auth.Token) error); ok {
		r1 = rf(ctx, eia)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) FindUser(ctx context.Context, email string) (sessions.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindUser")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByAPIToken provides a mock function with given fields: ctx, apiToken
func (_m *AuthenticationProvider) FindUserByAPIToken(ctx context.Context, apiToken string) (sessions.User, error) {
	ret := _m.Called(ctx, apiToken)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByAPIToken")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (sessions.User, error)); ok {
		return rf(ctx, apiToken)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) sessions.User); ok {
		r0 = rf(ctx, apiToken)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, apiToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserWebAuthn provides a mock function with given fields: ctx, email
func (_m *AuthenticationProvider) GetUserWebAuthn(ctx context.Context, email string) ([]sessions.WebAuthn, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserWebAuthn")
	}

	var r0 []sessions.WebAuthn
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]sessions.WebAuthn, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []sessions.WebAuthn); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.WebAuthn)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx
func (_m *AuthenticationProvider) ListUsers(ctx context.Context) ([]sessions.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]sessions.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []sessions.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveWebAuthn provides a mock function with given fields: ctx, token
func (_m *AuthenticationProvider) SaveWebAuthn(ctx context.Context, token *sessions.WebAuthn) error {
	ret := _m.Called(ctx, token)

	if len(ret) == 0 {
		panic("no return value specified for SaveWebAuthn")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.WebAuthn) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sessions provides a mock function with given fields: ctx, offset, limit
func (_m *AuthenticationProvider) Sessions(ctx context.Context, offset int, limit int) ([]sessions.Session, error) {
	ret := _m.Called(ctx, offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for Sessions")
	}

	var r0 []sessions.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]sessions.Session, error)); ok {
		return rf(ctx, offset, limit)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []sessions.Session); ok {
		r0 = rf(ctx, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]sessions.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAuthToken provides a mock function with given fields: ctx, user, token
func (_m *AuthenticationProvider) SetAuthToken(ctx context.Context, user *sessions.User, token *auth.Token) error {
	ret := _m.Called(ctx, user, token)

	if len(ret) == 0 {
		panic("no return value specified for SetAuthToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User, *auth.Token) error); ok {
		r0 = rf(ctx, user, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetPassword provides a mock function with given fields: ctx, user, newPassword
func (_m *AuthenticationProvider) SetPassword(ctx context.Context, user *sessions.User, newPassword string) error {
	ret := _m.Called(ctx, user, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for SetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sessions.User, string) error); ok {
		r0 = rf(ctx, user, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TestPassword provides a mock function with given fields: ctx, email, password
func (_m *AuthenticationProvider) TestPassword(ctx context.Context, email string, password string) error {
	ret := _m.Called(ctx, email, password)

	if len(ret) == 0 {
		panic("no return value specified for TestPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, email, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateRole provides a mock function with given fields: ctx, email, newRole
func (_m *AuthenticationProvider) UpdateRole(ctx context.Context, email string, newRole string) (sessions.User, error) {
	ret := _m.Called(ctx, email, newRole)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 sessions.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (sessions.User, error)); ok {
		return rf(ctx, email, newRole)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) sessions.User); ok {
		r0 = rf(ctx, email, newRole)
	} else {
		r0 = ret.Get(0).(sessions.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, newRole)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthenticationProvider creates a new instance of AuthenticationProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthenticationProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthenticationProvider {
	mock := &AuthenticationProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
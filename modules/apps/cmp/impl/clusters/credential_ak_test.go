// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda-proto-go/core/token/pb (interfaces: TokenServiceServer)

// Package clusters is a generated GoMock package.
package clusters

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/erda-project/erda-proto-go/core/token/pb"
)

// MockTokenServiceServer is a mock of TokenServiceServer interface.
type MockTokenServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceServerMockRecorder
}

// MockTokenServiceServerMockRecorder is the mock recorder for MockTokenServiceServer.
type MockTokenServiceServerMockRecorder struct {
	mock *MockTokenServiceServer
}

// NewMockTokenServiceServer creates a new mock instance.
func NewMockTokenServiceServer(ctrl *gomock.Controller) *MockTokenServiceServer {
	mock := &MockTokenServiceServer{ctrl: ctrl}
	mock.recorder = &MockTokenServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenServiceServer) EXPECT() *MockTokenServiceServerMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MockTokenServiceServer) CreateToken(arg0 context.Context, arg1 *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", arg0, arg1)
	ret0, _ := ret[0].(*pb.CreateTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockTokenServiceServerMockRecorder) CreateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockTokenServiceServer)(nil).CreateToken), arg0, arg1)
}

// DeleteToken mocks base method.
func (m *MockTokenServiceServer) DeleteToken(arg0 context.Context, arg1 *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteToken", arg0, arg1)
	ret0, _ := ret[0].(*pb.DeleteTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteToken indicates an expected call of DeleteToken.
func (mr *MockTokenServiceServerMockRecorder) DeleteToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteToken", reflect.TypeOf((*MockTokenServiceServer)(nil).DeleteToken), arg0, arg1)
}

// GetToken mocks base method.
func (m *MockTokenServiceServer) GetToken(arg0 context.Context, arg1 *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetToken", arg0, arg1)
	ret0, _ := ret[0].(*pb.GetTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetToken indicates an expected call of GetToken.
func (mr *MockTokenServiceServerMockRecorder) GetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetToken", reflect.TypeOf((*MockTokenServiceServer)(nil).GetToken), arg0, arg1)
}

// QueryTokens mocks base method.
func (m *MockTokenServiceServer) QueryTokens(arg0 context.Context, arg1 *pb.QueryTokensRequest) (*pb.QueryTokensResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTokens", arg0, arg1)
	ret0, _ := ret[0].(*pb.QueryTokensResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTokens indicates an expected call of QueryTokens.
func (mr *MockTokenServiceServerMockRecorder) QueryTokens(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTokens", reflect.TypeOf((*MockTokenServiceServer)(nil).QueryTokens), arg0, arg1)
}

// UpdateToken mocks base method.
func (m *MockTokenServiceServer) UpdateToken(arg0 context.Context, arg1 *pb.UpdateTokenRequest) (*pb.UpdateTokenResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateToken", arg0, arg1)
	ret0, _ := ret[0].(*pb.UpdateTokenResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateToken indicates an expected call of UpdateToken.
func (mr *MockTokenServiceServerMockRecorder) UpdateToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateToken", reflect.TypeOf((*MockTokenServiceServer)(nil).UpdateToken), arg0, arg1)
}

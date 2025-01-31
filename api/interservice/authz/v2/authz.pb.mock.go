// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: api/interservice/authz/v2/authz.proto

package v2

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the AuthorizationServer interface (at compile time)
var _ AuthorizationServer = &AuthorizationServerMock{}

// NewAuthorizationServerMock gives you a fresh instance of AuthorizationServerMock.
func NewAuthorizationServerMock() *AuthorizationServerMock {
	return &AuthorizationServerMock{validateRequests: true}
}

// NewAuthorizationServerMockWithoutValidation gives you a fresh instance of
// AuthorizationServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewAuthorizationServerMockWithoutValidation() *AuthorizationServerMock {
	return &AuthorizationServerMock{}
}

// AuthorizationServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type AuthorizationServerMock struct {
	validateRequests              bool
	IsAuthorizedFunc              func(context.Context, *IsAuthorizedReq) (*IsAuthorizedResp, error)
	FilterAuthorizedPairsFunc     func(context.Context, *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error)
	FilterAuthorizedProjectsFunc  func(context.Context, *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error)
	ProjectsAuthorizedFunc        func(context.Context, *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error)
	ValidateProjectAssignmentFunc func(context.Context, *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error)
}

func (m *AuthorizationServerMock) IsAuthorized(ctx context.Context, req *IsAuthorizedReq) (*IsAuthorizedResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.IsAuthorizedFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'IsAuthorized' not implemented")
}

func (m *AuthorizationServerMock) FilterAuthorizedPairs(ctx context.Context, req *FilterAuthorizedPairsReq) (*FilterAuthorizedPairsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.FilterAuthorizedPairsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'FilterAuthorizedPairs' not implemented")
}

func (m *AuthorizationServerMock) FilterAuthorizedProjects(ctx context.Context, req *FilterAuthorizedProjectsReq) (*FilterAuthorizedProjectsResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.FilterAuthorizedProjectsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'FilterAuthorizedProjects' not implemented")
}

func (m *AuthorizationServerMock) ProjectsAuthorized(ctx context.Context, req *ProjectsAuthorizedReq) (*ProjectsAuthorizedResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ProjectsAuthorizedFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ProjectsAuthorized' not implemented")
}

func (m *AuthorizationServerMock) ValidateProjectAssignment(ctx context.Context, req *ValidateProjectAssignmentReq) (*ValidateProjectAssignmentResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ValidateProjectAssignmentFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ValidateProjectAssignment' not implemented")
}

// Reset resets all overridden functions
func (m *AuthorizationServerMock) Reset() {
	m.IsAuthorizedFunc = nil
	m.FilterAuthorizedPairsFunc = nil
	m.FilterAuthorizedProjectsFunc = nil
	m.ProjectsAuthorizedFunc = nil
	m.ValidateProjectAssignmentFunc = nil
}

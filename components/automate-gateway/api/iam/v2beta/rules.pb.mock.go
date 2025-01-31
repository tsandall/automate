// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2beta/rules.proto

package v2beta

import (
	"context"

	request "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2beta/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the RulesServer interface (at compile time)
var _ RulesServer = &RulesServerMock{}

// NewRulesServerMock gives you a fresh instance of RulesServerMock.
func NewRulesServerMock() *RulesServerMock {
	return &RulesServerMock{validateRequests: true}
}

// NewRulesServerMockWithoutValidation gives you a fresh instance of
// RulesServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewRulesServerMockWithoutValidation() *RulesServerMock {
	return &RulesServerMock{}
}

// RulesServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type RulesServerMock struct {
	validateRequests        bool
	CreateRuleFunc          func(context.Context, *request.CreateRuleReq) (*response.CreateRuleResp, error)
	UpdateRuleFunc          func(context.Context, *request.UpdateRuleReq) (*response.UpdateRuleResp, error)
	GetRuleFunc             func(context.Context, *request.GetRuleReq) (*response.GetRuleResp, error)
	ListRulesForProjectFunc func(context.Context, *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error)
	DeleteRuleFunc          func(context.Context, *request.DeleteRuleReq) (*response.DeleteRuleResp, error)
	ApplyRulesStartFunc     func(context.Context, *request.ApplyRulesStartReq) (*response.ApplyRulesStartResp, error)
	ApplyRulesCancelFunc    func(context.Context, *request.ApplyRulesCancelReq) (*response.ApplyRulesCancelResp, error)
	ApplyRulesStatusFunc    func(context.Context, *request.ApplyRulesStatusReq) (*response.ApplyRulesStatusResp, error)
}

func (m *RulesServerMock) CreateRule(ctx context.Context, req *request.CreateRuleReq) (*response.CreateRuleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateRuleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateRule' not implemented")
}

func (m *RulesServerMock) UpdateRule(ctx context.Context, req *request.UpdateRuleReq) (*response.UpdateRuleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateRuleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateRule' not implemented")
}

func (m *RulesServerMock) GetRule(ctx context.Context, req *request.GetRuleReq) (*response.GetRuleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetRuleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetRule' not implemented")
}

func (m *RulesServerMock) ListRulesForProject(ctx context.Context, req *request.ListRulesForProjectReq) (*response.ListRulesForProjectResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ListRulesForProjectFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ListRulesForProject' not implemented")
}

func (m *RulesServerMock) DeleteRule(ctx context.Context, req *request.DeleteRuleReq) (*response.DeleteRuleResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteRuleFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteRule' not implemented")
}

func (m *RulesServerMock) ApplyRulesStart(ctx context.Context, req *request.ApplyRulesStartReq) (*response.ApplyRulesStartResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ApplyRulesStartFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ApplyRulesStart' not implemented")
}

func (m *RulesServerMock) ApplyRulesCancel(ctx context.Context, req *request.ApplyRulesCancelReq) (*response.ApplyRulesCancelResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ApplyRulesCancelFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ApplyRulesCancel' not implemented")
}

func (m *RulesServerMock) ApplyRulesStatus(ctx context.Context, req *request.ApplyRulesStatusReq) (*response.ApplyRulesStatusResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.ApplyRulesStatusFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'ApplyRulesStatus' not implemented")
}

// Reset resets all overridden functions
func (m *RulesServerMock) Reset() {
	m.CreateRuleFunc = nil
	m.UpdateRuleFunc = nil
	m.GetRuleFunc = nil
	m.ListRulesForProjectFunc = nil
	m.DeleteRuleFunc = nil
	m.ApplyRulesStartFunc = nil
	m.ApplyRulesCancelFunc = nil
	m.ApplyRulesStatusFunc = nil
}

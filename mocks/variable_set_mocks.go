// Code generated by MockGen. DO NOT EDIT.
// Source: variable_set.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockVariableSets is a mock of VariableSets interface.
type MockVariableSets struct {
	ctrl     *gomock.Controller
	recorder *MockVariableSetsMockRecorder
}

// MockVariableSetsMockRecorder is the mock recorder for MockVariableSets.
type MockVariableSetsMockRecorder struct {
	mock *MockVariableSets
}

// NewMockVariableSets creates a new mock instance.
func NewMockVariableSets(ctrl *gomock.Controller) *MockVariableSets {
	mock := &MockVariableSets{ctrl: ctrl}
	mock.recorder = &MockVariableSetsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVariableSets) EXPECT() *MockVariableSetsMockRecorder {
	return m.recorder
}

// ApplyToProjects mocks base method.
func (m *MockVariableSets) ApplyToProjects(ctx context.Context, variableSetID string, options tfe.VariableSetApplyToProjectsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyToProjects", ctx, variableSetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyToProjects indicates an expected call of ApplyToProjects.
func (mr *MockVariableSetsMockRecorder) ApplyToProjects(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyToProjects", reflect.TypeOf((*MockVariableSets)(nil).ApplyToProjects), ctx, variableSetID, options)
}

// ApplyToWorkspaces mocks base method.
func (m *MockVariableSets) ApplyToWorkspaces(ctx context.Context, variableSetID string, options *tfe.VariableSetApplyToWorkspacesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyToWorkspaces", ctx, variableSetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyToWorkspaces indicates an expected call of ApplyToWorkspaces.
func (mr *MockVariableSetsMockRecorder) ApplyToWorkspaces(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyToWorkspaces", reflect.TypeOf((*MockVariableSets)(nil).ApplyToWorkspaces), ctx, variableSetID, options)
}

// Create mocks base method.
func (m *MockVariableSets) Create(ctx context.Context, organization string, options *tfe.VariableSetCreateOptions) (*tfe.VariableSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.VariableSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVariableSetsMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVariableSets)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockVariableSets) Delete(ctx context.Context, variableSetID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, variableSetID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVariableSetsMockRecorder) Delete(ctx, variableSetID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVariableSets)(nil).Delete), ctx, variableSetID)
}

// List mocks base method.
func (m *MockVariableSets) List(ctx context.Context, organization string, options *tfe.VariableSetListOptions) (*tfe.VariableSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.VariableSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVariableSetsMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVariableSets)(nil).List), ctx, organization, options)
}

// ListForProject mocks base method.
func (m *MockVariableSets) ListForProject(ctx context.Context, projectID string, options *tfe.VariableSetListOptions) (*tfe.VariableSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForProject", ctx, projectID, options)
	ret0, _ := ret[0].(*tfe.VariableSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForProject indicates an expected call of ListForProject.
func (mr *MockVariableSetsMockRecorder) ListForProject(ctx, projectID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForProject", reflect.TypeOf((*MockVariableSets)(nil).ListForProject), ctx, projectID, options)
}

// ListForWorkspace mocks base method.
func (m *MockVariableSets) ListForWorkspace(ctx context.Context, workspaceID string, options *tfe.VariableSetListOptions) (*tfe.VariableSetList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListForWorkspace", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.VariableSetList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListForWorkspace indicates an expected call of ListForWorkspace.
func (mr *MockVariableSetsMockRecorder) ListForWorkspace(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListForWorkspace", reflect.TypeOf((*MockVariableSets)(nil).ListForWorkspace), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockVariableSets) Read(ctx context.Context, variableSetID string, options *tfe.VariableSetReadOptions) (*tfe.VariableSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, variableSetID, options)
	ret0, _ := ret[0].(*tfe.VariableSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockVariableSetsMockRecorder) Read(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockVariableSets)(nil).Read), ctx, variableSetID, options)
}

// RemoveFromProjects mocks base method.
func (m *MockVariableSets) RemoveFromProjects(ctx context.Context, variableSetID string, options tfe.VariableSetRemoveFromProjectsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromProjects", ctx, variableSetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromProjects indicates an expected call of RemoveFromProjects.
func (mr *MockVariableSetsMockRecorder) RemoveFromProjects(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromProjects", reflect.TypeOf((*MockVariableSets)(nil).RemoveFromProjects), ctx, variableSetID, options)
}

// RemoveFromWorkspaces mocks base method.
func (m *MockVariableSets) RemoveFromWorkspaces(ctx context.Context, variableSetID string, options *tfe.VariableSetRemoveFromWorkspacesOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromWorkspaces", ctx, variableSetID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromWorkspaces indicates an expected call of RemoveFromWorkspaces.
func (mr *MockVariableSetsMockRecorder) RemoveFromWorkspaces(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromWorkspaces", reflect.TypeOf((*MockVariableSets)(nil).RemoveFromWorkspaces), ctx, variableSetID, options)
}

// Update mocks base method.
func (m *MockVariableSets) Update(ctx context.Context, variableSetID string, options *tfe.VariableSetUpdateOptions) (*tfe.VariableSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, variableSetID, options)
	ret0, _ := ret[0].(*tfe.VariableSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockVariableSetsMockRecorder) Update(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVariableSets)(nil).Update), ctx, variableSetID, options)
}

// UpdateWorkspaces mocks base method.
func (m *MockVariableSets) UpdateWorkspaces(ctx context.Context, variableSetID string, options *tfe.VariableSetUpdateWorkspacesOptions) (*tfe.VariableSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWorkspaces", ctx, variableSetID, options)
	ret0, _ := ret[0].(*tfe.VariableSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWorkspaces indicates an expected call of UpdateWorkspaces.
func (mr *MockVariableSetsMockRecorder) UpdateWorkspaces(ctx, variableSetID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWorkspaces", reflect.TypeOf((*MockVariableSets)(nil).UpdateWorkspaces), ctx, variableSetID, options)
}

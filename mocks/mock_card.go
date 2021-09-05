// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonathan-buttner/game-framework/internal/deck (interfaces: Card,CardAction)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	deck "github.com/jonathan-buttner/game-framework/internal/deck"
	resource "github.com/jonathan-buttner/game-framework/internal/resource"
)

// MockCard is a mock of Card interface.
type MockCard struct {
	ctrl     *gomock.Controller
	recorder *MockCardMockRecorder
}

// MockCardMockRecorder is the mock recorder for MockCard.
type MockCardMockRecorder struct {
	mock *MockCard
}

// NewMockCard creates a new mock instance.
func NewMockCard(ctrl *gomock.Controller) *MockCard {
	mock := &MockCard{ctrl: ctrl}
	mock.recorder = &MockCardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCard) EXPECT() *MockCardMockRecorder {
	return m.recorder
}

// Cost mocks base method.
func (m *MockCard) Cost() resource.GroupedResources {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cost")
	ret0, _ := ret[0].(resource.GroupedResources)
	return ret0
}

// Cost indicates an expected call of Cost.
func (mr *MockCardMockRecorder) Cost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cost", reflect.TypeOf((*MockCard)(nil).Cost))
}

// GetOrientationAction mocks base method.
func (m *MockCard) GetOrientationAction(arg0 deck.CardOrientation) deck.CardAction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrientationAction", arg0)
	ret0, _ := ret[0].(deck.CardAction)
	return ret0
}

// GetOrientationAction indicates an expected call of GetOrientationAction.
func (mr *MockCardMockRecorder) GetOrientationAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrientationAction", reflect.TypeOf((*MockCard)(nil).GetOrientationAction), arg0)
}

// ID mocks base method.
func (m *MockCard) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockCardMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockCard)(nil).ID))
}

// IsOrientationValid mocks base method.
func (m *MockCard) IsOrientationValid(arg0 deck.CardOrientation) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOrientationValid", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOrientationValid indicates an expected call of IsOrientationValid.
func (mr *MockCardMockRecorder) IsOrientationValid(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOrientationValid", reflect.TypeOf((*MockCard)(nil).IsOrientationValid), arg0)
}

// MockCardAction is a mock of CardAction interface.
type MockCardAction struct {
	ctrl     *gomock.Controller
	recorder *MockCardActionMockRecorder
}

// MockCardActionMockRecorder is the mock recorder for MockCardAction.
type MockCardActionMockRecorder struct {
	mock *MockCardAction
}

// NewMockCardAction creates a new mock instance.
func NewMockCardAction(ctrl *gomock.Controller) *MockCardAction {
	mock := &MockCardAction{ctrl: ctrl}
	mock.recorder = &MockCardActionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCardAction) EXPECT() *MockCardActionMockRecorder {
	return m.recorder
}

// PerformEndRoundAction mocks base method.
func (m *MockCardAction) PerformEndRoundAction(arg0 deck.Game) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PerformEndRoundAction", arg0)
}

// PerformEndRoundAction indicates an expected call of PerformEndRoundAction.
func (mr *MockCardActionMockRecorder) PerformEndRoundAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformEndRoundAction", reflect.TypeOf((*MockCardAction)(nil).PerformEndRoundAction), arg0)
}

// PerformEndTurnAction mocks base method.
func (m *MockCardAction) PerformEndTurnAction(arg0 deck.Game) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PerformEndTurnAction", arg0)
}

// PerformEndTurnAction indicates an expected call of PerformEndTurnAction.
func (mr *MockCardActionMockRecorder) PerformEndTurnAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformEndTurnAction", reflect.TypeOf((*MockCardAction)(nil).PerformEndTurnAction), arg0)
}

// PerformPlayToTableaAction mocks base method.
func (m *MockCardAction) PerformPlayToTableaAction(arg0 deck.Game) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PerformPlayToTableaAction", arg0)
}

// PerformPlayToTableaAction indicates an expected call of PerformPlayToTableaAction.
func (mr *MockCardActionMockRecorder) PerformPlayToTableaAction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformPlayToTableaAction", reflect.TypeOf((*MockCardAction)(nil).PerformPlayToTableaAction), arg0)
}

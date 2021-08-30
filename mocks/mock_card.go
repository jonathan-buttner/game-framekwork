// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonathan-buttner/game-framework/internal/deck (interfaces: Card)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	core "github.com/jonathan-buttner/game-framework/internal/core"
	deck "github.com/jonathan-buttner/game-framework/internal/deck"
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

// PerformEndRoundAction mocks base method.
func (m *MockCard) PerformEndRoundAction(arg0 core.Game, arg1 deck.CardOrientation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PerformEndRoundAction", arg0, arg1)
}

// PerformEndRoundAction indicates an expected call of PerformEndRoundAction.
func (mr *MockCardMockRecorder) PerformEndRoundAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformEndRoundAction", reflect.TypeOf((*MockCard)(nil).PerformEndRoundAction), arg0, arg1)
}

// PerformEndTurnAction mocks base method.
func (m *MockCard) PerformEndTurnAction(arg0 core.Game, arg1 deck.CardOrientation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PerformEndTurnAction", arg0, arg1)
}

// PerformEndTurnAction indicates an expected call of PerformEndTurnAction.
func (mr *MockCardMockRecorder) PerformEndTurnAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PerformEndTurnAction", reflect.TypeOf((*MockCard)(nil).PerformEndTurnAction), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonathan-buttner/game-framework/internal/deck (interfaces: Player)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	deck "github.com/jonathan-buttner/game-framework/internal/deck"
)

// MockPlayer is a mock of Player interface.
type MockPlayer struct {
	ctrl     *gomock.Controller
	recorder *MockPlayerMockRecorder
}

// MockPlayerMockRecorder is the mock recorder for MockPlayer.
type MockPlayerMockRecorder struct {
	mock *MockPlayer
}

// NewMockPlayer creates a new mock instance.
func NewMockPlayer(ctrl *gomock.Controller) *MockPlayer {
	mock := &MockPlayer{ctrl: ctrl}
	mock.recorder = &MockPlayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlayer) EXPECT() *MockPlayerMockRecorder {
	return m.recorder
}

// GetHand mocks base method.
func (m *MockPlayer) GetHand() deck.Cards {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHand")
	ret0, _ := ret[0].(deck.Cards)
	return ret0
}

// GetHand indicates an expected call of GetHand.
func (mr *MockPlayerMockRecorder) GetHand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHand", reflect.TypeOf((*MockPlayer)(nil).GetHand))
}

// SetHand mocks base method.
func (m *MockPlayer) SetHand(arg0 []deck.Card) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHand", arg0)
}

// SetHand indicates an expected call of SetHand.
func (mr *MockPlayerMockRecorder) SetHand(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHand", reflect.TypeOf((*MockPlayer)(nil).SetHand), arg0)
}

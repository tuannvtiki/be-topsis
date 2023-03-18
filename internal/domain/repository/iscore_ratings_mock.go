// Code generated by MockGen. DO NOT EDIT.
// Source: iscore_ratings.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	model "topsis/internal/domain/model"

	gomock "github.com/golang/mock/gomock"
)

// MockScoreRatingRepositoryInterface is a mock of ScoreRatingRepositoryInterface interface.
type MockScoreRatingRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockScoreRatingRepositoryInterfaceMockRecorder
}

// MockScoreRatingRepositoryInterfaceMockRecorder is the mock recorder for MockScoreRatingRepositoryInterface.
type MockScoreRatingRepositoryInterfaceMockRecorder struct {
	mock *MockScoreRatingRepositoryInterface
}

// NewMockScoreRatingRepositoryInterface creates a new mock instance.
func NewMockScoreRatingRepositoryInterface(ctrl *gomock.Controller) *MockScoreRatingRepositoryInterface {
	mock := &MockScoreRatingRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockScoreRatingRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScoreRatingRepositoryInterface) EXPECT() *MockScoreRatingRepositoryInterfaceMockRecorder {
	return m.recorder
}

// BulkCreateScoreRating mocks base method.
func (m *MockScoreRatingRepositoryInterface) BulkCreateScoreRating(ctx context.Context, scoreRatings []*model.ScoreRating) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateScoreRating", ctx, scoreRatings)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateScoreRating indicates an expected call of BulkCreateScoreRating.
func (mr *MockScoreRatingRepositoryInterfaceMockRecorder) BulkCreateScoreRating(ctx, scoreRatings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateScoreRating", reflect.TypeOf((*MockScoreRatingRepositoryInterface)(nil).BulkCreateScoreRating), ctx, scoreRatings)
}

// DeleteScoreRatingByQueries mocks base method.
func (m *MockScoreRatingRepositoryInterface) DeleteScoreRatingByQueries(ctx context.Context, queries map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScoreRatingByQueries", ctx, queries)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScoreRatingByQueries indicates an expected call of DeleteScoreRatingByQueries.
func (mr *MockScoreRatingRepositoryInterfaceMockRecorder) DeleteScoreRatingByQueries(ctx, queries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScoreRatingByQueries", reflect.TypeOf((*MockScoreRatingRepositoryInterface)(nil).DeleteScoreRatingByQueries), ctx, queries)
}

// GetScoreRatingByListQueries mocks base method.
func (m *MockScoreRatingRepositoryInterface) GetScoreRatingByListQueries(ctx context.Context, queries map[string]interface{}, sort []string) ([]*model.ScoreRating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScoreRatingByListQueries", ctx, queries, sort)
	ret0, _ := ret[0].([]*model.ScoreRating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScoreRatingByListQueries indicates an expected call of GetScoreRatingByListQueries.
func (mr *MockScoreRatingRepositoryInterfaceMockRecorder) GetScoreRatingByListQueries(ctx, queries, sort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScoreRatingByListQueries", reflect.TypeOf((*MockScoreRatingRepositoryInterface)(nil).GetScoreRatingByListQueries), ctx, queries, sort)
}

// UpdateScoreRatingWithMap mocks base method.
func (m *MockScoreRatingRepositoryInterface) UpdateScoreRatingWithMap(ctx context.Context, scoreRating *model.ScoreRating, data map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScoreRatingWithMap", ctx, scoreRating, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateScoreRatingWithMap indicates an expected call of UpdateScoreRatingWithMap.
func (mr *MockScoreRatingRepositoryInterfaceMockRecorder) UpdateScoreRatingWithMap(ctx, scoreRating, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScoreRatingWithMap", reflect.TypeOf((*MockScoreRatingRepositoryInterface)(nil).UpdateScoreRatingWithMap), ctx, scoreRating, data)
}

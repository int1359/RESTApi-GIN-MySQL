// Code generated by MockGen. DO NOT EDIT.
// Source: rest/gin/dao (interfaces: EmpDao)

// Package mock_dao is a generated GoMock package.
package service

import (
	reflect "reflect"
	models "rest/gin/models"

	gomock "github.com/golang/mock/gomock"
)

// MockEmpDao is a mock of EmpDao interface.
type MockEmpDao struct {
	ctrl     *gomock.Controller
	recorder *MockEmpDaoMockRecorder
}

// MockEmpDaoMockRecorder is the mock recorder for MockEmpDao.
type MockEmpDaoMockRecorder struct {
	mock *MockEmpDao
}

// NewMockEmpDao creates a new mock instance.
func NewMockEmpDao(ctrl *gomock.Controller) *MockEmpDao {
	mock := &MockEmpDao{ctrl: ctrl}
	mock.recorder = &MockEmpDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmpDao) EXPECT() *MockEmpDaoMockRecorder {
	return m.recorder
}

// CreateEmployees mocks base method.
func (m *MockEmpDao) CreateEmployees(arg0 *models.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEmployees", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEmployees indicates an expected call of CreateEmployees.
func (mr *MockEmpDaoMockRecorder) CreateEmployees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEmployees", reflect.TypeOf((*MockEmpDao)(nil).CreateEmployees), arg0)
}

// GetAllEmployees mocks base method.
func (m *MockEmpDao) GetAllEmployees(arg0 *[]models.Employee) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployees", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllEmployees indicates an expected call of GetAllEmployees.
func (mr *MockEmpDaoMockRecorder) GetAllEmployees(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployees", reflect.TypeOf((*MockEmpDao)(nil).GetAllEmployees), arg0)
}

// GetEmployeeById mocks base method.
func (m *MockEmpDao) GetEmployeeById(arg0 *models.Employee, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployeeById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetEmployeeById indicates an expected call of GetEmployeeById.
func (mr *MockEmpDaoMockRecorder) GetEmployeeById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployeeById", reflect.TypeOf((*MockEmpDao)(nil).GetEmployeeById), arg0, arg1)
}

// UpdateEmployee mocks base method.
func (m *MockEmpDao) UpdateEmployee(arg0 *models.Employee, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmployee", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmployee indicates an expected call of UpdateEmployee.
func (mr *MockEmpDaoMockRecorder) UpdateEmployee(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmployee", reflect.TypeOf((*MockEmpDao)(nil).UpdateEmployee), arg0, arg1)
}

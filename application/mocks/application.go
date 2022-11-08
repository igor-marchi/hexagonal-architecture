// Code generated by MockGen. DO NOT EDIT.
// Source: application/product.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	application "github.com/igor-marchi/go-hexagonal/application"
)

// MockIProductInterface is a mock of IProductInterface interface.
type MockIProductInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIProductInterfaceMockRecorder
}

// MockIProductInterfaceMockRecorder is the mock recorder for MockIProductInterface.
type MockIProductInterfaceMockRecorder struct {
	mock *MockIProductInterface
}

// NewMockIProductInterface creates a new mock instance.
func NewMockIProductInterface(ctrl *gomock.Controller) *MockIProductInterface {
	mock := &MockIProductInterface{ctrl: ctrl}
	mock.recorder = &MockIProductInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductInterface) EXPECT() *MockIProductInterfaceMockRecorder {
	return m.recorder
}

// Disable mocks base method.
func (m *MockIProductInterface) Disable() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disable indicates an expected call of Disable.
func (mr *MockIProductInterfaceMockRecorder) Disable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIProductInterface)(nil).Disable))
}

// EnableD mocks base method.
func (m *MockIProductInterface) EnableD() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableD")
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableD indicates an expected call of EnableD.
func (mr *MockIProductInterfaceMockRecorder) EnableD() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableD", reflect.TypeOf((*MockIProductInterface)(nil).EnableD))
}

// GetId mocks base method.
func (m *MockIProductInterface) GetId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockIProductInterfaceMockRecorder) GetId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockIProductInterface)(nil).GetId))
}

// GetName mocks base method.
func (m *MockIProductInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockIProductInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockIProductInterface)(nil).GetName))
}

// GetPrice mocks base method.
func (m *MockIProductInterface) GetPrice() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrice")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetPrice indicates an expected call of GetPrice.
func (mr *MockIProductInterfaceMockRecorder) GetPrice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrice", reflect.TypeOf((*MockIProductInterface)(nil).GetPrice))
}

// GetStatus mocks base method.
func (m *MockIProductInterface) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockIProductInterfaceMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockIProductInterface)(nil).GetStatus))
}

// IsValid mocks base method.
func (m *MockIProductInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockIProductInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockIProductInterface)(nil).IsValid))
}

// MockIProductServiceInterface is a mock of IProductServiceInterface interface.
type MockIProductServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIProductServiceInterfaceMockRecorder
}

// MockIProductServiceInterfaceMockRecorder is the mock recorder for MockIProductServiceInterface.
type MockIProductServiceInterfaceMockRecorder struct {
	mock *MockIProductServiceInterface
}

// NewMockIProductServiceInterface creates a new mock instance.
func NewMockIProductServiceInterface(ctrl *gomock.Controller) *MockIProductServiceInterface {
	mock := &MockIProductServiceInterface{ctrl: ctrl}
	mock.recorder = &MockIProductServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductServiceInterface) EXPECT() *MockIProductServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIProductServiceInterface) Create(name string, price float64) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, price)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIProductServiceInterfaceMockRecorder) Create(name, price interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIProductServiceInterface)(nil).Create), name, price)
}

// Disable mocks base method.
func (m *MockIProductServiceInterface) Disable(product application.IProductInterface) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disable", product)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Disable indicates an expected call of Disable.
func (mr *MockIProductServiceInterfaceMockRecorder) Disable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockIProductServiceInterface)(nil).Disable), product)
}

// Enable mocks base method.
func (m *MockIProductServiceInterface) Enable(product application.IProductInterface) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", product)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enable indicates an expected call of Enable.
func (mr *MockIProductServiceInterfaceMockRecorder) Enable(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockIProductServiceInterface)(nil).Enable), product)
}

// Get mocks base method.
func (m *MockIProductServiceInterface) Get(id string) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIProductServiceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIProductServiceInterface)(nil).Get), id)
}

// MockIProductReaderInterface is a mock of IProductReaderInterface interface.
type MockIProductReaderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIProductReaderInterfaceMockRecorder
}

// MockIProductReaderInterfaceMockRecorder is the mock recorder for MockIProductReaderInterface.
type MockIProductReaderInterfaceMockRecorder struct {
	mock *MockIProductReaderInterface
}

// NewMockIProductReaderInterface creates a new mock instance.
func NewMockIProductReaderInterface(ctrl *gomock.Controller) *MockIProductReaderInterface {
	mock := &MockIProductReaderInterface{ctrl: ctrl}
	mock.recorder = &MockIProductReaderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductReaderInterface) EXPECT() *MockIProductReaderInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIProductReaderInterface) Get(id string) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIProductReaderInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIProductReaderInterface)(nil).Get), id)
}

// MockIProductWriterInterface is a mock of IProductWriterInterface interface.
type MockIProductWriterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIProductWriterInterfaceMockRecorder
}

// MockIProductWriterInterfaceMockRecorder is the mock recorder for MockIProductWriterInterface.
type MockIProductWriterInterfaceMockRecorder struct {
	mock *MockIProductWriterInterface
}

// NewMockIProductWriterInterface creates a new mock instance.
func NewMockIProductWriterInterface(ctrl *gomock.Controller) *MockIProductWriterInterface {
	mock := &MockIProductWriterInterface{ctrl: ctrl}
	mock.recorder = &MockIProductWriterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductWriterInterface) EXPECT() *MockIProductWriterInterfaceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockIProductWriterInterface) Save(product application.IProductInterface) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIProductWriterInterfaceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIProductWriterInterface)(nil).Save), product)
}

// MockIProductPersistenceInterface is a mock of IProductPersistenceInterface interface.
type MockIProductPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIProductPersistenceInterfaceMockRecorder
}

// MockIProductPersistenceInterfaceMockRecorder is the mock recorder for MockIProductPersistenceInterface.
type MockIProductPersistenceInterfaceMockRecorder struct {
	mock *MockIProductPersistenceInterface
}

// NewMockIProductPersistenceInterface creates a new mock instance.
func NewMockIProductPersistenceInterface(ctrl *gomock.Controller) *MockIProductPersistenceInterface {
	mock := &MockIProductPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockIProductPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductPersistenceInterface) EXPECT() *MockIProductPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockIProductPersistenceInterface) Get(id string) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockIProductPersistenceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIProductPersistenceInterface)(nil).Get), id)
}

// Save mocks base method.
func (m *MockIProductPersistenceInterface) Save(product application.IProductInterface) (application.IProductInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", product)
	ret0, _ := ret[0].(application.IProductInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockIProductPersistenceInterfaceMockRecorder) Save(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIProductPersistenceInterface)(nil).Save), product)
}

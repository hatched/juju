// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/description (interfaces: ExternalController,OfferConnection,RemoteEntity,RelationNetwork,RemoteApplication,RemoteSpace,Status)

// Package migrations is a generated GoMock package.
package migrations

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	description "github.com/juju/description/v2"
	names_v3 "github.com/juju/names/v4"
)

// MockExternalController is a mock of ExternalController interface
type MockExternalController struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerMockRecorder
}

// MockExternalControllerMockRecorder is the mock recorder for MockExternalController
type MockExternalControllerMockRecorder struct {
	mock *MockExternalController
}

// NewMockExternalController creates a new mock instance
func NewMockExternalController(ctrl *gomock.Controller) *MockExternalController {
	mock := &MockExternalController{ctrl: ctrl}
	mock.recorder = &MockExternalControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalController) EXPECT() *MockExternalControllerMockRecorder {
	return m.recorder
}

// Addrs mocks base method
func (m *MockExternalController) Addrs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addrs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Addrs indicates an expected call of Addrs
func (mr *MockExternalControllerMockRecorder) Addrs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addrs", reflect.TypeOf((*MockExternalController)(nil).Addrs))
}

// Alias mocks base method
func (m *MockExternalController) Alias() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alias")
	ret0, _ := ret[0].(string)
	return ret0
}

// Alias indicates an expected call of Alias
func (mr *MockExternalControllerMockRecorder) Alias() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alias", reflect.TypeOf((*MockExternalController)(nil).Alias))
}

// CACert mocks base method
func (m *MockExternalController) CACert() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CACert")
	ret0, _ := ret[0].(string)
	return ret0
}

// CACert indicates an expected call of CACert
func (mr *MockExternalControllerMockRecorder) CACert() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CACert", reflect.TypeOf((*MockExternalController)(nil).CACert))
}

// ID mocks base method
func (m *MockExternalController) ID() names_v3.ControllerTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(names_v3.ControllerTag)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockExternalControllerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockExternalController)(nil).ID))
}

// Models mocks base method
func (m *MockExternalController) Models() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Models")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Models indicates an expected call of Models
func (mr *MockExternalControllerMockRecorder) Models() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Models", reflect.TypeOf((*MockExternalController)(nil).Models))
}

// MockOfferConnection is a mock of OfferConnection interface
type MockOfferConnection struct {
	ctrl     *gomock.Controller
	recorder *MockOfferConnectionMockRecorder
}

// MockOfferConnectionMockRecorder is the mock recorder for MockOfferConnection
type MockOfferConnectionMockRecorder struct {
	mock *MockOfferConnection
}

// NewMockOfferConnection creates a new mock instance
func NewMockOfferConnection(ctrl *gomock.Controller) *MockOfferConnection {
	mock := &MockOfferConnection{ctrl: ctrl}
	mock.recorder = &MockOfferConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOfferConnection) EXPECT() *MockOfferConnectionMockRecorder {
	return m.recorder
}

// OfferUUID mocks base method
func (m *MockOfferConnection) OfferUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// OfferUUID indicates an expected call of OfferUUID
func (mr *MockOfferConnectionMockRecorder) OfferUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferUUID", reflect.TypeOf((*MockOfferConnection)(nil).OfferUUID))
}

// RelationID mocks base method
func (m *MockOfferConnection) RelationID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationID")
	ret0, _ := ret[0].(int)
	return ret0
}

// RelationID indicates an expected call of RelationID
func (mr *MockOfferConnectionMockRecorder) RelationID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationID", reflect.TypeOf((*MockOfferConnection)(nil).RelationID))
}

// RelationKey mocks base method
func (m *MockOfferConnection) RelationKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// RelationKey indicates an expected call of RelationKey
func (mr *MockOfferConnectionMockRecorder) RelationKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationKey", reflect.TypeOf((*MockOfferConnection)(nil).RelationKey))
}

// SourceModelUUID mocks base method
func (m *MockOfferConnection) SourceModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SourceModelUUID indicates an expected call of SourceModelUUID
func (mr *MockOfferConnectionMockRecorder) SourceModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceModelUUID", reflect.TypeOf((*MockOfferConnection)(nil).SourceModelUUID))
}

// UserName mocks base method
func (m *MockOfferConnection) UserName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UserName indicates an expected call of UserName
func (mr *MockOfferConnectionMockRecorder) UserName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserName", reflect.TypeOf((*MockOfferConnection)(nil).UserName))
}

// MockRemoteEntity is a mock of RemoteEntity interface
type MockRemoteEntity struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteEntityMockRecorder
}

// MockRemoteEntityMockRecorder is the mock recorder for MockRemoteEntity
type MockRemoteEntityMockRecorder struct {
	mock *MockRemoteEntity
}

// NewMockRemoteEntity creates a new mock instance
func NewMockRemoteEntity(ctrl *gomock.Controller) *MockRemoteEntity {
	mock := &MockRemoteEntity{ctrl: ctrl}
	mock.recorder = &MockRemoteEntityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteEntity) EXPECT() *MockRemoteEntityMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockRemoteEntity) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockRemoteEntityMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRemoteEntity)(nil).ID))
}

// Macaroon mocks base method
func (m *MockRemoteEntity) Macaroon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Macaroon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Macaroon indicates an expected call of Macaroon
func (mr *MockRemoteEntityMockRecorder) Macaroon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Macaroon", reflect.TypeOf((*MockRemoteEntity)(nil).Macaroon))
}

// Token mocks base method
func (m *MockRemoteEntity) Token() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(string)
	return ret0
}

// Token indicates an expected call of Token
func (mr *MockRemoteEntityMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockRemoteEntity)(nil).Token))
}

// MockRelationNetwork is a mock of RelationNetwork interface
type MockRelationNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockRelationNetworkMockRecorder
}

// MockRelationNetworkMockRecorder is the mock recorder for MockRelationNetwork
type MockRelationNetworkMockRecorder struct {
	mock *MockRelationNetwork
}

// NewMockRelationNetwork creates a new mock instance
func NewMockRelationNetwork(ctrl *gomock.Controller) *MockRelationNetwork {
	mock := &MockRelationNetwork{ctrl: ctrl}
	mock.recorder = &MockRelationNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationNetwork) EXPECT() *MockRelationNetworkMockRecorder {
	return m.recorder
}

// CIDRS mocks base method
func (m *MockRelationNetwork) CIDRS() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CIDRS")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CIDRS indicates an expected call of CIDRS
func (mr *MockRelationNetworkMockRecorder) CIDRS() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CIDRS", reflect.TypeOf((*MockRelationNetwork)(nil).CIDRS))
}

// ID mocks base method
func (m *MockRelationNetwork) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockRelationNetworkMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockRelationNetwork)(nil).ID))
}

// RelationKey mocks base method
func (m *MockRelationNetwork) RelationKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// RelationKey indicates an expected call of RelationKey
func (mr *MockRelationNetworkMockRecorder) RelationKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationKey", reflect.TypeOf((*MockRelationNetwork)(nil).RelationKey))
}

// MockRemoteApplication is a mock of RemoteApplication interface
type MockRemoteApplication struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteApplicationMockRecorder
}

// MockRemoteApplicationMockRecorder is the mock recorder for MockRemoteApplication
type MockRemoteApplicationMockRecorder struct {
	mock *MockRemoteApplication
}

// NewMockRemoteApplication creates a new mock instance
func NewMockRemoteApplication(ctrl *gomock.Controller) *MockRemoteApplication {
	mock := &MockRemoteApplication{ctrl: ctrl}
	mock.recorder = &MockRemoteApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteApplication) EXPECT() *MockRemoteApplicationMockRecorder {
	return m.recorder
}

// AddEndpoint mocks base method
func (m *MockRemoteApplication) AddEndpoint(arg0 description.RemoteEndpointArgs) description.RemoteEndpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEndpoint", arg0)
	ret0, _ := ret[0].(description.RemoteEndpoint)
	return ret0
}

// AddEndpoint indicates an expected call of AddEndpoint
func (mr *MockRemoteApplicationMockRecorder) AddEndpoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEndpoint", reflect.TypeOf((*MockRemoteApplication)(nil).AddEndpoint), arg0)
}

// AddSpace mocks base method
func (m *MockRemoteApplication) AddSpace(arg0 description.RemoteSpaceArgs) description.RemoteSpace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSpace", arg0)
	ret0, _ := ret[0].(description.RemoteSpace)
	return ret0
}

// AddSpace indicates an expected call of AddSpace
func (mr *MockRemoteApplicationMockRecorder) AddSpace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSpace", reflect.TypeOf((*MockRemoteApplication)(nil).AddSpace), arg0)
}

// Bindings mocks base method
func (m *MockRemoteApplication) Bindings() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bindings")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Bindings indicates an expected call of Bindings
func (mr *MockRemoteApplicationMockRecorder) Bindings() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bindings", reflect.TypeOf((*MockRemoteApplication)(nil).Bindings))
}

// Endpoints mocks base method
func (m *MockRemoteApplication) Endpoints() []description.RemoteEndpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Endpoints")
	ret0, _ := ret[0].([]description.RemoteEndpoint)
	return ret0
}

// Endpoints indicates an expected call of Endpoints
func (mr *MockRemoteApplicationMockRecorder) Endpoints() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Endpoints", reflect.TypeOf((*MockRemoteApplication)(nil).Endpoints))
}

// IsConsumerProxy mocks base method
func (m *MockRemoteApplication) IsConsumerProxy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsConsumerProxy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsConsumerProxy indicates an expected call of IsConsumerProxy
func (mr *MockRemoteApplicationMockRecorder) IsConsumerProxy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsConsumerProxy", reflect.TypeOf((*MockRemoteApplication)(nil).IsConsumerProxy))
}

// Macaroon mocks base method
func (m *MockRemoteApplication) Macaroon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Macaroon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Macaroon indicates an expected call of Macaroon
func (mr *MockRemoteApplicationMockRecorder) Macaroon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Macaroon", reflect.TypeOf((*MockRemoteApplication)(nil).Macaroon))
}

// Name mocks base method
func (m *MockRemoteApplication) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockRemoteApplicationMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRemoteApplication)(nil).Name))
}

// OfferUUID mocks base method
func (m *MockRemoteApplication) OfferUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// OfferUUID indicates an expected call of OfferUUID
func (mr *MockRemoteApplicationMockRecorder) OfferUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferUUID", reflect.TypeOf((*MockRemoteApplication)(nil).OfferUUID))
}

// SetStatus mocks base method
func (m *MockRemoteApplication) SetStatus(arg0 description.StatusArgs) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStatus", arg0)
}

// SetStatus indicates an expected call of SetStatus
func (mr *MockRemoteApplicationMockRecorder) SetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockRemoteApplication)(nil).SetStatus), arg0)
}

// SourceModelTag mocks base method
func (m *MockRemoteApplication) SourceModelTag() names_v3.ModelTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceModelTag")
	ret0, _ := ret[0].(names_v3.ModelTag)
	return ret0
}

// SourceModelTag indicates an expected call of SourceModelTag
func (mr *MockRemoteApplicationMockRecorder) SourceModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceModelTag", reflect.TypeOf((*MockRemoteApplication)(nil).SourceModelTag))
}

// Spaces mocks base method
func (m *MockRemoteApplication) Spaces() []description.RemoteSpace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spaces")
	ret0, _ := ret[0].([]description.RemoteSpace)
	return ret0
}

// Spaces indicates an expected call of Spaces
func (mr *MockRemoteApplicationMockRecorder) Spaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spaces", reflect.TypeOf((*MockRemoteApplication)(nil).Spaces))
}

// Status mocks base method
func (m *MockRemoteApplication) Status() description.Status {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(description.Status)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockRemoteApplicationMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockRemoteApplication)(nil).Status))
}

// Tag mocks base method
func (m *MockRemoteApplication) Tag() names_v3.ApplicationTag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tag")
	ret0, _ := ret[0].(names_v3.ApplicationTag)
	return ret0
}

// Tag indicates an expected call of Tag
func (mr *MockRemoteApplicationMockRecorder) Tag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tag", reflect.TypeOf((*MockRemoteApplication)(nil).Tag))
}

// URL mocks base method
func (m *MockRemoteApplication) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockRemoteApplicationMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockRemoteApplication)(nil).URL))
}

// MockRemoteSpace is a mock of RemoteSpace interface
type MockRemoteSpace struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteSpaceMockRecorder
}

// MockRemoteSpaceMockRecorder is the mock recorder for MockRemoteSpace
type MockRemoteSpaceMockRecorder struct {
	mock *MockRemoteSpace
}

// NewMockRemoteSpace creates a new mock instance
func NewMockRemoteSpace(ctrl *gomock.Controller) *MockRemoteSpace {
	mock := &MockRemoteSpace{ctrl: ctrl}
	mock.recorder = &MockRemoteSpaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteSpace) EXPECT() *MockRemoteSpaceMockRecorder {
	return m.recorder
}

// AddSubnet mocks base method
func (m *MockRemoteSpace) AddSubnet(arg0 description.SubnetArgs) description.Subnet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSubnet", arg0)
	ret0, _ := ret[0].(description.Subnet)
	return ret0
}

// AddSubnet indicates an expected call of AddSubnet
func (mr *MockRemoteSpaceMockRecorder) AddSubnet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSubnet", reflect.TypeOf((*MockRemoteSpace)(nil).AddSubnet), arg0)
}

// CloudType mocks base method
func (m *MockRemoteSpace) CloudType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudType")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudType indicates an expected call of CloudType
func (mr *MockRemoteSpaceMockRecorder) CloudType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudType", reflect.TypeOf((*MockRemoteSpace)(nil).CloudType))
}

// Name mocks base method
func (m *MockRemoteSpace) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockRemoteSpaceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockRemoteSpace)(nil).Name))
}

// ProviderAttributes mocks base method
func (m *MockRemoteSpace) ProviderAttributes() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderAttributes")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// ProviderAttributes indicates an expected call of ProviderAttributes
func (mr *MockRemoteSpaceMockRecorder) ProviderAttributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderAttributes", reflect.TypeOf((*MockRemoteSpace)(nil).ProviderAttributes))
}

// ProviderId mocks base method
func (m *MockRemoteSpace) ProviderId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProviderId")
	ret0, _ := ret[0].(string)
	return ret0
}

// ProviderId indicates an expected call of ProviderId
func (mr *MockRemoteSpaceMockRecorder) ProviderId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProviderId", reflect.TypeOf((*MockRemoteSpace)(nil).ProviderId))
}

// Subnets mocks base method
func (m *MockRemoteSpace) Subnets() []description.Subnet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subnets")
	ret0, _ := ret[0].([]description.Subnet)
	return ret0
}

// Subnets indicates an expected call of Subnets
func (mr *MockRemoteSpaceMockRecorder) Subnets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subnets", reflect.TypeOf((*MockRemoteSpace)(nil).Subnets))
}

// MockStatus is a mock of Status interface
type MockStatus struct {
	ctrl     *gomock.Controller
	recorder *MockStatusMockRecorder
}

// MockStatusMockRecorder is the mock recorder for MockStatus
type MockStatusMockRecorder struct {
	mock *MockStatus
}

// NewMockStatus creates a new mock instance
func NewMockStatus(ctrl *gomock.Controller) *MockStatus {
	mock := &MockStatus{ctrl: ctrl}
	mock.recorder = &MockStatusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStatus) EXPECT() *MockStatusMockRecorder {
	return m.recorder
}

// Data mocks base method
func (m *MockStatus) Data() map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].(map[string]interface{})
	return ret0
}

// Data indicates an expected call of Data
func (mr *MockStatusMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockStatus)(nil).Data))
}

// Message mocks base method
func (m *MockStatus) Message() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(string)
	return ret0
}

// Message indicates an expected call of Message
func (mr *MockStatusMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockStatus)(nil).Message))
}

// NeverSet mocks base method
func (m *MockStatus) NeverSet() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeverSet")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeverSet indicates an expected call of NeverSet
func (mr *MockStatusMockRecorder) NeverSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeverSet", reflect.TypeOf((*MockStatus)(nil).NeverSet))
}

// Updated mocks base method
func (m *MockStatus) Updated() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Updated")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Updated indicates an expected call of Updated
func (mr *MockStatusMockRecorder) Updated() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Updated", reflect.TypeOf((*MockStatus)(nil).Updated))
}

// Value mocks base method
func (m *MockStatus) Value() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(string)
	return ret0
}

// Value indicates an expected call of Value
func (mr *MockStatusMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockStatus)(nil).Value))
}

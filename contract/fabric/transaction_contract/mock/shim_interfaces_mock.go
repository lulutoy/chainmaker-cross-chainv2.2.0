// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/shawnshen/workspace/ChainMaker/chainmaker-cross-chain/contract/fabric/transaction_contract/vendor/github.com/hyperledger/fabric-chaincode-go/shim/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	shim "github.com/hyperledger/fabric-chaincode-go/shim"
	queryresult "github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	peer "github.com/hyperledger/fabric-protos-go/peer"
)

// MockChaincode is a mock of Chaincode interface.
type MockChaincode struct {
	ctrl     *gomock.Controller
	recorder *MockChaincodeMockRecorder
}

// MockChaincodeMockRecorder is the mock recorder for MockChaincode.
type MockChaincodeMockRecorder struct {
	mock *MockChaincode
}

// NewMockChaincode creates a new mock instance.
func NewMockChaincode(ctrl *gomock.Controller) *MockChaincode {
	mock := &MockChaincode{ctrl: ctrl}
	mock.recorder = &MockChaincodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChaincode) EXPECT() *MockChaincodeMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", stub)
	ret0, _ := ret[0].(peer.Response)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockChaincodeMockRecorder) Init(stub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockChaincode)(nil).Init), stub)
}

// Invoke mocks base method.
func (m *MockChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", stub)
	ret0, _ := ret[0].(peer.Response)
	return ret0
}

// Invoke indicates an expected call of Invoke.
func (mr *MockChaincodeMockRecorder) Invoke(stub interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockChaincode)(nil).Invoke), stub)
}

// MockChaincodeStubInterface is a mock of ChaincodeStubInterface interface.
type MockChaincodeStubInterface struct {
	ctrl     *gomock.Controller
	recorder *MockChaincodeStubInterfaceMockRecorder
}

// MockChaincodeStubInterfaceMockRecorder is the mock recorder for MockChaincodeStubInterface.
type MockChaincodeStubInterfaceMockRecorder struct {
	mock *MockChaincodeStubInterface
}

// NewMockChaincodeStubInterface creates a new mock instance.
func NewMockChaincodeStubInterface(ctrl *gomock.Controller) *MockChaincodeStubInterface {
	mock := &MockChaincodeStubInterface{ctrl: ctrl}
	mock.recorder = &MockChaincodeStubInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChaincodeStubInterface) EXPECT() *MockChaincodeStubInterfaceMockRecorder {
	return m.recorder
}

// CreateCompositeKey mocks base method.
func (m *MockChaincodeStubInterface) CreateCompositeKey(objectType string, attributes []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompositeKey", objectType, attributes)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompositeKey indicates an expected call of CreateCompositeKey.
func (mr *MockChaincodeStubInterfaceMockRecorder) CreateCompositeKey(objectType, attributes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompositeKey", reflect.TypeOf((*MockChaincodeStubInterface)(nil).CreateCompositeKey), objectType, attributes)
}

// DelPrivateData mocks base method.
func (m *MockChaincodeStubInterface) DelPrivateData(collection, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelPrivateData", collection, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelPrivateData indicates an expected call of DelPrivateData.
func (mr *MockChaincodeStubInterfaceMockRecorder) DelPrivateData(collection, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelPrivateData", reflect.TypeOf((*MockChaincodeStubInterface)(nil).DelPrivateData), collection, key)
}

// DelState mocks base method.
func (m *MockChaincodeStubInterface) DelState(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelState", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelState indicates an expected call of DelState.
func (mr *MockChaincodeStubInterfaceMockRecorder) DelState(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelState", reflect.TypeOf((*MockChaincodeStubInterface)(nil).DelState), key)
}

// GetArgs mocks base method.
func (m *MockChaincodeStubInterface) GetArgs() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArgs")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// GetArgs indicates an expected call of GetArgs.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetArgs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArgs", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetArgs))
}

// GetArgsSlice mocks base method.
func (m *MockChaincodeStubInterface) GetArgsSlice() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArgsSlice")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArgsSlice indicates an expected call of GetArgsSlice.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetArgsSlice() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArgsSlice", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetArgsSlice))
}

// GetBinding mocks base method.
func (m *MockChaincodeStubInterface) GetBinding() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBinding")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBinding indicates an expected call of GetBinding.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetBinding() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBinding", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetBinding))
}

// GetChannelID mocks base method.
func (m *MockChaincodeStubInterface) GetChannelID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetChannelID indicates an expected call of GetChannelID.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetChannelID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelID", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetChannelID))
}

// GetCreator mocks base method.
func (m *MockChaincodeStubInterface) GetCreator() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetCreator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetCreator))
}

// GetDecorations mocks base method.
func (m *MockChaincodeStubInterface) GetDecorations() map[string][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDecorations")
	ret0, _ := ret[0].(map[string][]byte)
	return ret0
}

// GetDecorations indicates an expected call of GetDecorations.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetDecorations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDecorations", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetDecorations))
}

// GetFunctionAndParameters mocks base method.
func (m *MockChaincodeStubInterface) GetFunctionAndParameters() (string, []string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunctionAndParameters")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	return ret0, ret1
}

// GetFunctionAndParameters indicates an expected call of GetFunctionAndParameters.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetFunctionAndParameters() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunctionAndParameters", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetFunctionAndParameters))
}

// GetHistoryForKey mocks base method.
func (m *MockChaincodeStubInterface) GetHistoryForKey(key string) (shim.HistoryQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryForKey", key)
	ret0, _ := ret[0].(shim.HistoryQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHistoryForKey indicates an expected call of GetHistoryForKey.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetHistoryForKey(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryForKey", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetHistoryForKey), key)
}

// GetPrivateData mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateData(collection, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateData", collection, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateData indicates an expected call of GetPrivateData.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateData(collection, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateData", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateData), collection, key)
}

// GetPrivateDataByPartialCompositeKey mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateDataByPartialCompositeKey(collection, objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDataByPartialCompositeKey", collection, objectType, keys)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateDataByPartialCompositeKey indicates an expected call of GetPrivateDataByPartialCompositeKey.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateDataByPartialCompositeKey(collection, objectType, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDataByPartialCompositeKey", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateDataByPartialCompositeKey), collection, objectType, keys)
}

// GetPrivateDataByRange mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateDataByRange(collection, startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDataByRange", collection, startKey, endKey)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateDataByRange indicates an expected call of GetPrivateDataByRange.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateDataByRange(collection, startKey, endKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDataByRange", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateDataByRange), collection, startKey, endKey)
}

// GetPrivateDataHash mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateDataHash(collection, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDataHash", collection, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateDataHash indicates an expected call of GetPrivateDataHash.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateDataHash(collection, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDataHash", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateDataHash), collection, key)
}

// GetPrivateDataQueryResult mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateDataQueryResult(collection, query string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDataQueryResult", collection, query)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateDataQueryResult indicates an expected call of GetPrivateDataQueryResult.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateDataQueryResult(collection, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDataQueryResult", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateDataQueryResult), collection, query)
}

// GetPrivateDataValidationParameter mocks base method.
func (m *MockChaincodeStubInterface) GetPrivateDataValidationParameter(collection, key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrivateDataValidationParameter", collection, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPrivateDataValidationParameter indicates an expected call of GetPrivateDataValidationParameter.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetPrivateDataValidationParameter(collection, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrivateDataValidationParameter", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetPrivateDataValidationParameter), collection, key)
}

// GetQueryResult mocks base method.
func (m *MockChaincodeStubInterface) GetQueryResult(query string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryResult", query)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetQueryResult indicates an expected call of GetQueryResult.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetQueryResult(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResult", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetQueryResult), query)
}

// GetQueryResultWithPagination mocks base method.
func (m *MockChaincodeStubInterface) GetQueryResultWithPagination(query string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryResultWithPagination", query, pageSize, bookmark)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(*peer.QueryResponseMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetQueryResultWithPagination indicates an expected call of GetQueryResultWithPagination.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetQueryResultWithPagination(query, pageSize, bookmark interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryResultWithPagination", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetQueryResultWithPagination), query, pageSize, bookmark)
}

// GetSignedProposal mocks base method.
func (m *MockChaincodeStubInterface) GetSignedProposal() (*peer.SignedProposal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignedProposal")
	ret0, _ := ret[0].(*peer.SignedProposal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignedProposal indicates an expected call of GetSignedProposal.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetSignedProposal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignedProposal", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetSignedProposal))
}

// GetState mocks base method.
func (m *MockChaincodeStubInterface) GetState(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetState", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetState indicates an expected call of GetState.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetState(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetState", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetState), key)
}

// GetStateByPartialCompositeKey mocks base method.
func (m *MockChaincodeStubInterface) GetStateByPartialCompositeKey(objectType string, keys []string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateByPartialCompositeKey", objectType, keys)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateByPartialCompositeKey indicates an expected call of GetStateByPartialCompositeKey.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStateByPartialCompositeKey(objectType, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateByPartialCompositeKey", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStateByPartialCompositeKey), objectType, keys)
}

// GetStateByPartialCompositeKeyWithPagination mocks base method.
func (m *MockChaincodeStubInterface) GetStateByPartialCompositeKeyWithPagination(objectType string, keys []string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateByPartialCompositeKeyWithPagination", objectType, keys, pageSize, bookmark)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(*peer.QueryResponseMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStateByPartialCompositeKeyWithPagination indicates an expected call of GetStateByPartialCompositeKeyWithPagination.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStateByPartialCompositeKeyWithPagination(objectType, keys, pageSize, bookmark interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateByPartialCompositeKeyWithPagination", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStateByPartialCompositeKeyWithPagination), objectType, keys, pageSize, bookmark)
}

// GetStateByRange mocks base method.
func (m *MockChaincodeStubInterface) GetStateByRange(startKey, endKey string) (shim.StateQueryIteratorInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateByRange", startKey, endKey)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateByRange indicates an expected call of GetStateByRange.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStateByRange(startKey, endKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateByRange", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStateByRange), startKey, endKey)
}

// GetStateByRangeWithPagination mocks base method.
func (m *MockChaincodeStubInterface) GetStateByRangeWithPagination(startKey, endKey string, pageSize int32, bookmark string) (shim.StateQueryIteratorInterface, *peer.QueryResponseMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateByRangeWithPagination", startKey, endKey, pageSize, bookmark)
	ret0, _ := ret[0].(shim.StateQueryIteratorInterface)
	ret1, _ := ret[1].(*peer.QueryResponseMetadata)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetStateByRangeWithPagination indicates an expected call of GetStateByRangeWithPagination.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStateByRangeWithPagination(startKey, endKey, pageSize, bookmark interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateByRangeWithPagination", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStateByRangeWithPagination), startKey, endKey, pageSize, bookmark)
}

// GetStateValidationParameter mocks base method.
func (m *MockChaincodeStubInterface) GetStateValidationParameter(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStateValidationParameter", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStateValidationParameter indicates an expected call of GetStateValidationParameter.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStateValidationParameter(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStateValidationParameter", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStateValidationParameter), key)
}

// GetStringArgs mocks base method.
func (m *MockChaincodeStubInterface) GetStringArgs() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringArgs")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetStringArgs indicates an expected call of GetStringArgs.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetStringArgs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringArgs", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetStringArgs))
}

// GetTransient mocks base method.
func (m *MockChaincodeStubInterface) GetTransient() (map[string][]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransient")
	ret0, _ := ret[0].(map[string][]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransient indicates an expected call of GetTransient.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetTransient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransient", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetTransient))
}

// GetTxID mocks base method.
func (m *MockChaincodeStubInterface) GetTxID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTxID indicates an expected call of GetTxID.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetTxID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxID", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetTxID))
}

// GetTxTimestamp mocks base method.
func (m *MockChaincodeStubInterface) GetTxTimestamp() (*timestamp.Timestamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxTimestamp")
	ret0, _ := ret[0].(*timestamp.Timestamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxTimestamp indicates an expected call of GetTxTimestamp.
func (mr *MockChaincodeStubInterfaceMockRecorder) GetTxTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxTimestamp", reflect.TypeOf((*MockChaincodeStubInterface)(nil).GetTxTimestamp))
}

// InvokeChaincode mocks base method.
func (m *MockChaincodeStubInterface) InvokeChaincode(chaincodeName string, args [][]byte, channel string) peer.Response {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeChaincode", chaincodeName, args, channel)
	ret0, _ := ret[0].(peer.Response)
	return ret0
}

// InvokeChaincode indicates an expected call of InvokeChaincode.
func (mr *MockChaincodeStubInterfaceMockRecorder) InvokeChaincode(chaincodeName, args, channel interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeChaincode", reflect.TypeOf((*MockChaincodeStubInterface)(nil).InvokeChaincode), chaincodeName, args, channel)
}

// PutPrivateData mocks base method.
func (m *MockChaincodeStubInterface) PutPrivateData(collection, key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutPrivateData", collection, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutPrivateData indicates an expected call of PutPrivateData.
func (mr *MockChaincodeStubInterfaceMockRecorder) PutPrivateData(collection, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutPrivateData", reflect.TypeOf((*MockChaincodeStubInterface)(nil).PutPrivateData), collection, key, value)
}

// PutState mocks base method.
func (m *MockChaincodeStubInterface) PutState(key string, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutState", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutState indicates an expected call of PutState.
func (mr *MockChaincodeStubInterfaceMockRecorder) PutState(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutState", reflect.TypeOf((*MockChaincodeStubInterface)(nil).PutState), key, value)
}

// SetEvent mocks base method.
func (m *MockChaincodeStubInterface) SetEvent(name string, payload []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEvent", name, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEvent indicates an expected call of SetEvent.
func (mr *MockChaincodeStubInterfaceMockRecorder) SetEvent(name, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEvent", reflect.TypeOf((*MockChaincodeStubInterface)(nil).SetEvent), name, payload)
}

// SetPrivateDataValidationParameter mocks base method.
func (m *MockChaincodeStubInterface) SetPrivateDataValidationParameter(collection, key string, ep []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPrivateDataValidationParameter", collection, key, ep)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPrivateDataValidationParameter indicates an expected call of SetPrivateDataValidationParameter.
func (mr *MockChaincodeStubInterfaceMockRecorder) SetPrivateDataValidationParameter(collection, key, ep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPrivateDataValidationParameter", reflect.TypeOf((*MockChaincodeStubInterface)(nil).SetPrivateDataValidationParameter), collection, key, ep)
}

// SetStateValidationParameter mocks base method.
func (m *MockChaincodeStubInterface) SetStateValidationParameter(key string, ep []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetStateValidationParameter", key, ep)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetStateValidationParameter indicates an expected call of SetStateValidationParameter.
func (mr *MockChaincodeStubInterfaceMockRecorder) SetStateValidationParameter(key, ep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateValidationParameter", reflect.TypeOf((*MockChaincodeStubInterface)(nil).SetStateValidationParameter), key, ep)
}

// SplitCompositeKey mocks base method.
func (m *MockChaincodeStubInterface) SplitCompositeKey(compositeKey string) (string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SplitCompositeKey", compositeKey)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SplitCompositeKey indicates an expected call of SplitCompositeKey.
func (mr *MockChaincodeStubInterfaceMockRecorder) SplitCompositeKey(compositeKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SplitCompositeKey", reflect.TypeOf((*MockChaincodeStubInterface)(nil).SplitCompositeKey), compositeKey)
}

// MockCommonIteratorInterface is a mock of CommonIteratorInterface interface.
type MockCommonIteratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCommonIteratorInterfaceMockRecorder
}

// MockCommonIteratorInterfaceMockRecorder is the mock recorder for MockCommonIteratorInterface.
type MockCommonIteratorInterfaceMockRecorder struct {
	mock *MockCommonIteratorInterface
}

// NewMockCommonIteratorInterface creates a new mock instance.
func NewMockCommonIteratorInterface(ctrl *gomock.Controller) *MockCommonIteratorInterface {
	mock := &MockCommonIteratorInterface{ctrl: ctrl}
	mock.recorder = &MockCommonIteratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommonIteratorInterface) EXPECT() *MockCommonIteratorInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCommonIteratorInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCommonIteratorInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCommonIteratorInterface)(nil).Close))
}

// HasNext mocks base method.
func (m *MockCommonIteratorInterface) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockCommonIteratorInterfaceMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockCommonIteratorInterface)(nil).HasNext))
}

// MockStateQueryIteratorInterface is a mock of StateQueryIteratorInterface interface.
type MockStateQueryIteratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStateQueryIteratorInterfaceMockRecorder
}

// MockStateQueryIteratorInterfaceMockRecorder is the mock recorder for MockStateQueryIteratorInterface.
type MockStateQueryIteratorInterfaceMockRecorder struct {
	mock *MockStateQueryIteratorInterface
}

// NewMockStateQueryIteratorInterface creates a new mock instance.
func NewMockStateQueryIteratorInterface(ctrl *gomock.Controller) *MockStateQueryIteratorInterface {
	mock := &MockStateQueryIteratorInterface{ctrl: ctrl}
	mock.recorder = &MockStateQueryIteratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateQueryIteratorInterface) EXPECT() *MockStateQueryIteratorInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStateQueryIteratorInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStateQueryIteratorInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStateQueryIteratorInterface)(nil).Close))
}

// HasNext mocks base method.
func (m *MockStateQueryIteratorInterface) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockStateQueryIteratorInterfaceMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockStateQueryIteratorInterface)(nil).HasNext))
}

// Next mocks base method.
func (m *MockStateQueryIteratorInterface) Next() (*queryresult.KV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*queryresult.KV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockStateQueryIteratorInterfaceMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockStateQueryIteratorInterface)(nil).Next))
}

// MockHistoryQueryIteratorInterface is a mock of HistoryQueryIteratorInterface interface.
type MockHistoryQueryIteratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryQueryIteratorInterfaceMockRecorder
}

// MockHistoryQueryIteratorInterfaceMockRecorder is the mock recorder for MockHistoryQueryIteratorInterface.
type MockHistoryQueryIteratorInterfaceMockRecorder struct {
	mock *MockHistoryQueryIteratorInterface
}

// NewMockHistoryQueryIteratorInterface creates a new mock instance.
func NewMockHistoryQueryIteratorInterface(ctrl *gomock.Controller) *MockHistoryQueryIteratorInterface {
	mock := &MockHistoryQueryIteratorInterface{ctrl: ctrl}
	mock.recorder = &MockHistoryQueryIteratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistoryQueryIteratorInterface) EXPECT() *MockHistoryQueryIteratorInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockHistoryQueryIteratorInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockHistoryQueryIteratorInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHistoryQueryIteratorInterface)(nil).Close))
}

// HasNext mocks base method.
func (m *MockHistoryQueryIteratorInterface) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockHistoryQueryIteratorInterfaceMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockHistoryQueryIteratorInterface)(nil).HasNext))
}

// Next mocks base method.
func (m *MockHistoryQueryIteratorInterface) Next() (*queryresult.KeyModification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*queryresult.KeyModification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockHistoryQueryIteratorInterfaceMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockHistoryQueryIteratorInterface)(nil).Next))
}

// MockMockQueryIteratorInterface is a mock of MockQueryIteratorInterface interface.
type MockMockQueryIteratorInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMockQueryIteratorInterfaceMockRecorder
}

// MockMockQueryIteratorInterfaceMockRecorder is the mock recorder for MockMockQueryIteratorInterface.
type MockMockQueryIteratorInterfaceMockRecorder struct {
	mock *MockMockQueryIteratorInterface
}

// NewMockMockQueryIteratorInterface creates a new mock instance.
func NewMockMockQueryIteratorInterface(ctrl *gomock.Controller) *MockMockQueryIteratorInterface {
	mock := &MockMockQueryIteratorInterface{ctrl: ctrl}
	mock.recorder = &MockMockQueryIteratorInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMockQueryIteratorInterface) EXPECT() *MockMockQueryIteratorInterfaceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMockQueryIteratorInterface) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMockQueryIteratorInterfaceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMockQueryIteratorInterface)(nil).Close))
}

// HasNext mocks base method.
func (m *MockMockQueryIteratorInterface) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockMockQueryIteratorInterfaceMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockMockQueryIteratorInterface)(nil).HasNext))
}

// Next mocks base method.
func (m *MockMockQueryIteratorInterface) Next() (*queryresult.KV, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*queryresult.KV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockMockQueryIteratorInterfaceMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockMockQueryIteratorInterface)(nil).Next))
}

// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-k8s-cni/pkg/httpwrapper (interfaces: HTTPRESP)

// Package mock_httpwrapper is a generated GoMock package.
package mock_httpwrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHTTPRESP is a mock of HTTPRESP interface
type MockHTTPRESP struct {
	ctrl     *gomock.Controller
	recorder *MockHTTPRESPMockRecorder
}

// MockHTTPRESPMockRecorder is the mock recorder for MockHTTPRESP
type MockHTTPRESPMockRecorder struct {
	mock *MockHTTPRESP
}

// NewMockHTTPRESP creates a new mock instance
func NewMockHTTPRESP(ctrl *gomock.Controller) *MockHTTPRESP {
	mock := &MockHTTPRESP{ctrl: ctrl}
	mock.recorder = &MockHTTPRESPMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHTTPRESP) EXPECT() *MockHTTPRESPMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockHTTPRESP) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHTTPRESPMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHTTPRESP)(nil).Close))
}

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsData struct {
	Child IAdsDataChild
}

// The corresponding interface
type IAdsData interface {
	// CommandId returns CommandId
	CommandId() CommandId
	// Response returns Response
	Response() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IAdsDataParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error
	GetTypeName() string
}

type IAdsDataChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *AdsData)
	GetTypeName() string
	IAdsData
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAdsData factory function for AdsData
func NewAdsData() *AdsData {
	return &AdsData{}
}

func CastAdsData(structType interface{}) *AdsData {
	if casted, ok := structType.(AdsData); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsData); ok {
		return casted
	}
	return nil
}

func (m *AdsData) GetTypeName() string {
	return "AdsData"
}

func (m *AdsData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsData) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *AdsData) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *AdsData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDataParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *AdsData
	var typeSwitchError error
	switch {
	case commandId == CommandId_INVALID && response == bool(false): // AdsInvalidRequest
		_parent, typeSwitchError = AdsInvalidRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_INVALID && response == bool(true): // AdsInvalidResponse
		_parent, typeSwitchError = AdsInvalidResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(false): // AdsReadDeviceInfoRequest
		_parent, typeSwitchError = AdsReadDeviceInfoRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_DEVICE_INFO && response == bool(true): // AdsReadDeviceInfoResponse
		_parent, typeSwitchError = AdsReadDeviceInfoResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ && response == bool(false): // AdsReadRequest
		_parent, typeSwitchError = AdsReadRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ && response == bool(true): // AdsReadResponse
		_parent, typeSwitchError = AdsReadResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE && response == bool(false): // AdsWriteRequest
		_parent, typeSwitchError = AdsWriteRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE && response == bool(true): // AdsWriteResponse
		_parent, typeSwitchError = AdsWriteResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_STATE && response == bool(false): // AdsReadStateRequest
		_parent, typeSwitchError = AdsReadStateRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_STATE && response == bool(true): // AdsReadStateResponse
		_parent, typeSwitchError = AdsReadStateResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE_CONTROL && response == bool(false): // AdsWriteControlRequest
		_parent, typeSwitchError = AdsWriteControlRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_WRITE_CONTROL && response == bool(true): // AdsWriteControlResponse
		_parent, typeSwitchError = AdsWriteControlResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(false): // AdsAddDeviceNotificationRequest
		_parent, typeSwitchError = AdsAddDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_ADD_DEVICE_NOTIFICATION && response == bool(true): // AdsAddDeviceNotificationResponse
		_parent, typeSwitchError = AdsAddDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(false): // AdsDeleteDeviceNotificationRequest
		_parent, typeSwitchError = AdsDeleteDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DELETE_DEVICE_NOTIFICATION && response == bool(true): // AdsDeleteDeviceNotificationResponse
		_parent, typeSwitchError = AdsDeleteDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(false): // AdsDeviceNotificationRequest
		_parent, typeSwitchError = AdsDeviceNotificationRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_DEVICE_NOTIFICATION && response == bool(true): // AdsDeviceNotificationResponse
		_parent, typeSwitchError = AdsDeviceNotificationResponseParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_WRITE && response == bool(false): // AdsReadWriteRequest
		_parent, typeSwitchError = AdsReadWriteRequestParse(readBuffer, commandId, response)
	case commandId == CommandId_ADS_READ_WRITE && response == bool(true): // AdsReadWriteResponse
		_parent, typeSwitchError = AdsReadWriteResponseParse(readBuffer, commandId, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("AdsData"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *AdsData) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *AdsData) SerializeParent(writeBuffer utils.WriteBuffer, child IAdsData, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("AdsData"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("AdsData"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

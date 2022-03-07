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
type ApduDataExt struct {

	// Arguments.
	Length uint8
	Child  IApduDataExtChild
}

// The corresponding interface
type IApduDataExt interface {
	// ExtApciType returns ExtApciType
	ExtApciType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IApduDataExtParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduDataExtChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ApduDataExt)
	GetTypeName() string
	IApduDataExt
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataExt factory function for ApduDataExt
func NewApduDataExt(length uint8) *ApduDataExt {
	return &ApduDataExt{Length: length}
}

func CastApduDataExt(structType interface{}) *ApduDataExt {
	if casted, ok := structType.(ApduDataExt); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return casted
	}
	return nil
}

func (m *ApduDataExt) GetTypeName() string {
	return "ApduDataExt"
}

func (m *ApduDataExt) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExt) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *ApduDataExt) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (extApciType)
	lengthInBits += 6

	return lengthInBits
}

func (m *ApduDataExt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExt"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType, _extApciTypeErr := readBuffer.ReadUint8("extApciType", 6)
	if _extApciTypeErr != nil {
		return nil, errors.Wrap(_extApciTypeErr, "Error parsing 'extApciType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduDataExt
	var typeSwitchError error
	switch {
	case extApciType == 0x00: // ApduDataExtOpenRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtOpenRoutingTableRequestParse(readBuffer, length)
	case extApciType == 0x01: // ApduDataExtReadRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtReadRoutingTableRequestParse(readBuffer, length)
	case extApciType == 0x02: // ApduDataExtReadRoutingTableResponse
		_parent, typeSwitchError = ApduDataExtReadRoutingTableResponseParse(readBuffer, length)
	case extApciType == 0x03: // ApduDataExtWriteRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtWriteRoutingTableRequestParse(readBuffer, length)
	case extApciType == 0x08: // ApduDataExtReadRouterMemoryRequest
		_parent, typeSwitchError = ApduDataExtReadRouterMemoryRequestParse(readBuffer, length)
	case extApciType == 0x09: // ApduDataExtReadRouterMemoryResponse
		_parent, typeSwitchError = ApduDataExtReadRouterMemoryResponseParse(readBuffer, length)
	case extApciType == 0x0A: // ApduDataExtWriteRouterMemoryRequest
		_parent, typeSwitchError = ApduDataExtWriteRouterMemoryRequestParse(readBuffer, length)
	case extApciType == 0x0D: // ApduDataExtReadRouterStatusRequest
		_parent, typeSwitchError = ApduDataExtReadRouterStatusRequestParse(readBuffer, length)
	case extApciType == 0x0E: // ApduDataExtReadRouterStatusResponse
		_parent, typeSwitchError = ApduDataExtReadRouterStatusResponseParse(readBuffer, length)
	case extApciType == 0x0F: // ApduDataExtWriteRouterStatusRequest
		_parent, typeSwitchError = ApduDataExtWriteRouterStatusRequestParse(readBuffer, length)
	case extApciType == 0x10: // ApduDataExtMemoryBitWrite
		_parent, typeSwitchError = ApduDataExtMemoryBitWriteParse(readBuffer, length)
	case extApciType == 0x11: // ApduDataExtAuthorizeRequest
		_parent, typeSwitchError = ApduDataExtAuthorizeRequestParse(readBuffer, length)
	case extApciType == 0x12: // ApduDataExtAuthorizeResponse
		_parent, typeSwitchError = ApduDataExtAuthorizeResponseParse(readBuffer, length)
	case extApciType == 0x13: // ApduDataExtKeyWrite
		_parent, typeSwitchError = ApduDataExtKeyWriteParse(readBuffer, length)
	case extApciType == 0x14: // ApduDataExtKeyResponse
		_parent, typeSwitchError = ApduDataExtKeyResponseParse(readBuffer, length)
	case extApciType == 0x15: // ApduDataExtPropertyValueRead
		_parent, typeSwitchError = ApduDataExtPropertyValueReadParse(readBuffer, length)
	case extApciType == 0x16: // ApduDataExtPropertyValueResponse
		_parent, typeSwitchError = ApduDataExtPropertyValueResponseParse(readBuffer, length)
	case extApciType == 0x17: // ApduDataExtPropertyValueWrite
		_parent, typeSwitchError = ApduDataExtPropertyValueWriteParse(readBuffer, length)
	case extApciType == 0x18: // ApduDataExtPropertyDescriptionRead
		_parent, typeSwitchError = ApduDataExtPropertyDescriptionReadParse(readBuffer, length)
	case extApciType == 0x19: // ApduDataExtPropertyDescriptionResponse
		_parent, typeSwitchError = ApduDataExtPropertyDescriptionResponseParse(readBuffer, length)
	case extApciType == 0x1A: // ApduDataExtNetworkParameterRead
		_parent, typeSwitchError = ApduDataExtNetworkParameterReadParse(readBuffer, length)
	case extApciType == 0x1B: // ApduDataExtNetworkParameterResponse
		_parent, typeSwitchError = ApduDataExtNetworkParameterResponseParse(readBuffer, length)
	case extApciType == 0x1C: // ApduDataExtIndividualAddressSerialNumberRead
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberReadParse(readBuffer, length)
	case extApciType == 0x1D: // ApduDataExtIndividualAddressSerialNumberResponse
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberResponseParse(readBuffer, length)
	case extApciType == 0x1E: // ApduDataExtIndividualAddressSerialNumberWrite
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberWriteParse(readBuffer, length)
	case extApciType == 0x20: // ApduDataExtDomainAddressWrite
		_parent, typeSwitchError = ApduDataExtDomainAddressWriteParse(readBuffer, length)
	case extApciType == 0x21: // ApduDataExtDomainAddressRead
		_parent, typeSwitchError = ApduDataExtDomainAddressReadParse(readBuffer, length)
	case extApciType == 0x22: // ApduDataExtDomainAddressResponse
		_parent, typeSwitchError = ApduDataExtDomainAddressResponseParse(readBuffer, length)
	case extApciType == 0x23: // ApduDataExtDomainAddressSelectiveRead
		_parent, typeSwitchError = ApduDataExtDomainAddressSelectiveReadParse(readBuffer, length)
	case extApciType == 0x24: // ApduDataExtNetworkParameterWrite
		_parent, typeSwitchError = ApduDataExtNetworkParameterWriteParse(readBuffer, length)
	case extApciType == 0x25: // ApduDataExtLinkRead
		_parent, typeSwitchError = ApduDataExtLinkReadParse(readBuffer, length)
	case extApciType == 0x26: // ApduDataExtLinkResponse
		_parent, typeSwitchError = ApduDataExtLinkResponseParse(readBuffer, length)
	case extApciType == 0x27: // ApduDataExtLinkWrite
		_parent, typeSwitchError = ApduDataExtLinkWriteParse(readBuffer, length)
	case extApciType == 0x28: // ApduDataExtGroupPropertyValueRead
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueReadParse(readBuffer, length)
	case extApciType == 0x29: // ApduDataExtGroupPropertyValueResponse
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueResponseParse(readBuffer, length)
	case extApciType == 0x2A: // ApduDataExtGroupPropertyValueWrite
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueWriteParse(readBuffer, length)
	case extApciType == 0x2B: // ApduDataExtGroupPropertyValueInfoReport
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueInfoReportParse(readBuffer, length)
	case extApciType == 0x2C: // ApduDataExtDomainAddressSerialNumberRead
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberReadParse(readBuffer, length)
	case extApciType == 0x2D: // ApduDataExtDomainAddressSerialNumberResponse
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberResponseParse(readBuffer, length)
	case extApciType == 0x2E: // ApduDataExtDomainAddressSerialNumberWrite
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberWriteParse(readBuffer, length)
	case extApciType == 0x30: // ApduDataExtFileStreamInfoReport
		_parent, typeSwitchError = ApduDataExtFileStreamInfoReportParse(readBuffer, length)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExt"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduDataExt) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ApduDataExt) SerializeParent(writeBuffer utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ApduDataExt"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType := uint8(child.ExtApciType())
	_extApciTypeErr := writeBuffer.WriteUint8("extApciType", 6, (extApciType))

	if _extApciTypeErr != nil {
		return errors.Wrap(_extApciTypeErr, "Error serializing 'extApciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduDataExt"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ApduDataExt) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

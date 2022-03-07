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
type ApduDataGroupValueResponse struct {
	*ApduData
	DataFirstByte int8
	Data          []byte

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataGroupValueResponse interface {
	IApduData
	// GetDataFirstByte returns DataFirstByte
	GetDataFirstByte() int8
	// GetData returns Data
	GetData() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueResponse) ApciType() uint8 {
	return 0x1
}

func (m *ApduDataGroupValueResponse) GetApciType() uint8 {
	return 0x1
}

func (m *ApduDataGroupValueResponse) InitializeParent(parent *ApduData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataGroupValueResponse) GetDataFirstByte() int8 {
	return m.DataFirstByte
}

func (m *ApduDataGroupValueResponse) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataGroupValueResponse factory function for ApduDataGroupValueResponse
func NewApduDataGroupValueResponse(dataFirstByte int8, data []byte, dataLength uint8) *ApduData {
	child := &ApduDataGroupValueResponse{
		DataFirstByte: dataFirstByte,
		Data:          data,
		ApduData:      NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataGroupValueResponse(structType interface{}) *ApduDataGroupValueResponse {
	if casted, ok := structType.(ApduDataGroupValueResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataGroupValueResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataGroupValueResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataGroupValueResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataGroupValueResponse) GetTypeName() string {
	return "ApduDataGroupValueResponse"
}

func (m *ApduDataGroupValueResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataGroupValueResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dataFirstByte)
	lengthInBits += 6

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataGroupValueResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataGroupValueResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataGroupValueResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (dataFirstByte)
	_dataFirstByte, _dataFirstByteErr := readBuffer.ReadInt8("dataFirstByte", 6)
	if _dataFirstByteErr != nil {
		return nil, errors.Wrap(_dataFirstByteErr, "Error parsing 'dataFirstByte' field")
	}
	dataFirstByte := _dataFirstByte
	// Byte Array field (data)
	numberOfBytesdata := int(utils.InlineIf(bool(bool((dataLength) < (1))), func() interface{} { return uint16(uint16(0)) }, func() interface{} { return uint16(uint16(dataLength) - uint16(uint16(1))) }).(uint16))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueResponse{
		DataFirstByte: dataFirstByte,
		Data:          data,
		ApduData:      &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataGroupValueResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (dataFirstByte)
		dataFirstByte := int8(m.DataFirstByte)
		_dataFirstByteErr := writeBuffer.WriteInt8("dataFirstByte", 6, (dataFirstByte))
		if _dataFirstByteErr != nil {
			return errors.Wrap(_dataFirstByteErr, "Error serializing 'dataFirstByte' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataGroupValueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

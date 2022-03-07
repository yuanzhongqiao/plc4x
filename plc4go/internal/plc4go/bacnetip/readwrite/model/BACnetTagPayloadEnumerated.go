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
type BACnetTagPayloadEnumerated struct {
	Data []byte

	// Arguments.
	ActualLength uint32
}

// The corresponding interface
type IBACnetTagPayloadEnumerated interface {
	// GetData returns Data
	GetData() []byte
	// GetActualValue returns ActualValue
	GetActualValue() uint32
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadEnumerated) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadEnumerated) GetActualValue() uint32 {
	return ParseVarUint(m.GetData())
}

// NewBACnetTagPayloadEnumerated factory function for BACnetTagPayloadEnumerated
func NewBACnetTagPayloadEnumerated(data []byte, actualLength uint32) *BACnetTagPayloadEnumerated {
	return &BACnetTagPayloadEnumerated{Data: data, ActualLength: actualLength}
}

func CastBACnetTagPayloadEnumerated(structType interface{}) *BACnetTagPayloadEnumerated {
	if casted, ok := structType.(BACnetTagPayloadEnumerated); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadEnumerated); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadEnumerated) GetTypeName() string {
	return "BACnetTagPayloadEnumerated"
}

func (m *BACnetTagPayloadEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTagPayloadEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadEnumeratedParse(readBuffer utils.ReadBuffer, actualLength uint32) (*BACnetTagPayloadEnumerated, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadEnumerated"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	// Byte Array field (data)
	numberOfBytesdata := int(actualLength)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	// Virtual field
	_actualValue := ParseVarUint(data)
	actualValue := uint32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadEnumerated"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadEnumerated(data, actualLength), nil
}

func (m *BACnetTagPayloadEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadEnumerated"); pushErr != nil {
		return pushErr
	}

	// Array Field (data)
	if m.Data != nil {
		// Byte Array field (data)
		_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
		}
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadEnumerated"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

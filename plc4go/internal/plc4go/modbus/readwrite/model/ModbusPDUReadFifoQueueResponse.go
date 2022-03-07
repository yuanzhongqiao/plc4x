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
type ModbusPDUReadFifoQueueResponse struct {
	*ModbusPDU
	FifoValue []uint16
}

// The corresponding interface
type IModbusPDUReadFifoQueueResponse interface {
	IModbusPDU
	// GetFifoValue returns FifoValue
	GetFifoValue() []uint16
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
func (m *ModbusPDUReadFifoQueueResponse) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFifoQueueResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFifoQueueResponse) FunctionFlag() uint8 {
	return 0x18
}

func (m *ModbusPDUReadFifoQueueResponse) GetFunctionFlag() uint8 {
	return 0x18
}

func (m *ModbusPDUReadFifoQueueResponse) Response() bool {
	return bool(true)
}

func (m *ModbusPDUReadFifoQueueResponse) GetResponse() bool {
	return bool(true)
}

func (m *ModbusPDUReadFifoQueueResponse) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadFifoQueueResponse) GetFifoValue() []uint16 {
	return m.FifoValue
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUReadFifoQueueResponse factory function for ModbusPDUReadFifoQueueResponse
func NewModbusPDUReadFifoQueueResponse(fifoValue []uint16) *ModbusPDU {
	child := &ModbusPDUReadFifoQueueResponse{
		FifoValue: fifoValue,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadFifoQueueResponse(structType interface{}) *ModbusPDUReadFifoQueueResponse {
	if casted, ok := structType.(ModbusPDUReadFifoQueueResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadFifoQueueResponse); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadFifoQueueResponse(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadFifoQueueResponse(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadFifoQueueResponse) GetTypeName() string {
	return "ModbusPDUReadFifoQueueResponse"
}

func (m *ModbusPDUReadFifoQueueResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadFifoQueueResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 16

	// Implicit Field (fifoCount)
	lengthInBits += 16

	// Array field
	if len(m.FifoValue) > 0 {
		lengthInBits += 16 * uint16(len(m.FifoValue))
	}

	return lengthInBits
}

func (m *ModbusPDUReadFifoQueueResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFifoQueueResponseParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadFifoQueueResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint16("byteCount", 16)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Implicit Field (fifoCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	fifoCount, _fifoCountErr := readBuffer.ReadUint16("fifoCount", 16)
	_ = fifoCount
	if _fifoCountErr != nil {
		return nil, errors.Wrap(_fifoCountErr, "Error parsing 'fifoCount' field")
	}

	// Array field (fifoValue)
	if pullErr := readBuffer.PullContext("fifoValue", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	fifoValue := make([]uint16, fifoCount)
	{
		for curItem := uint16(0); curItem < uint16(fifoCount); curItem++ {
			_item, _err := readBuffer.ReadUint16("", 16)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'fifoValue' field")
			}
			fifoValue[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("fifoValue", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFifoQueueResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadFifoQueueResponse{
		FifoValue: fifoValue,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadFifoQueueResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFifoQueueResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint16(uint16(uint16(uint16(uint16(len(m.GetFifoValue())))*uint16(uint16(2)))) + uint16(uint16(2)))
		_byteCountErr := writeBuffer.WriteUint16("byteCount", 16, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
		}

		// Implicit Field (fifoCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		fifoCount := uint16(uint16(uint16(uint16(uint16(len(m.GetFifoValue())))*uint16(uint16(2)))) / uint16(uint16(2)))
		_fifoCountErr := writeBuffer.WriteUint16("fifoCount", 16, (fifoCount))
		if _fifoCountErr != nil {
			return errors.Wrap(_fifoCountErr, "Error serializing 'fifoCount' field")
		}

		// Array Field (fifoValue)
		if m.FifoValue != nil {
			if pushErr := writeBuffer.PushContext("fifoValue", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.FifoValue {
				_elementErr := writeBuffer.WriteUint16("", 16, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'fifoValue' field")
				}
			}
			if popErr := writeBuffer.PopContext("fifoValue", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadFifoQueueResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadFifoQueueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

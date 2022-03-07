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
type CALDataRequestRecall struct {
	*CALData
	ParamNo uint8
	Count   uint8
}

// The corresponding interface
type ICALDataRequestRecall interface {
	ICALData
	// GetParamNo returns ParamNo
	GetParamNo() uint8
	// GetCount returns Count
	GetCount() uint8
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
func (m *CALDataRequestRecall) CommandType() CALCommandType {
	return CALCommandType_RECALL
}

func (m *CALDataRequestRecall) GetCommandType() CALCommandType {
	return CALCommandType_RECALL
}

func (m *CALDataRequestRecall) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CALDataRequestRecall) GetParamNo() uint8 {
	return m.ParamNo
}

func (m *CALDataRequestRecall) GetCount() uint8 {
	return m.Count
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCALDataRequestRecall factory function for CALDataRequestRecall
func NewCALDataRequestRecall(paramNo uint8, count uint8, commandTypeContainer CALCommandTypeContainer) *CALData {
	child := &CALDataRequestRecall{
		ParamNo: paramNo,
		Count:   count,
		CALData: NewCALData(commandTypeContainer),
	}
	child.Child = child
	return child.CALData
}

func CastCALDataRequestRecall(structType interface{}) *CALDataRequestRecall {
	if casted, ok := structType.(CALDataRequestRecall); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataRequestRecall); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataRequestRecall(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataRequestRecall(casted.Child)
	}
	return nil
}

func (m *CALDataRequestRecall) GetTypeName() string {
	return "CALDataRequestRecall"
}

func (m *CALDataRequestRecall) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataRequestRecall) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALDataRequestRecall) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRequestRecallParse(readBuffer utils.ReadBuffer) (*CALData, error) {
	if pullErr := readBuffer.PullContext("CALDataRequestRecall"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	_paramNo, _paramNoErr := readBuffer.ReadUint8("paramNo", 8)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field")
	}
	paramNo := _paramNo

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("CALDataRequestRecall"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALDataRequestRecall{
		ParamNo: paramNo,
		Count:   count,
		CALData: &CALData{},
	}
	_child.CALData.Child = _child
	return _child.CALData, nil
}

func (m *CALDataRequestRecall) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRequestRecall"); pushErr != nil {
			return pushErr
		}

		// Simple Field (paramNo)
		paramNo := uint8(m.ParamNo)
		_paramNoErr := writeBuffer.WriteUint8("paramNo", 8, (paramNo))
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (count)
		count := uint8(m.Count)
		_countErr := writeBuffer.WriteUint8("count", 8, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRequestRecall"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataRequestRecall) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

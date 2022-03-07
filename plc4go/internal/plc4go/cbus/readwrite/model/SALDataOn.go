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
type SALDataOn struct {
	*SALData
	Group byte
}

// The corresponding interface
type ISALDataOn interface {
	ISALData
	// GetGroup returns Group
	GetGroup() byte
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
func (m *SALDataOn) CommandType() SALCommandType {
	return SALCommandType_ON
}

func (m *SALDataOn) GetCommandType() SALCommandType {
	return SALCommandType_ON
}

func (m *SALDataOn) InitializeParent(parent *SALData, commandTypeContainer SALCommandTypeContainer) {
	m.SALData.CommandTypeContainer = commandTypeContainer
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *SALDataOn) GetGroup() byte {
	return m.Group
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewSALDataOn factory function for SALDataOn
func NewSALDataOn(group byte, commandTypeContainer SALCommandTypeContainer) *SALData {
	child := &SALDataOn{
		Group:   group,
		SALData: NewSALData(commandTypeContainer),
	}
	child.Child = child
	return child.SALData
}

func CastSALDataOn(structType interface{}) *SALDataOn {
	if casted, ok := structType.(SALDataOn); ok {
		return &casted
	}
	if casted, ok := structType.(*SALDataOn); ok {
		return casted
	}
	if casted, ok := structType.(SALData); ok {
		return CastSALDataOn(casted.Child)
	}
	if casted, ok := structType.(*SALData); ok {
		return CastSALDataOn(casted.Child)
	}
	return nil
}

func (m *SALDataOn) GetTypeName() string {
	return "SALDataOn"
}

func (m *SALDataOn) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *SALDataOn) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *SALDataOn) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataOnParse(readBuffer utils.ReadBuffer) (*SALData, error) {
	if pullErr := readBuffer.PullContext("SALDataOn"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field")
	}
	group := _group

	if closeErr := readBuffer.CloseContext("SALDataOn"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &SALDataOn{
		Group:   group,
		SALData: &SALData{},
	}
	_child.SALData.Child = _child
	return _child.SALData, nil
}

func (m *SALDataOn) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataOn"); pushErr != nil {
			return pushErr
		}

		// Simple Field (group)
		group := byte(m.Group)
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("SALDataOn"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *SALDataOn) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type IdentifyReplyCommandTerminalLevels struct {
	*IdentifyReplyCommand
}

// The corresponding interface
type IIdentifyReplyCommandTerminalLevels interface {
	IIdentifyReplyCommand
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
func (m *IdentifyReplyCommandTerminalLevels) Attribute() Attribute {
	return Attribute_TerminalLevel
}

func (m *IdentifyReplyCommandTerminalLevels) GetAttribute() Attribute {
	return Attribute_TerminalLevel
}

func (m *IdentifyReplyCommandTerminalLevels) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandTerminalLevels factory function for IdentifyReplyCommandTerminalLevels
func NewIdentifyReplyCommandTerminalLevels() *IdentifyReplyCommand {
	child := &IdentifyReplyCommandTerminalLevels{
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandTerminalLevels(structType interface{}) *IdentifyReplyCommandTerminalLevels {
	if casted, ok := structType.(IdentifyReplyCommandTerminalLevels); ok {
		return &casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandTerminalLevels(casted.Child)
	}
	if casted, ok := structType.(*IdentifyReplyCommand); ok {
		return CastIdentifyReplyCommandTerminalLevels(casted.Child)
	}
	return nil
}

func (m *IdentifyReplyCommandTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandTerminalLevels"
}

func (m *IdentifyReplyCommandTerminalLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandTerminalLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *IdentifyReplyCommandTerminalLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandTerminalLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandTerminalLevels"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandTerminalLevels"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandTerminalLevels{
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandTerminalLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandTerminalLevels"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandTerminalLevels"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandTerminalLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

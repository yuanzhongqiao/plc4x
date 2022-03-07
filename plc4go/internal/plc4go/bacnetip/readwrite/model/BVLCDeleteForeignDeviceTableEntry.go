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
type BVLCDeleteForeignDeviceTableEntry struct {
	*BVLC
}

// The corresponding interface
type IBVLCDeleteForeignDeviceTableEntry interface {
	IBVLC
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
func (m *BVLCDeleteForeignDeviceTableEntry) BvlcFunction() uint8 {
	return 0x08
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetBvlcFunction() uint8 {
	return 0x08
}

func (m *BVLCDeleteForeignDeviceTableEntry) InitializeParent(parent *BVLC) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBVLCDeleteForeignDeviceTableEntry factory function for BVLCDeleteForeignDeviceTableEntry
func NewBVLCDeleteForeignDeviceTableEntry() *BVLC {
	child := &BVLCDeleteForeignDeviceTableEntry{
		BVLC: NewBVLC(),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCDeleteForeignDeviceTableEntry(structType interface{}) *BVLCDeleteForeignDeviceTableEntry {
	if casted, ok := structType.(BVLCDeleteForeignDeviceTableEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCDeleteForeignDeviceTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
	}
	return nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetTypeName() string {
	return "BVLCDeleteForeignDeviceTableEntry"
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BVLCDeleteForeignDeviceTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCDeleteForeignDeviceTableEntryParse(readBuffer utils.ReadBuffer) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCDeleteForeignDeviceTableEntry"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("BVLCDeleteForeignDeviceTableEntry"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCDeleteForeignDeviceTableEntry{
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCDeleteForeignDeviceTableEntry"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BVLCDeleteForeignDeviceTableEntry"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCDeleteForeignDeviceTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

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
type ComObjectTableRealisationType2 struct {
	*ComObjectTable
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []*GroupObjectDescriptorRealisationType2
}

// The corresponding interface
type IComObjectTableRealisationType2 interface {
	IComObjectTable
	// GetNumEntries returns NumEntries
	GetNumEntries() uint8
	// GetRamFlagsTablePointer returns RamFlagsTablePointer
	GetRamFlagsTablePointer() uint8
	// GetComObjectDescriptors returns ComObjectDescriptors
	GetComObjectDescriptors() []*GroupObjectDescriptorRealisationType2
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
func (m *ComObjectTableRealisationType2) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_2
}

func (m *ComObjectTableRealisationType2) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_2
}

func (m *ComObjectTableRealisationType2) InitializeParent(parent *ComObjectTable) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType2) GetNumEntries() uint8 {
	return m.NumEntries
}

func (m *ComObjectTableRealisationType2) GetRamFlagsTablePointer() uint8 {
	return m.RamFlagsTablePointer
}

func (m *ComObjectTableRealisationType2) GetComObjectDescriptors() []*GroupObjectDescriptorRealisationType2 {
	return m.ComObjectDescriptors
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType2 factory function for ComObjectTableRealisationType2
func NewComObjectTableRealisationType2(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []*GroupObjectDescriptorRealisationType2) *ComObjectTable {
	child := &ComObjectTableRealisationType2{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		ComObjectTable:       NewComObjectTable(),
	}
	child.Child = child
	return child.ComObjectTable
}

func CastComObjectTableRealisationType2(structType interface{}) *ComObjectTableRealisationType2 {
	if casted, ok := structType.(ComObjectTableRealisationType2); ok {
		return &casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType2); ok {
		return casted
	}
	if casted, ok := structType.(ComObjectTable); ok {
		return CastComObjectTableRealisationType2(casted.Child)
	}
	if casted, ok := structType.(*ComObjectTable); ok {
		return CastComObjectTableRealisationType2(casted.Child)
	}
	return nil
}

func (m *ComObjectTableRealisationType2) GetTypeName() string {
	return "ComObjectTableRealisationType2"
}

func (m *ComObjectTableRealisationType2) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ComObjectTableRealisationType2) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for i, element := range m.ComObjectDescriptors {
			last := i == len(m.ComObjectDescriptors)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *ComObjectTableRealisationType2) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ComObjectTableRealisationType2Parse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (*ComObjectTable, error) {
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType2"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (numEntries)
	_numEntries, _numEntriesErr := readBuffer.ReadUint8("numEntries", 8)
	if _numEntriesErr != nil {
		return nil, errors.Wrap(_numEntriesErr, "Error parsing 'numEntries' field")
	}
	numEntries := _numEntries

	// Simple Field (ramFlagsTablePointer)
	_ramFlagsTablePointer, _ramFlagsTablePointerErr := readBuffer.ReadUint8("ramFlagsTablePointer", 8)
	if _ramFlagsTablePointerErr != nil {
		return nil, errors.Wrap(_ramFlagsTablePointerErr, "Error parsing 'ramFlagsTablePointer' field")
	}
	ramFlagsTablePointer := _ramFlagsTablePointer

	// Array field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	comObjectDescriptors := make([]*GroupObjectDescriptorRealisationType2, numEntries)
	{
		for curItem := uint16(0); curItem < uint16(numEntries); curItem++ {
			_item, _err := GroupObjectDescriptorRealisationType2Parse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'comObjectDescriptors' field")
			}
			comObjectDescriptors[curItem] = CastGroupObjectDescriptorRealisationType2(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("comObjectDescriptors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType2"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType2{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		ComObjectTable:       &ComObjectTable{},
	}
	_child.ComObjectTable.Child = _child
	return _child.ComObjectTable, nil
}

func (m *ComObjectTableRealisationType2) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType2"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numEntries)
		numEntries := uint8(m.NumEntries)
		_numEntriesErr := writeBuffer.WriteUint8("numEntries", 8, (numEntries))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.RamFlagsTablePointer)
		_ramFlagsTablePointerErr := writeBuffer.WriteUint8("ramFlagsTablePointer", 8, (ramFlagsTablePointer))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if m.ComObjectDescriptors != nil {
			if pushErr := writeBuffer.PushContext("comObjectDescriptors", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ComObjectDescriptors {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
				}
			}
			if popErr := writeBuffer.PopContext("comObjectDescriptors", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType2"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ComObjectTableRealisationType2) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

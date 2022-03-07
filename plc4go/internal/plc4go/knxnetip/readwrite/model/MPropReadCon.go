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
type MPropReadCon struct {
	*CEMI
	InterfaceObjectType uint16
	ObjectInstance      uint8
	PropertyId          uint8
	NumberOfElements    uint8
	StartIndex          uint16
	Data                uint16

	// Arguments.
	Size uint16
}

// The corresponding interface
type IMPropReadCon interface {
	ICEMI
	// GetInterfaceObjectType returns InterfaceObjectType
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex
	GetStartIndex() uint16
	// GetData returns Data
	GetData() uint16
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
func (m *MPropReadCon) MessageCode() uint8 {
	return 0xFB
}

func (m *MPropReadCon) GetMessageCode() uint8 {
	return 0xFB
}

func (m *MPropReadCon) InitializeParent(parent *CEMI) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *MPropReadCon) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *MPropReadCon) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *MPropReadCon) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *MPropReadCon) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *MPropReadCon) GetStartIndex() uint16 {
	return m.StartIndex
}

func (m *MPropReadCon) GetData() uint16 {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewMPropReadCon factory function for MPropReadCon
func NewMPropReadCon(interfaceObjectType uint16, objectInstance uint8, propertyId uint8, numberOfElements uint8, startIndex uint16, data uint16, size uint16) *CEMI {
	child := &MPropReadCon{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Data:                data,
		CEMI:                NewCEMI(size),
	}
	child.Child = child
	return child.CEMI
}

func CastMPropReadCon(structType interface{}) *MPropReadCon {
	if casted, ok := structType.(MPropReadCon); ok {
		return &casted
	}
	if casted, ok := structType.(*MPropReadCon); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastMPropReadCon(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastMPropReadCon(casted.Child)
	}
	return nil
}

func (m *MPropReadCon) GetTypeName() string {
	return "MPropReadCon"
}

func (m *MPropReadCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *MPropReadCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (interfaceObjectType)
	lengthInBits += 16

	// Simple field (objectInstance)
	lengthInBits += 8

	// Simple field (propertyId)
	lengthInBits += 8

	// Simple field (numberOfElements)
	lengthInBits += 4

	// Simple field (startIndex)
	lengthInBits += 12

	// Simple field (data)
	lengthInBits += 16

	return lengthInBits
}

func (m *MPropReadCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropReadConParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("MPropReadCon"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (interfaceObjectType)
	_interfaceObjectType, _interfaceObjectTypeErr := readBuffer.ReadUint16("interfaceObjectType", 16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.Wrap(_interfaceObjectTypeErr, "Error parsing 'interfaceObjectType' field")
	}
	interfaceObjectType := _interfaceObjectType

	// Simple Field (objectInstance)
	_objectInstance, _objectInstanceErr := readBuffer.ReadUint8("objectInstance", 8)
	if _objectInstanceErr != nil {
		return nil, errors.Wrap(_objectInstanceErr, "Error parsing 'objectInstance' field")
	}
	objectInstance := _objectInstance

	// Simple Field (propertyId)
	_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field")
	}
	propertyId := _propertyId

	// Simple Field (numberOfElements)
	_numberOfElements, _numberOfElementsErr := readBuffer.ReadUint8("numberOfElements", 4)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field")
	}
	numberOfElements := _numberOfElements

	// Simple Field (startIndex)
	_startIndex, _startIndexErr := readBuffer.ReadUint16("startIndex", 12)
	if _startIndexErr != nil {
		return nil, errors.Wrap(_startIndexErr, "Error parsing 'startIndex' field")
	}
	startIndex := _startIndex

	// Simple Field (data)
	_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("MPropReadCon"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MPropReadCon{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance:      objectInstance,
		PropertyId:          propertyId,
		NumberOfElements:    numberOfElements,
		StartIndex:          startIndex,
		Data:                data,
		CEMI:                &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child.CEMI, nil
}

func (m *MPropReadCon) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadCon"); pushErr != nil {
			return pushErr
		}

		// Simple Field (interfaceObjectType)
		interfaceObjectType := uint16(m.InterfaceObjectType)
		_interfaceObjectTypeErr := writeBuffer.WriteUint16("interfaceObjectType", 16, (interfaceObjectType))
		if _interfaceObjectTypeErr != nil {
			return errors.Wrap(_interfaceObjectTypeErr, "Error serializing 'interfaceObjectType' field")
		}

		// Simple Field (objectInstance)
		objectInstance := uint8(m.ObjectInstance)
		_objectInstanceErr := writeBuffer.WriteUint8("objectInstance", 8, (objectInstance))
		if _objectInstanceErr != nil {
			return errors.Wrap(_objectInstanceErr, "Error serializing 'objectInstance' field")
		}

		// Simple Field (propertyId)
		propertyId := uint8(m.PropertyId)
		_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
		if _propertyIdErr != nil {
			return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
		}

		// Simple Field (numberOfElements)
		numberOfElements := uint8(m.NumberOfElements)
		_numberOfElementsErr := writeBuffer.WriteUint8("numberOfElements", 4, (numberOfElements))
		if _numberOfElementsErr != nil {
			return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
		}

		// Simple Field (startIndex)
		startIndex := uint16(m.StartIndex)
		_startIndexErr := writeBuffer.WriteUint16("startIndex", 12, (startIndex))
		if _startIndexErr != nil {
			return errors.Wrap(_startIndexErr, "Error serializing 'startIndex' field")
		}

		// Simple Field (data)
		data := uint16(m.Data)
		_dataErr := writeBuffer.WriteUint16("data", 16, (data))
		if _dataErr != nil {
			return errors.Wrap(_dataErr, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("MPropReadCon"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MPropReadCon) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

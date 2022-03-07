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
type BACnetNotificationParametersComplexEventType struct {
	*BACnetNotificationParameters
	ListOfValues *BACnetPropertyValues

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// The corresponding interface
type IBACnetNotificationParametersComplexEventType interface {
	IBACnetNotificationParameters
	// GetListOfValues returns ListOfValues
	GetListOfValues() *BACnetPropertyValues
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
func (m *BACnetNotificationParametersComplexEventType) PeekedTagNumber() uint8 {
	return uint8(6)
}

func (m *BACnetNotificationParametersComplexEventType) GetPeekedTagNumber() uint8 {
	return uint8(6)
}

func (m *BACnetNotificationParametersComplexEventType) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersComplexEventType) GetListOfValues() *BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersComplexEventType factory function for BACnetNotificationParametersComplexEventType
func NewBACnetNotificationParametersComplexEventType(listOfValues *BACnetPropertyValues, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParameters {
	child := &BACnetNotificationParametersComplexEventType{
		ListOfValues:                 listOfValues,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	child.Child = child
	return child.BACnetNotificationParameters
}

func CastBACnetNotificationParametersComplexEventType(structType interface{}) *BACnetNotificationParametersComplexEventType {
	if casted, ok := structType.(BACnetNotificationParametersComplexEventType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersComplexEventType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersComplexEventType(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersComplexEventType(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersComplexEventType) GetTypeName() string {
	return "BACnetNotificationParametersComplexEventType"
}

func (m *BACnetNotificationParametersComplexEventType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersComplexEventType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersComplexEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersComplexEventTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersComplexEventType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParse(readBuffer, uint8(peekedTagNumber), BACnetObjectType(objectType))
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field")
	}
	listOfValues := CastBACnetPropertyValues(_listOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersComplexEventType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersComplexEventType{
		ListOfValues:                 CastBACnetPropertyValues(listOfValues),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child.BACnetNotificationParameters, nil
}

func (m *BACnetNotificationParametersComplexEventType) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersComplexEventType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (listOfValues)
		if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
			return pushErr
		}
		_listOfValuesErr := m.ListOfValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
			return popErr
		}
		if _listOfValuesErr != nil {
			return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersComplexEventType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersComplexEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

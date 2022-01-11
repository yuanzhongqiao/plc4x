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
type BACnetNotificationParametersChangeOfValueNewValueChangedValue struct {
	*BACnetNotificationParametersChangeOfValueNewValue
	ChangedValue *BACnetContextTagReal
}

// The corresponding interface
type IBACnetNotificationParametersChangeOfValueNewValueChangedValue interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) InitializeParent(parent *BACnetNotificationParametersChangeOfValueNewValue, openingTag *BACnetOpeningTag, peekedTagNumber uint8, closingTag *BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagNumber = peekedTagNumber
	m.ClosingTag = closingTag
}

func NewBACnetNotificationParametersChangeOfValueNewValueChangedValue(changedValue *BACnetContextTagReal, openingTag *BACnetOpeningTag, peekedTagNumber uint8, closingTag *BACnetClosingTag) *BACnetNotificationParametersChangeOfValueNewValue {
	child := &BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: changedValue,
		BACnetNotificationParametersChangeOfValueNewValue: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagNumber, closingTag),
	}
	child.Child = child
	return child.BACnetNotificationParametersChangeOfValueNewValue
}

func CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(structType interface{}) *BACnetNotificationParametersChangeOfValueNewValueChangedValue {
	castFunc := func(typ interface{}) *BACnetNotificationParametersChangeOfValueNewValueChangedValue {
		if casted, ok := typ.(BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersChangeOfValueNewValueChangedValue); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParametersChangeOfValueNewValue); ok {
			return CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParametersChangeOfValueNewValue); ok {
			return CastBACnetNotificationParametersChangeOfValueNewValueChangedValue(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedValue"
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (changedValue)
	lengthInBits += m.ChangedValue.LengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetNotificationParametersChangeOfValueNewValueChangedValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfValueNewValue, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (changedValue)
	if pullErr := readBuffer.PullContext("changedValue"); pullErr != nil {
		return nil, pullErr
	}
	_changedValue, _changedValueErr := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_REAL)
	if _changedValueErr != nil {
		return nil, errors.Wrap(_changedValueErr, "Error parsing 'changedValue' field")
	}
	changedValue := CastBACnetContextTagReal(_changedValue)
	if closeErr := readBuffer.CloseContext("changedValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfValueNewValueChangedValue{
		ChangedValue: CastBACnetContextTagReal(changedValue),
		BACnetNotificationParametersChangeOfValueNewValue: &BACnetNotificationParametersChangeOfValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfValueNewValue.Child = _child
	return _child.BACnetNotificationParametersChangeOfValueNewValue, nil
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (changedValue)
		if pushErr := writeBuffer.PushContext("changedValue"); pushErr != nil {
			return pushErr
		}
		_changedValueErr := m.ChangedValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changedValue"); popErr != nil {
			return popErr
		}
		if _changedValueErr != nil {
			return errors.Wrap(_changedValueErr, "Error serializing 'changedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfValueNewValueChangedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}

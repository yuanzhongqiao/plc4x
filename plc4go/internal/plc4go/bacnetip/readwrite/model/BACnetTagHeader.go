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
type BACnetTagHeader struct {
	TagNumber       uint8
	TagClass        TagClass
	LengthValueType uint8
	ExtTagNumber    *uint8
	ExtLength       *uint8
	ExtExtLength    *uint16
	ExtExtExtLength *uint32
}

// The corresponding interface
type IBACnetTagHeader interface {
	// GetTagNumber returns TagNumber
	GetTagNumber() uint8
	// GetTagClass returns TagClass
	GetTagClass() TagClass
	// GetLengthValueType returns LengthValueType
	GetLengthValueType() uint8
	// GetExtTagNumber returns ExtTagNumber
	GetExtTagNumber() *uint8
	// GetExtLength returns ExtLength
	GetExtLength() *uint8
	// GetExtExtLength returns ExtExtLength
	GetExtExtLength() *uint16
	// GetExtExtExtLength returns ExtExtExtLength
	GetExtExtExtLength() *uint32
	// GetActualTagNumber returns ActualTagNumber
	GetActualTagNumber() uint8
	// GetIsBoolean returns IsBoolean
	GetIsBoolean() bool
	// GetIsConstructed returns IsConstructed
	GetIsConstructed() bool
	// GetIsPrimitiveAndNotBoolean returns IsPrimitiveAndNotBoolean
	GetIsPrimitiveAndNotBoolean() bool
	// GetActualLength returns ActualLength
	GetActualLength() uint32
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
func (m *BACnetTagHeader) GetTagNumber() uint8 {
	return m.TagNumber
}

func (m *BACnetTagHeader) GetTagClass() TagClass {
	return m.TagClass
}

func (m *BACnetTagHeader) GetLengthValueType() uint8 {
	return m.LengthValueType
}

func (m *BACnetTagHeader) GetExtTagNumber() *uint8 {
	return m.ExtTagNumber
}

func (m *BACnetTagHeader) GetExtLength() *uint8 {
	return m.ExtLength
}

func (m *BACnetTagHeader) GetExtExtLength() *uint16 {
	return m.ExtExtLength
}

func (m *BACnetTagHeader) GetExtExtExtLength() *uint32 {
	return m.ExtExtExtLength
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagHeader) GetActualTagNumber() uint8 {
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return utils.InlineIf(bool((m.GetTagNumber()) < (15)), func() interface{} { return uint8(m.GetTagNumber()) }, func() interface{} { return uint8((*m.GetExtTagNumber())) }).(uint8)
}

func (m *BACnetTagHeader) GetIsBoolean() bool {
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(bool((m.GetTagNumber()) == (1))) && bool(bool((m.GetTagClass()) == (TagClass_APPLICATION_TAGS)))
}

func (m *BACnetTagHeader) GetIsConstructed() bool {
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(bool((m.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((m.GetLengthValueType()) == (6)))
}

func (m *BACnetTagHeader) GetIsPrimitiveAndNotBoolean() bool {
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return bool(!(m.GetIsConstructed())) && bool(!(m.GetIsBoolean()))
}

func (m *BACnetTagHeader) GetActualLength() uint32 {
	extTagNumber := m.ExtTagNumber
	_ = extTagNumber
	extLength := m.ExtLength
	_ = extLength
	extExtLength := m.ExtExtLength
	_ = extExtLength
	extExtExtLength := m.ExtExtExtLength
	_ = extExtExtLength
	return utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (255))), func() interface{} { return uint32((*m.GetExtExtExtLength())) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(bool(bool((m.GetLengthValueType()) == (5))) && bool(bool((*m.GetExtLength()) == (254))), func() interface{} { return uint32((*m.GetExtExtLength())) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(bool((m.GetLengthValueType()) == (5)), func() interface{} { return uint32((*m.GetExtLength())) }, func() interface{} { return uint32(m.GetLengthValueType()) }).(uint32)))
		}).(uint32)))
	}).(uint32)
}

// NewBACnetTagHeader factory function for BACnetTagHeader
func NewBACnetTagHeader(tagNumber uint8, tagClass TagClass, lengthValueType uint8, extTagNumber *uint8, extLength *uint8, extExtLength *uint16, extExtExtLength *uint32) *BACnetTagHeader {
	return &BACnetTagHeader{TagNumber: tagNumber, TagClass: tagClass, LengthValueType: lengthValueType, ExtTagNumber: extTagNumber, ExtLength: extLength, ExtExtLength: extExtLength, ExtExtExtLength: extExtExtLength}
}

func CastBACnetTagHeader(structType interface{}) *BACnetTagHeader {
	if casted, ok := structType.(BACnetTagHeader); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagHeader); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagHeader) GetTypeName() string {
	return "BACnetTagHeader"
}

func (m *BACnetTagHeader) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagHeader) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (tagNumber)
	lengthInBits += 4

	// Simple field (tagClass)
	lengthInBits += 1

	// Simple field (lengthValueType)
	lengthInBits += 3

	// Optional Field (extTagNumber)
	if m.ExtTagNumber != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (extLength)
	if m.ExtLength != nil {
		lengthInBits += 8
	}

	// Optional Field (extExtLength)
	if m.ExtExtLength != nil {
		lengthInBits += 16
	}

	// Optional Field (extExtExtLength)
	if m.ExtExtExtLength != nil {
		lengthInBits += 32
	}

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTagHeader) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagHeaderParse(readBuffer utils.ReadBuffer) (*BACnetTagHeader, error) {
	if pullErr := readBuffer.PullContext("BACnetTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (tagNumber)
	_tagNumber, _tagNumberErr := readBuffer.ReadUint8("tagNumber", 4)
	if _tagNumberErr != nil {
		return nil, errors.Wrap(_tagNumberErr, "Error parsing 'tagNumber' field")
	}
	tagNumber := _tagNumber

	// Simple Field (tagClass)
	if pullErr := readBuffer.PullContext("tagClass"); pullErr != nil {
		return nil, pullErr
	}
	_tagClass, _tagClassErr := TagClassParse(readBuffer)
	if _tagClassErr != nil {
		return nil, errors.Wrap(_tagClassErr, "Error parsing 'tagClass' field")
	}
	tagClass := _tagClass
	if closeErr := readBuffer.CloseContext("tagClass"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lengthValueType)
	_lengthValueType, _lengthValueTypeErr := readBuffer.ReadUint8("lengthValueType", 3)
	if _lengthValueTypeErr != nil {
		return nil, errors.Wrap(_lengthValueTypeErr, "Error parsing 'lengthValueType' field")
	}
	lengthValueType := _lengthValueType

	// Optional Field (extTagNumber) (Can be skipped, if a given expression evaluates to false)
	var extTagNumber *uint8 = nil
	if bool((tagNumber) == (15)) {
		_val, _err := readBuffer.ReadUint8("extTagNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extTagNumber' field")
		}
		extTagNumber = &_val
	}

	// Virtual field
	_actualTagNumber := utils.InlineIf(bool((tagNumber) < (15)), func() interface{} { return uint8(tagNumber) }, func() interface{} { return uint8((*extTagNumber)) }).(uint8)
	actualTagNumber := uint8(_actualTagNumber)
	_ = actualTagNumber

	// Virtual field
	_isBoolean := bool(bool((tagNumber) == (1))) && bool(bool((tagClass) == (TagClass_APPLICATION_TAGS)))
	isBoolean := bool(_isBoolean)
	_ = isBoolean

	// Virtual field
	_isConstructed := bool(bool((tagClass) == (TagClass_CONTEXT_SPECIFIC_TAGS))) && bool(bool((lengthValueType) == (6)))
	isConstructed := bool(_isConstructed)
	_ = isConstructed

	// Virtual field
	_isPrimitiveAndNotBoolean := bool(!(isConstructed)) && bool(!(isBoolean))
	isPrimitiveAndNotBoolean := bool(_isPrimitiveAndNotBoolean)
	_ = isPrimitiveAndNotBoolean

	// Optional Field (extLength) (Can be skipped, if a given expression evaluates to false)
	var extLength *uint8 = nil
	if bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5))) {
		_val, _err := readBuffer.ReadUint8("extLength", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extLength' field")
		}
		extLength = &_val
	}

	// Optional Field (extExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtLength *uint16 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (254))) {
		_val, _err := readBuffer.ReadUint16("extExtLength", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtLength' field")
		}
		extExtLength = &_val
	}

	// Optional Field (extExtExtLength) (Can be skipped, if a given expression evaluates to false)
	var extExtExtLength *uint32 = nil
	if bool(bool(isPrimitiveAndNotBoolean) && bool(bool((lengthValueType) == (5)))) && bool(bool((*extLength) == (255))) {
		_val, _err := readBuffer.ReadUint32("extExtExtLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'extExtExtLength' field")
		}
		extExtExtLength = &_val
	}

	// Virtual field
	_actualLength := utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (255))), func() interface{} { return uint32((*extExtExtLength)) }, func() interface{} {
		return uint32(uint32(utils.InlineIf(bool(bool((lengthValueType) == (5))) && bool(bool((*extLength) == (254))), func() interface{} { return uint32((*extExtLength)) }, func() interface{} {
			return uint32(uint32(utils.InlineIf(bool((lengthValueType) == (5)), func() interface{} { return uint32((*extLength)) }, func() interface{} { return uint32(lengthValueType) }).(uint32)))
		}).(uint32)))
	}).(uint32)
	actualLength := uint32(_actualLength)
	_ = actualLength

	if closeErr := readBuffer.CloseContext("BACnetTagHeader"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagHeader(tagNumber, tagClass, lengthValueType, extTagNumber, extLength, extExtLength, extExtExtLength), nil
}

func (m *BACnetTagHeader) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagHeader"); pushErr != nil {
		return pushErr
	}

	// Simple Field (tagNumber)
	tagNumber := uint8(m.TagNumber)
	_tagNumberErr := writeBuffer.WriteUint8("tagNumber", 4, (tagNumber))
	if _tagNumberErr != nil {
		return errors.Wrap(_tagNumberErr, "Error serializing 'tagNumber' field")
	}

	// Simple Field (tagClass)
	if pushErr := writeBuffer.PushContext("tagClass"); pushErr != nil {
		return pushErr
	}
	_tagClassErr := m.TagClass.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("tagClass"); popErr != nil {
		return popErr
	}
	if _tagClassErr != nil {
		return errors.Wrap(_tagClassErr, "Error serializing 'tagClass' field")
	}

	// Simple Field (lengthValueType)
	lengthValueType := uint8(m.LengthValueType)
	_lengthValueTypeErr := writeBuffer.WriteUint8("lengthValueType", 3, (lengthValueType))
	if _lengthValueTypeErr != nil {
		return errors.Wrap(_lengthValueTypeErr, "Error serializing 'lengthValueType' field")
	}

	// Optional Field (extTagNumber) (Can be skipped, if the value is null)
	var extTagNumber *uint8 = nil
	if m.ExtTagNumber != nil {
		extTagNumber = m.ExtTagNumber
		_extTagNumberErr := writeBuffer.WriteUint8("extTagNumber", 8, *(extTagNumber))
		if _extTagNumberErr != nil {
			return errors.Wrap(_extTagNumberErr, "Error serializing 'extTagNumber' field")
		}
	}
	// Virtual field
	if _actualTagNumberErr := writeBuffer.WriteVirtual("actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	if _isBooleanErr := writeBuffer.WriteVirtual("isBoolean", m.GetIsBoolean()); _isBooleanErr != nil {
		return errors.Wrap(_isBooleanErr, "Error serializing 'isBoolean' field")
	}
	// Virtual field
	if _isConstructedErr := writeBuffer.WriteVirtual("isConstructed", m.GetIsConstructed()); _isConstructedErr != nil {
		return errors.Wrap(_isConstructedErr, "Error serializing 'isConstructed' field")
	}
	// Virtual field
	if _isPrimitiveAndNotBooleanErr := writeBuffer.WriteVirtual("isPrimitiveAndNotBoolean", m.GetIsPrimitiveAndNotBoolean()); _isPrimitiveAndNotBooleanErr != nil {
		return errors.Wrap(_isPrimitiveAndNotBooleanErr, "Error serializing 'isPrimitiveAndNotBoolean' field")
	}

	// Optional Field (extLength) (Can be skipped, if the value is null)
	var extLength *uint8 = nil
	if m.ExtLength != nil {
		extLength = m.ExtLength
		_extLengthErr := writeBuffer.WriteUint8("extLength", 8, *(extLength))
		if _extLengthErr != nil {
			return errors.Wrap(_extLengthErr, "Error serializing 'extLength' field")
		}
	}

	// Optional Field (extExtLength) (Can be skipped, if the value is null)
	var extExtLength *uint16 = nil
	if m.ExtExtLength != nil {
		extExtLength = m.ExtExtLength
		_extExtLengthErr := writeBuffer.WriteUint16("extExtLength", 16, *(extExtLength))
		if _extExtLengthErr != nil {
			return errors.Wrap(_extExtLengthErr, "Error serializing 'extExtLength' field")
		}
	}

	// Optional Field (extExtExtLength) (Can be skipped, if the value is null)
	var extExtExtLength *uint32 = nil
	if m.ExtExtExtLength != nil {
		extExtExtLength = m.ExtExtExtLength
		_extExtExtLengthErr := writeBuffer.WriteUint32("extExtExtLength", 32, *(extExtExtLength))
		if _extExtExtLengthErr != nil {
			return errors.Wrap(_extExtExtLengthErr, "Error serializing 'extExtExtLength' field")
		}
	}
	// Virtual field
	if _actualLengthErr := writeBuffer.WriteVirtual("actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagHeader"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagHeader) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

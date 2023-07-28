/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// TrustListDataType is the corresponding interface of TrustListDataType
type TrustListDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSpecifiedLists returns SpecifiedLists (property field)
	GetSpecifiedLists() uint32
	// GetNoOfTrustedCertificates returns NoOfTrustedCertificates (property field)
	GetNoOfTrustedCertificates() int32
	// GetTrustedCertificates returns TrustedCertificates (property field)
	GetTrustedCertificates() []PascalByteString
	// GetNoOfTrustedCrls returns NoOfTrustedCrls (property field)
	GetNoOfTrustedCrls() int32
	// GetTrustedCrls returns TrustedCrls (property field)
	GetTrustedCrls() []PascalByteString
	// GetNoOfIssuerCertificates returns NoOfIssuerCertificates (property field)
	GetNoOfIssuerCertificates() int32
	// GetIssuerCertificates returns IssuerCertificates (property field)
	GetIssuerCertificates() []PascalByteString
	// GetNoOfIssuerCrls returns NoOfIssuerCrls (property field)
	GetNoOfIssuerCrls() int32
	// GetIssuerCrls returns IssuerCrls (property field)
	GetIssuerCrls() []PascalByteString
}

// TrustListDataTypeExactly can be used when we want exactly this type and not a type which fulfills TrustListDataType.
// This is useful for switch cases.
type TrustListDataTypeExactly interface {
	TrustListDataType
	isTrustListDataType() bool
}

// _TrustListDataType is the data-structure of this message
type _TrustListDataType struct {
	*_ExtensionObjectDefinition
	SpecifiedLists          uint32
	NoOfTrustedCertificates int32
	TrustedCertificates     []PascalByteString
	NoOfTrustedCrls         int32
	TrustedCrls             []PascalByteString
	NoOfIssuerCertificates  int32
	IssuerCertificates      []PascalByteString
	NoOfIssuerCrls          int32
	IssuerCrls              []PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TrustListDataType) GetIdentifier() string {
	return "12556"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TrustListDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TrustListDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TrustListDataType) GetSpecifiedLists() uint32 {
	return m.SpecifiedLists
}

func (m *_TrustListDataType) GetNoOfTrustedCertificates() int32 {
	return m.NoOfTrustedCertificates
}

func (m *_TrustListDataType) GetTrustedCertificates() []PascalByteString {
	return m.TrustedCertificates
}

func (m *_TrustListDataType) GetNoOfTrustedCrls() int32 {
	return m.NoOfTrustedCrls
}

func (m *_TrustListDataType) GetTrustedCrls() []PascalByteString {
	return m.TrustedCrls
}

func (m *_TrustListDataType) GetNoOfIssuerCertificates() int32 {
	return m.NoOfIssuerCertificates
}

func (m *_TrustListDataType) GetIssuerCertificates() []PascalByteString {
	return m.IssuerCertificates
}

func (m *_TrustListDataType) GetNoOfIssuerCrls() int32 {
	return m.NoOfIssuerCrls
}

func (m *_TrustListDataType) GetIssuerCrls() []PascalByteString {
	return m.IssuerCrls
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTrustListDataType factory function for _TrustListDataType
func NewTrustListDataType(specifiedLists uint32, noOfTrustedCertificates int32, trustedCertificates []PascalByteString, noOfTrustedCrls int32, trustedCrls []PascalByteString, noOfIssuerCertificates int32, issuerCertificates []PascalByteString, noOfIssuerCrls int32, issuerCrls []PascalByteString) *_TrustListDataType {
	_result := &_TrustListDataType{
		SpecifiedLists:             specifiedLists,
		NoOfTrustedCertificates:    noOfTrustedCertificates,
		TrustedCertificates:        trustedCertificates,
		NoOfTrustedCrls:            noOfTrustedCrls,
		TrustedCrls:                trustedCrls,
		NoOfIssuerCertificates:     noOfIssuerCertificates,
		IssuerCertificates:         issuerCertificates,
		NoOfIssuerCrls:             noOfIssuerCrls,
		IssuerCrls:                 issuerCrls,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTrustListDataType(structType any) TrustListDataType {
	if casted, ok := structType.(TrustListDataType); ok {
		return casted
	}
	if casted, ok := structType.(*TrustListDataType); ok {
		return *casted
	}
	return nil
}

func (m *_TrustListDataType) GetTypeName() string {
	return "TrustListDataType"
}

func (m *_TrustListDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (specifiedLists)
	lengthInBits += 32

	// Simple field (noOfTrustedCertificates)
	lengthInBits += 32

	// Array field
	if len(m.TrustedCertificates) > 0 {
		for _curItem, element := range m.TrustedCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TrustedCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfTrustedCrls)
	lengthInBits += 32

	// Array field
	if len(m.TrustedCrls) > 0 {
		for _curItem, element := range m.TrustedCrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TrustedCrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfIssuerCertificates)
	lengthInBits += 32

	// Array field
	if len(m.IssuerCertificates) > 0 {
		for _curItem, element := range m.IssuerCertificates {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.IssuerCertificates), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfIssuerCrls)
	lengthInBits += 32

	// Array field
	if len(m.IssuerCrls) > 0 {
		for _curItem, element := range m.IssuerCrls {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.IssuerCrls), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TrustListDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TrustListDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (TrustListDataType, error) {
	return TrustListDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TrustListDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TrustListDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TrustListDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TrustListDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (specifiedLists)
	_specifiedLists, _specifiedListsErr := readBuffer.ReadUint32("specifiedLists", 32)
	if _specifiedListsErr != nil {
		return nil, errors.Wrap(_specifiedListsErr, "Error parsing 'specifiedLists' field of TrustListDataType")
	}
	specifiedLists := _specifiedLists

	// Simple Field (noOfTrustedCertificates)
	_noOfTrustedCertificates, _noOfTrustedCertificatesErr := readBuffer.ReadInt32("noOfTrustedCertificates", 32)
	if _noOfTrustedCertificatesErr != nil {
		return nil, errors.Wrap(_noOfTrustedCertificatesErr, "Error parsing 'noOfTrustedCertificates' field of TrustListDataType")
	}
	noOfTrustedCertificates := _noOfTrustedCertificates

	// Array field (trustedCertificates)
	if pullErr := readBuffer.PullContext("trustedCertificates", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for trustedCertificates")
	}
	// Count array
	trustedCertificates := make([]PascalByteString, noOfTrustedCertificates)
	// This happens when the size is set conditional to 0
	if len(trustedCertificates) == 0 {
		trustedCertificates = nil
	}
	{
		_numItems := uint16(noOfTrustedCertificates)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalByteStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'trustedCertificates' field of TrustListDataType")
			}
			trustedCertificates[_curItem] = _item.(PascalByteString)
		}
	}
	if closeErr := readBuffer.CloseContext("trustedCertificates", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for trustedCertificates")
	}

	// Simple Field (noOfTrustedCrls)
	_noOfTrustedCrls, _noOfTrustedCrlsErr := readBuffer.ReadInt32("noOfTrustedCrls", 32)
	if _noOfTrustedCrlsErr != nil {
		return nil, errors.Wrap(_noOfTrustedCrlsErr, "Error parsing 'noOfTrustedCrls' field of TrustListDataType")
	}
	noOfTrustedCrls := _noOfTrustedCrls

	// Array field (trustedCrls)
	if pullErr := readBuffer.PullContext("trustedCrls", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for trustedCrls")
	}
	// Count array
	trustedCrls := make([]PascalByteString, noOfTrustedCrls)
	// This happens when the size is set conditional to 0
	if len(trustedCrls) == 0 {
		trustedCrls = nil
	}
	{
		_numItems := uint16(noOfTrustedCrls)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalByteStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'trustedCrls' field of TrustListDataType")
			}
			trustedCrls[_curItem] = _item.(PascalByteString)
		}
	}
	if closeErr := readBuffer.CloseContext("trustedCrls", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for trustedCrls")
	}

	// Simple Field (noOfIssuerCertificates)
	_noOfIssuerCertificates, _noOfIssuerCertificatesErr := readBuffer.ReadInt32("noOfIssuerCertificates", 32)
	if _noOfIssuerCertificatesErr != nil {
		return nil, errors.Wrap(_noOfIssuerCertificatesErr, "Error parsing 'noOfIssuerCertificates' field of TrustListDataType")
	}
	noOfIssuerCertificates := _noOfIssuerCertificates

	// Array field (issuerCertificates)
	if pullErr := readBuffer.PullContext("issuerCertificates", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for issuerCertificates")
	}
	// Count array
	issuerCertificates := make([]PascalByteString, noOfIssuerCertificates)
	// This happens when the size is set conditional to 0
	if len(issuerCertificates) == 0 {
		issuerCertificates = nil
	}
	{
		_numItems := uint16(noOfIssuerCertificates)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalByteStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'issuerCertificates' field of TrustListDataType")
			}
			issuerCertificates[_curItem] = _item.(PascalByteString)
		}
	}
	if closeErr := readBuffer.CloseContext("issuerCertificates", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for issuerCertificates")
	}

	// Simple Field (noOfIssuerCrls)
	_noOfIssuerCrls, _noOfIssuerCrlsErr := readBuffer.ReadInt32("noOfIssuerCrls", 32)
	if _noOfIssuerCrlsErr != nil {
		return nil, errors.Wrap(_noOfIssuerCrlsErr, "Error parsing 'noOfIssuerCrls' field of TrustListDataType")
	}
	noOfIssuerCrls := _noOfIssuerCrls

	// Array field (issuerCrls)
	if pullErr := readBuffer.PullContext("issuerCrls", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for issuerCrls")
	}
	// Count array
	issuerCrls := make([]PascalByteString, noOfIssuerCrls)
	// This happens when the size is set conditional to 0
	if len(issuerCrls) == 0 {
		issuerCrls = nil
	}
	{
		_numItems := uint16(noOfIssuerCrls)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalByteStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'issuerCrls' field of TrustListDataType")
			}
			issuerCrls[_curItem] = _item.(PascalByteString)
		}
	}
	if closeErr := readBuffer.CloseContext("issuerCrls", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for issuerCrls")
	}

	if closeErr := readBuffer.CloseContext("TrustListDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TrustListDataType")
	}

	// Create a partially initialized instance
	_child := &_TrustListDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SpecifiedLists:             specifiedLists,
		NoOfTrustedCertificates:    noOfTrustedCertificates,
		TrustedCertificates:        trustedCertificates,
		NoOfTrustedCrls:            noOfTrustedCrls,
		TrustedCrls:                trustedCrls,
		NoOfIssuerCertificates:     noOfIssuerCertificates,
		IssuerCertificates:         issuerCertificates,
		NoOfIssuerCrls:             noOfIssuerCrls,
		IssuerCrls:                 issuerCrls,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TrustListDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TrustListDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TrustListDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TrustListDataType")
		}

		// Simple Field (specifiedLists)
		specifiedLists := uint32(m.GetSpecifiedLists())
		_specifiedListsErr := writeBuffer.WriteUint32("specifiedLists", 32, (specifiedLists))
		if _specifiedListsErr != nil {
			return errors.Wrap(_specifiedListsErr, "Error serializing 'specifiedLists' field")
		}

		// Simple Field (noOfTrustedCertificates)
		noOfTrustedCertificates := int32(m.GetNoOfTrustedCertificates())
		_noOfTrustedCertificatesErr := writeBuffer.WriteInt32("noOfTrustedCertificates", 32, (noOfTrustedCertificates))
		if _noOfTrustedCertificatesErr != nil {
			return errors.Wrap(_noOfTrustedCertificatesErr, "Error serializing 'noOfTrustedCertificates' field")
		}

		// Array Field (trustedCertificates)
		if pushErr := writeBuffer.PushContext("trustedCertificates", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for trustedCertificates")
		}
		for _curItem, _element := range m.GetTrustedCertificates() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetTrustedCertificates()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'trustedCertificates' field")
			}
		}
		if popErr := writeBuffer.PopContext("trustedCertificates", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for trustedCertificates")
		}

		// Simple Field (noOfTrustedCrls)
		noOfTrustedCrls := int32(m.GetNoOfTrustedCrls())
		_noOfTrustedCrlsErr := writeBuffer.WriteInt32("noOfTrustedCrls", 32, (noOfTrustedCrls))
		if _noOfTrustedCrlsErr != nil {
			return errors.Wrap(_noOfTrustedCrlsErr, "Error serializing 'noOfTrustedCrls' field")
		}

		// Array Field (trustedCrls)
		if pushErr := writeBuffer.PushContext("trustedCrls", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for trustedCrls")
		}
		for _curItem, _element := range m.GetTrustedCrls() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetTrustedCrls()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'trustedCrls' field")
			}
		}
		if popErr := writeBuffer.PopContext("trustedCrls", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for trustedCrls")
		}

		// Simple Field (noOfIssuerCertificates)
		noOfIssuerCertificates := int32(m.GetNoOfIssuerCertificates())
		_noOfIssuerCertificatesErr := writeBuffer.WriteInt32("noOfIssuerCertificates", 32, (noOfIssuerCertificates))
		if _noOfIssuerCertificatesErr != nil {
			return errors.Wrap(_noOfIssuerCertificatesErr, "Error serializing 'noOfIssuerCertificates' field")
		}

		// Array Field (issuerCertificates)
		if pushErr := writeBuffer.PushContext("issuerCertificates", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for issuerCertificates")
		}
		for _curItem, _element := range m.GetIssuerCertificates() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetIssuerCertificates()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'issuerCertificates' field")
			}
		}
		if popErr := writeBuffer.PopContext("issuerCertificates", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for issuerCertificates")
		}

		// Simple Field (noOfIssuerCrls)
		noOfIssuerCrls := int32(m.GetNoOfIssuerCrls())
		_noOfIssuerCrlsErr := writeBuffer.WriteInt32("noOfIssuerCrls", 32, (noOfIssuerCrls))
		if _noOfIssuerCrlsErr != nil {
			return errors.Wrap(_noOfIssuerCrlsErr, "Error serializing 'noOfIssuerCrls' field")
		}

		// Array Field (issuerCrls)
		if pushErr := writeBuffer.PushContext("issuerCrls", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for issuerCrls")
		}
		for _curItem, _element := range m.GetIssuerCrls() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetIssuerCrls()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'issuerCrls' field")
			}
		}
		if popErr := writeBuffer.PopContext("issuerCrls", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for issuerCrls")
		}

		if popErr := writeBuffer.PopContext("TrustListDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TrustListDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TrustListDataType) isTrustListDataType() bool {
	return true
}

func (m *_TrustListDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
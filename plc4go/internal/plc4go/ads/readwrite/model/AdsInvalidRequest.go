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
type AdsInvalidRequest struct {
	*AdsData
}

// The corresponding interface
type IAdsInvalidRequest interface {
	IAdsData
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
func (m *AdsInvalidRequest) CommandId() CommandId {
	return CommandId_INVALID
}

func (m *AdsInvalidRequest) GetCommandId() CommandId {
	return CommandId_INVALID
}

func (m *AdsInvalidRequest) Response() bool {
	return bool(false)
}

func (m *AdsInvalidRequest) GetResponse() bool {
	return bool(false)
}

func (m *AdsInvalidRequest) InitializeParent(parent *AdsData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAdsInvalidRequest factory function for AdsInvalidRequest
func NewAdsInvalidRequest() *AdsData {
	child := &AdsInvalidRequest{
		AdsData: NewAdsData(),
	}
	child.Child = child
	return child.AdsData
}

func CastAdsInvalidRequest(structType interface{}) *AdsInvalidRequest {
	if casted, ok := structType.(AdsInvalidRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsInvalidRequest); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsInvalidRequest(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsInvalidRequest(casted.Child)
	}
	return nil
}

func (m *AdsInvalidRequest) GetTypeName() string {
	return "AdsInvalidRequest"
}

func (m *AdsInvalidRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsInvalidRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *AdsInvalidRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsInvalidRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsInvalidRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsInvalidRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsInvalidRequest{
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child.AdsData, nil
}

func (m *AdsInvalidRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsInvalidRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("AdsInvalidRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsInvalidRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

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
type CIPEncapsulationReadRequest struct {
	*CIPEncapsulationPacket
	Request *DF1RequestMessage
}

// The corresponding interface
type ICIPEncapsulationReadRequest interface {
	ICIPEncapsulationPacket
	// GetRequest returns Request
	GetRequest() *DF1RequestMessage
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
func (m *CIPEncapsulationReadRequest) CommandType() uint16 {
	return 0x0107
}

func (m *CIPEncapsulationReadRequest) GetCommandType() uint16 {
	return 0x0107
}

func (m *CIPEncapsulationReadRequest) InitializeParent(parent *CIPEncapsulationPacket, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) {
	m.CIPEncapsulationPacket.SessionHandle = sessionHandle
	m.CIPEncapsulationPacket.Status = status
	m.CIPEncapsulationPacket.SenderContext = senderContext
	m.CIPEncapsulationPacket.Options = options
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CIPEncapsulationReadRequest) GetRequest() *DF1RequestMessage {
	return m.Request
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCIPEncapsulationReadRequest factory function for CIPEncapsulationReadRequest
func NewCIPEncapsulationReadRequest(request *DF1RequestMessage, sessionHandle uint32, status uint32, senderContext []uint8, options uint32) *CIPEncapsulationPacket {
	child := &CIPEncapsulationReadRequest{
		Request:                request,
		CIPEncapsulationPacket: NewCIPEncapsulationPacket(sessionHandle, status, senderContext, options),
	}
	child.Child = child
	return child.CIPEncapsulationPacket
}

func CastCIPEncapsulationReadRequest(structType interface{}) *CIPEncapsulationReadRequest {
	if casted, ok := structType.(CIPEncapsulationReadRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*CIPEncapsulationReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(CIPEncapsulationPacket); ok {
		return CastCIPEncapsulationReadRequest(casted.Child)
	}
	if casted, ok := structType.(*CIPEncapsulationPacket); ok {
		return CastCIPEncapsulationReadRequest(casted.Child)
	}
	return nil
}

func (m *CIPEncapsulationReadRequest) GetTypeName() string {
	return "CIPEncapsulationReadRequest"
}

func (m *CIPEncapsulationReadRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CIPEncapsulationReadRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits()

	return lengthInBits
}

func (m *CIPEncapsulationReadRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CIPEncapsulationReadRequestParse(readBuffer utils.ReadBuffer) (*CIPEncapsulationPacket, error) {
	if pullErr := readBuffer.PullContext("CIPEncapsulationReadRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (request)
	if pullErr := readBuffer.PullContext("request"); pullErr != nil {
		return nil, pullErr
	}
	_request, _requestErr := DF1RequestMessageParse(readBuffer)
	if _requestErr != nil {
		return nil, errors.Wrap(_requestErr, "Error parsing 'request' field")
	}
	request := CastDF1RequestMessage(_request)
	if closeErr := readBuffer.CloseContext("request"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("CIPEncapsulationReadRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CIPEncapsulationReadRequest{
		Request:                CastDF1RequestMessage(request),
		CIPEncapsulationPacket: &CIPEncapsulationPacket{},
	}
	_child.CIPEncapsulationPacket.Child = _child
	return _child.CIPEncapsulationPacket, nil
}

func (m *CIPEncapsulationReadRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CIPEncapsulationReadRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (request)
		if pushErr := writeBuffer.PushContext("request"); pushErr != nil {
			return pushErr
		}
		_requestErr := m.Request.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("request"); popErr != nil {
			return popErr
		}
		if _requestErr != nil {
			return errors.Wrap(_requestErr, "Error serializing 'request' field")
		}

		if popErr := writeBuffer.PopContext("CIPEncapsulationReadRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CIPEncapsulationReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const StandardFormatStatusReply_CR byte = 0x0D
const StandardFormatStatusReply_LF byte = 0x0A

// The data-structure of this message
type StandardFormatStatusReply struct {
	StatusHeader *StatusHeader
	Application  ApplicationIdContainer
	BlockStart   uint8
	StatusBytes  []*StatusByte
	Crc          *Checksum
}

// The corresponding interface
type IStandardFormatStatusReply interface {
	// GetStatusHeader returns StatusHeader
	GetStatusHeader() *StatusHeader
	// GetApplication returns Application
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart
	GetBlockStart() uint8
	// GetStatusBytes returns StatusBytes
	GetStatusBytes() []*StatusByte
	// GetCrc returns Crc
	GetCrc() *Checksum
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
func (m *StandardFormatStatusReply) GetStatusHeader() *StatusHeader {
	return m.StatusHeader
}

func (m *StandardFormatStatusReply) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *StandardFormatStatusReply) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *StandardFormatStatusReply) GetStatusBytes() []*StatusByte {
	return m.StatusBytes
}

func (m *StandardFormatStatusReply) GetCrc() *Checksum {
	return m.Crc
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewStandardFormatStatusReply factory function for StandardFormatStatusReply
func NewStandardFormatStatusReply(statusHeader *StatusHeader, application ApplicationIdContainer, blockStart uint8, statusBytes []*StatusByte, crc *Checksum) *StandardFormatStatusReply {
	return &StandardFormatStatusReply{StatusHeader: statusHeader, Application: application, BlockStart: blockStart, StatusBytes: statusBytes, Crc: crc}
}

func CastStandardFormatStatusReply(structType interface{}) *StandardFormatStatusReply {
	if casted, ok := structType.(StandardFormatStatusReply); ok {
		return &casted
	}
	if casted, ok := structType.(*StandardFormatStatusReply); ok {
		return casted
	}
	return nil
}

func (m *StandardFormatStatusReply) GetTypeName() string {
	return "StandardFormatStatusReply"
}

func (m *StandardFormatStatusReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StandardFormatStatusReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (statusHeader)
	lengthInBits += m.StatusHeader.GetLengthInBits()

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// Array field
	if len(m.StatusBytes) > 0 {
		for i, element := range m.StatusBytes {
			last := i == len(m.StatusBytes)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	// Simple field (crc)
	lengthInBits += m.Crc.GetLengthInBits()

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *StandardFormatStatusReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StandardFormatStatusReplyParse(readBuffer utils.ReadBuffer) (*StandardFormatStatusReply, error) {
	if pullErr := readBuffer.PullContext("StandardFormatStatusReply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (statusHeader)
	if pullErr := readBuffer.PullContext("statusHeader"); pullErr != nil {
		return nil, pullErr
	}
	_statusHeader, _statusHeaderErr := StatusHeaderParse(readBuffer)
	if _statusHeaderErr != nil {
		return nil, errors.Wrap(_statusHeaderErr, "Error parsing 'statusHeader' field")
	}
	statusHeader := CastStatusHeader(_statusHeader)
	if closeErr := readBuffer.CloseContext("statusHeader"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, pullErr
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (blockStart)
	_blockStart, _blockStartErr := readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field")
	}
	blockStart := _blockStart

	// Array field (statusBytes)
	if pullErr := readBuffer.PullContext("statusBytes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	statusBytes := make([]*StatusByte, uint16(statusHeader.GetNumberOfCharacterPairs())-uint16(uint16(2)))
	{
		for curItem := uint16(0); curItem < uint16(uint16(statusHeader.GetNumberOfCharacterPairs())-uint16(uint16(2))); curItem++ {
			_item, _err := StatusByteParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'statusBytes' field")
			}
			statusBytes[curItem] = CastStatusByte(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("statusBytes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (crc)
	if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
		return nil, pullErr
	}
	_crc, _crcErr := ChecksumParse(readBuffer)
	if _crcErr != nil {
		return nil, errors.Wrap(_crcErr, "Error parsing 'crc' field")
	}
	crc := CastChecksum(_crc)
	if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != StandardFormatStatusReply_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", StandardFormatStatusReply_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field")
	}
	if lf != StandardFormatStatusReply_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", StandardFormatStatusReply_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("StandardFormatStatusReply"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewStandardFormatStatusReply(statusHeader, application, blockStart, statusBytes, crc), nil
}

func (m *StandardFormatStatusReply) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("StandardFormatStatusReply"); pushErr != nil {
		return pushErr
	}

	// Simple Field (statusHeader)
	if pushErr := writeBuffer.PushContext("statusHeader"); pushErr != nil {
		return pushErr
	}
	_statusHeaderErr := m.StatusHeader.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("statusHeader"); popErr != nil {
		return popErr
	}
	if _statusHeaderErr != nil {
		return errors.Wrap(_statusHeaderErr, "Error serializing 'statusHeader' field")
	}

	// Simple Field (application)
	if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
		return pushErr
	}
	_applicationErr := m.Application.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("application"); popErr != nil {
		return popErr
	}
	if _applicationErr != nil {
		return errors.Wrap(_applicationErr, "Error serializing 'application' field")
	}

	// Simple Field (blockStart)
	blockStart := uint8(m.BlockStart)
	_blockStartErr := writeBuffer.WriteUint8("blockStart", 8, (blockStart))
	if _blockStartErr != nil {
		return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
	}

	// Array Field (statusBytes)
	if m.StatusBytes != nil {
		if pushErr := writeBuffer.PushContext("statusBytes", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.StatusBytes {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'statusBytes' field")
			}
		}
		if popErr := writeBuffer.PopContext("statusBytes", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (crc)
	if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
		return pushErr
	}
	_crcErr := m.Crc.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("crc"); popErr != nil {
		return popErr
	}
	if _crcErr != nil {
		return errors.Wrap(_crcErr, "Error serializing 'crc' field")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("StandardFormatStatusReply"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *StandardFormatStatusReply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

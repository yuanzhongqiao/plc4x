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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AmsTCPPacket struct {
	Userdata *AmsPacket
}

// The corresponding interface
type IAmsTCPPacket interface {
	// GetUserdata returns Userdata
	GetUserdata() *AmsPacket
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
func (m *AmsTCPPacket) GetUserdata() *AmsPacket {
	return m.Userdata
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAmsTCPPacket factory function for AmsTCPPacket
func NewAmsTCPPacket(userdata *AmsPacket) *AmsTCPPacket {
	return &AmsTCPPacket{Userdata: userdata}
}

func CastAmsTCPPacket(structType interface{}) *AmsTCPPacket {
	if casted, ok := structType.(AmsTCPPacket); ok {
		return &casted
	}
	if casted, ok := structType.(*AmsTCPPacket); ok {
		return casted
	}
	return nil
}

func (m *AmsTCPPacket) GetTypeName() string {
	return "AmsTCPPacket"
}

func (m *AmsTCPPacket) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AmsTCPPacket) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Simple field (userdata)
	lengthInBits += m.Userdata.GetLengthInBits()

	return lengthInBits
}

func (m *AmsTCPPacket) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AmsTCPPacketParse(readBuffer utils.ReadBuffer) (*AmsTCPPacket, error) {
	if pullErr := readBuffer.PullContext("AmsTCPPacket"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0000) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0000),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}

	// Simple Field (userdata)
	if pullErr := readBuffer.PullContext("userdata"); pullErr != nil {
		return nil, pullErr
	}
	_userdata, _userdataErr := AmsPacketParse(readBuffer)
	if _userdataErr != nil {
		return nil, errors.Wrap(_userdataErr, "Error parsing 'userdata' field")
	}
	userdata := CastAmsPacket(_userdata)
	if closeErr := readBuffer.CloseContext("userdata"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AmsTCPPacket"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAmsTCPPacket(userdata), nil
}

func (m *AmsTCPPacket) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AmsTCPPacket"); pushErr != nil {
		return pushErr
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0000))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(m.GetUserdata().GetLengthInBytes())
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Simple Field (userdata)
	if pushErr := writeBuffer.PushContext("userdata"); pushErr != nil {
		return pushErr
	}
	_userdataErr := m.Userdata.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("userdata"); popErr != nil {
		return popErr
	}
	if _userdataErr != nil {
		return errors.Wrap(_userdataErr, "Error serializing 'userdata' field")
	}

	if popErr := writeBuffer.PopContext("AmsTCPPacket"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AmsTCPPacket) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

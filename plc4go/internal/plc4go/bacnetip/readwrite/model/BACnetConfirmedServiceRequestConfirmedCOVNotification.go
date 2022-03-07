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
type BACnetConfirmedServiceRequestConfirmedCOVNotification struct {
	*BACnetConfirmedServiceRequest
	SubscriberProcessIdentifier *BACnetContextTagUnsignedInteger
	InitiatingDeviceIdentifier  *BACnetContextTagObjectIdentifier
	MonitoredObjectIdentifier   *BACnetContextTagObjectIdentifier
	LifetimeInSeconds           *BACnetContextTagUnsignedInteger
	ListOfValues                *BACnetPropertyValues

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetConfirmedServiceRequestConfirmedCOVNotification interface {
	IBACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier
	GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetInitiatingDeviceIdentifier returns InitiatingDeviceIdentifier
	GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier
	GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetLifetimeInSeconds returns LifetimeInSeconds
	GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger
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
func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) ServiceChoice() uint8 {
	return 0x01
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetServiceChoice() uint8 {
	return 0x01
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetSubscriberProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetInitiatingDeviceIdentifier() *BACnetContextTagObjectIdentifier {
	return m.InitiatingDeviceIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetMonitoredObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLifetimeInSeconds() *BACnetContextTagUnsignedInteger {
	return m.LifetimeInSeconds
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetListOfValues() *BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestConfirmedCOVNotification factory function for BACnetConfirmedServiceRequestConfirmedCOVNotification
func NewBACnetConfirmedServiceRequestConfirmedCOVNotification(subscriberProcessIdentifier *BACnetContextTagUnsignedInteger, initiatingDeviceIdentifier *BACnetContextTagObjectIdentifier, monitoredObjectIdentifier *BACnetContextTagObjectIdentifier, lifetimeInSeconds *BACnetContextTagUnsignedInteger, listOfValues *BACnetPropertyValues, len uint16) *BACnetConfirmedServiceRequest {
	child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:   subscriberProcessIdentifier,
		InitiatingDeviceIdentifier:    initiatingDeviceIdentifier,
		MonitoredObjectIdentifier:     monitoredObjectIdentifier,
		LifetimeInSeconds:             lifetimeInSeconds,
		ListOfValues:                  listOfValues,
		BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetConfirmedServiceRequest
}

func CastBACnetConfirmedServiceRequestConfirmedCOVNotification(structType interface{}) *BACnetConfirmedServiceRequestConfirmedCOVNotification {
	if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedCOVNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestConfirmedCOVNotification(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedCOVNotification"
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits()

	// Simple field (initiatingDeviceIdentifier)
	lengthInBits += m.InitiatingDeviceIdentifier.GetLengthInBits()

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (lifetimeInSeconds)
	lengthInBits += m.LifetimeInSeconds.GetLengthInBits()

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedCOVNotificationParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetConfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (subscriberProcessIdentifier)
	if pullErr := readBuffer.PullContext("subscriberProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_subscriberProcessIdentifier, _subscriberProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _subscriberProcessIdentifierErr != nil {
		return nil, errors.Wrap(_subscriberProcessIdentifierErr, "Error parsing 'subscriberProcessIdentifier' field")
	}
	subscriberProcessIdentifier := CastBACnetContextTagUnsignedInteger(_subscriberProcessIdentifier)
	if closeErr := readBuffer.CloseContext("subscriberProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (initiatingDeviceIdentifier)
	if pullErr := readBuffer.PullContext("initiatingDeviceIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_initiatingDeviceIdentifier, _initiatingDeviceIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _initiatingDeviceIdentifierErr != nil {
		return nil, errors.Wrap(_initiatingDeviceIdentifierErr, "Error parsing 'initiatingDeviceIdentifier' field")
	}
	initiatingDeviceIdentifier := CastBACnetContextTagObjectIdentifier(_initiatingDeviceIdentifier)
	if closeErr := readBuffer.CloseContext("initiatingDeviceIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field")
	}
	monitoredObjectIdentifier := CastBACnetContextTagObjectIdentifier(_monitoredObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lifetimeInSeconds)
	if pullErr := readBuffer.PullContext("lifetimeInSeconds"); pullErr != nil {
		return nil, pullErr
	}
	_lifetimeInSeconds, _lifetimeInSecondsErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _lifetimeInSecondsErr != nil {
		return nil, errors.Wrap(_lifetimeInSecondsErr, "Error parsing 'lifetimeInSeconds' field")
	}
	lifetimeInSeconds := CastBACnetContextTagUnsignedInteger(_lifetimeInSeconds)
	if closeErr := readBuffer.CloseContext("lifetimeInSeconds"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfValues, _listOfValuesErr := BACnetPropertyValuesParse(readBuffer, uint8(uint8(4)), BACnetObjectType(monitoredObjectIdentifier.GetObjectType()))
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field")
	}
	listOfValues := CastBACnetPropertyValues(_listOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestConfirmedCOVNotification{
		SubscriberProcessIdentifier:   CastBACnetContextTagUnsignedInteger(subscriberProcessIdentifier),
		InitiatingDeviceIdentifier:    CastBACnetContextTagObjectIdentifier(initiatingDeviceIdentifier),
		MonitoredObjectIdentifier:     CastBACnetContextTagObjectIdentifier(monitoredObjectIdentifier),
		LifetimeInSeconds:             CastBACnetContextTagUnsignedInteger(lifetimeInSeconds),
		ListOfValues:                  CastBACnetPropertyValues(listOfValues),
		BACnetConfirmedServiceRequest: &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child.BACnetConfirmedServiceRequest, nil
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (subscriberProcessIdentifier)
		if pushErr := writeBuffer.PushContext("subscriberProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_subscriberProcessIdentifierErr := m.SubscriberProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("subscriberProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _subscriberProcessIdentifierErr != nil {
			return errors.Wrap(_subscriberProcessIdentifierErr, "Error serializing 'subscriberProcessIdentifier' field")
		}

		// Simple Field (initiatingDeviceIdentifier)
		if pushErr := writeBuffer.PushContext("initiatingDeviceIdentifier"); pushErr != nil {
			return pushErr
		}
		_initiatingDeviceIdentifierErr := m.InitiatingDeviceIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("initiatingDeviceIdentifier"); popErr != nil {
			return popErr
		}
		if _initiatingDeviceIdentifierErr != nil {
			return errors.Wrap(_initiatingDeviceIdentifierErr, "Error serializing 'initiatingDeviceIdentifier' field")
		}

		// Simple Field (monitoredObjectIdentifier)
		if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_monitoredObjectIdentifierErr := m.MonitoredObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _monitoredObjectIdentifierErr != nil {
			return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
		}

		// Simple Field (lifetimeInSeconds)
		if pushErr := writeBuffer.PushContext("lifetimeInSeconds"); pushErr != nil {
			return pushErr
		}
		_lifetimeInSecondsErr := m.LifetimeInSeconds.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lifetimeInSeconds"); popErr != nil {
			return popErr
		}
		if _lifetimeInSecondsErr != nil {
			return errors.Wrap(_lifetimeInSecondsErr, "Error serializing 'lifetimeInSeconds' field")
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

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedCOVNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestConfirmedCOVNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}

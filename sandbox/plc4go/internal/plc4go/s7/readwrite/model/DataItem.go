//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
    "errors"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/model/values"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/model/values/iec61131"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    api "plc4x.apache.org/plc4go-modbus-driver/v0/pkg/plc4go/values"
)

func DataItemParse(io *spi.ReadBuffer, dataProtocolId uint8, stringLength int32) (api.PlcValue, error) {
    switch {
        case dataProtocolId == 01: // BOOL


            // Simple Field (value)
            value, _valueErr := io.ReadBit()
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcBOOL(value), nil
        case dataProtocolId == 11: // BOOL

            // Array Field (value)
            var value []api.PlcValue
            return values.NewPlcList(value), nil
        case dataProtocolId == 12: // BOOL

            // Array Field (value)
            var value []api.PlcValue
            return values.NewPlcList(value), nil
        case dataProtocolId == 13: // BOOL

            // Array Field (value)
            var value []api.PlcValue
            return values.NewPlcList(value), nil
        case dataProtocolId == 14: // BOOL

            // Array Field (value)
            var value []api.PlcValue
            return values.NewPlcList(value), nil
        case dataProtocolId == 21: // SINT

            // Simple Field (value)
            value, _valueErr := io.ReadInt8(8)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcSINT(value), nil
        case dataProtocolId == 22: // USINT

            // Simple Field (value)
            value, _valueErr := io.ReadUint8(8)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcUSINT(value), nil
        case dataProtocolId == 23: // INT

            // Simple Field (value)
            value, _valueErr := io.ReadInt16(16)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcINT(value), nil
        case dataProtocolId == 24: // UINT

            // Simple Field (value)
            value, _valueErr := io.ReadUint16(16)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcUINT(value), nil
        case dataProtocolId == 25: // DINT

            // Simple Field (value)
            value, _valueErr := io.ReadInt32(32)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcDINT(value), nil
        case dataProtocolId == 26: // UDINT

            // Simple Field (value)
            value, _valueErr := io.ReadUint32(32)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcUDINT(value), nil
        case dataProtocolId == 27: // LINT

            // Simple Field (value)
            value, _valueErr := io.ReadInt64(64)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcLINT(value), nil
        case dataProtocolId == 28: // ULINT

            // Simple Field (value)
            value, _valueErr := io.ReadUint64(64)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcULINT(value), nil
        case dataProtocolId == 31: // REAL

            // Simple Field (value)
            value, _valueErr := io.ReadFloat32(32)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcREAL(value), nil
        case dataProtocolId == 32: // LREAL

            // Simple Field (value)
            value, _valueErr := io.ReadFloat64(64)
            if _valueErr != nil {
                return nil, errors.New("Error parsing 'value' field " + _valueErr.Error())
            }
            return iec61131.NewPlcLREAL(value), nil
        case dataProtocolId == 41: // STRING
            //return iec61131.NewPlcSTRING(value), nil
        case dataProtocolId == 42: // STRING
            //return iec61131.NewPlcSTRING(value), nil
        case dataProtocolId == 43: // STRING
            //return iec61131.NewPlcSTRING(value), nil
        case dataProtocolId == 44: // STRING
            //return iec61131.NewPlcSTRING(value), nil
        case dataProtocolId == 51: // Time
            //return new PlcTime(LocalTime.of((int) hours, (int) minutes, (int) seconds))
        case dataProtocolId == 52: // Time
            //return new PlcTime(LocalTime.of((int) hours, (int) minutes, (int) seconds))
        case dataProtocolId == 53: // Time
            //return new PlcTime(LocalTime.of((int) hours, (int) minutes, (int) seconds))
        case dataProtocolId == 54: // Date
            //return new PlcDate(LocalDate.of((int) year, (int) month, (int) day))
        case dataProtocolId == 55: // Time
            //return new PlcTime(LocalTime.of((int) hours, (int) minutes, (int) seconds))
        case dataProtocolId == 56: // DateTime
            //return new PlcDateTime(LocalDateTime.of((int) year, (int) month, (int) day, (int) hours, (int) minutes, (int) seconds))
    }
    return nil, errors.New("unsupported type")
}

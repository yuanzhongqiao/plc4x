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
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.WithOption;
import org.apache.plc4x.java.spi.generation.ByteOrder;
import org.apache.plc4x.java.spi.generation.EvaluationHelper;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.SerializationException;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.values.*;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

// Code generated by code-generation. DO NOT EDIT.

public class DataItem {

  private static final Logger LOGGER = LoggerFactory.getLogger(DataItem.class);

  public static PlcValue staticParse(
      ReadBuffer readBuffer, String dataProtocolId, Integer stringLength) throws ParseException {
    if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BOOL")) { // BOOL

      // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
      {
        short reserved = /*TODO: migrate me*/ /*TODO: migrate me*/
            readBuffer.readUnsignedShort("", 7);
        if (reserved != (short) 0x00) {
          LOGGER.info(
              "Expected constant value " + 0x00 + " but got " + reserved + " for reserved field.");
        }
      }

      // Simple Field (value)
      Boolean value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readBit("");

      return new PlcBOOL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BYTE")) { // BYTE

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcBYTE(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WORD")) { // WORD

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DWORD")) { // DWORD

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 32);

      return new PlcDWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LWORD")) { // LWORD

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return new PlcLWORD(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_SINT")) { // SINT

      // Simple Field (value)
      Byte value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readSignedByte("", 8);

      return new PlcSINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_USINT")) { // USINT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      return new PlcUSINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_INT")) { // INT

      // Simple Field (value)
      Short value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readShort("", 16);

      return new PlcINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UINT")) { // UINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      return new PlcUINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DINT")) { // DINT

      // Simple Field (value)
      Integer value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readInt("", 32);

      return new PlcDINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UDINT")) { // UDINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedLong("", 32);

      return new PlcUDINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LINT")) { // LINT

      // Simple Field (value)
      Long value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readLong("", 64);

      return new PlcLINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_ULINT")) { // ULINT

      // Simple Field (value)
      BigInteger value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return new PlcULINT(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_REAL")) { // REAL

      // Simple Field (value)
      Float value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readFloat("", 32);

      return new PlcREAL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LREAL")) { // LREAL

      // Simple Field (value)
      Double value = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readDouble("", 64);

      return new PlcLREAL(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_CHAR")) { // CHAR

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", 8, WithOption.WithEncoding("UTF-8"));

      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WCHAR")) { // CHAR

      // Simple Field (value)
      String value = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readString("", 16, WithOption.WithEncoding("UTF-16"));

      return new PlcCHAR(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_STRING")) { // STRING

      // Manual Field (value)
      String value =
          org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseS7String(
              readBuffer, stringLength, "UTF-8");

      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WSTRING")) { // STRING

      // Manual Field (value)
      String value =
          org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.parseS7String(
              readBuffer, stringLength, "UTF-16");

      return new PlcSTRING(value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME")) { // TIME

      // Simple Field (milliseconds)
      Long milliseconds = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedLong("", 32);

      return PlcTIME.ofMilliseconds(milliseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME")) { // LTIME

      // Simple Field (nanoseconds)
      BigInteger nanoseconds = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return PlcLTIME.ofNanoseconds(nanoseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE")) { // DATE

      // Simple Field (daysSinceSiemensEpoch)
      Integer daysSinceSiemensEpoch = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedInt("", 16);

      return PlcDATE.ofDaysSinceSiemensEpoch(daysSinceSiemensEpoch);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY

      // Simple Field (millisecondsSinceMidnight)
      Long millisecondsSinceMidnight = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedLong("", 32);

      return PlcTIME_OF_DAY.ofMillisecondsSinceMidnight(millisecondsSinceMidnight);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY

      // Simple Field (nanosecondsSinceMidnight)
      BigInteger nanosecondsSinceMidnight = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedBigInteger("", 64);

      return PlcLTIME_OF_DAY.ofNanosecondsSinceMidnight(nanosecondsSinceMidnight);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME

      // Simple Field (year)
      Integer year = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedInt("", 16);

      // Simple Field (month)
      Short month = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      // Simple Field (day)
      Short day = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      // Simple Field (dayOfWeek)
      Short dayOfWeek = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedShort("", 8);

      // Simple Field (hour)
      Short hour = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      // Simple Field (minutes)
      Short minutes = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      // Simple Field (seconds)
      Short seconds = /*TODO: migrate me*/ /*TODO: migrate me*/ readBuffer.readUnsignedShort("", 8);

      // Simple Field (nanoseconds)
      Long nanoseconds = /*TODO: migrate me*/ /*TODO: migrate me*/
          readBuffer.readUnsignedLong("", 32);

      return PlcDATE_AND_TIME.ofSegments(
          year.intValue(),
          (month == 0) ? 1 : month.intValue(),
          (day == 0) ? 1 : day.intValue(),
          hour.intValue(),
          minutes.intValue(),
          seconds.intValue(),
          nanoseconds.intValue());
    }
    return null;
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer, PlcValue _value, String dataProtocolId, Integer stringLength)
      throws SerializationException {
    staticSerialize(writeBuffer, _value, dataProtocolId, stringLength, ByteOrder.BIG_ENDIAN);
  }

  public static void staticSerialize(
      WriteBuffer writeBuffer,
      PlcValue _value,
      String dataProtocolId,
      Integer stringLength,
      ByteOrder byteOrder)
      throws SerializationException {
    if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BOOL")) { // BOOL
      // Reserved Field
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort(
          "", 7, ((Number) (short) 0x00).shortValue());
      // Simple Field (value)
      boolean value = _value.getBoolean();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeBit("", value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BYTE")) { // BYTE
      // Simple Field (value)
      short value = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WORD")) { // WORD
      // Simple Field (value)
      int value = _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DWORD")) { // DWORD
      // Simple Field (value)
      long value = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LWORD")) { // LWORD
      // Simple Field (value)
      BigInteger value = _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_SINT")) { // SINT
      // Simple Field (value)
      byte value = _value.getByte();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeSignedByte("", 8, ((Number) (value)).byteValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_USINT")) { // USINT
      // Simple Field (value)
      short value = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_INT")) { // INT
      // Simple Field (value)
      short value = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeShort("", 16, ((Number) (value)).shortValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UINT")) { // UINT
      // Simple Field (value)
      int value = _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DINT")) { // DINT
      // Simple Field (value)
      int value = _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeInt("", 32, ((Number) (value)).intValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UDINT")) { // UDINT
      // Simple Field (value)
      long value = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong("", 32, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LINT")) { // LINT
      // Simple Field (value)
      long value = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeLong("", 64, ((Number) (value)).longValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_ULINT")) { // ULINT
      // Simple Field (value)
      BigInteger value = _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, value);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_REAL")) { // REAL
      // Simple Field (value)
      float value = _value.getFloat();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeFloat("", 32, (value));
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LREAL")) { // LREAL
      // Simple Field (value)
      double value = _value.getDouble();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeDouble("", 64, (value));
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_CHAR")) { // CHAR
      // Simple Field (value)
      String value = _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString(
          "", 8, value, WithOption.WithEncoding("UTF-8"));
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WCHAR")) { // CHAR
      // Simple Field (value)
      String value = _value.getString();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeString(
          "", 16, value, WithOption.WithEncoding("UTF-16"));
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_STRING")) { // STRING
      // Manual Field (value)
      org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeS7String(
          writeBuffer, _value, stringLength, "UTF-8");
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WSTRING")) { // STRING
      // Manual Field (value)
      org.apache.plc4x.java.s7.readwrite.utils.StaticHelper.serializeS7String(
          writeBuffer, _value, stringLength, "UTF-16");
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME")) { // TIME
      // Simple Field (milliseconds)
      long milliseconds = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong(
          "", 32, ((Number) (milliseconds)).longValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME")) { // LTIME
      // Simple Field (nanoseconds)
      BigInteger nanoseconds = _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger("", 64, nanoseconds);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE")) { // DATE
      // Simple Field (daysSinceSiemensEpoch)
      int daysSinceSiemensEpoch = _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt(
          "", 16, ((Number) (daysSinceSiemensEpoch)).intValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY
      // Simple Field (millisecondsSinceMidnight)
      long millisecondsSinceMidnight = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong(
          "", 32, ((Number) (millisecondsSinceMidnight)).longValue());
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY
      // Simple Field (nanosecondsSinceMidnight)
      BigInteger nanosecondsSinceMidnight = _value.getBigInteger();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedBigInteger(
          "", 64, nanosecondsSinceMidnight);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME
      // Simple Field (year)
      int year = _value.getInt();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedInt("", 16, ((Number) (year)).intValue());
      // Simple Field (month)
      short month = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (month)).shortValue());
      // Simple Field (day)
      short day = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (day)).shortValue());
      // Simple Field (dayOfWeek)
      short dayOfWeek = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort(
          "", 8, ((Number) (dayOfWeek)).shortValue());
      // Simple Field (hour)
      short hour = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (hour)).shortValue());
      // Simple Field (minutes)
      short minutes = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (minutes)).shortValue());
      // Simple Field (seconds)
      short seconds = _value.getShort();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedShort("", 8, ((Number) (seconds)).shortValue());
      // Simple Field (nanoseconds)
      long nanoseconds = _value.getLong();
      /*TODO: migrate me*/
      /*TODO: migrate me*/ writeBuffer.writeUnsignedLong(
          "", 32, ((Number) (nanoseconds)).longValue());
    }
  }

  public static int getLengthInBytes(PlcValue _value, String dataProtocolId, Integer stringLength) {
    return (int) Math.ceil((float) getLengthInBits(_value, dataProtocolId, stringLength) / 8.0);
  }

  public static int getLengthInBits(PlcValue _value, String dataProtocolId, Integer stringLength) {
    int sizeInBits = 0;
    if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BOOL")) { // BOOL
      // Reserved Field
      sizeInBits += 7;
      // Simple Field (value)
      sizeInBits += 1;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_BYTE")) { // BYTE
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WORD")) { // WORD
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DWORD")) { // DWORD
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LWORD")) { // LWORD
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_SINT")) { // SINT
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_USINT")) { // USINT
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_INT")) { // INT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UINT")) { // UINT
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DINT")) { // DINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_UDINT")) { // UDINT
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LINT")) { // LINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_ULINT")) { // ULINT
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_REAL")) { // REAL
      // Simple Field (value)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LREAL")) { // LREAL
      // Simple Field (value)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_CHAR")) { // CHAR
      // Simple Field (value)
      sizeInBits += 8;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WCHAR")) { // CHAR
      // Simple Field (value)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_STRING")) { // STRING
      // Manual Field (value)
      sizeInBits += (STR_LEN(_value)) + (2);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_WSTRING")) { // STRING
      // Manual Field (value)
      sizeInBits += (((STR_LEN(_value)) * (2))) + (2);
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME")) { // TIME
      // Simple Field (milliseconds)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME")) { // LTIME
      // Simple Field (nanoseconds)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE")) { // DATE
      // Simple Field (daysSinceSiemensEpoch)
      sizeInBits += 16;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_TIME_OF_DAY")) { // TIME_OF_DAY
      // Simple Field (millisecondsSinceMidnight)
      sizeInBits += 32;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_LTIME_OF_DAY")) { // LTIME_OF_DAY
      // Simple Field (nanosecondsSinceMidnight)
      sizeInBits += 64;
    } else if (EvaluationHelper.equals(dataProtocolId, "IEC61131_DATE_AND_TIME")) { // DATE_AND_TIME
      // Simple Field (year)
      sizeInBits += 16;
      // Simple Field (month)
      sizeInBits += 8;
      // Simple Field (day)
      sizeInBits += 8;
      // Simple Field (dayOfWeek)
      sizeInBits += 8;
      // Simple Field (hour)
      sizeInBits += 8;
      // Simple Field (minutes)
      sizeInBits += 8;
      // Simple Field (seconds)
      sizeInBits += 8;
      // Simple Field (nanoseconds)
      sizeInBits += 32;
    }
    return sizeInBits;
  }
}

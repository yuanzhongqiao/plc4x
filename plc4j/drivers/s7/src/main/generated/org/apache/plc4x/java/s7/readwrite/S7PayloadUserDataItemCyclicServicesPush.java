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

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class S7PayloadUserDataItemCyclicServicesPush extends S7PayloadUserDataItem
    implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x02;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x00;
  }

  public Short getCpuSubfunction() {
    return (short) 0x01;
  }

  // Properties.
  protected final int itemsCount;
  protected final List<AssociatedValueType> items;

  public S7PayloadUserDataItemCyclicServicesPush(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      int itemsCount,
      List<AssociatedValueType> items) {
    super(returnCode, transportSize, dataLength);
    this.itemsCount = itemsCount;
    this.items = items;
  }

  public int getItemsCount() {
    return itemsCount;
  }

  public List<AssociatedValueType> getItems() {
    return items;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadUserDataItemCyclicServicesPush");

    // Simple Field (itemsCount)
    writeSimpleField("itemsCount", itemsCount, writeUnsignedInt(writeBuffer, 16));

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("S7PayloadUserDataItemCyclicServicesPush");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCyclicServicesPush _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (itemsCount)
    lengthInBits += 16;

    // Array field
    if (items != null) {
      int i = 0;
      for (AssociatedValueType element : items) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= items.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCyclicServicesPush");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int itemsCount = readSimpleField("itemsCount", readUnsignedInt(readBuffer, 16));

    List<AssociatedValueType> items =
        readCountArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> AssociatedValueType.staticParse(readBuffer), readBuffer),
            itemsCount);

    readBuffer.closeContext("S7PayloadUserDataItemCyclicServicesPush");
    // Create the instance
    return new S7PayloadUserDataItemCyclicServicesPushBuilderImpl(itemsCount, items);
  }

  public static class S7PayloadUserDataItemCyclicServicesPushBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final int itemsCount;
    private final List<AssociatedValueType> items;

    public S7PayloadUserDataItemCyclicServicesPushBuilderImpl(
        int itemsCount, List<AssociatedValueType> items) {
      this.itemsCount = itemsCount;
      this.items = items;
    }

    public S7PayloadUserDataItemCyclicServicesPush build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemCyclicServicesPush s7PayloadUserDataItemCyclicServicesPush =
          new S7PayloadUserDataItemCyclicServicesPush(
              returnCode, transportSize, dataLength, itemsCount, items);
      return s7PayloadUserDataItemCyclicServicesPush;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCyclicServicesPush)) {
      return false;
    }
    S7PayloadUserDataItemCyclicServicesPush that = (S7PayloadUserDataItemCyclicServicesPush) o;
    return (getItemsCount() == that.getItemsCount())
        && (getItems() == that.getItems())
        && super.equals(that);
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItemsCount(), getItems());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}

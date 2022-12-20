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
package org.apache.plc4x.java.opcua.readwrite;

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

public class FindServersOnNetworkResponse extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "12193";
  }

  // Properties.
  protected final ExtensionObjectDefinition responseHeader;
  protected final long lastCounterResetTime;
  protected final int noOfServers;
  protected final List<ExtensionObjectDefinition> servers;

  public FindServersOnNetworkResponse(
      ExtensionObjectDefinition responseHeader,
      long lastCounterResetTime,
      int noOfServers,
      List<ExtensionObjectDefinition> servers) {
    super();
    this.responseHeader = responseHeader;
    this.lastCounterResetTime = lastCounterResetTime;
    this.noOfServers = noOfServers;
    this.servers = servers;
  }

  public ExtensionObjectDefinition getResponseHeader() {
    return responseHeader;
  }

  public long getLastCounterResetTime() {
    return lastCounterResetTime;
  }

  public int getNoOfServers() {
    return noOfServers;
  }

  public List<ExtensionObjectDefinition> getServers() {
    return servers;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("FindServersOnNetworkResponse");

    // Simple Field (responseHeader)
    writeSimpleField("responseHeader", responseHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (lastCounterResetTime)
    writeSimpleField(
        "lastCounterResetTime", lastCounterResetTime, writeSignedLong(writeBuffer, 64));

    // Simple Field (noOfServers)
    writeSimpleField("noOfServers", noOfServers, writeSignedInt(writeBuffer, 32));

    // Array Field (servers)
    writeComplexTypeArrayField("servers", servers, writeBuffer);

    writeBuffer.popContext("FindServersOnNetworkResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FindServersOnNetworkResponse _value = this;

    // Simple field (responseHeader)
    lengthInBits += responseHeader.getLengthInBits();

    // Simple field (lastCounterResetTime)
    lengthInBits += 64;

    // Simple field (noOfServers)
    lengthInBits += 32;

    // Array field
    if (servers != null) {
      int i = 0;
      for (ExtensionObjectDefinition element : servers) {
        boolean last = ++i >= servers.size();
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static FindServersOnNetworkResponseBuilder staticParseBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("FindServersOnNetworkResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;

    ExtensionObjectDefinition responseHeader =
        readSimpleField(
            "responseHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("394")),
                readBuffer));

    long lastCounterResetTime =
        readSimpleField("lastCounterResetTime", readSignedLong(readBuffer, 64));

    int noOfServers = readSimpleField("noOfServers", readSignedInt(readBuffer, 32));

    List<ExtensionObjectDefinition> servers =
        readCountArrayField(
            "servers",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("12191")),
                readBuffer),
            noOfServers);

    readBuffer.closeContext("FindServersOnNetworkResponse");
    // Create the instance
    return new FindServersOnNetworkResponseBuilder(
        responseHeader, lastCounterResetTime, noOfServers, servers);
  }

  public static class FindServersOnNetworkResponseBuilder
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition responseHeader;
    private final long lastCounterResetTime;
    private final int noOfServers;
    private final List<ExtensionObjectDefinition> servers;

    public FindServersOnNetworkResponseBuilder(
        ExtensionObjectDefinition responseHeader,
        long lastCounterResetTime,
        int noOfServers,
        List<ExtensionObjectDefinition> servers) {

      this.responseHeader = responseHeader;
      this.lastCounterResetTime = lastCounterResetTime;
      this.noOfServers = noOfServers;
      this.servers = servers;
    }

    public FindServersOnNetworkResponse build() {
      FindServersOnNetworkResponse findServersOnNetworkResponse =
          new FindServersOnNetworkResponse(
              responseHeader, lastCounterResetTime, noOfServers, servers);
      return findServersOnNetworkResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FindServersOnNetworkResponse)) {
      return false;
    }
    FindServersOnNetworkResponse that = (FindServersOnNetworkResponse) o;
    return (getResponseHeader() == that.getResponseHeader())
        && (getLastCounterResetTime() == that.getLastCounterResetTime())
        && (getNoOfServers() == that.getNoOfServers())
        && (getServers() == that.getServers())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getResponseHeader(),
        getLastCounterResetTime(),
        getNoOfServers(),
        getServers());
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
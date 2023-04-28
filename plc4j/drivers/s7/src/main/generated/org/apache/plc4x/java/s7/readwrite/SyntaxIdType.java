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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum SyntaxIdType {
  S7ANY((short) 0x01),
  PBC_ID((short) 0x13),
  ALARM_LOCKFREESET((short) 0x15),
  ALARM_INDSET((short) 0x16),
  ALARM_ACKSET((short) 0x19),
  ALARM_QUERYREQSET((short) 0x1A),
  NOTIFY_INDSET((short) 0x1C),
  NCK((short) 0x82),
  NCK_METRIC((short) 0x83),
  NCK_INCH((short) 0x84),
  DRIVEESANY((short) 0xA2),
  SYM1200((short) 0xB2),
  DBREAD((short) 0xB0);
  private static final Map<Short, SyntaxIdType> map;

  static {
    map = new HashMap<>();
    for (SyntaxIdType value : SyntaxIdType.values()) {
      map.put(value.getValue(), value);
    }
  }

  private final short value;

  SyntaxIdType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static SyntaxIdType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}

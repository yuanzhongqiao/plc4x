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

public enum SzlModuleTypeClass {
  CPU((byte) 0x0),
  IM((byte) 0x4),
  FM((byte) 0x8),
  CP((byte) 0xC);
  private static final Map<Byte, SzlModuleTypeClass> map;

  static {
    map = new HashMap<>();
    for (SzlModuleTypeClass value : SzlModuleTypeClass.values()) {
      map.put(value.getValue(), value);
    }
  }

  private final byte value;

  SzlModuleTypeClass(byte value) {
    this.value = value;
  }

  public byte getValue() {
    return value;
  }

  public static SzlModuleTypeClass enumForValue(byte value) {
    return map.get(value);
  }

  public static Boolean isDefined(byte value) {
    return map.containsKey(value);
  }
}

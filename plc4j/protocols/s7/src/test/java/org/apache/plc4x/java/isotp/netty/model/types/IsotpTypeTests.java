package org.apache.plc4x.java.isotp.netty.model.types;

import org.junit.jupiter.api.Tag;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class IsotpTypeTests {

    @Test
    @Tag("fast")
    void deviceGroup() {
        DeviceGroup deviceGroup;

        deviceGroup = DeviceGroup.PG_OR_PC;
        assertTrue(deviceGroup == DeviceGroup.PG_OR_PC, "1 incorrectly mapped");
        assertTrue(deviceGroup.getCode() == (byte)1, "code is not 1");

        deviceGroup = DeviceGroup.OS;
        assertTrue(deviceGroup == DeviceGroup.OS, "2 incorrectly mapped");
        assertTrue(deviceGroup.getCode() == (byte)2, "code is not 2");

        deviceGroup = DeviceGroup.OTHERS;
        assertTrue(deviceGroup == DeviceGroup.OTHERS, "3 incorrectly mapped");
        assertTrue(deviceGroup.getCode() == (byte)3, "code is not 3");
    }

    @Test
    @Tag("fast")
    void deviceGroupUnknown() {
        DeviceGroup deviceGroup = DeviceGroup.valueOf((byte)0x40);

        assertNull(deviceGroup, "expected device group to be null");
    }


    @Test
    @Tag("fast")
    void disconnectReason() {
        DisconnectReason disconnectReason = DisconnectReason.ADDRESS_UNKNOWN;

        assertTrue(DisconnectReason.valueOf((byte)3) == DisconnectReason.ADDRESS_UNKNOWN, "3 incorrectly mapped");
        assertTrue(disconnectReason.getCode() == (byte)3, "code is not 3");
    }

    @Test
    @Tag("fast")
    void diosconectReasonUnknown() {
        DisconnectReason disconnectReason = DisconnectReason.valueOf((byte)4);

        assertNull(disconnectReason, "expected disconnect reason to be null");
    }

    @Test
    @Tag("fast")
    void parameterCode() {
        ParameterCode parameterCode = ParameterCode.CALLING_TSAP;

        assertTrue(ParameterCode.valueOf((byte)0xC1) == ParameterCode.CALLING_TSAP, "0xC1 incorrectly mapped");
        assertTrue(parameterCode.getCode() == (byte)0xC1, "code is not 0xC1");
    }

    @Test
    @Tag("fast")
    void parameterCodeUnknown() {
        ParameterCode parameterCode = ParameterCode.valueOf((byte)0x90);

        assertNull(parameterCode, "expected parameter code to be null");
    }

    @Test
    @Tag("fast")
    void protocolClass() {
        ProtocolClass protocolClass;

        protocolClass = ProtocolClass.CLASS_1;
        assertTrue(protocolClass == ProtocolClass.CLASS_1, "0x10 incorrectly mapped");
        assertTrue(protocolClass.getCode() == (byte)0x10, "code is not 0x10");

        protocolClass = ProtocolClass.CLASS_2;
        assertTrue(protocolClass == ProtocolClass.CLASS_2, "0x20 incorrectly mapped");
        assertTrue(protocolClass.getCode() == (byte)0x20, "code is not 0x20");

        protocolClass = ProtocolClass.CLASS_3;
        assertTrue(protocolClass == ProtocolClass.CLASS_3, "0x30 incorrectly mapped");
        assertTrue(protocolClass.getCode() == (byte)0x30, "code is not 0x30");

        protocolClass = ProtocolClass.CLASS_4;
        assertTrue(protocolClass == ProtocolClass.CLASS_4, "0x40 incorrectly mapped");
        assertTrue(protocolClass.getCode() == (byte)0x40, "code is not 0x40");
    }

    @Test
    @Tag("fast")
    void protocolClassUnknown() {
        ProtocolClass protocolClass = ProtocolClass.valueOf((byte)0x50);

        assertNull(protocolClass, "expected protocol class to be null");
    }

    @Test
    @Tag("fast")
    void rejectCause() {
        RejectCause rejectCause = RejectCause.INVALID_PARAMETER_TYPE;

        assertTrue(RejectCause.valueOf((byte)0x03) == RejectCause.INVALID_PARAMETER_TYPE, "0x03 incorrectly mapped");
        assertTrue(rejectCause.getCode() == (byte)0x03, "code is not 0x03");
    }

    @Test
    @Tag("fast")
    void rejectClauseUnknown() {
        RejectCause rejectCause = RejectCause.valueOf((byte)0x90);

        assertNull(rejectCause, "expected reject cause to be null");
    }

    @Test
    @Tag("fast")
    void tpduCode() {
        TpduCode tpduCode = TpduCode.DATA;

        assertTrue(TpduCode.valueOf((byte)0xF0) == TpduCode.DATA, "0xF0 incorrectly mapped");
        assertTrue(tpduCode.getCode() == (byte)0xF0, "code is not 0xF0");
    }

    @Test
    @Tag("fast")
    void tpduCodeUnknown() {
        TpduCode tpduCode = TpduCode.valueOf((byte)0x01);

        assertNull(tpduCode, "expected tpdu code to be null");
    }

    @Test
    @Tag("fast")
    void typduSize() {
        TpduSize tpduSize = TpduSize.SIZE_128;

        assertTrue(TpduSize.valueOf((byte)0x07) == TpduSize.SIZE_128, "0x07 incorrectly mapped");
        assertTrue(tpduSize.getCode() == (byte)0x07, "code is not 0x07");
    }

    @Test
    @Tag("fast")
    void tpduSizeUnknown() {
        TpduSize tpduSize = TpduSize.valueOf((byte)0x06);

        assertNull(tpduSize, "expected tpdu size to be null");
    }


}
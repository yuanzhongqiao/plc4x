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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableTrust {
  TrustListType_Size((int) 12523L),
  TrustListType_OpenCount((int) 12526L),
  TrustListType_Open_InputArguments((int) 12528L),
  TrustListType_Open_OutputArguments((int) 12529L),
  TrustListType_Close_InputArguments((int) 12531L),
  TrustListType_Read_InputArguments((int) 12533L),
  TrustListType_Read_OutputArguments((int) 12534L),
  TrustListType_Write_InputArguments((int) 12536L),
  TrustListType_GetPosition_InputArguments((int) 12538L),
  TrustListType_GetPosition_OutputArguments((int) 12539L),
  TrustListType_SetPosition_InputArguments((int) 12541L),
  TrustListType_LastUpdateTime((int) 12542L),
  TrustListType_OpenWithMasks_InputArguments((int) 12544L),
  TrustListType_OpenWithMasks_OutputArguments((int) 12545L),
  TrustListType_CloseAndUpdate_OutputArguments((int) 12547L),
  TrustListType_AddCertificate_InputArguments((int) 12549L),
  TrustListType_RemoveCertificate_InputArguments((int) 12551L),
  TrustListMasks_EnumValues((int) 12553L),
  TrustListUpdatedAuditEventType_EventId((int) 12562L),
  TrustListUpdatedAuditEventType_EventType((int) 12563L),
  TrustListUpdatedAuditEventType_SourceNode((int) 12564L),
  TrustListUpdatedAuditEventType_SourceName((int) 12565L),
  TrustListUpdatedAuditEventType_Time((int) 12566L),
  TrustListUpdatedAuditEventType_ReceiveTime((int) 12567L),
  TrustListUpdatedAuditEventType_LocalTime((int) 12568L),
  TrustListUpdatedAuditEventType_Message((int) 12569L),
  TrustListUpdatedAuditEventType_Severity((int) 12570L),
  TrustListUpdatedAuditEventType_ActionTimeStamp((int) 12571L),
  TrustListUpdatedAuditEventType_Status((int) 12572L),
  TrustListUpdatedAuditEventType_ServerId((int) 12573L),
  TrustListUpdatedAuditEventType_ClientAuditEntryId((int) 12574L),
  TrustListUpdatedAuditEventType_ClientUserId((int) 12575L),
  TrustListType_Writable((int) 12698L),
  TrustListType_UserWritable((int) 12699L),
  TrustListType_CloseAndUpdate_InputArguments((int) 12705L),
  TrustListType_MimeType((int) 13403L),
  TrustListType_UpdateFrequency((int) 19296L),
  TrustListOutOfDateAlarmType_EventId((int) 19298L),
  TrustListOutOfDateAlarmType_EventType((int) 19299L),
  TrustListOutOfDateAlarmType_SourceNode((int) 19300L),
  TrustListOutOfDateAlarmType_SourceName((int) 19301L),
  TrustListOutOfDateAlarmType_Time((int) 19302L),
  TrustListOutOfDateAlarmType_ReceiveTime((int) 19303L),
  TrustListOutOfDateAlarmType_LocalTime((int) 19304L),
  TrustListOutOfDateAlarmType_Message((int) 19305L),
  TrustListOutOfDateAlarmType_Severity((int) 19306L),
  TrustListOutOfDateAlarmType_ConditionClassId((int) 19307L),
  TrustListOutOfDateAlarmType_ConditionClassName((int) 19308L),
  TrustListOutOfDateAlarmType_ConditionSubClassId((int) 19309L),
  TrustListOutOfDateAlarmType_ConditionSubClassName((int) 19310L),
  TrustListOutOfDateAlarmType_ConditionName((int) 19311L),
  TrustListOutOfDateAlarmType_BranchId((int) 19312L),
  TrustListOutOfDateAlarmType_Retain((int) 19313L),
  TrustListOutOfDateAlarmType_EnabledState((int) 19314L),
  TrustListOutOfDateAlarmType_EnabledState_Id((int) 19315L),
  TrustListOutOfDateAlarmType_EnabledState_Name((int) 19316L),
  TrustListOutOfDateAlarmType_EnabledState_Number((int) 19317L),
  TrustListOutOfDateAlarmType_EnabledState_EffectiveDisplayName((int) 19318L),
  TrustListOutOfDateAlarmType_EnabledState_TransitionTime((int) 19319L),
  TrustListOutOfDateAlarmType_EnabledState_EffectiveTransitionTime((int) 19320L),
  TrustListOutOfDateAlarmType_EnabledState_TrueState((int) 19321L),
  TrustListOutOfDateAlarmType_EnabledState_FalseState((int) 19322L),
  TrustListOutOfDateAlarmType_Quality((int) 19323L),
  TrustListOutOfDateAlarmType_Quality_SourceTimestamp((int) 19324L),
  TrustListOutOfDateAlarmType_LastSeverity((int) 19325L),
  TrustListOutOfDateAlarmType_LastSeverity_SourceTimestamp((int) 19326L),
  TrustListOutOfDateAlarmType_Comment((int) 19327L),
  TrustListOutOfDateAlarmType_Comment_SourceTimestamp((int) 19328L),
  TrustListOutOfDateAlarmType_ClientUserId((int) 19329L),
  TrustListOutOfDateAlarmType_AddComment_InputArguments((int) 19333L),
  TrustListOutOfDateAlarmType_ConditionRefresh_InputArguments((int) 19335L),
  TrustListOutOfDateAlarmType_ConditionRefresh2_InputArguments((int) 19337L),
  TrustListOutOfDateAlarmType_AckedState((int) 19338L),
  TrustListOutOfDateAlarmType_AckedState_Id((int) 19339L),
  TrustListOutOfDateAlarmType_AckedState_Name((int) 19340L),
  TrustListOutOfDateAlarmType_AckedState_Number((int) 19341L),
  TrustListOutOfDateAlarmType_AckedState_EffectiveDisplayName((int) 19342L),
  TrustListOutOfDateAlarmType_AckedState_TransitionTime((int) 19343L),
  TrustListOutOfDateAlarmType_AckedState_EffectiveTransitionTime((int) 19344L),
  TrustListOutOfDateAlarmType_AckedState_TrueState((int) 19345L),
  TrustListOutOfDateAlarmType_AckedState_FalseState((int) 19346L),
  TrustListOutOfDateAlarmType_ConfirmedState((int) 19347L),
  TrustListOutOfDateAlarmType_ConfirmedState_Id((int) 19348L),
  TrustListOutOfDateAlarmType_ConfirmedState_Name((int) 19349L),
  TrustListOutOfDateAlarmType_ConfirmedState_Number((int) 19350L),
  TrustListOutOfDateAlarmType_ConfirmedState_EffectiveDisplayName((int) 19351L),
  TrustListOutOfDateAlarmType_ConfirmedState_TransitionTime((int) 19352L),
  TrustListOutOfDateAlarmType_ConfirmedState_EffectiveTransitionTime((int) 19353L),
  TrustListOutOfDateAlarmType_ConfirmedState_TrueState((int) 19354L),
  TrustListOutOfDateAlarmType_ConfirmedState_FalseState((int) 19355L),
  TrustListOutOfDateAlarmType_Acknowledge_InputArguments((int) 19357L),
  TrustListOutOfDateAlarmType_Confirm_InputArguments((int) 19359L),
  TrustListOutOfDateAlarmType_ActiveState((int) 19360L),
  TrustListOutOfDateAlarmType_ActiveState_Id((int) 19361L),
  TrustListOutOfDateAlarmType_ActiveState_Name((int) 19362L),
  TrustListOutOfDateAlarmType_ActiveState_Number((int) 19363L),
  TrustListOutOfDateAlarmType_ActiveState_EffectiveDisplayName((int) 19364L),
  TrustListOutOfDateAlarmType_ActiveState_TransitionTime((int) 19365L),
  TrustListOutOfDateAlarmType_ActiveState_EffectiveTransitionTime((int) 19366L),
  TrustListOutOfDateAlarmType_ActiveState_TrueState((int) 19367L),
  TrustListOutOfDateAlarmType_ActiveState_FalseState((int) 19368L),
  TrustListOutOfDateAlarmType_InputNode((int) 19369L),
  TrustListOutOfDateAlarmType_SuppressedState((int) 19370L),
  TrustListOutOfDateAlarmType_SuppressedState_Id((int) 19371L),
  TrustListOutOfDateAlarmType_SuppressedState_Name((int) 19372L),
  TrustListOutOfDateAlarmType_SuppressedState_Number((int) 19373L),
  TrustListOutOfDateAlarmType_SuppressedState_EffectiveDisplayName((int) 19374L),
  TrustListOutOfDateAlarmType_SuppressedState_TransitionTime((int) 19375L),
  TrustListOutOfDateAlarmType_SuppressedState_EffectiveTransitionTime((int) 19376L),
  TrustListOutOfDateAlarmType_SuppressedState_TrueState((int) 19377L),
  TrustListOutOfDateAlarmType_SuppressedState_FalseState((int) 19378L),
  TrustListOutOfDateAlarmType_OutOfServiceState((int) 19379L),
  TrustListOutOfDateAlarmType_OutOfServiceState_Id((int) 19380L),
  TrustListOutOfDateAlarmType_OutOfServiceState_Name((int) 19381L),
  TrustListOutOfDateAlarmType_OutOfServiceState_Number((int) 19382L),
  TrustListOutOfDateAlarmType_OutOfServiceState_EffectiveDisplayName((int) 19383L),
  TrustListOutOfDateAlarmType_OutOfServiceState_TransitionTime((int) 19384L),
  TrustListOutOfDateAlarmType_OutOfServiceState_EffectiveTransitionTime((int) 19385L),
  TrustListOutOfDateAlarmType_OutOfServiceState_TrueState((int) 19386L),
  TrustListOutOfDateAlarmType_OutOfServiceState_FalseState((int) 19387L),
  TrustListOutOfDateAlarmType_ShelvingState_CurrentState((int) 19389L),
  TrustListOutOfDateAlarmType_ShelvingState_CurrentState_Id((int) 19390L),
  TrustListOutOfDateAlarmType_ShelvingState_CurrentState_Name((int) 19391L),
  TrustListOutOfDateAlarmType_ShelvingState_CurrentState_Number((int) 19392L),
  TrustListOutOfDateAlarmType_ShelvingState_CurrentState_EffectiveDisplayName((int) 19393L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition((int) 19394L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition_Id((int) 19395L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition_Name((int) 19396L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition_Number((int) 19397L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition_TransitionTime((int) 19398L),
  TrustListOutOfDateAlarmType_ShelvingState_LastTransition_EffectiveTransitionTime((int) 19399L),
  TrustListOutOfDateAlarmType_ShelvingState_AvailableStates((int) 19400L),
  TrustListOutOfDateAlarmType_ShelvingState_AvailableTransitions((int) 19401L),
  TrustListOutOfDateAlarmType_ShelvingState_UnshelveTime((int) 19402L),
  TrustListOutOfDateAlarmType_ShelvingState_TimedShelve_InputArguments((int) 19404L),
  TrustListOutOfDateAlarmType_SuppressedOrShelved((int) 19407L),
  TrustListOutOfDateAlarmType_MaxTimeShelved((int) 19408L),
  TrustListOutOfDateAlarmType_AudibleEnabled((int) 19409L),
  TrustListOutOfDateAlarmType_AudibleSound((int) 19410L),
  TrustListOutOfDateAlarmType_AudibleSound_ListId((int) 19411L),
  TrustListOutOfDateAlarmType_AudibleSound_AgencyId((int) 19412L),
  TrustListOutOfDateAlarmType_AudibleSound_VersionId((int) 19413L),
  TrustListOutOfDateAlarmType_SilenceState((int) 19414L),
  TrustListOutOfDateAlarmType_SilenceState_Id((int) 19415L),
  TrustListOutOfDateAlarmType_SilenceState_Name((int) 19416L),
  TrustListOutOfDateAlarmType_SilenceState_Number((int) 19417L),
  TrustListOutOfDateAlarmType_SilenceState_EffectiveDisplayName((int) 19418L),
  TrustListOutOfDateAlarmType_SilenceState_TransitionTime((int) 19419L),
  TrustListOutOfDateAlarmType_SilenceState_EffectiveTransitionTime((int) 19420L),
  TrustListOutOfDateAlarmType_SilenceState_TrueState((int) 19421L),
  TrustListOutOfDateAlarmType_SilenceState_FalseState((int) 19422L),
  TrustListOutOfDateAlarmType_OnDelay((int) 19423L),
  TrustListOutOfDateAlarmType_OffDelay((int) 19424L),
  TrustListOutOfDateAlarmType_FirstInGroupFlag((int) 19425L),
  TrustListOutOfDateAlarmType_LatchedState((int) 19427L),
  TrustListOutOfDateAlarmType_LatchedState_Id((int) 19428L),
  TrustListOutOfDateAlarmType_LatchedState_Name((int) 19429L),
  TrustListOutOfDateAlarmType_LatchedState_Number((int) 19430L),
  TrustListOutOfDateAlarmType_LatchedState_EffectiveDisplayName((int) 19431L),
  TrustListOutOfDateAlarmType_LatchedState_TransitionTime((int) 19432L),
  TrustListOutOfDateAlarmType_LatchedState_EffectiveTransitionTime((int) 19433L),
  TrustListOutOfDateAlarmType_LatchedState_TrueState((int) 19434L),
  TrustListOutOfDateAlarmType_LatchedState_FalseState((int) 19435L),
  TrustListOutOfDateAlarmType_ReAlarmTime((int) 19437L),
  TrustListOutOfDateAlarmType_ReAlarmRepeatCount((int) 19438L),
  TrustListOutOfDateAlarmType_NormalState((int) 19445L),
  TrustListOutOfDateAlarmType_TrustListId((int) 19446L),
  TrustListOutOfDateAlarmType_LastUpdateTime((int) 19447L),
  TrustListOutOfDateAlarmType_UpdateFrequency((int) 19448L),
  TrustListType_DefaultValidationOptions((int) 23563L),
  TrustListValidationOptions_OptionSetValues((int) 23565L),
  TrustListType_MaxByteStringLength((int) 24250L),
  TrustListOutOfDateAlarmType_Suppress2_InputArguments((int) 24509L),
  TrustListOutOfDateAlarmType_Unsuppress2_InputArguments((int) 24511L),
  TrustListOutOfDateAlarmType_RemoveFromService2_InputArguments((int) 24513L),
  TrustListOutOfDateAlarmType_PlaceInService2_InputArguments((int) 24515L),
  TrustListOutOfDateAlarmType_Reset2_InputArguments((int) 24517L),
  TrustListOutOfDateAlarmType_ShelvingState_TimedShelve2_InputArguments((int) 24971L),
  TrustListOutOfDateAlarmType_ShelvingState_Unshelve2_InputArguments((int) 24973L),
  TrustListOutOfDateAlarmType_ShelvingState_OneShotShelve2_InputArguments((int) 24975L),
  TrustListOutOfDateAlarmType_GetGroupMemberships_OutputArguments((int) 25176L),
  TrustListType_LastModifiedTime((int) 25206L),
  TrustListUpdatedAuditEventType_ConditionClassId((int) 32027L),
  TrustListUpdatedAuditEventType_ConditionClassName((int) 32028L),
  TrustListUpdatedAuditEventType_ConditionSubClassId((int) 32029L),
  TrustListUpdatedAuditEventType_ConditionSubClassName((int) 32030L),
  TrustListOutOfDateAlarmType_SupportsFilteredRetain((int) 32253L),
  TrustListType_ActivityTimeout((int) 32254L),
  TrustListUpdateRequestedAuditEventType_EventId((int) 32261L),
  TrustListUpdateRequestedAuditEventType_EventType((int) 32262L),
  TrustListUpdateRequestedAuditEventType_SourceNode((int) 32263L),
  TrustListUpdateRequestedAuditEventType_SourceName((int) 32264L),
  TrustListUpdateRequestedAuditEventType_Time((int) 32265L),
  TrustListUpdateRequestedAuditEventType_ReceiveTime((int) 32266L),
  TrustListUpdateRequestedAuditEventType_LocalTime((int) 32267L),
  TrustListUpdateRequestedAuditEventType_Message((int) 32268L),
  TrustListUpdateRequestedAuditEventType_Severity((int) 32269L),
  TrustListUpdateRequestedAuditEventType_ConditionClassId((int) 32270L),
  TrustListUpdateRequestedAuditEventType_ConditionClassName((int) 32271L),
  TrustListUpdateRequestedAuditEventType_ConditionSubClassId((int) 32272L),
  TrustListUpdateRequestedAuditEventType_ConditionSubClassName((int) 32273L),
  TrustListUpdateRequestedAuditEventType_ActionTimeStamp((int) 32274L),
  TrustListUpdateRequestedAuditEventType_Status((int) 32275L),
  TrustListUpdateRequestedAuditEventType_ServerId((int) 32276L),
  TrustListUpdateRequestedAuditEventType_ClientAuditEntryId((int) 32277L),
  TrustListUpdateRequestedAuditEventType_ClientUserId((int) 32278L),
  TrustListUpdateRequestedAuditEventType_MethodId((int) 32279L),
  TrustListUpdateRequestedAuditEventType_InputArguments((int) 32280L),
  TrustListUpdatedAuditEventType_TrustListId((int) 32281L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableTrust> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableTrust value : OpcuaNodeIdServicesVariableTrust.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableTrust(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableTrust enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
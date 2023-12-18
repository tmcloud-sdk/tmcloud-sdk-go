package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ScalingPolicyExecuteLogList struct {
	Status *ScalingPolicyExecuteLogListStatus `json:"status,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	ExecuteType *ScalingPolicyExecuteLogListExecuteType `json:"execute_type,omitempty"`

	ExecuteTime *string `json:"execute_time,omitempty"`

	Id *string `json:"id,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`

	ScalingPolicyId *string `json:"scaling_policy_id,omitempty"`

	ScalingResourceType *ScalingPolicyExecuteLogListScalingResourceType `json:"scaling_resource_type,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	OldValue *string `json:"old_value,omitempty"`

	DesireValue *string `json:"desire_value,omitempty"`

	LimitValue *string `json:"limit_value,omitempty"`

	Type *ScalingPolicyExecuteLogListType `json:"type,omitempty"`

	JobRecords *[]JobRecords `json:"job_records,omitempty"`

	MetaData *EipMetaData `json:"meta_data,omitempty"`
}

func (o ScalingPolicyExecuteLogList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyExecuteLogList struct{}"
	}

	return strings.Join([]string{"ScalingPolicyExecuteLogList", string(data)}, " ")
}

type ScalingPolicyExecuteLogListStatus struct {
	value string
}

type ScalingPolicyExecuteLogListStatusEnum struct {
	SUCCESS   ScalingPolicyExecuteLogListStatus
	FAIL      ScalingPolicyExecuteLogListStatus
	EXECUTING ScalingPolicyExecuteLogListStatus
}

func GetScalingPolicyExecuteLogListStatusEnum() ScalingPolicyExecuteLogListStatusEnum {
	return ScalingPolicyExecuteLogListStatusEnum{
		SUCCESS: ScalingPolicyExecuteLogListStatus{
			value: "SUCCESS",
		},
		FAIL: ScalingPolicyExecuteLogListStatus{
			value: "FAIL",
		},
		EXECUTING: ScalingPolicyExecuteLogListStatus{
			value: "EXECUTING",
		},
	}
}

func (c ScalingPolicyExecuteLogListStatus) Value() string {
	return c.value
}

func (c ScalingPolicyExecuteLogListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyExecuteLogListStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScalingPolicyExecuteLogListExecuteType struct {
	value string
}

type ScalingPolicyExecuteLogListExecuteTypeEnum struct {
	SCHEDULE   ScalingPolicyExecuteLogListExecuteType
	RECURRENCE ScalingPolicyExecuteLogListExecuteType
	ALARM      ScalingPolicyExecuteLogListExecuteType
	MANUAL     ScalingPolicyExecuteLogListExecuteType
}

func GetScalingPolicyExecuteLogListExecuteTypeEnum() ScalingPolicyExecuteLogListExecuteTypeEnum {
	return ScalingPolicyExecuteLogListExecuteTypeEnum{
		SCHEDULE: ScalingPolicyExecuteLogListExecuteType{
			value: "SCHEDULE",
		},
		RECURRENCE: ScalingPolicyExecuteLogListExecuteType{
			value: "RECURRENCE",
		},
		ALARM: ScalingPolicyExecuteLogListExecuteType{
			value: "ALARM",
		},
		MANUAL: ScalingPolicyExecuteLogListExecuteType{
			value: "MANUAL",
		},
	}
}

func (c ScalingPolicyExecuteLogListExecuteType) Value() string {
	return c.value
}

func (c ScalingPolicyExecuteLogListExecuteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyExecuteLogListExecuteType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScalingPolicyExecuteLogListScalingResourceType struct {
	value string
}

type ScalingPolicyExecuteLogListScalingResourceTypeEnum struct {
	SCALING_GROUP ScalingPolicyExecuteLogListScalingResourceType
	BANDWIDTH     ScalingPolicyExecuteLogListScalingResourceType
}

func GetScalingPolicyExecuteLogListScalingResourceTypeEnum() ScalingPolicyExecuteLogListScalingResourceTypeEnum {
	return ScalingPolicyExecuteLogListScalingResourceTypeEnum{
		SCALING_GROUP: ScalingPolicyExecuteLogListScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ScalingPolicyExecuteLogListScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ScalingPolicyExecuteLogListScalingResourceType) Value() string {
	return c.value
}

func (c ScalingPolicyExecuteLogListScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyExecuteLogListScalingResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScalingPolicyExecuteLogListType struct {
	value string
}

type ScalingPolicyExecuteLogListTypeEnum struct {
	ADD    ScalingPolicyExecuteLogListType
	REMOVE ScalingPolicyExecuteLogListType
	SET    ScalingPolicyExecuteLogListType
}

func GetScalingPolicyExecuteLogListTypeEnum() ScalingPolicyExecuteLogListTypeEnum {
	return ScalingPolicyExecuteLogListTypeEnum{
		ADD: ScalingPolicyExecuteLogListType{
			value: "ADD",
		},
		REMOVE: ScalingPolicyExecuteLogListType{
			value: "REMOVE",
		},
		SET: ScalingPolicyExecuteLogListType{
			value: "SET",
		},
	}
}

func (c ScalingPolicyExecuteLogListType) Value() string {
	return c.value
}

func (c ScalingPolicyExecuteLogListType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingPolicyExecuteLogListType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

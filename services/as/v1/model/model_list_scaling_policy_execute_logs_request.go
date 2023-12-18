package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListScalingPolicyExecuteLogsRequest struct {
	ScalingPolicyId string `json:"scaling_policy_id"`

	LogId *string `json:"log_id,omitempty"`

	ScalingResourceType *ListScalingPolicyExecuteLogsRequestScalingResourceType `json:"scaling_resource_type,omitempty"`

	ScalingResourceId *string `json:"scaling_resource_id,omitempty"`

	ExecuteType *ListScalingPolicyExecuteLogsRequestExecuteType `json:"execute_type,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`

	StartNumber *int32 `json:"start_number,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScalingPolicyExecuteLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingPolicyExecuteLogsRequest struct{}"
	}

	return strings.Join([]string{"ListScalingPolicyExecuteLogsRequest", string(data)}, " ")
}

type ListScalingPolicyExecuteLogsRequestScalingResourceType struct {
	value string
}

type ListScalingPolicyExecuteLogsRequestScalingResourceTypeEnum struct {
	SCALING_GROUP ListScalingPolicyExecuteLogsRequestScalingResourceType
	BANDWIDTH     ListScalingPolicyExecuteLogsRequestScalingResourceType
}

func GetListScalingPolicyExecuteLogsRequestScalingResourceTypeEnum() ListScalingPolicyExecuteLogsRequestScalingResourceTypeEnum {
	return ListScalingPolicyExecuteLogsRequestScalingResourceTypeEnum{
		SCALING_GROUP: ListScalingPolicyExecuteLogsRequestScalingResourceType{
			value: "SCALING_GROUP",
		},
		BANDWIDTH: ListScalingPolicyExecuteLogsRequestScalingResourceType{
			value: "BANDWIDTH",
		},
	}
}

func (c ListScalingPolicyExecuteLogsRequestScalingResourceType) Value() string {
	return c.value
}

func (c ListScalingPolicyExecuteLogsRequestScalingResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingPolicyExecuteLogsRequestScalingResourceType) UnmarshalJSON(b []byte) error {
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

type ListScalingPolicyExecuteLogsRequestExecuteType struct {
	value string
}

type ListScalingPolicyExecuteLogsRequestExecuteTypeEnum struct {
	SCHEDULED  ListScalingPolicyExecuteLogsRequestExecuteType
	RECURRENCE ListScalingPolicyExecuteLogsRequestExecuteType
	ALARM      ListScalingPolicyExecuteLogsRequestExecuteType
	MANUAL     ListScalingPolicyExecuteLogsRequestExecuteType
}

func GetListScalingPolicyExecuteLogsRequestExecuteTypeEnum() ListScalingPolicyExecuteLogsRequestExecuteTypeEnum {
	return ListScalingPolicyExecuteLogsRequestExecuteTypeEnum{
		SCHEDULED: ListScalingPolicyExecuteLogsRequestExecuteType{
			value: "SCHEDULED",
		},
		RECURRENCE: ListScalingPolicyExecuteLogsRequestExecuteType{
			value: "RECURRENCE",
		},
		ALARM: ListScalingPolicyExecuteLogsRequestExecuteType{
			value: "ALARM",
		},
		MANUAL: ListScalingPolicyExecuteLogsRequestExecuteType{
			value: "MANUAL",
		},
	}
}

func (c ListScalingPolicyExecuteLogsRequestExecuteType) Value() string {
	return c.value
}

func (c ListScalingPolicyExecuteLogsRequestExecuteType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScalingPolicyExecuteLogsRequestExecuteType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ScheduledPolicy struct {
	LaunchTime string `json:"launch_time"`

	RecurrenceType *ScheduledPolicyRecurrenceType `json:"recurrence_type,omitempty"`

	RecurrenceValue *string `json:"recurrence_value,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	EndTime *string `json:"end_time,omitempty"`
}

func (o ScheduledPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledPolicy struct{}"
	}

	return strings.Join([]string{"ScheduledPolicy", string(data)}, " ")
}

type ScheduledPolicyRecurrenceType struct {
	value string
}

type ScheduledPolicyRecurrenceTypeEnum struct {
	DAILY   ScheduledPolicyRecurrenceType
	WEEKLY  ScheduledPolicyRecurrenceType
	MONTHLY ScheduledPolicyRecurrenceType
}

func GetScheduledPolicyRecurrenceTypeEnum() ScheduledPolicyRecurrenceTypeEnum {
	return ScheduledPolicyRecurrenceTypeEnum{
		DAILY: ScheduledPolicyRecurrenceType{
			value: "Daily",
		},
		WEEKLY: ScheduledPolicyRecurrenceType{
			value: "Weekly",
		},
		MONTHLY: ScheduledPolicyRecurrenceType{
			value: "Monthly",
		},
	}
}

func (c ScheduledPolicyRecurrenceType) Value() string {
	return c.value
}

func (c ScheduledPolicyRecurrenceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledPolicyRecurrenceType) UnmarshalJSON(b []byte) error {
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

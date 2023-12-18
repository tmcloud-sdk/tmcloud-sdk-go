package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchResumeScalingPoliciesOption struct {
	ScalingPolicyId []string `json:"scaling_policy_id"`

	ForceDelete *BatchResumeScalingPoliciesOptionForceDelete `json:"force_delete,omitempty"`

	Action BatchResumeScalingPoliciesOptionAction `json:"action"`

	DeleteAlarm *string `json:"delete_alarm,omitempty"`
}

func (o BatchResumeScalingPoliciesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResumeScalingPoliciesOption struct{}"
	}

	return strings.Join([]string{"BatchResumeScalingPoliciesOption", string(data)}, " ")
}

type BatchResumeScalingPoliciesOptionForceDelete struct {
	value string
}

type BatchResumeScalingPoliciesOptionForceDeleteEnum struct {
	NO  BatchResumeScalingPoliciesOptionForceDelete
	YES BatchResumeScalingPoliciesOptionForceDelete
}

func GetBatchResumeScalingPoliciesOptionForceDeleteEnum() BatchResumeScalingPoliciesOptionForceDeleteEnum {
	return BatchResumeScalingPoliciesOptionForceDeleteEnum{
		NO: BatchResumeScalingPoliciesOptionForceDelete{
			value: "no",
		},
		YES: BatchResumeScalingPoliciesOptionForceDelete{
			value: "yes",
		},
	}
}

func (c BatchResumeScalingPoliciesOptionForceDelete) Value() string {
	return c.value
}

func (c BatchResumeScalingPoliciesOptionForceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchResumeScalingPoliciesOptionForceDelete) UnmarshalJSON(b []byte) error {
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

type BatchResumeScalingPoliciesOptionAction struct {
	value string
}

type BatchResumeScalingPoliciesOptionActionEnum struct {
	RESUME BatchResumeScalingPoliciesOptionAction
}

func GetBatchResumeScalingPoliciesOptionActionEnum() BatchResumeScalingPoliciesOptionActionEnum {
	return BatchResumeScalingPoliciesOptionActionEnum{
		RESUME: BatchResumeScalingPoliciesOptionAction{
			value: "resume",
		},
	}
}

func (c BatchResumeScalingPoliciesOptionAction) Value() string {
	return c.value
}

func (c BatchResumeScalingPoliciesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchResumeScalingPoliciesOptionAction) UnmarshalJSON(b []byte) error {
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

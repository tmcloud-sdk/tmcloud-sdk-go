package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchDeleteScalingPoliciesOption struct {
	ScalingPolicyId []string `json:"scaling_policy_id"`

	ForceDelete *BatchDeleteScalingPoliciesOptionForceDelete `json:"force_delete,omitempty"`

	Action BatchDeleteScalingPoliciesOptionAction `json:"action"`

	DeleteAlarm *string `json:"delete_alarm,omitempty"`
}

func (o BatchDeleteScalingPoliciesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteScalingPoliciesOption struct{}"
	}

	return strings.Join([]string{"BatchDeleteScalingPoliciesOption", string(data)}, " ")
}

type BatchDeleteScalingPoliciesOptionForceDelete struct {
	value string
}

type BatchDeleteScalingPoliciesOptionForceDeleteEnum struct {
	NO  BatchDeleteScalingPoliciesOptionForceDelete
	YES BatchDeleteScalingPoliciesOptionForceDelete
}

func GetBatchDeleteScalingPoliciesOptionForceDeleteEnum() BatchDeleteScalingPoliciesOptionForceDeleteEnum {
	return BatchDeleteScalingPoliciesOptionForceDeleteEnum{
		NO: BatchDeleteScalingPoliciesOptionForceDelete{
			value: "no",
		},
		YES: BatchDeleteScalingPoliciesOptionForceDelete{
			value: "yes",
		},
	}
}

func (c BatchDeleteScalingPoliciesOptionForceDelete) Value() string {
	return c.value
}

func (c BatchDeleteScalingPoliciesOptionForceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteScalingPoliciesOptionForceDelete) UnmarshalJSON(b []byte) error {
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

type BatchDeleteScalingPoliciesOptionAction struct {
	value string
}

type BatchDeleteScalingPoliciesOptionActionEnum struct {
	DELETE BatchDeleteScalingPoliciesOptionAction
}

func GetBatchDeleteScalingPoliciesOptionActionEnum() BatchDeleteScalingPoliciesOptionActionEnum {
	return BatchDeleteScalingPoliciesOptionActionEnum{
		DELETE: BatchDeleteScalingPoliciesOptionAction{
			value: "delete",
		},
	}
}

func (c BatchDeleteScalingPoliciesOptionAction) Value() string {
	return c.value
}

func (c BatchDeleteScalingPoliciesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteScalingPoliciesOptionAction) UnmarshalJSON(b []byte) error {
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

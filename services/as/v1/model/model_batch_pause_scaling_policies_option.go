package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchPauseScalingPoliciesOption struct {
	ScalingPolicyId []string `json:"scaling_policy_id"`

	ForceDelete *BatchPauseScalingPoliciesOptionForceDelete `json:"force_delete,omitempty"`

	Action BatchPauseScalingPoliciesOptionAction `json:"action"`

	DeleteAlarm *string `json:"delete_alarm,omitempty"`
}

func (o BatchPauseScalingPoliciesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseScalingPoliciesOption struct{}"
	}

	return strings.Join([]string{"BatchPauseScalingPoliciesOption", string(data)}, " ")
}

type BatchPauseScalingPoliciesOptionForceDelete struct {
	value string
}

type BatchPauseScalingPoliciesOptionForceDeleteEnum struct {
	NO  BatchPauseScalingPoliciesOptionForceDelete
	YES BatchPauseScalingPoliciesOptionForceDelete
}

func GetBatchPauseScalingPoliciesOptionForceDeleteEnum() BatchPauseScalingPoliciesOptionForceDeleteEnum {
	return BatchPauseScalingPoliciesOptionForceDeleteEnum{
		NO: BatchPauseScalingPoliciesOptionForceDelete{
			value: "no",
		},
		YES: BatchPauseScalingPoliciesOptionForceDelete{
			value: "yes",
		},
	}
}

func (c BatchPauseScalingPoliciesOptionForceDelete) Value() string {
	return c.value
}

func (c BatchPauseScalingPoliciesOptionForceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchPauseScalingPoliciesOptionForceDelete) UnmarshalJSON(b []byte) error {
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

type BatchPauseScalingPoliciesOptionAction struct {
	value string
}

type BatchPauseScalingPoliciesOptionActionEnum struct {
	PAUSE BatchPauseScalingPoliciesOptionAction
}

func GetBatchPauseScalingPoliciesOptionActionEnum() BatchPauseScalingPoliciesOptionActionEnum {
	return BatchPauseScalingPoliciesOptionActionEnum{
		PAUSE: BatchPauseScalingPoliciesOptionAction{
			value: "pause",
		},
	}
}

func (c BatchPauseScalingPoliciesOptionAction) Value() string {
	return c.value
}

func (c BatchPauseScalingPoliciesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchPauseScalingPoliciesOptionAction) UnmarshalJSON(b []byte) error {
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

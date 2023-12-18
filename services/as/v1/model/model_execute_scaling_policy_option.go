package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ExecuteScalingPolicyOption struct {
	Action ExecuteScalingPolicyOptionAction `json:"action"`
}

func (o ExecuteScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"ExecuteScalingPolicyOption", string(data)}, " ")
}

type ExecuteScalingPolicyOptionAction struct {
	value string
}

type ExecuteScalingPolicyOptionActionEnum struct {
	EXECUTE ExecuteScalingPolicyOptionAction
}

func GetExecuteScalingPolicyOptionActionEnum() ExecuteScalingPolicyOptionActionEnum {
	return ExecuteScalingPolicyOptionActionEnum{
		EXECUTE: ExecuteScalingPolicyOptionAction{
			value: "execute",
		},
	}
}

func (c ExecuteScalingPolicyOptionAction) Value() string {
	return c.value
}

func (c ExecuteScalingPolicyOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteScalingPolicyOptionAction) UnmarshalJSON(b []byte) error {
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

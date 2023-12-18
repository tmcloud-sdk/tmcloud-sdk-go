package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ResumeScalingPolicyOption struct {
	Action ResumeScalingPolicyOptionAction `json:"action"`
}

func (o ResumeScalingPolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingPolicyOption struct{}"
	}

	return strings.Join([]string{"ResumeScalingPolicyOption", string(data)}, " ")
}

type ResumeScalingPolicyOptionAction struct {
	value string
}

type ResumeScalingPolicyOptionActionEnum struct {
	RESUME ResumeScalingPolicyOptionAction
}

func GetResumeScalingPolicyOptionActionEnum() ResumeScalingPolicyOptionActionEnum {
	return ResumeScalingPolicyOptionActionEnum{
		RESUME: ResumeScalingPolicyOptionAction{
			value: "resume",
		},
	}
}

func (c ResumeScalingPolicyOptionAction) Value() string {
	return c.value
}

func (c ResumeScalingPolicyOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResumeScalingPolicyOptionAction) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ResumeScalingGroupOption struct {
	Action ResumeScalingGroupOptionAction `json:"action"`
}

func (o ResumeScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeScalingGroupOption struct{}"
	}

	return strings.Join([]string{"ResumeScalingGroupOption", string(data)}, " ")
}

type ResumeScalingGroupOptionAction struct {
	value string
}

type ResumeScalingGroupOptionActionEnum struct {
	RESUME ResumeScalingGroupOptionAction
}

func GetResumeScalingGroupOptionActionEnum() ResumeScalingGroupOptionActionEnum {
	return ResumeScalingGroupOptionActionEnum{
		RESUME: ResumeScalingGroupOptionAction{
			value: "resume",
		},
	}
}

func (c ResumeScalingGroupOptionAction) Value() string {
	return c.value
}

func (c ResumeScalingGroupOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResumeScalingGroupOptionAction) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PauseScalingGroupOption struct {
	Action PauseScalingGroupOptionAction `json:"action"`
}

func (o PauseScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseScalingGroupOption struct{}"
	}

	return strings.Join([]string{"PauseScalingGroupOption", string(data)}, " ")
}

type PauseScalingGroupOptionAction struct {
	value string
}

type PauseScalingGroupOptionActionEnum struct {
	PAUSE PauseScalingGroupOptionAction
}

func GetPauseScalingGroupOptionActionEnum() PauseScalingGroupOptionActionEnum {
	return PauseScalingGroupOptionActionEnum{
		PAUSE: PauseScalingGroupOptionAction{
			value: "pause",
		},
	}
}

func (c PauseScalingGroupOptionAction) Value() string {
	return c.value
}

func (c PauseScalingGroupOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PauseScalingGroupOptionAction) UnmarshalJSON(b []byte) error {
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

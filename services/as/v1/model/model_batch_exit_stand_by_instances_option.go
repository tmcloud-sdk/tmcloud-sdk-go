package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchExitStandByInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchExitStandByInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchExitStandByInstancesOptionAction `json:"action"`

	InstanceAppend *BatchExitStandByInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchExitStandByInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExitStandByInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchExitStandByInstancesOption", string(data)}, " ")
}

type BatchExitStandByInstancesOptionInstanceDelete struct {
	value string
}

type BatchExitStandByInstancesOptionInstanceDeleteEnum struct {
	YES BatchExitStandByInstancesOptionInstanceDelete
	NO  BatchExitStandByInstancesOptionInstanceDelete
}

func GetBatchExitStandByInstancesOptionInstanceDeleteEnum() BatchExitStandByInstancesOptionInstanceDeleteEnum {
	return BatchExitStandByInstancesOptionInstanceDeleteEnum{
		YES: BatchExitStandByInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchExitStandByInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchExitStandByInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchExitStandByInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchExitStandByInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchExitStandByInstancesOptionAction struct {
	value string
}

type BatchExitStandByInstancesOptionActionEnum struct {
	EXIT_STANDBY BatchExitStandByInstancesOptionAction
}

func GetBatchExitStandByInstancesOptionActionEnum() BatchExitStandByInstancesOptionActionEnum {
	return BatchExitStandByInstancesOptionActionEnum{
		EXIT_STANDBY: BatchExitStandByInstancesOptionAction{
			value: "EXIT_STANDBY",
		},
	}
}

func (c BatchExitStandByInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchExitStandByInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchExitStandByInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchExitStandByInstancesOptionInstanceAppend struct {
	value string
}

type BatchExitStandByInstancesOptionInstanceAppendEnum struct {
	NO  BatchExitStandByInstancesOptionInstanceAppend
	YES BatchExitStandByInstancesOptionInstanceAppend
}

func GetBatchExitStandByInstancesOptionInstanceAppendEnum() BatchExitStandByInstancesOptionInstanceAppendEnum {
	return BatchExitStandByInstancesOptionInstanceAppendEnum{
		NO: BatchExitStandByInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchExitStandByInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchExitStandByInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchExitStandByInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchExitStandByInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

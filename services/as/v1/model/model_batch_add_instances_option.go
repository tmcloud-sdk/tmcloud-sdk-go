package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchAddInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchAddInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchAddInstancesOptionAction `json:"action"`

	InstanceAppend *BatchAddInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchAddInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchAddInstancesOption", string(data)}, " ")
}

type BatchAddInstancesOptionInstanceDelete struct {
	value string
}

type BatchAddInstancesOptionInstanceDeleteEnum struct {
	YES BatchAddInstancesOptionInstanceDelete
	NO  BatchAddInstancesOptionInstanceDelete
}

func GetBatchAddInstancesOptionInstanceDeleteEnum() BatchAddInstancesOptionInstanceDeleteEnum {
	return BatchAddInstancesOptionInstanceDeleteEnum{
		YES: BatchAddInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchAddInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchAddInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchAddInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchAddInstancesOptionAction struct {
	value string
}

type BatchAddInstancesOptionActionEnum struct {
	ADD BatchAddInstancesOptionAction
}

func GetBatchAddInstancesOptionActionEnum() BatchAddInstancesOptionActionEnum {
	return BatchAddInstancesOptionActionEnum{
		ADD: BatchAddInstancesOptionAction{
			value: "ADD",
		},
	}
}

func (c BatchAddInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchAddInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchAddInstancesOptionInstanceAppend struct {
	value string
}

type BatchAddInstancesOptionInstanceAppendEnum struct {
	NO  BatchAddInstancesOptionInstanceAppend
	YES BatchAddInstancesOptionInstanceAppend
}

func GetBatchAddInstancesOptionInstanceAppendEnum() BatchAddInstancesOptionInstanceAppendEnum {
	return BatchAddInstancesOptionInstanceAppendEnum{
		NO: BatchAddInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchAddInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchAddInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchAddInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchAddInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

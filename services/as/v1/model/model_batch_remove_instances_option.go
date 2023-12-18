package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchRemoveInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchRemoveInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchRemoveInstancesOptionAction `json:"action"`

	InstanceAppend *BatchRemoveInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchRemoveInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchRemoveInstancesOption", string(data)}, " ")
}

type BatchRemoveInstancesOptionInstanceDelete struct {
	value string
}

type BatchRemoveInstancesOptionInstanceDeleteEnum struct {
	YES BatchRemoveInstancesOptionInstanceDelete
	NO  BatchRemoveInstancesOptionInstanceDelete
}

func GetBatchRemoveInstancesOptionInstanceDeleteEnum() BatchRemoveInstancesOptionInstanceDeleteEnum {
	return BatchRemoveInstancesOptionInstanceDeleteEnum{
		YES: BatchRemoveInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchRemoveInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchRemoveInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchRemoveInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRemoveInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchRemoveInstancesOptionAction struct {
	value string
}

type BatchRemoveInstancesOptionActionEnum struct {
	REMOVE BatchRemoveInstancesOptionAction
}

func GetBatchRemoveInstancesOptionActionEnum() BatchRemoveInstancesOptionActionEnum {
	return BatchRemoveInstancesOptionActionEnum{
		REMOVE: BatchRemoveInstancesOptionAction{
			value: "REMOVE",
		},
	}
}

func (c BatchRemoveInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchRemoveInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRemoveInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchRemoveInstancesOptionInstanceAppend struct {
	value string
}

type BatchRemoveInstancesOptionInstanceAppendEnum struct {
	NO  BatchRemoveInstancesOptionInstanceAppend
	YES BatchRemoveInstancesOptionInstanceAppend
}

func GetBatchRemoveInstancesOptionInstanceAppendEnum() BatchRemoveInstancesOptionInstanceAppendEnum {
	return BatchRemoveInstancesOptionInstanceAppendEnum{
		NO: BatchRemoveInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchRemoveInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchRemoveInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchRemoveInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchRemoveInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

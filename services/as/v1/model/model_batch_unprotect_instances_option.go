package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchUnprotectInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchUnprotectInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchUnprotectInstancesOptionAction `json:"action"`

	InstanceAppend *BatchUnprotectInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchUnprotectInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnprotectInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchUnprotectInstancesOption", string(data)}, " ")
}

type BatchUnprotectInstancesOptionInstanceDelete struct {
	value string
}

type BatchUnprotectInstancesOptionInstanceDeleteEnum struct {
	YES BatchUnprotectInstancesOptionInstanceDelete
	NO  BatchUnprotectInstancesOptionInstanceDelete
}

func GetBatchUnprotectInstancesOptionInstanceDeleteEnum() BatchUnprotectInstancesOptionInstanceDeleteEnum {
	return BatchUnprotectInstancesOptionInstanceDeleteEnum{
		YES: BatchUnprotectInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchUnprotectInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchUnprotectInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchUnprotectInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUnprotectInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchUnprotectInstancesOptionAction struct {
	value string
}

type BatchUnprotectInstancesOptionActionEnum struct {
	UNPROTECT BatchUnprotectInstancesOptionAction
}

func GetBatchUnprotectInstancesOptionActionEnum() BatchUnprotectInstancesOptionActionEnum {
	return BatchUnprotectInstancesOptionActionEnum{
		UNPROTECT: BatchUnprotectInstancesOptionAction{
			value: "UNPROTECT",
		},
	}
}

func (c BatchUnprotectInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchUnprotectInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUnprotectInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchUnprotectInstancesOptionInstanceAppend struct {
	value string
}

type BatchUnprotectInstancesOptionInstanceAppendEnum struct {
	NO  BatchUnprotectInstancesOptionInstanceAppend
	YES BatchUnprotectInstancesOptionInstanceAppend
}

func GetBatchUnprotectInstancesOptionInstanceAppendEnum() BatchUnprotectInstancesOptionInstanceAppendEnum {
	return BatchUnprotectInstancesOptionInstanceAppendEnum{
		NO: BatchUnprotectInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchUnprotectInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchUnprotectInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchUnprotectInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUnprotectInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

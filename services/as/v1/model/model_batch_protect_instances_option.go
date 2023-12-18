package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchProtectInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchProtectInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchProtectInstancesOptionAction `json:"action"`

	InstanceAppend *BatchProtectInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchProtectInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchProtectInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchProtectInstancesOption", string(data)}, " ")
}

type BatchProtectInstancesOptionInstanceDelete struct {
	value string
}

type BatchProtectInstancesOptionInstanceDeleteEnum struct {
	YES BatchProtectInstancesOptionInstanceDelete
	NO  BatchProtectInstancesOptionInstanceDelete
}

func GetBatchProtectInstancesOptionInstanceDeleteEnum() BatchProtectInstancesOptionInstanceDeleteEnum {
	return BatchProtectInstancesOptionInstanceDeleteEnum{
		YES: BatchProtectInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchProtectInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchProtectInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchProtectInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchProtectInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchProtectInstancesOptionAction struct {
	value string
}

type BatchProtectInstancesOptionActionEnum struct {
	PROTECT BatchProtectInstancesOptionAction
}

func GetBatchProtectInstancesOptionActionEnum() BatchProtectInstancesOptionActionEnum {
	return BatchProtectInstancesOptionActionEnum{
		PROTECT: BatchProtectInstancesOptionAction{
			value: "PROTECT",
		},
	}
}

func (c BatchProtectInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchProtectInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchProtectInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchProtectInstancesOptionInstanceAppend struct {
	value string
}

type BatchProtectInstancesOptionInstanceAppendEnum struct {
	NO  BatchProtectInstancesOptionInstanceAppend
	YES BatchProtectInstancesOptionInstanceAppend
}

func GetBatchProtectInstancesOptionInstanceAppendEnum() BatchProtectInstancesOptionInstanceAppendEnum {
	return BatchProtectInstancesOptionInstanceAppendEnum{
		NO: BatchProtectInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchProtectInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchProtectInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchProtectInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchProtectInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

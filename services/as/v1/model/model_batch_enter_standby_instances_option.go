package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchEnterStandbyInstancesOption struct {
	InstancesId []string `json:"instances_id"`

	InstanceDelete *BatchEnterStandbyInstancesOptionInstanceDelete `json:"instance_delete,omitempty"`

	Action BatchEnterStandbyInstancesOptionAction `json:"action"`

	InstanceAppend *BatchEnterStandbyInstancesOptionInstanceAppend `json:"instance_append,omitempty"`
}

func (o BatchEnterStandbyInstancesOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchEnterStandbyInstancesOption struct{}"
	}

	return strings.Join([]string{"BatchEnterStandbyInstancesOption", string(data)}, " ")
}

type BatchEnterStandbyInstancesOptionInstanceDelete struct {
	value string
}

type BatchEnterStandbyInstancesOptionInstanceDeleteEnum struct {
	YES BatchEnterStandbyInstancesOptionInstanceDelete
	NO  BatchEnterStandbyInstancesOptionInstanceDelete
}

func GetBatchEnterStandbyInstancesOptionInstanceDeleteEnum() BatchEnterStandbyInstancesOptionInstanceDeleteEnum {
	return BatchEnterStandbyInstancesOptionInstanceDeleteEnum{
		YES: BatchEnterStandbyInstancesOptionInstanceDelete{
			value: "yes",
		},
		NO: BatchEnterStandbyInstancesOptionInstanceDelete{
			value: "no",
		},
	}
}

func (c BatchEnterStandbyInstancesOptionInstanceDelete) Value() string {
	return c.value
}

func (c BatchEnterStandbyInstancesOptionInstanceDelete) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchEnterStandbyInstancesOptionInstanceDelete) UnmarshalJSON(b []byte) error {
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

type BatchEnterStandbyInstancesOptionAction struct {
	value string
}

type BatchEnterStandbyInstancesOptionActionEnum struct {
	ENTER_STANDBY BatchEnterStandbyInstancesOptionAction
}

func GetBatchEnterStandbyInstancesOptionActionEnum() BatchEnterStandbyInstancesOptionActionEnum {
	return BatchEnterStandbyInstancesOptionActionEnum{
		ENTER_STANDBY: BatchEnterStandbyInstancesOptionAction{
			value: "ENTER_STANDBY",
		},
	}
}

func (c BatchEnterStandbyInstancesOptionAction) Value() string {
	return c.value
}

func (c BatchEnterStandbyInstancesOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchEnterStandbyInstancesOptionAction) UnmarshalJSON(b []byte) error {
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

type BatchEnterStandbyInstancesOptionInstanceAppend struct {
	value string
}

type BatchEnterStandbyInstancesOptionInstanceAppendEnum struct {
	NO  BatchEnterStandbyInstancesOptionInstanceAppend
	YES BatchEnterStandbyInstancesOptionInstanceAppend
}

func GetBatchEnterStandbyInstancesOptionInstanceAppendEnum() BatchEnterStandbyInstancesOptionInstanceAppendEnum {
	return BatchEnterStandbyInstancesOptionInstanceAppendEnum{
		NO: BatchEnterStandbyInstancesOptionInstanceAppend{
			value: "no",
		},
		YES: BatchEnterStandbyInstancesOptionInstanceAppend{
			value: "yes",
		},
	}
}

func (c BatchEnterStandbyInstancesOptionInstanceAppend) Value() string {
	return c.value
}

func (c BatchEnterStandbyInstancesOptionInstanceAppend) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchEnterStandbyInstancesOptionInstanceAppend) UnmarshalJSON(b []byte) error {
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

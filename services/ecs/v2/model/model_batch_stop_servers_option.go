package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchStopServersOption struct {
	Servers []ServerId `json:"servers"`

	Type *BatchStopServersOptionType `json:"type,omitempty"`
}

func (o BatchStopServersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopServersOption struct{}"
	}

	return strings.Join([]string{"BatchStopServersOption", string(data)}, " ")
}

type BatchStopServersOptionType struct {
	value string
}

type BatchStopServersOptionTypeEnum struct {
	SOFT BatchStopServersOptionType
	HARD BatchStopServersOptionType
}

func GetBatchStopServersOptionTypeEnum() BatchStopServersOptionTypeEnum {
	return BatchStopServersOptionTypeEnum{
		SOFT: BatchStopServersOptionType{
			value: "SOFT",
		},
		HARD: BatchStopServersOptionType{
			value: "HARD",
		},
	}
}

func (c BatchStopServersOptionType) Value() string {
	return c.value
}

func (c BatchStopServersOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchStopServersOptionType) UnmarshalJSON(b []byte) error {
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

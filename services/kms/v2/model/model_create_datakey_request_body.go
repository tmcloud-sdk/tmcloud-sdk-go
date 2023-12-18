package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateDatakeyRequestBody struct {
	KeyId string `json:"key_id"`

	KeySpec *CreateDatakeyRequestBodyKeySpec `json:"key_spec,omitempty"`

	DatakeyLength *string `json:"datakey_length,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o CreateDatakeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDatakeyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDatakeyRequestBody", string(data)}, " ")
}

type CreateDatakeyRequestBodyKeySpec struct {
	value string
}

type CreateDatakeyRequestBodyKeySpecEnum struct {
	AES_256 CreateDatakeyRequestBodyKeySpec
	AES_128 CreateDatakeyRequestBodyKeySpec
}

func GetCreateDatakeyRequestBodyKeySpecEnum() CreateDatakeyRequestBodyKeySpecEnum {
	return CreateDatakeyRequestBodyKeySpecEnum{
		AES_256: CreateDatakeyRequestBodyKeySpec{
			value: "AES_256",
		},
		AES_128: CreateDatakeyRequestBodyKeySpec{
			value: "AES_128",
		},
	}
}

func (c CreateDatakeyRequestBodyKeySpec) Value() string {
	return c.value
}

func (c CreateDatakeyRequestBodyKeySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDatakeyRequestBodyKeySpec) UnmarshalJSON(b []byte) error {
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

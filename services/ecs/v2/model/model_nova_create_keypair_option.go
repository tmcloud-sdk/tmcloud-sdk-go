package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaCreateKeypairOption struct {
	PublicKey *string `json:"public_key,omitempty"`

	Type *NovaCreateKeypairOptionType `json:"type,omitempty"`

	Name string `json:"name"`

	UserId *string `json:"user_id,omitempty"`
}

func (o NovaCreateKeypairOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairOption struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairOption", string(data)}, " ")
}

type NovaCreateKeypairOptionType struct {
	value string
}

type NovaCreateKeypairOptionTypeEnum struct {
	SSH  NovaCreateKeypairOptionType
	X509 NovaCreateKeypairOptionType
}

func GetNovaCreateKeypairOptionTypeEnum() NovaCreateKeypairOptionTypeEnum {
	return NovaCreateKeypairOptionTypeEnum{
		SSH: NovaCreateKeypairOptionType{
			value: "ssh",
		},
		X509: NovaCreateKeypairOptionType{
			value: "x509",
		},
	}
}

func (c NovaCreateKeypairOptionType) Value() string {
	return c.value
}

func (c NovaCreateKeypairOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaCreateKeypairOptionType) UnmarshalJSON(b []byte) error {
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

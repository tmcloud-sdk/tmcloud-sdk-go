package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ExtraDhcpOption struct {
	OptName ExtraDhcpOptionOptName `json:"opt_name"`

	OptValue *string `json:"opt_value,omitempty"`
}

func (o ExtraDhcpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtraDhcpOption struct{}"
	}

	return strings.Join([]string{"ExtraDhcpOption", string(data)}, " ")
}

type ExtraDhcpOptionOptName struct {
	value string
}

type ExtraDhcpOptionOptNameEnum struct {
	NTP         ExtraDhcpOptionOptName
	ADDRESSTIME ExtraDhcpOptionOptName
}

func GetExtraDhcpOptionOptNameEnum() ExtraDhcpOptionOptNameEnum {
	return ExtraDhcpOptionOptNameEnum{
		NTP: ExtraDhcpOptionOptName{
			value: "ntp",
		},
		ADDRESSTIME: ExtraDhcpOptionOptName{
			value: "addresstime",
		},
	}
}

func (c ExtraDhcpOptionOptName) Value() string {
	return c.value
}

func (c ExtraDhcpOptionOptName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExtraDhcpOptionOptName) UnmarshalJSON(b []byte) error {
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

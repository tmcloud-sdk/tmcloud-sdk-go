package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ConfigurationParameter struct {
	Name string `json:"name"`

	Value string `json:"value"`

	RestartRequired bool `json:"restart_required"`

	Readonly bool `json:"readonly"`

	ValueRange string `json:"value_range"`

	Type ConfigurationParameterType `json:"type"`

	Description string `json:"description"`
}

func (o ConfigurationParameter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameter struct{}"
	}

	return strings.Join([]string{"ConfigurationParameter", string(data)}, " ")
}

type ConfigurationParameterType struct {
	value string
}

type ConfigurationParameterTypeEnum struct {
	STRING  ConfigurationParameterType
	INTEGER ConfigurationParameterType
	BOOLEAN ConfigurationParameterType
	LIST    ConfigurationParameterType
	FLOAT   ConfigurationParameterType
}

func GetConfigurationParameterTypeEnum() ConfigurationParameterTypeEnum {
	return ConfigurationParameterTypeEnum{
		STRING: ConfigurationParameterType{
			value: "string",
		},
		INTEGER: ConfigurationParameterType{
			value: "integer",
		},
		BOOLEAN: ConfigurationParameterType{
			value: "boolean",
		},
		LIST: ConfigurationParameterType{
			value: "list",
		},
		FLOAT: ConfigurationParameterType{
			value: "float",
		},
	}
}

func (c ConfigurationParameterType) Value() string {
	return c.value
}

func (c ConfigurationParameterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationParameterType) UnmarshalJSON(b []byte) error {
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

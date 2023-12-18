package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateConfigurationRequest struct {
	XLanguage *CreateConfigurationRequestXLanguage `json:"X-Language,omitempty"`

	Body *ConfigurationForCreation `json:"body,omitempty"`
}

func (o CreateConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationRequest struct{}"
	}

	return strings.Join([]string{"CreateConfigurationRequest", string(data)}, " ")
}

type CreateConfigurationRequestXLanguage struct {
	value string
}

type CreateConfigurationRequestXLanguageEnum struct {
	ZH_CN CreateConfigurationRequestXLanguage
	EN_US CreateConfigurationRequestXLanguage
}

func GetCreateConfigurationRequestXLanguageEnum() CreateConfigurationRequestXLanguageEnum {
	return CreateConfigurationRequestXLanguageEnum{
		ZH_CN: CreateConfigurationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: CreateConfigurationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c CreateConfigurationRequestXLanguage) Value() string {
	return c.value
}

func (c CreateConfigurationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConfigurationRequestXLanguage) UnmarshalJSON(b []byte) error {
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

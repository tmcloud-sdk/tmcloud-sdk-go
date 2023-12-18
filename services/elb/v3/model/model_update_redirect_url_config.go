package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateRedirectUrlConfig struct {
	Protocol *UpdateRedirectUrlConfigProtocol `json:"protocol,omitempty"`

	Host *string `json:"host,omitempty"`

	Port *string `json:"port,omitempty"`

	Path *string `json:"path,omitempty"`

	Query *string `json:"query,omitempty"`

	StatusCode *UpdateRedirectUrlConfigStatusCode `json:"status_code,omitempty"`
}

func (o UpdateRedirectUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedirectUrlConfig struct{}"
	}

	return strings.Join([]string{"UpdateRedirectUrlConfig", string(data)}, " ")
}

type UpdateRedirectUrlConfigProtocol struct {
	value string
}

type UpdateRedirectUrlConfigProtocolEnum struct {
	HTTP      UpdateRedirectUrlConfigProtocol
	HTTPS     UpdateRedirectUrlConfigProtocol
	_PROTOCOL UpdateRedirectUrlConfigProtocol
}

func GetUpdateRedirectUrlConfigProtocolEnum() UpdateRedirectUrlConfigProtocolEnum {
	return UpdateRedirectUrlConfigProtocolEnum{
		HTTP: UpdateRedirectUrlConfigProtocol{
			value: "HTTP",
		},
		HTTPS: UpdateRedirectUrlConfigProtocol{
			value: "HTTPS",
		},
		_PROTOCOL: UpdateRedirectUrlConfigProtocol{
			value: "${protocol}",
		},
	}
}

func (c UpdateRedirectUrlConfigProtocol) Value() string {
	return c.value
}

func (c UpdateRedirectUrlConfigProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRedirectUrlConfigProtocol) UnmarshalJSON(b []byte) error {
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

type UpdateRedirectUrlConfigStatusCode struct {
	value string
}

type UpdateRedirectUrlConfigStatusCodeEnum struct {
	E_301 UpdateRedirectUrlConfigStatusCode
	E_302 UpdateRedirectUrlConfigStatusCode
	E_303 UpdateRedirectUrlConfigStatusCode
	E_307 UpdateRedirectUrlConfigStatusCode
	E_308 UpdateRedirectUrlConfigStatusCode
}

func GetUpdateRedirectUrlConfigStatusCodeEnum() UpdateRedirectUrlConfigStatusCodeEnum {
	return UpdateRedirectUrlConfigStatusCodeEnum{
		E_301: UpdateRedirectUrlConfigStatusCode{
			value: "301",
		},
		E_302: UpdateRedirectUrlConfigStatusCode{
			value: "302",
		},
		E_303: UpdateRedirectUrlConfigStatusCode{
			value: "303",
		},
		E_307: UpdateRedirectUrlConfigStatusCode{
			value: "307",
		},
		E_308: UpdateRedirectUrlConfigStatusCode{
			value: "308",
		},
	}
}

func (c UpdateRedirectUrlConfigStatusCode) Value() string {
	return c.value
}

func (c UpdateRedirectUrlConfigStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRedirectUrlConfigStatusCode) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateRedirectUrlConfig struct {
	Protocol *CreateRedirectUrlConfigProtocol `json:"protocol,omitempty"`

	Host *string `json:"host,omitempty"`

	Port *string `json:"port,omitempty"`

	Path *string `json:"path,omitempty"`

	Query *string `json:"query,omitempty"`

	StatusCode CreateRedirectUrlConfigStatusCode `json:"status_code"`
}

func (o CreateRedirectUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedirectUrlConfig struct{}"
	}

	return strings.Join([]string{"CreateRedirectUrlConfig", string(data)}, " ")
}

type CreateRedirectUrlConfigProtocol struct {
	value string
}

type CreateRedirectUrlConfigProtocolEnum struct {
	HTTP      CreateRedirectUrlConfigProtocol
	HTTPS     CreateRedirectUrlConfigProtocol
	_PROTOCOL CreateRedirectUrlConfigProtocol
}

func GetCreateRedirectUrlConfigProtocolEnum() CreateRedirectUrlConfigProtocolEnum {
	return CreateRedirectUrlConfigProtocolEnum{
		HTTP: CreateRedirectUrlConfigProtocol{
			value: "HTTP",
		},
		HTTPS: CreateRedirectUrlConfigProtocol{
			value: "HTTPS",
		},
		_PROTOCOL: CreateRedirectUrlConfigProtocol{
			value: "${protocol}",
		},
	}
}

func (c CreateRedirectUrlConfigProtocol) Value() string {
	return c.value
}

func (c CreateRedirectUrlConfigProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRedirectUrlConfigProtocol) UnmarshalJSON(b []byte) error {
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

type CreateRedirectUrlConfigStatusCode struct {
	value string
}

type CreateRedirectUrlConfigStatusCodeEnum struct {
	E_301 CreateRedirectUrlConfigStatusCode
	E_302 CreateRedirectUrlConfigStatusCode
	E_303 CreateRedirectUrlConfigStatusCode
	E_307 CreateRedirectUrlConfigStatusCode
	E_308 CreateRedirectUrlConfigStatusCode
}

func GetCreateRedirectUrlConfigStatusCodeEnum() CreateRedirectUrlConfigStatusCodeEnum {
	return CreateRedirectUrlConfigStatusCodeEnum{
		E_301: CreateRedirectUrlConfigStatusCode{
			value: "301",
		},
		E_302: CreateRedirectUrlConfigStatusCode{
			value: "302",
		},
		E_303: CreateRedirectUrlConfigStatusCode{
			value: "303",
		},
		E_307: CreateRedirectUrlConfigStatusCode{
			value: "307",
		},
		E_308: CreateRedirectUrlConfigStatusCode{
			value: "308",
		},
	}
}

func (c CreateRedirectUrlConfigStatusCode) Value() string {
	return c.value
}

func (c CreateRedirectUrlConfigStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRedirectUrlConfigStatusCode) UnmarshalJSON(b []byte) error {
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

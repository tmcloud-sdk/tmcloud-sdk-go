package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type RedirectUrlConfig struct {
	Protocol RedirectUrlConfigProtocol `json:"protocol"`

	Host string `json:"host"`

	Port string `json:"port"`

	Path string `json:"path"`

	Query string `json:"query"`

	StatusCode RedirectUrlConfigStatusCode `json:"status_code"`
}

func (o RedirectUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RedirectUrlConfig struct{}"
	}

	return strings.Join([]string{"RedirectUrlConfig", string(data)}, " ")
}

type RedirectUrlConfigProtocol struct {
	value string
}

type RedirectUrlConfigProtocolEnum struct {
	HTTP      RedirectUrlConfigProtocol
	HTTPS     RedirectUrlConfigProtocol
	_PROTOCOL RedirectUrlConfigProtocol
}

func GetRedirectUrlConfigProtocolEnum() RedirectUrlConfigProtocolEnum {
	return RedirectUrlConfigProtocolEnum{
		HTTP: RedirectUrlConfigProtocol{
			value: "HTTP",
		},
		HTTPS: RedirectUrlConfigProtocol{
			value: "HTTPS",
		},
		_PROTOCOL: RedirectUrlConfigProtocol{
			value: "${protocol}",
		},
	}
}

func (c RedirectUrlConfigProtocol) Value() string {
	return c.value
}

func (c RedirectUrlConfigProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RedirectUrlConfigProtocol) UnmarshalJSON(b []byte) error {
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

type RedirectUrlConfigStatusCode struct {
	value string
}

type RedirectUrlConfigStatusCodeEnum struct {
	E_301 RedirectUrlConfigStatusCode
	E_302 RedirectUrlConfigStatusCode
	E_303 RedirectUrlConfigStatusCode
	E_307 RedirectUrlConfigStatusCode
	E_308 RedirectUrlConfigStatusCode
}

func GetRedirectUrlConfigStatusCodeEnum() RedirectUrlConfigStatusCodeEnum {
	return RedirectUrlConfigStatusCodeEnum{
		E_301: RedirectUrlConfigStatusCode{
			value: "301",
		},
		E_302: RedirectUrlConfigStatusCode{
			value: "302",
		},
		E_303: RedirectUrlConfigStatusCode{
			value: "303",
		},
		E_307: RedirectUrlConfigStatusCode{
			value: "307",
		},
		E_308: RedirectUrlConfigStatusCode{
			value: "308",
		},
	}
}

func (c RedirectUrlConfigStatusCode) Value() string {
	return c.value
}

func (c RedirectUrlConfigStatusCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RedirectUrlConfigStatusCode) UnmarshalJSON(b []byte) error {
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

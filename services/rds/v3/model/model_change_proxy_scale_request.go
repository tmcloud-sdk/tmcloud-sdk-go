package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ChangeProxyScaleRequest struct {
	XLanguage *ChangeProxyScaleRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ScaleProxyRequestBody `json:"body,omitempty"`
}

func (o ChangeProxyScaleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeProxyScaleRequest struct{}"
	}

	return strings.Join([]string{"ChangeProxyScaleRequest", string(data)}, " ")
}

type ChangeProxyScaleRequestXLanguage struct {
	value string
}

type ChangeProxyScaleRequestXLanguageEnum struct {
	ZH_CN ChangeProxyScaleRequestXLanguage
	EN_US ChangeProxyScaleRequestXLanguage
}

func GetChangeProxyScaleRequestXLanguageEnum() ChangeProxyScaleRequestXLanguageEnum {
	return ChangeProxyScaleRequestXLanguageEnum{
		ZH_CN: ChangeProxyScaleRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeProxyScaleRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeProxyScaleRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeProxyScaleRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeProxyScaleRequestXLanguage) UnmarshalJSON(b []byte) error {
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

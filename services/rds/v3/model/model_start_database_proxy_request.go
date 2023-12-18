package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StartDatabaseProxyRequest struct {
	XLanguage *StartDatabaseProxyRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *OpenProxyRequest `json:"body,omitempty"`
}

func (o StartDatabaseProxyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDatabaseProxyRequest struct{}"
	}

	return strings.Join([]string{"StartDatabaseProxyRequest", string(data)}, " ")
}

type StartDatabaseProxyRequestXLanguage struct {
	value string
}

type StartDatabaseProxyRequestXLanguageEnum struct {
	ZH_CN StartDatabaseProxyRequestXLanguage
	EN_US StartDatabaseProxyRequestXLanguage
}

func GetStartDatabaseProxyRequestXLanguageEnum() StartDatabaseProxyRequestXLanguageEnum {
	return StartDatabaseProxyRequestXLanguageEnum{
		ZH_CN: StartDatabaseProxyRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartDatabaseProxyRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartDatabaseProxyRequestXLanguage) Value() string {
	return c.value
}

func (c StartDatabaseProxyRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartDatabaseProxyRequestXLanguage) UnmarshalJSON(b []byte) error {
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

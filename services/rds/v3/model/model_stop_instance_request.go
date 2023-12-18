package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StopInstanceRequest struct {
	XLanguage *StopInstanceRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o StopInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInstanceRequest struct{}"
	}

	return strings.Join([]string{"StopInstanceRequest", string(data)}, " ")
}

type StopInstanceRequestXLanguage struct {
	value string
}

type StopInstanceRequestXLanguageEnum struct {
	ZH_CN StopInstanceRequestXLanguage
	EN_US StopInstanceRequestXLanguage
}

func GetStopInstanceRequestXLanguageEnum() StopInstanceRequestXLanguageEnum {
	return StopInstanceRequestXLanguageEnum{
		ZH_CN: StopInstanceRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StopInstanceRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StopInstanceRequestXLanguage) Value() string {
	return c.value
}

func (c StopInstanceRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopInstanceRequestXLanguage) UnmarshalJSON(b []byte) error {
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

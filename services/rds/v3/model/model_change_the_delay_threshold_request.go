package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ChangeTheDelayThresholdRequest struct {
	XLanguage *ChangeTheDelayThresholdRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ChangingTheDelayThresholdRequestBody `json:"body,omitempty"`
}

func (o ChangeTheDelayThresholdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeTheDelayThresholdRequest struct{}"
	}

	return strings.Join([]string{"ChangeTheDelayThresholdRequest", string(data)}, " ")
}

type ChangeTheDelayThresholdRequestXLanguage struct {
	value string
}

type ChangeTheDelayThresholdRequestXLanguageEnum struct {
	ZH_CN ChangeTheDelayThresholdRequestXLanguage
	EN_US ChangeTheDelayThresholdRequestXLanguage
}

func GetChangeTheDelayThresholdRequestXLanguageEnum() ChangeTheDelayThresholdRequestXLanguageEnum {
	return ChangeTheDelayThresholdRequestXLanguageEnum{
		ZH_CN: ChangeTheDelayThresholdRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeTheDelayThresholdRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeTheDelayThresholdRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeTheDelayThresholdRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeTheDelayThresholdRequestXLanguage) UnmarshalJSON(b []byte) error {
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

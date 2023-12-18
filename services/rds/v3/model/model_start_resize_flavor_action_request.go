package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StartResizeFlavorActionRequest struct {
	XLanguage *StartResizeFlavorActionRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ResizeFlavorRequest `json:"body,omitempty"`
}

func (o StartResizeFlavorActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResizeFlavorActionRequest struct{}"
	}

	return strings.Join([]string{"StartResizeFlavorActionRequest", string(data)}, " ")
}

type StartResizeFlavorActionRequestXLanguage struct {
	value string
}

type StartResizeFlavorActionRequestXLanguageEnum struct {
	ZH_CN StartResizeFlavorActionRequestXLanguage
	EN_US StartResizeFlavorActionRequestXLanguage
}

func GetStartResizeFlavorActionRequestXLanguageEnum() StartResizeFlavorActionRequestXLanguageEnum {
	return StartResizeFlavorActionRequestXLanguageEnum{
		ZH_CN: StartResizeFlavorActionRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: StartResizeFlavorActionRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c StartResizeFlavorActionRequestXLanguage) Value() string {
	return c.value
}

func (c StartResizeFlavorActionRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartResizeFlavorActionRequestXLanguage) UnmarshalJSON(b []byte) error {
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

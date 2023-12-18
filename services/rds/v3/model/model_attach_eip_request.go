package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type AttachEipRequest struct {
	XLanguage *AttachEipRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *BindEipRequest `json:"body,omitempty"`
}

func (o AttachEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachEipRequest struct{}"
	}

	return strings.Join([]string{"AttachEipRequest", string(data)}, " ")
}

type AttachEipRequestXLanguage struct {
	value string
}

type AttachEipRequestXLanguageEnum struct {
	ZH_CN AttachEipRequestXLanguage
	EN_US AttachEipRequestXLanguage
}

func GetAttachEipRequestXLanguageEnum() AttachEipRequestXLanguageEnum {
	return AttachEipRequestXLanguageEnum{
		ZH_CN: AttachEipRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: AttachEipRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c AttachEipRequestXLanguage) Value() string {
	return c.value
}

func (c AttachEipRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachEipRequestXLanguage) UnmarshalJSON(b []byte) error {
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

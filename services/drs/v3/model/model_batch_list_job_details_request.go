package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchListJobDetailsRequest struct {
	XLanguage *BatchListJobDetailsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryJobReqPage `json:"body,omitempty"`
}

func (o BatchListJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"BatchListJobDetailsRequest", string(data)}, " ")
}

type BatchListJobDetailsRequestXLanguage struct {
	value string
}

type BatchListJobDetailsRequestXLanguageEnum struct {
	EN_US BatchListJobDetailsRequestXLanguage
	ZH_CN BatchListJobDetailsRequestXLanguage
}

func GetBatchListJobDetailsRequestXLanguageEnum() BatchListJobDetailsRequestXLanguageEnum {
	return BatchListJobDetailsRequestXLanguageEnum{
		EN_US: BatchListJobDetailsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListJobDetailsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListJobDetailsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListJobDetailsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListJobDetailsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

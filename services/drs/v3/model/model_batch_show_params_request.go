package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchShowParamsRequest struct {
	XLanguage *BatchShowParamsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryParamReq `json:"body,omitempty"`
}

func (o BatchShowParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowParamsRequest struct{}"
	}

	return strings.Join([]string{"BatchShowParamsRequest", string(data)}, " ")
}

type BatchShowParamsRequestXLanguage struct {
	value string
}

type BatchShowParamsRequestXLanguageEnum struct {
	EN_US BatchShowParamsRequestXLanguage
	ZH_CN BatchShowParamsRequestXLanguage
}

func GetBatchShowParamsRequestXLanguageEnum() BatchShowParamsRequestXLanguageEnum {
	return BatchShowParamsRequestXLanguageEnum{
		EN_US: BatchShowParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchShowParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchShowParamsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchShowParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchShowParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

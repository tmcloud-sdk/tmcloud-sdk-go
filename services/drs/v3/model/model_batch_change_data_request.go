package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchChangeDataRequest struct {
	XLanguage *BatchChangeDataRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchDataTransformationReq `json:"body,omitempty"`
}

func (o BatchChangeDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeDataRequest struct{}"
	}

	return strings.Join([]string{"BatchChangeDataRequest", string(data)}, " ")
}

type BatchChangeDataRequestXLanguage struct {
	value string
}

type BatchChangeDataRequestXLanguageEnum struct {
	EN_US BatchChangeDataRequestXLanguage
	ZH_CN BatchChangeDataRequestXLanguage
}

func GetBatchChangeDataRequestXLanguageEnum() BatchChangeDataRequestXLanguageEnum {
	return BatchChangeDataRequestXLanguageEnum{
		EN_US: BatchChangeDataRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchChangeDataRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchChangeDataRequestXLanguage) Value() string {
	return c.value
}

func (c BatchChangeDataRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchChangeDataRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchCheckResultsRequest struct {
	XLanguage *BatchCheckResultsRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryPrecheckResultReq `json:"body,omitempty"`
}

func (o BatchCheckResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckResultsRequest struct{}"
	}

	return strings.Join([]string{"BatchCheckResultsRequest", string(data)}, " ")
}

type BatchCheckResultsRequestXLanguage struct {
	value string
}

type BatchCheckResultsRequestXLanguageEnum struct {
	EN_US BatchCheckResultsRequestXLanguage
	ZH_CN BatchCheckResultsRequestXLanguage
}

func GetBatchCheckResultsRequestXLanguageEnum() BatchCheckResultsRequestXLanguageEnum {
	return BatchCheckResultsRequestXLanguageEnum{
		EN_US: BatchCheckResultsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchCheckResultsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchCheckResultsRequestXLanguage) Value() string {
	return c.value
}

func (c BatchCheckResultsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCheckResultsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

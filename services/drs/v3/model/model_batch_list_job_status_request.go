package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchListJobStatusRequest struct {
	XLanguage *BatchListJobStatusRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchQueryJobReqPage `json:"body,omitempty"`
}

func (o BatchListJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobStatusRequest struct{}"
	}

	return strings.Join([]string{"BatchListJobStatusRequest", string(data)}, " ")
}

type BatchListJobStatusRequestXLanguage struct {
	value string
}

type BatchListJobStatusRequestXLanguageEnum struct {
	EN_US BatchListJobStatusRequestXLanguage
	ZH_CN BatchListJobStatusRequestXLanguage
}

func GetBatchListJobStatusRequestXLanguageEnum() BatchListJobStatusRequestXLanguageEnum {
	return BatchListJobStatusRequestXLanguageEnum{
		EN_US: BatchListJobStatusRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchListJobStatusRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchListJobStatusRequestXLanguage) Value() string {
	return c.value
}

func (c BatchListJobStatusRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListJobStatusRequestXLanguage) UnmarshalJSON(b []byte) error {
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

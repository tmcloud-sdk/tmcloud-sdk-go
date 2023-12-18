package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchSwitchoverRequest struct {
	XLanguage *BatchSwitchoverRequestXLanguage `json:"X-Language,omitempty"`

	Body *BatchSwitchoverReq `json:"body,omitempty"`
}

func (o BatchSwitchoverRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSwitchoverRequest struct{}"
	}

	return strings.Join([]string{"BatchSwitchoverRequest", string(data)}, " ")
}

type BatchSwitchoverRequestXLanguage struct {
	value string
}

type BatchSwitchoverRequestXLanguageEnum struct {
	EN_US BatchSwitchoverRequestXLanguage
	ZH_CN BatchSwitchoverRequestXLanguage
}

func GetBatchSwitchoverRequestXLanguageEnum() BatchSwitchoverRequestXLanguageEnum {
	return BatchSwitchoverRequestXLanguageEnum{
		EN_US: BatchSwitchoverRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: BatchSwitchoverRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c BatchSwitchoverRequestXLanguage) Value() string {
	return c.value
}

func (c BatchSwitchoverRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchSwitchoverRequestXLanguage) UnmarshalJSON(b []byte) error {
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

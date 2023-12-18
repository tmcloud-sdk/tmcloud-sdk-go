package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListSlowlogForLtsRequest struct {
	InstanceId string `json:"instance_id"`

	XLanguage *ListSlowlogForLtsRequestXLanguage `json:"X-Language,omitempty"`

	Body *SlowlogForLtsRequest `json:"body,omitempty"`
}

func (o ListSlowlogForLtsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogForLtsRequest struct{}"
	}

	return strings.Join([]string{"ListSlowlogForLtsRequest", string(data)}, " ")
}

type ListSlowlogForLtsRequestXLanguage struct {
	value string
}

type ListSlowlogForLtsRequestXLanguageEnum struct {
	ZH_CN ListSlowlogForLtsRequestXLanguage
	EN_US ListSlowlogForLtsRequestXLanguage
}

func GetListSlowlogForLtsRequestXLanguageEnum() ListSlowlogForLtsRequestXLanguageEnum {
	return ListSlowlogForLtsRequestXLanguageEnum{
		ZH_CN: ListSlowlogForLtsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListSlowlogForLtsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListSlowlogForLtsRequestXLanguage) Value() string {
	return c.value
}

func (c ListSlowlogForLtsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSlowlogForLtsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

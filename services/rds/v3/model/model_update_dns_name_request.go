package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateDnsNameRequest struct {
	XLanguage *UpdateDnsNameRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *ModifyDnsNameRequestBody `json:"body,omitempty"`
}

func (o UpdateDnsNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateDnsNameRequest", string(data)}, " ")
}

type UpdateDnsNameRequestXLanguage struct {
	value string
}

type UpdateDnsNameRequestXLanguageEnum struct {
	ZH_CN UpdateDnsNameRequestXLanguage
	EN_US UpdateDnsNameRequestXLanguage
}

func GetUpdateDnsNameRequestXLanguageEnum() UpdateDnsNameRequestXLanguageEnum {
	return UpdateDnsNameRequestXLanguageEnum{
		ZH_CN: UpdateDnsNameRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: UpdateDnsNameRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c UpdateDnsNameRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateDnsNameRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDnsNameRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ChangeOpsWindowRequest struct {
	XLanguage *ChangeOpsWindowRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *OpsWindowRequest `json:"body,omitempty"`
}

func (o ChangeOpsWindowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOpsWindowRequest struct{}"
	}

	return strings.Join([]string{"ChangeOpsWindowRequest", string(data)}, " ")
}

type ChangeOpsWindowRequestXLanguage struct {
	value string
}

type ChangeOpsWindowRequestXLanguageEnum struct {
	ZH_CN ChangeOpsWindowRequestXLanguage
	EN_US ChangeOpsWindowRequestXLanguage
}

func GetChangeOpsWindowRequestXLanguageEnum() ChangeOpsWindowRequestXLanguageEnum {
	return ChangeOpsWindowRequestXLanguageEnum{
		ZH_CN: ChangeOpsWindowRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ChangeOpsWindowRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ChangeOpsWindowRequestXLanguage) Value() string {
	return c.value
}

func (c ChangeOpsWindowRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeOpsWindowRequestXLanguage) UnmarshalJSON(b []byte) error {
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

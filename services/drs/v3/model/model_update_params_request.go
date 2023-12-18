package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateParamsRequest struct {
	JobId string `json:"job_id"`

	XLanguage *UpdateParamsRequestXLanguage `json:"X-Language,omitempty"`

	Body *ModifyTargetParamsReq `json:"body,omitempty"`
}

func (o UpdateParamsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateParamsRequest struct{}"
	}

	return strings.Join([]string{"UpdateParamsRequest", string(data)}, " ")
}

type UpdateParamsRequestXLanguage struct {
	value string
}

type UpdateParamsRequestXLanguageEnum struct {
	EN_US UpdateParamsRequestXLanguage
	ZH_CN UpdateParamsRequestXLanguage
}

func GetUpdateParamsRequestXLanguageEnum() UpdateParamsRequestXLanguageEnum {
	return UpdateParamsRequestXLanguageEnum{
		EN_US: UpdateParamsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: UpdateParamsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c UpdateParamsRequestXLanguage) Value() string {
	return c.value
}

func (c UpdateParamsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateParamsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

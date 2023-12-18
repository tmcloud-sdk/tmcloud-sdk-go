package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SetSecurityGroupRequest struct {
	XLanguage *SetSecurityGroupRequestXLanguage `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SecurityGroupRequest `json:"body,omitempty"`
}

func (o SetSecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetSecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"SetSecurityGroupRequest", string(data)}, " ")
}

type SetSecurityGroupRequestXLanguage struct {
	value string
}

type SetSecurityGroupRequestXLanguageEnum struct {
	ZH_CN SetSecurityGroupRequestXLanguage
	EN_US SetSecurityGroupRequestXLanguage
}

func GetSetSecurityGroupRequestXLanguageEnum() SetSecurityGroupRequestXLanguageEnum {
	return SetSecurityGroupRequestXLanguageEnum{
		ZH_CN: SetSecurityGroupRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetSecurityGroupRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetSecurityGroupRequestXLanguage) Value() string {
	return c.value
}

func (c SetSecurityGroupRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetSecurityGroupRequestXLanguage) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateListenerIpGroupOption struct {
	IpgroupId *string `json:"ipgroup_id,omitempty"`

	EnableIpgroup *bool `json:"enable_ipgroup,omitempty"`

	Type *UpdateListenerIpGroupOptionType `json:"type,omitempty"`
}

func (o UpdateListenerIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerIpGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerIpGroupOption", string(data)}, " ")
}

type UpdateListenerIpGroupOptionType struct {
	value string
}

type UpdateListenerIpGroupOptionTypeEnum struct {
	WHITE UpdateListenerIpGroupOptionType
	BLACK UpdateListenerIpGroupOptionType
}

func GetUpdateListenerIpGroupOptionTypeEnum() UpdateListenerIpGroupOptionTypeEnum {
	return UpdateListenerIpGroupOptionTypeEnum{
		WHITE: UpdateListenerIpGroupOptionType{
			value: "white",
		},
		BLACK: UpdateListenerIpGroupOptionType{
			value: "black",
		},
	}
}

func (c UpdateListenerIpGroupOptionType) Value() string {
	return c.value
}

func (c UpdateListenerIpGroupOptionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateListenerIpGroupOptionType) UnmarshalJSON(b []byte) error {
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

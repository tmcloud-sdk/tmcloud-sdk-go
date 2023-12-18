package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateNatGatewayOption struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Spec *UpdateNatGatewayOptionSpec `json:"spec,omitempty"`
}

func (o UpdateNatGatewayOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayOption", string(data)}, " ")
}

type UpdateNatGatewayOptionSpec struct {
	value string
}

type UpdateNatGatewayOptionSpecEnum struct {
	E_1 UpdateNatGatewayOptionSpec
	E_2 UpdateNatGatewayOptionSpec
	E_3 UpdateNatGatewayOptionSpec
	E_4 UpdateNatGatewayOptionSpec
}

func GetUpdateNatGatewayOptionSpecEnum() UpdateNatGatewayOptionSpecEnum {
	return UpdateNatGatewayOptionSpecEnum{
		E_1: UpdateNatGatewayOptionSpec{
			value: "1",
		},
		E_2: UpdateNatGatewayOptionSpec{
			value: "2",
		},
		E_3: UpdateNatGatewayOptionSpec{
			value: "3",
		},
		E_4: UpdateNatGatewayOptionSpec{
			value: "4",
		},
	}
}

func (c UpdateNatGatewayOptionSpec) Value() string {
	return c.value
}

func (c UpdateNatGatewayOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNatGatewayOptionSpec) UnmarshalJSON(b []byte) error {
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

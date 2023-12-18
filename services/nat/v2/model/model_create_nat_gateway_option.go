package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateNatGatewayOption struct {
	Name string `json:"name"`

	RouterId string `json:"router_id"`

	InternalNetworkId string `json:"internal_network_id"`

	Description *string `json:"description,omitempty"`

	Spec CreateNatGatewayOptionSpec `json:"spec"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateNatGatewayOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayOption struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayOption", string(data)}, " ")
}

type CreateNatGatewayOptionSpec struct {
	value string
}

type CreateNatGatewayOptionSpecEnum struct {
	E_1 CreateNatGatewayOptionSpec
	E_2 CreateNatGatewayOptionSpec
	E_3 CreateNatGatewayOptionSpec
	E_4 CreateNatGatewayOptionSpec
}

func GetCreateNatGatewayOptionSpecEnum() CreateNatGatewayOptionSpecEnum {
	return CreateNatGatewayOptionSpecEnum{
		E_1: CreateNatGatewayOptionSpec{
			value: "1",
		},
		E_2: CreateNatGatewayOptionSpec{
			value: "2",
		},
		E_3: CreateNatGatewayOptionSpec{
			value: "3",
		},
		E_4: CreateNatGatewayOptionSpec{
			value: "4",
		},
	}
}

func (c CreateNatGatewayOptionSpec) Value() string {
	return c.value
}

func (c CreateNatGatewayOptionSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNatGatewayOptionSpec) UnmarshalJSON(b []byte) error {
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

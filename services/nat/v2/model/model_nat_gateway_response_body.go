package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NatGatewayResponseBody struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Spec NatGatewayResponseBodySpec `json:"spec"`

	Status NatGatewayResponseBodyStatus `json:"status"`

	AdminStateUp bool `json:"admin_state_up"`

	CreatedAt string `json:"created_at"`

	RouterId string `json:"router_id"`

	InternalNetworkId string `json:"internal_network_id"`

	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o NatGatewayResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewayResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewayResponseBody", string(data)}, " ")
}

type NatGatewayResponseBodySpec struct {
	value string
}

type NatGatewayResponseBodySpecEnum struct {
	E_1 NatGatewayResponseBodySpec
	E_2 NatGatewayResponseBodySpec
	E_3 NatGatewayResponseBodySpec
	E_4 NatGatewayResponseBodySpec
}

func GetNatGatewayResponseBodySpecEnum() NatGatewayResponseBodySpecEnum {
	return NatGatewayResponseBodySpecEnum{
		E_1: NatGatewayResponseBodySpec{
			value: "1",
		},
		E_2: NatGatewayResponseBodySpec{
			value: "2",
		},
		E_3: NatGatewayResponseBodySpec{
			value: "3",
		},
		E_4: NatGatewayResponseBodySpec{
			value: "4",
		},
	}
}

func (c NatGatewayResponseBodySpec) Value() string {
	return c.value
}

func (c NatGatewayResponseBodySpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayResponseBodySpec) UnmarshalJSON(b []byte) error {
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

type NatGatewayResponseBodyStatus struct {
	value string
}

type NatGatewayResponseBodyStatusEnum struct {
	ACTIVE         NatGatewayResponseBodyStatus
	PENDING_CREATE NatGatewayResponseBodyStatus
	PENDING_UPDATE NatGatewayResponseBodyStatus
	PENDING_DELETE NatGatewayResponseBodyStatus
	INACTIVE       NatGatewayResponseBodyStatus
}

func GetNatGatewayResponseBodyStatusEnum() NatGatewayResponseBodyStatusEnum {
	return NatGatewayResponseBodyStatusEnum{
		ACTIVE: NatGatewayResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewayResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewayResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewayResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		INACTIVE: NatGatewayResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewayResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewayResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

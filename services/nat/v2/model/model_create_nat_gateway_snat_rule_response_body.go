package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateNatGatewaySnatRuleResponseBody struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	NatGatewayId string `json:"nat_gateway_id"`

	Cidr string `json:"cidr"`

	SourceType int32 `json:"source_type"`

	FloatingIpId string `json:"floating_ip_id"`

	Description string `json:"description"`

	Status CreateNatGatewaySnatRuleResponseBodyStatus `json:"status"`

	CreatedAt string `json:"created_at"`

	NetworkId string `json:"network_id"`

	AdminStateUp bool `json:"admin_state_up"`

	FloatingIpAddress string `json:"floating_ip_address"`
}

func (o CreateNatGatewaySnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewaySnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"CreateNatGatewaySnatRuleResponseBody", string(data)}, " ")
}

type CreateNatGatewaySnatRuleResponseBodyStatus struct {
	value string
}

type CreateNatGatewaySnatRuleResponseBodyStatusEnum struct {
	ACTIVE         CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_CREATE CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_UPDATE CreateNatGatewaySnatRuleResponseBodyStatus
	PENDING_DELETE CreateNatGatewaySnatRuleResponseBodyStatus
	EIP_FREEZED    CreateNatGatewaySnatRuleResponseBodyStatus
	INACTIVE       CreateNatGatewaySnatRuleResponseBodyStatus
}

func GetCreateNatGatewaySnatRuleResponseBodyStatusEnum() CreateNatGatewaySnatRuleResponseBodyStatusEnum {
	return CreateNatGatewaySnatRuleResponseBodyStatusEnum{
		ACTIVE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: CreateNatGatewaySnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c CreateNatGatewaySnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c CreateNatGatewaySnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNatGatewaySnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

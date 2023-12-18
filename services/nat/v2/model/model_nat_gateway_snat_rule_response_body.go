package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NatGatewaySnatRuleResponseBody struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	NatGatewayId string `json:"nat_gateway_id"`

	Cidr string `json:"cidr"`

	SourceType int32 `json:"source_type"`

	FloatingIpId string `json:"floating_ip_id"`

	Description string `json:"description"`

	Status NatGatewaySnatRuleResponseBodyStatus `json:"status"`

	CreatedAt string `json:"created_at"`

	NetworkId string `json:"network_id"`

	AdminStateUp bool `json:"admin_state_up"`

	FloatingIpAddress string `json:"floating_ip_address"`

	FreezedIpAddress string `json:"freezed_ip_address"`
}

func (o NatGatewaySnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewaySnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewaySnatRuleResponseBody", string(data)}, " ")
}

type NatGatewaySnatRuleResponseBodyStatus struct {
	value string
}

type NatGatewaySnatRuleResponseBodyStatusEnum struct {
	ACTIVE         NatGatewaySnatRuleResponseBodyStatus
	PENDING_CREATE NatGatewaySnatRuleResponseBodyStatus
	PENDING_UPDATE NatGatewaySnatRuleResponseBodyStatus
	PENDING_DELETE NatGatewaySnatRuleResponseBodyStatus
	EIP_FREEZED    NatGatewaySnatRuleResponseBodyStatus
	INACTIVE       NatGatewaySnatRuleResponseBodyStatus
}

func GetNatGatewaySnatRuleResponseBodyStatusEnum() NatGatewaySnatRuleResponseBodyStatusEnum {
	return NatGatewaySnatRuleResponseBodyStatusEnum{
		ACTIVE: NatGatewaySnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewaySnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: NatGatewaySnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: NatGatewaySnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewaySnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewaySnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewaySnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

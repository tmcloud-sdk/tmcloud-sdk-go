package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NatGatewayUpdateSnatRuleResponseBody struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	NatGatewayId string `json:"nat_gateway_id"`

	SourceType int32 `json:"source_type"`

	Cidr string `json:"cidr"`

	FloatingIpId string `json:"floating_ip_id"`

	Description string `json:"description"`

	Status NatGatewayUpdateSnatRuleResponseBodyStatus `json:"status"`

	CreatedAt string `json:"created_at"`

	NetworkId string `json:"network_id"`

	AdminStateUp bool `json:"admin_state_up"`

	FloatingIpAddress string `json:"floating_ip_address"`

	PublicIpAddress string `json:"public_ip_address"`
}

func (o NatGatewayUpdateSnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewayUpdateSnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewayUpdateSnatRuleResponseBody", string(data)}, " ")
}

type NatGatewayUpdateSnatRuleResponseBodyStatus struct {
	value string
}

type NatGatewayUpdateSnatRuleResponseBodyStatusEnum struct {
	ACTIVE         NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_CREATE NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_UPDATE NatGatewayUpdateSnatRuleResponseBodyStatus
	PENDING_DELETE NatGatewayUpdateSnatRuleResponseBodyStatus
	EIP_FREEZED    NatGatewayUpdateSnatRuleResponseBodyStatus
	INACTIVE       NatGatewayUpdateSnatRuleResponseBodyStatus
}

func GetNatGatewayUpdateSnatRuleResponseBodyStatusEnum() NatGatewayUpdateSnatRuleResponseBodyStatusEnum {
	return NatGatewayUpdateSnatRuleResponseBodyStatusEnum{
		ACTIVE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: NatGatewayUpdateSnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewayUpdateSnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewayUpdateSnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayUpdateSnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

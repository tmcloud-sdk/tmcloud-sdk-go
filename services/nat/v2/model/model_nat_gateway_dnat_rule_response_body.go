package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NatGatewayDnatRuleResponseBody struct {
	Id string `json:"id"`

	TenantId string `json:"tenant_id"`

	Description string `json:"description"`

	PortId *string `json:"port_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	InternalServicePort int32 `json:"internal_service_port"`

	NatGatewayId string `json:"nat_gateway_id"`

	FloatingIpId string `json:"floating_ip_id"`

	FloatingIpAddress string `json:"floating_ip_address"`

	ExternalServicePort int32 `json:"external_service_port"`

	Status NatGatewayDnatRuleResponseBodyStatus `json:"status"`

	AdminStateUp bool `json:"admin_state_up"`

	InternalServicePortRange *string `json:"internal_service_port_range,omitempty"`

	ExternalServicePortRange *string `json:"external_service_port_range,omitempty"`

	Protocol NatGatewayDnatRuleResponseBodyProtocol `json:"protocol"`

	CreatedAt string `json:"created_at"`
}

func (o NatGatewayDnatRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NatGatewayDnatRuleResponseBody struct{}"
	}

	return strings.Join([]string{"NatGatewayDnatRuleResponseBody", string(data)}, " ")
}

type NatGatewayDnatRuleResponseBodyStatus struct {
	value string
}

type NatGatewayDnatRuleResponseBodyStatusEnum struct {
	ACTIVE         NatGatewayDnatRuleResponseBodyStatus
	PENDING_CREATE NatGatewayDnatRuleResponseBodyStatus
	PENDING_UPDATE NatGatewayDnatRuleResponseBodyStatus
	PENDING_DELETE NatGatewayDnatRuleResponseBodyStatus
	EIP_FREEZED    NatGatewayDnatRuleResponseBodyStatus
	INACTIVE       NatGatewayDnatRuleResponseBodyStatus
}

func GetNatGatewayDnatRuleResponseBodyStatusEnum() NatGatewayDnatRuleResponseBodyStatusEnum {
	return NatGatewayDnatRuleResponseBodyStatusEnum{
		ACTIVE: NatGatewayDnatRuleResponseBodyStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: NatGatewayDnatRuleResponseBodyStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: NatGatewayDnatRuleResponseBodyStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: NatGatewayDnatRuleResponseBodyStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: NatGatewayDnatRuleResponseBodyStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: NatGatewayDnatRuleResponseBodyStatus{
			value: "INACTIVE",
		},
	}
}

func (c NatGatewayDnatRuleResponseBodyStatus) Value() string {
	return c.value
}

func (c NatGatewayDnatRuleResponseBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayDnatRuleResponseBodyStatus) UnmarshalJSON(b []byte) error {
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

type NatGatewayDnatRuleResponseBodyProtocol struct {
	value string
}

type NatGatewayDnatRuleResponseBodyProtocolEnum struct {
	TCP NatGatewayDnatRuleResponseBodyProtocol
	UDP NatGatewayDnatRuleResponseBodyProtocol
	ANY NatGatewayDnatRuleResponseBodyProtocol
}

func GetNatGatewayDnatRuleResponseBodyProtocolEnum() NatGatewayDnatRuleResponseBodyProtocolEnum {
	return NatGatewayDnatRuleResponseBodyProtocolEnum{
		TCP: NatGatewayDnatRuleResponseBodyProtocol{
			value: "tcp",
		},
		UDP: NatGatewayDnatRuleResponseBodyProtocol{
			value: "udp",
		},
		ANY: NatGatewayDnatRuleResponseBodyProtocol{
			value: "any",
		},
	}
}

func (c NatGatewayDnatRuleResponseBodyProtocol) Value() string {
	return c.value
}

func (c NatGatewayDnatRuleResponseBodyProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NatGatewayDnatRuleResponseBodyProtocol) UnmarshalJSON(b []byte) error {
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

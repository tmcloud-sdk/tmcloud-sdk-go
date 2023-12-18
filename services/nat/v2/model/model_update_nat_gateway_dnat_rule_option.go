package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateNatGatewayDnatRuleOption struct {
	NatGatewayId string `json:"nat_gateway_id"`

	Description *string `json:"description,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	Protocol *UpdateNatGatewayDnatRuleOptionProtocol `json:"protocol,omitempty"`

	FloatingIpId *string `json:"floating_ip_id,omitempty"`

	InternalServicePort *int32 `json:"internal_service_port,omitempty"`

	ExternalServicePort *int32 `json:"external_service_port,omitempty"`

	InternalServicePortRange *string `json:"internal_service_port_range,omitempty"`

	ExternalServicePortRange *string `json:"external_service_port_range,omitempty"`
}

func (o UpdateNatGatewayDnatRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayDnatRuleOption struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayDnatRuleOption", string(data)}, " ")
}

type UpdateNatGatewayDnatRuleOptionProtocol struct {
	value string
}

type UpdateNatGatewayDnatRuleOptionProtocolEnum struct {
	TCP UpdateNatGatewayDnatRuleOptionProtocol
	UDP UpdateNatGatewayDnatRuleOptionProtocol
	ANY UpdateNatGatewayDnatRuleOptionProtocol
}

func GetUpdateNatGatewayDnatRuleOptionProtocolEnum() UpdateNatGatewayDnatRuleOptionProtocolEnum {
	return UpdateNatGatewayDnatRuleOptionProtocolEnum{
		TCP: UpdateNatGatewayDnatRuleOptionProtocol{
			value: "TCP",
		},
		UDP: UpdateNatGatewayDnatRuleOptionProtocol{
			value: "UDP",
		},
		ANY: UpdateNatGatewayDnatRuleOptionProtocol{
			value: "ANY",
		},
	}
}

func (c UpdateNatGatewayDnatRuleOptionProtocol) Value() string {
	return c.value
}

func (c UpdateNatGatewayDnatRuleOptionProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateNatGatewayDnatRuleOptionProtocol) UnmarshalJSON(b []byte) error {
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

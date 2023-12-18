package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NeutronCreateSecurityGroupRuleOption struct {
	Description *string `json:"description,omitempty"`

	Direction NeutronCreateSecurityGroupRuleOptionDirection `json:"direction"`

	Ethertype *NeutronCreateSecurityGroupRuleOptionEthertype `json:"ethertype,omitempty"`

	PortRangeMax *int32 `json:"port_range_max,omitempty"`

	PortRangeMin *int32 `json:"port_range_min,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	RemoteGroupId *string `json:"remote_group_id,omitempty"`

	RemoteIpPrefix *string `json:"remote_ip_prefix,omitempty"`

	SecurityGroupId string `json:"security_group_id"`
}

func (o NeutronCreateSecurityGroupRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreateSecurityGroupRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronCreateSecurityGroupRuleOption", string(data)}, " ")
}

type NeutronCreateSecurityGroupRuleOptionDirection struct {
	value string
}

type NeutronCreateSecurityGroupRuleOptionDirectionEnum struct {
	INGRESS NeutronCreateSecurityGroupRuleOptionDirection
	EGRESS  NeutronCreateSecurityGroupRuleOptionDirection
}

func GetNeutronCreateSecurityGroupRuleOptionDirectionEnum() NeutronCreateSecurityGroupRuleOptionDirectionEnum {
	return NeutronCreateSecurityGroupRuleOptionDirectionEnum{
		INGRESS: NeutronCreateSecurityGroupRuleOptionDirection{
			value: "ingress",
		},
		EGRESS: NeutronCreateSecurityGroupRuleOptionDirection{
			value: "egress",
		},
	}
}

func (c NeutronCreateSecurityGroupRuleOptionDirection) Value() string {
	return c.value
}

func (c NeutronCreateSecurityGroupRuleOptionDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronCreateSecurityGroupRuleOptionDirection) UnmarshalJSON(b []byte) error {
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

type NeutronCreateSecurityGroupRuleOptionEthertype struct {
	value string
}

type NeutronCreateSecurityGroupRuleOptionEthertypeEnum struct {
	I_PV4 NeutronCreateSecurityGroupRuleOptionEthertype
	I_PV6 NeutronCreateSecurityGroupRuleOptionEthertype
}

func GetNeutronCreateSecurityGroupRuleOptionEthertypeEnum() NeutronCreateSecurityGroupRuleOptionEthertypeEnum {
	return NeutronCreateSecurityGroupRuleOptionEthertypeEnum{
		I_PV4: NeutronCreateSecurityGroupRuleOptionEthertype{
			value: "IPv4",
		},
		I_PV6: NeutronCreateSecurityGroupRuleOptionEthertype{
			value: "IPv6",
		},
	}
}

func (c NeutronCreateSecurityGroupRuleOptionEthertype) Value() string {
	return c.value
}

func (c NeutronCreateSecurityGroupRuleOptionEthertype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronCreateSecurityGroupRuleOptionEthertype) UnmarshalJSON(b []byte) error {
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

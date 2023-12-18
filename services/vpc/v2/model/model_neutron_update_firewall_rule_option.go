package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NeutronUpdateFirewallRuleOption struct {
	Action *NeutronUpdateFirewallRuleOptionAction `json:"action,omitempty"`

	Description *string `json:"description,omitempty"`

	DestinationIpAddress *string `json:"destination_ip_address,omitempty"`

	DestinationPort *string `json:"destination_port,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	IpVersion *int32 `json:"ip_version,omitempty"`

	Name *string `json:"name,omitempty"`

	Protocol *string `json:"protocol,omitempty"`

	SourceIpAddress *string `json:"source_ip_address,omitempty"`

	SourcePort *string `json:"source_port,omitempty"`
}

func (o NeutronUpdateFirewallRuleOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateFirewallRuleOption struct{}"
	}

	return strings.Join([]string{"NeutronUpdateFirewallRuleOption", string(data)}, " ")
}

type NeutronUpdateFirewallRuleOptionAction struct {
	value string
}

type NeutronUpdateFirewallRuleOptionActionEnum struct {
	DENY  NeutronUpdateFirewallRuleOptionAction
	ALLOW NeutronUpdateFirewallRuleOptionAction
}

func GetNeutronUpdateFirewallRuleOptionActionEnum() NeutronUpdateFirewallRuleOptionActionEnum {
	return NeutronUpdateFirewallRuleOptionActionEnum{
		DENY: NeutronUpdateFirewallRuleOptionAction{
			value: "DENY",
		},
		ALLOW: NeutronUpdateFirewallRuleOptionAction{
			value: "ALLOW",
		},
	}
}

func (c NeutronUpdateFirewallRuleOptionAction) Value() string {
	return c.value
}

func (c NeutronUpdateFirewallRuleOptionAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronUpdateFirewallRuleOptionAction) UnmarshalJSON(b []byte) error {
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

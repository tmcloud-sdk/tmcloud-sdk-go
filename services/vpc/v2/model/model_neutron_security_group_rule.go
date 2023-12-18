package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type NeutronSecurityGroupRule struct {
	Description string `json:"description"`

	Direction NeutronSecurityGroupRuleDirection `json:"direction"`

	Ethertype string `json:"ethertype"`

	Id string `json:"id"`

	PortRangeMax int32 `json:"port_range_max"`

	PortRangeMin int32 `json:"port_range_min"`

	Protocol string `json:"protocol"`

	RemoteGroupId string `json:"remote_group_id"`

	RemoteIpPrefix string `json:"remote_ip_prefix"`

	SecurityGroupId string `json:"security_group_id"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronSecurityGroupRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronSecurityGroupRule struct{}"
	}

	return strings.Join([]string{"NeutronSecurityGroupRule", string(data)}, " ")
}

type NeutronSecurityGroupRuleDirection struct {
	value string
}

type NeutronSecurityGroupRuleDirectionEnum struct {
	INGRESS NeutronSecurityGroupRuleDirection
	EGRESS  NeutronSecurityGroupRuleDirection
}

func GetNeutronSecurityGroupRuleDirectionEnum() NeutronSecurityGroupRuleDirectionEnum {
	return NeutronSecurityGroupRuleDirectionEnum{
		INGRESS: NeutronSecurityGroupRuleDirection{
			value: "ingress",
		},
		EGRESS: NeutronSecurityGroupRuleDirection{
			value: "egress",
		},
	}
}

func (c NeutronSecurityGroupRuleDirection) Value() string {
	return c.value
}

func (c NeutronSecurityGroupRuleDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NeutronSecurityGroupRuleDirection) UnmarshalJSON(b []byte) error {
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

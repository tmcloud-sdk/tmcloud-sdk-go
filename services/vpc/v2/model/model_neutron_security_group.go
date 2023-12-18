package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronSecurityGroup struct {
	Description string `json:"description"`

	Id string `json:"id"`

	Name string `json:"name"`

	SecurityGroupRules []NeutronSecurityGroupRule `json:"security_group_rules"`

	TenantId string `json:"tenant_id"`

	ProjectId string `json:"project_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NeutronSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronSecurityGroup struct{}"
	}

	return strings.Join([]string{"NeutronSecurityGroup", string(data)}, " ")
}

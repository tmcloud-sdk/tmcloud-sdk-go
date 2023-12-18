package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroup struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Id string `json:"id"`

	VpcId *string `json:"vpc_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}

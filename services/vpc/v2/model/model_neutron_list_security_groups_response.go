package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronListSecurityGroupsResponse struct {
	SecurityGroups *[]NeutronSecurityGroup `json:"security_groups,omitempty"`

	SecurityGroupsLinks *[]NeutronPageLink `json:"security_groups_links,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o NeutronListSecurityGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronListSecurityGroupsResponse struct{}"
	}

	return strings.Join([]string{"NeutronListSecurityGroupsResponse", string(data)}, " ")
}

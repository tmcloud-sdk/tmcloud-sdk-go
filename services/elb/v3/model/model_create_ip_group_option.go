package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateIpGroupOption struct {
	ProjectId *string `json:"project_id,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	IpList []CreateIpGroupIpOption `json:"ip_list"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CreateIpGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpGroupOption struct{}"
	}

	return strings.Join([]string{"CreateIpGroupOption", string(data)}, " ")
}

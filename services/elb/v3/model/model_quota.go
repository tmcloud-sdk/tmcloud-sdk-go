package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Quota struct {
	ProjectId string `json:"project_id"`

	Loadbalancer int32 `json:"loadbalancer"`

	Certificate int32 `json:"certificate"`

	Listener int32 `json:"listener"`

	L7policy int32 `json:"l7policy"`

	Pool int32 `json:"pool"`

	Healthmonitor int32 `json:"healthmonitor"`

	Member int32 `json:"member"`

	MembersPerPool int32 `json:"members_per_pool"`

	Ipgroup int32 `json:"ipgroup"`

	SecurityPolicy int32 `json:"security_policy"`

	IpgroupBindings string `json:"ipgroup_bindings"`

	IpgroupMaxLength string `json:"ipgroup_max_length"`
}

func (o Quota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Quota struct{}"
	}

	return strings.Join([]string{"Quota", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusPolicy struct {
	Action *string `json:"action,omitempty"`

	Id *string `json:"id,omitempty"`

	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	Name *string `json:"name,omitempty"`

	Rules *[]LoadBalancerStatusL7Rule `json:"rules,omitempty"`
}

func (o LoadBalancerStatusPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusPolicy struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusPolicy", string(data)}, " ")
}

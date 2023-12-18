package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LoadBalancerStatusL7Rule struct {
	Id string `json:"id"`

	Type string `json:"type"`

	ProvisioningStatus string `json:"provisioning_status"`
}

func (o LoadBalancerStatusL7Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusL7Rule struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusL7Rule", string(data)}, " ")
}

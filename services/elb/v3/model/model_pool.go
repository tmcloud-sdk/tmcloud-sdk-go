package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Pool struct {
	AdminStateUp bool `json:"admin_state_up"`

	Description string `json:"description"`

	HealthmonitorId string `json:"healthmonitor_id"`

	Id string `json:"id"`

	LbAlgorithm string `json:"lb_algorithm"`

	Listeners []ListenerRef `json:"listeners"`

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	Members []MemberRef `json:"members"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	Protocol string `json:"protocol"`

	SessionPersistence *SessionPersistence `json:"session_persistence"`

	IpVersion string `json:"ip_version"`

	SlowStart *SlowStart `json:"slow_start"`

	MemberDeletionProtectionEnable bool `json:"member_deletion_protection_enable"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	VpcId string `json:"vpc_id"`

	Type string `json:"type"`
}

func (o Pool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Pool struct{}"
	}

	return strings.Join([]string{"Pool", string(data)}, " ")
}

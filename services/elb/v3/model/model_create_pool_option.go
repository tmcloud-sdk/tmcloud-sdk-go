package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreatePoolOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	LbAlgorithm string `json:"lb_algorithm"`

	ListenerId *string `json:"listener_id,omitempty"`

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Protocol string `json:"protocol"`

	SessionPersistence *CreatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *CreatePoolSlowStartOption `json:"slow_start,omitempty"`

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o CreatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoolOption struct{}"
	}

	return strings.Join([]string{"CreatePoolOption", string(data)}, " ")
}

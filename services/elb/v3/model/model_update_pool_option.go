package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePoolOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	LbAlgorithm *string `json:"lb_algorithm,omitempty"`

	Name *string `json:"name,omitempty"`

	SessionPersistence *UpdatePoolSessionPersistenceOption `json:"session_persistence,omitempty"`

	SlowStart *UpdatePoolSlowStartOption `json:"slow_start,omitempty"`

	MemberDeletionProtectionEnable *bool `json:"member_deletion_protection_enable,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o UpdatePoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePoolOption struct{}"
	}

	return strings.Join([]string{"UpdatePoolOption", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceBackupPolicy struct {
	BackupPolicyId *string `json:"backup_policy_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Policy *BackupPolicy `json:"policy,omitempty"`

	TenantId *string `json:"tenant_id,omitempty"`
}

func (o InstanceBackupPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceBackupPolicy struct{}"
	}

	return strings.Join([]string{"InstanceBackupPolicy", string(data)}, " ")
}

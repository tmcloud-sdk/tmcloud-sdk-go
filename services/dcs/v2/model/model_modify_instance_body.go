package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyInstanceBody struct {
	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Port *int32 `json:"port,omitempty"`

	RenameCommands *RenameCommandResp `json:"rename_commands,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	InstanceBackupPolicy *BackupPolicy `json:"instance_backup_policy,omitempty"`
}

func (o ModifyInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyInstanceBody struct{}"
	}

	return strings.Join([]string{"ModifyInstanceBody", string(data)}, " ")
}

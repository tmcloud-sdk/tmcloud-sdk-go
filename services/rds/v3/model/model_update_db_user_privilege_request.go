package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateDbUserPrivilegeRequest struct {
	InstanceId string `json:"instance_id"`

	Body *DbUserPrivilegeRequest `json:"body,omitempty"`
}

func (o UpdateDbUserPrivilegeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserPrivilegeRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbUserPrivilegeRequest", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`

	Body *DeleteServerGroupMemberRequestBody `json:"body,omitempty"`
}

func (o DeleteServerGroupMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServerGroupMemberRequest struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupMemberRequest", string(data)}, " ")
}

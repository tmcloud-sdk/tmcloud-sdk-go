package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateDbUserCommentRequest struct {
	InstanceId string `json:"instance_id"`

	UserName string `json:"user_name"`

	Body *UpdateDbUserReq `json:"body,omitempty"`
}

func (o UpdateDbUserCommentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDbUserCommentRequest struct{}"
	}

	return strings.Join([]string{"UpdateDbUserCommentRequest", string(data)}, " ")
}

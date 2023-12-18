package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddMembersRequestBody struct {
	Images []string `json:"images"`

	Projects []string `json:"projects"`
}

func (o BatchAddMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAddMembersRequestBody", string(data)}, " ")
}

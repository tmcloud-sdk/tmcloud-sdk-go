package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDeleteMembersState struct {
	Id string `json:"id"`

	RetStatus string `json:"ret_status"`
}

func (o BatchDeleteMembersState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMembersState struct{}"
	}

	return strings.Join([]string{"BatchDeleteMembersState", string(data)}, " ")
}

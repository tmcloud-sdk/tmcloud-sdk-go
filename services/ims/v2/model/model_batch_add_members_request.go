package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddMembersRequest struct {
	Body *BatchAddMembersRequestBody `json:"body,omitempty"`
}

func (o BatchAddMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddMembersRequest struct{}"
	}

	return strings.Join([]string{"BatchAddMembersRequest", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdatePriorityRequestBody struct {
	Id string `json:"id"`

	Priority int32 `json:"priority"`
}

func (o BatchUpdatePriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePriorityRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePriorityRequestBody", string(data)}, " ")
}

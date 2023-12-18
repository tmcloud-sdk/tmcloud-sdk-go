package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SecurityGroupsResult struct {
	Id *string `json:"id,omitempty"`
}

func (o SecurityGroupsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupsResult struct{}"
	}

	return strings.Join([]string{"SecurityGroupsResult", string(data)}, " ")
}

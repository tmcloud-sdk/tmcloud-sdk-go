package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FollowerMigrateRequest struct {
	NodeId string `json:"nodeId"`

	AzCode string `json:"azCode"`
}

func (o FollowerMigrateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FollowerMigrateRequest struct{}"
	}

	return strings.Join([]string{"FollowerMigrateRequest", string(data)}, " ")
}

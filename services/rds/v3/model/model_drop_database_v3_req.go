package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DropDatabaseV3Req struct {
	IsForceDelete *bool `json:"is_force_delete,omitempty"`
}

func (o DropDatabaseV3Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DropDatabaseV3Req struct{}"
	}

	return strings.Join([]string{"DropDatabaseV3Req", string(data)}, " ")
}

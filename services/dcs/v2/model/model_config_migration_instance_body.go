package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ConfigMigrationInstanceBody struct {
	Id *string `json:"id,omitempty"`

	Addrs *string `json:"addrs,omitempty"`

	Password *string `json:"password,omitempty"`
}

func (o ConfigMigrationInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMigrationInstanceBody struct{}"
	}

	return strings.Join([]string{"ConfigMigrationInstanceBody", string(data)}, " ")
}

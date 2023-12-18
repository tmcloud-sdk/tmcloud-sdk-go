package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GrantRequest struct {
	DbName string `json:"db_name"`

	Users []UserWithPrivilege `json:"users"`
}

func (o GrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GrantRequest struct{}"
	}

	return strings.Join([]string{"GrantRequest", string(data)}, " ")
}

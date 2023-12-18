package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SqlserverRevokeRequest struct {
	DbName string `json:"db_name"`

	Users []SqlserverUserWithPrivilege `json:"users"`
}

func (o SqlserverRevokeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlserverRevokeRequest struct{}"
	}

	return strings.Join([]string{"SqlserverRevokeRequest", string(data)}, " ")
}

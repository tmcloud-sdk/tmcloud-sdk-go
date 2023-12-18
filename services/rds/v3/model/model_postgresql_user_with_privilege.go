package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlUserWithPrivilege struct {
	Name string `json:"name"`

	Readonly bool `json:"readonly"`

	SchemaName string `json:"schema_name"`
}

func (o PostgresqlUserWithPrivilege) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlUserWithPrivilege struct{}"
	}

	return strings.Join([]string{"PostgresqlUserWithPrivilege", string(data)}, " ")
}

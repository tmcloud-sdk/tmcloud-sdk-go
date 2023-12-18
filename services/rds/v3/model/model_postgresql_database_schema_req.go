package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlDatabaseSchemaReq struct {
	DbName string `json:"db_name"`

	Schemas []PostgresqlCreateSchemaReq `json:"schemas"`
}

func (o PostgresqlDatabaseSchemaReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlDatabaseSchemaReq struct{}"
	}

	return strings.Join([]string{"PostgresqlDatabaseSchemaReq", string(data)}, " ")
}

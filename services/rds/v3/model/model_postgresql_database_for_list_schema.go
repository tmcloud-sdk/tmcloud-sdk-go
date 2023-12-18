package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlDatabaseForListSchema struct {
	SchemaName string `json:"schema_name"`

	Owner string `json:"owner"`
}

func (o PostgresqlDatabaseForListSchema) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlDatabaseForListSchema struct{}"
	}

	return strings.Join([]string{"PostgresqlDatabaseForListSchema", string(data)}, " ")
}

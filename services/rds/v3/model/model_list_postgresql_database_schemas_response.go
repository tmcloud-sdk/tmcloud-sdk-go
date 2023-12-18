package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListPostgresqlDatabaseSchemasResponse struct {
	DatabaseSchemas *[]PostgresqlDatabaseForListSchema `json:"database_schemas,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPostgresqlDatabaseSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostgresqlDatabaseSchemasResponse struct{}"
	}

	return strings.Join([]string{"ListPostgresqlDatabaseSchemasResponse", string(data)}, " ")
}

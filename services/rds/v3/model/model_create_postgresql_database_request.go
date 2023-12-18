package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreatePostgresqlDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *PostgresqlDatabaseForCreation `json:"body,omitempty"`
}

func (o CreatePostgresqlDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostgresqlDatabaseRequest struct{}"
	}

	return strings.Join([]string{"CreatePostgresqlDatabaseRequest", string(data)}, " ")
}

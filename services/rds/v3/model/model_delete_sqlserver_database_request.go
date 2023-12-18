package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteSqlserverDatabaseRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db_name"`

	Body *DropDatabaseV3Req `json:"body,omitempty"`
}

func (o DeleteSqlserverDatabaseRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSqlserverDatabaseRequest struct{}"
	}

	return strings.Join([]string{"DeleteSqlserverDatabaseRequest", string(data)}, " ")
}

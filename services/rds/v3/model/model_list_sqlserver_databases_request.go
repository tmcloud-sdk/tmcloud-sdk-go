package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSqlserverDatabasesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`

	DbName *string `json:"db-name,omitempty"`
}

func (o ListSqlserverDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlserverDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListSqlserverDatabasesRequest", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAuthorizedSqlserverDbUsersRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	DbName string `json:"db-name"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListAuthorizedSqlserverDbUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedSqlserverDbUsersRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizedSqlserverDbUsersRequest", string(data)}, " ")
}

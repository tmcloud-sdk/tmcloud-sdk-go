package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSqlserverDbUserRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *SqlserverUserForCreation `json:"body,omitempty"`
}

func (o CreateSqlserverDbUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlserverDbUserRequest struct{}"
	}

	return strings.Join([]string{"CreateSqlserverDbUserRequest", string(data)}, " ")
}

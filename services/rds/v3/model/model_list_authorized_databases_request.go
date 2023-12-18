package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAuthorizedDatabasesRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	UserName string `json:"user-name"`

	Page int32 `json:"page"`

	Limit int32 `json:"limit"`
}

func (o ListAuthorizedDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedDatabasesRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorizedDatabasesRequest", string(data)}, " ")
}

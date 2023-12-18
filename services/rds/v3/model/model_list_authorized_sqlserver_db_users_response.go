package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAuthorizedSqlserverDbUsersResponse struct {
	Users *[]UserWithPrivilege `json:"users,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizedSqlserverDbUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizedSqlserverDbUsersResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizedSqlserverDbUsersResponse", string(data)}, " ")
}

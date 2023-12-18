package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlUserForList struct {
	Name string `json:"name"`

	Attributes *interface{} `json:"attributes,omitempty"`

	Memberof *[]string `json:"memberof,omitempty"`
}

func (o PostgresqlUserForList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlUserForList struct{}"
	}

	return strings.Join([]string{"PostgresqlUserForList", string(data)}, " ")
}

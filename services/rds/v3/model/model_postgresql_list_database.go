package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostgresqlListDatabase struct {
	Name *string `json:"name,omitempty"`

	Owner *string `json:"owner,omitempty"`

	CharacterSet *string `json:"character_set,omitempty"`

	CollateSet *string `json:"collate_set,omitempty"`

	Size *int64 `json:"size,omitempty"`
}

func (o PostgresqlListDatabase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostgresqlListDatabase struct{}"
	}

	return strings.Join([]string{"PostgresqlListDatabase", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DefaultRootDb struct {
	DbName *string `json:"db_name,omitempty"`

	DbEncoding *string `json:"db_encoding,omitempty"`
}

func (o DefaultRootDb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefaultRootDb struct{}"
	}

	return strings.Join([]string{"DefaultRootDb", string(data)}, " ")
}

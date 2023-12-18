package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DatabaseObjectInfo struct {
	Id *string `json:"id,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	Type *string `json:"type,omitempty"`

	Name *string `json:"name,omitempty"`

	AliasName *string `json:"alias_name,omitempty"`
}

func (o DatabaseObjectInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectInfo struct{}"
	}

	return strings.Join([]string{"DatabaseObjectInfo", string(data)}, " ")
}

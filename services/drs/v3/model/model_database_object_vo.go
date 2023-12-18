package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DatabaseObjectVo struct {
	Id *string `json:"id,omitempty"`

	Select *string `json:"select,omitempty"`
}

func (o DatabaseObjectVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatabaseObjectVo struct{}"
	}

	return strings.Join([]string{"DatabaseObjectVo", string(data)}, " ")
}

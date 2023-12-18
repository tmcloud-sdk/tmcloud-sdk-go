package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type VmMetaData struct {
	AdminPass *string `json:"admin_pass,omitempty"`
}

func (o VmMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VmMetaData struct{}"
	}

	return strings.Join([]string{"VmMetaData", string(data)}, " ")
}

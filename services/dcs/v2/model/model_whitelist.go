package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Whitelist struct {
	GroupName string `json:"group_name"`

	IpList []string `json:"ip_list"`
}

func (o Whitelist) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Whitelist struct{}"
	}

	return strings.Join([]string{"Whitelist", string(data)}, " ")
}

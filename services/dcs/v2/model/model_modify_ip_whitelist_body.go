package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyIpWhitelistBody struct {
	InstanceId *string `json:"instance_id,omitempty"`

	EnableWhitelist bool `json:"enable_whitelist"`

	Whitelist []Whitelist `json:"whitelist"`
}

func (o ModifyIpWhitelistBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyIpWhitelistBody struct{}"
	}

	return strings.Join([]string{"ModifyIpWhitelistBody", string(data)}, " ")
}

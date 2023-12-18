package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowIpWhitelistResponse struct {
	InstanceId *string `json:"instance_id,omitempty"`

	EnableWhitelist *bool `json:"enable_whitelist,omitempty"`

	Whitelist      *[]Whitelist `json:"whitelist,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowIpWhitelistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpWhitelistResponse struct{}"
	}

	return strings.Join([]string{"ShowIpWhitelistResponse", string(data)}, " ")
}

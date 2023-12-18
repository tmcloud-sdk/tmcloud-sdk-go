package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListenerIpGroup struct {
	IpgroupId string `json:"ipgroup_id"`

	EnableIpgroup bool `json:"enable_ipgroup"`

	Type string `json:"type"`
}

func (o ListenerIpGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerIpGroup struct{}"
	}

	return strings.Join([]string{"ListenerIpGroup", string(data)}, " ")
}

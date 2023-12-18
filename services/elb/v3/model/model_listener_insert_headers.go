package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListenerInsertHeaders struct {
	XForwardedELBIP *bool `json:"X-Forwarded-ELB-IP,omitempty"`

	XForwardedPort *bool `json:"X-Forwarded-Port,omitempty"`

	XForwardedForPort *bool `json:"X-Forwarded-For-Port,omitempty"`

	XForwardedHost *bool `json:"X-Forwarded-Host,omitempty"`
}

func (o ListenerInsertHeaders) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerInsertHeaders struct{}"
	}

	return strings.Join([]string{"ListenerInsertHeaders", string(data)}, " ")
}

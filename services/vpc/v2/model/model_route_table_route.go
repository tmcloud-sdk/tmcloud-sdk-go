package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RouteTableRoute struct {
	Type string `json:"type"`

	Destination string `json:"destination"`

	Nexthop string `json:"nexthop"`

	Description *string `json:"description,omitempty"`
}

func (o RouteTableRoute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RouteTableRoute struct{}"
	}

	return strings.Join([]string{"RouteTableRoute", string(data)}, " ")
}

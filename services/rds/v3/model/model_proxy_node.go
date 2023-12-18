package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ProxyNode struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Role string `json:"role"`

	AzCode string `json:"az_code"`

	Status string `json:"status"`

	FrozenFlag int32 `json:"frozen_flag"`
}

func (o ProxyNode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyNode struct{}"
	}

	return strings.Join([]string{"ProxyNode", string(data)}, " ")
}

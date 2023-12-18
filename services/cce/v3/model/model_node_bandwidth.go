package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodeBandwidth struct {
	Chargemode *string `json:"chargemode,omitempty"`

	Size *int32 `json:"size,omitempty"`

	Sharetype *string `json:"sharetype,omitempty"`
}

func (o NodeBandwidth) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeBandwidth struct{}"
	}

	return strings.Join([]string{"NodeBandwidth", string(data)}, " ")
}

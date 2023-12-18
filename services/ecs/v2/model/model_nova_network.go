package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaNetwork struct {
	Addr string `json:"addr"`

	Version int32 `json:"version"`

	OSEXTIPSMACmacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`

	OSEXTIPStype string `json:"OS-EXT-IPS:type"`
}

func (o NovaNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaNetwork struct{}"
	}

	return strings.Join([]string{"NovaNetwork", string(data)}, " ")
}

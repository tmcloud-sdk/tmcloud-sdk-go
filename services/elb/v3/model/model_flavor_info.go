package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FlavorInfo struct {
	Connection int32 `json:"connection"`

	Cps int32 `json:"cps"`

	Qps *int32 `json:"qps,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	Lcu *int32 `json:"lcu,omitempty"`

	HttpsCps *int32 `json:"https_cps,omitempty"`
}

func (o FlavorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorInfo struct{}"
	}

	return strings.Join([]string{"FlavorInfo", string(data)}, " ")
}

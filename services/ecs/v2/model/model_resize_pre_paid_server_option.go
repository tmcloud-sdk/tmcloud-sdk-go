package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResizePrePaidServerOption struct {
	FlavorRef string `json:"flavorRef"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	Extendparam *ResizeServerExtendParam `json:"extendparam,omitempty"`

	Mode *string `json:"mode,omitempty"`
}

func (o ResizePrePaidServerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePrePaidServerOption struct{}"
	}

	return strings.Join([]string{"ResizePrePaidServerOption", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResizePostPaidServerOption struct {
	FlavorRef string `json:"flavorRef"`

	Mode *string `json:"mode,omitempty"`
}

func (o ResizePostPaidServerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerOption struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerOption", string(data)}, " ")
}

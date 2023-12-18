package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResizeFlavorRequest struct {
	ResizeFlavor *ResizeFlavorObject `json:"resize_flavor"`
}

func (o ResizeFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeFlavorRequest struct{}"
	}

	return strings.Join([]string{"ResizeFlavorRequest", string(data)}, " ")
}

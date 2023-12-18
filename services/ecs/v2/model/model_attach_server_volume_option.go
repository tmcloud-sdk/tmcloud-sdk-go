package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AttachServerVolumeOption struct {
	Device *string `json:"device,omitempty"`

	VolumeId string `json:"volumeId"`

	VolumeType *string `json:"volume_type,omitempty"`

	Count *int32 `json:"count,omitempty"`

	Hwpassthrough *string `json:"hw:passthrough,omitempty"`
}

func (o AttachServerVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeOption struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeOption", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerExtendVolumeAttachment struct {
	Id string `json:"id"`

	DeleteOnTermination string `json:"delete_on_termination"`

	BootIndex *string `json:"bootIndex,omitempty"`

	Device string `json:"device"`
}

func (o ServerExtendVolumeAttachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerExtendVolumeAttachment struct{}"
	}

	return strings.Join([]string{"ServerExtendVolumeAttachment", string(data)}, " ")
}

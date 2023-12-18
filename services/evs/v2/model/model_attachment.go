package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Attachment struct {
	AttachedAt string `json:"attached_at"`

	AttachmentId string `json:"attachment_id"`

	Device string `json:"device"`

	HostName string `json:"host_name"`

	Id string `json:"id"`

	ServerId string `json:"server_id"`

	VolumeId string `json:"volume_id"`
}

func (o Attachment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Attachment struct{}"
	}

	return strings.Join([]string{"Attachment", string(data)}, " ")
}

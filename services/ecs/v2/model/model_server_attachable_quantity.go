package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerAttachableQuantity struct {
	FreeScsi int32 `json:"free_scsi"`

	FreeBlk int32 `json:"free_blk"`

	FreeDisk int32 `json:"free_disk"`

	FreeNic int32 `json:"free_nic"`
}

func (o ServerAttachableQuantity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAttachableQuantity struct{}"
	}

	return strings.Join([]string{"ServerAttachableQuantity", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type VolumeTransferSummary struct {
	Id string `json:"id"`

	Links []Link `json:"links"`

	Name string `json:"name"`

	VolumeId string `json:"volume_id"`
}

func (o VolumeTransferSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTransferSummary struct{}"
	}

	return strings.Join([]string{"VolumeTransferSummary", string(data)}, " ")
}

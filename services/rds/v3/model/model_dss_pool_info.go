package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DssPoolInfo struct {
	AzName string `json:"az_name"`

	FreeCapacityGb string `json:"free_capacity_gb"`

	DsspoolVolumeType string `json:"dsspool_volume_type"`

	DsspoolId string `json:"dsspool_id"`

	DsspoolStatus string `json:"dsspool_status"`
}

func (o DssPoolInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DssPoolInfo struct{}"
	}

	return strings.Join([]string{"DssPoolInfo", string(data)}, " ")
}

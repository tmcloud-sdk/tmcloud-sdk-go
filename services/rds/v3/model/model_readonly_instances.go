package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ReadonlyInstances struct {
	Id string `json:"id"`

	Status string `json:"status"`

	Name string `json:"name"`

	Weight int32 `json:"weight"`

	AvailableZones []AvailableZone `json:"available_zones"`

	CpuNum int32 `json:"cpu_num"`
}

func (o ReadonlyInstances) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadonlyInstances struct{}"
	}

	return strings.Join([]string{"ReadonlyInstances", string(data)}, " ")
}

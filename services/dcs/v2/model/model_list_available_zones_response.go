package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAvailableZonesResponse struct {
	RegionId *string `json:"region_id,omitempty"`

	AvailableZones *[]AvailableZones `json:"available_zones,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}

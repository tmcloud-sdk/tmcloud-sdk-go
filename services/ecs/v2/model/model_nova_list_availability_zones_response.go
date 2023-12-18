package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaListAvailabilityZonesResponse struct {
	AvailabilityZoneInfo *[]NovaAvailabilityZone `json:"availabilityZoneInfo,omitempty"`
	HttpStatusCode       int                     `json:"-"`
}

func (o NovaListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"NovaListAvailabilityZonesResponse", string(data)}, " ")
}

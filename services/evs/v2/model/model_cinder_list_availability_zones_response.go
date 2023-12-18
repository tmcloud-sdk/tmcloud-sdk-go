package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderListAvailabilityZonesResponse struct {
	AvailabilityZoneInfo *[]AzInfo `json:"availabilityZoneInfo,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o CinderListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"CinderListAvailabilityZonesResponse", string(data)}, " ")
}

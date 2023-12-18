package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaShowServerRequest struct {
	ServerId string `json:"server_id"`

	OpenStackAPIVersion *string `json:"OpenStack-API-Version,omitempty"`
}

func (o NovaShowServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowServerRequest struct{}"
	}

	return strings.Join([]string{"NovaShowServerRequest", string(data)}, " ")
}

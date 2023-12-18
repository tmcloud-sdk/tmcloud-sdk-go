package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAvailableZoneResponse struct {
	AzInfos        *[]AzInfoResp `json:"az_infos,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAvailableZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZoneResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZoneResponse", string(data)}, " ")
}

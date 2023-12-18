package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowServerPasswordRequest struct {
	ServerId string `json:"server_id"`
}

func (o ShowServerPasswordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"ShowServerPasswordRequest", string(data)}, " ")
}

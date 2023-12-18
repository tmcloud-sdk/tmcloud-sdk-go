package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderDeleteVolumeTransferRequest struct {
	TransferId string `json:"transfer_id"`
}

func (o CinderDeleteVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderDeleteVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderDeleteVolumeTransferRequest", string(data)}, " ")
}

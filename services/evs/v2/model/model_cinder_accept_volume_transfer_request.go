package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderAcceptVolumeTransferRequest struct {
	TransferId string `json:"transfer_id"`

	Body *CinderAcceptVolumeTransferRequestBody `json:"body,omitempty"`
}

func (o CinderAcceptVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferRequest", string(data)}, " ")
}

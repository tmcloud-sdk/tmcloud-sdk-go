package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderCreateVolumeTransferRequest struct {
	Body *CinderCreateVolumeTransferRequestBody `json:"body,omitempty"`
}

func (o CinderCreateVolumeTransferRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderCreateVolumeTransferRequest struct{}"
	}

	return strings.Join([]string{"CinderCreateVolumeTransferRequest", string(data)}, " ")
}

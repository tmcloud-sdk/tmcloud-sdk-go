package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderAcceptVolumeTransferRequestBody struct {
	Accept *CinderAcceptVolumeTransferOption `json:"accept"`
}

func (o CinderAcceptVolumeTransferRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderAcceptVolumeTransferRequestBody struct{}"
	}

	return strings.Join([]string{"CinderAcceptVolumeTransferRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CinderListVolumeTransfersRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Offset *int32 `json:"offset,omitempty"`
}

func (o CinderListVolumeTransfersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListVolumeTransfersRequest struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTransfersRequest", string(data)}, " ")
}

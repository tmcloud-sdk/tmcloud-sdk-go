package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAttachSharableVolumesRequest struct {
	VolumeId string `json:"volume_id"`

	Body *BatchAttachSharableVolumesRequestBody `json:"body,omitempty"`
}

func (o BatchAttachSharableVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequest struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequest", string(data)}, " ")
}

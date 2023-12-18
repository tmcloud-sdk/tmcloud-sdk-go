package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAttachSharableVolumesRequestBody struct {
	Serverinfo []BatchAttachSharableVolumesOption `json:"serverinfo"`
}

func (o BatchAttachSharableVolumesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesRequestBody struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesRequestBody", string(data)}, " ")
}

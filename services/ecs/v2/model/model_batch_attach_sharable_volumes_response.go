package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAttachSharableVolumesResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAttachSharableVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAttachSharableVolumesResponse struct{}"
	}

	return strings.Join([]string{"BatchAttachSharableVolumesResponse", string(data)}, " ")
}

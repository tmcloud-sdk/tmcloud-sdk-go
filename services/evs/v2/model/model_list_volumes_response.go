package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListVolumesResponse struct {
	Count *int32 `json:"count,omitempty"`

	VolumesLinks *[]Link `json:"volumes_links,omitempty"`

	Volumes        *[]VolumeDetail `json:"volumes,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}

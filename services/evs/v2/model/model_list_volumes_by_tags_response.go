package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListVolumesByTagsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Resources      *[]Resource `json:"resources,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListVolumesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesByTagsResponse", string(data)}, " ")
}

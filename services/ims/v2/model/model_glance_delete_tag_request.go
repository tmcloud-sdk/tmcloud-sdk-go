package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteTagRequest struct {
	ImageId string `json:"image_id"`

	Tag string `json:"tag"`
}

func (o GlanceDeleteTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteTagRequest", string(data)}, " ")
}

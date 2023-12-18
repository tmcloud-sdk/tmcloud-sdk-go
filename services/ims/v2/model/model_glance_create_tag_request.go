package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceCreateTagRequest struct {
	ImageId string `json:"image_id"`

	Tag string `json:"tag"`
}

func (o GlanceCreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceCreateTagRequest struct{}"
	}

	return strings.Join([]string{"GlanceCreateTagRequest", string(data)}, " ")
}

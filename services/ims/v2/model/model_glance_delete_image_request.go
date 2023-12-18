package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceDeleteImageRequest struct {
	ImageId string `json:"image_id"`

	Body *GlanceDeleteImageRequestBody `json:"body,omitempty"`
}

func (o GlanceDeleteImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceDeleteImageRequest struct{}"
	}

	return strings.Join([]string{"GlanceDeleteImageRequest", string(data)}, " ")
}

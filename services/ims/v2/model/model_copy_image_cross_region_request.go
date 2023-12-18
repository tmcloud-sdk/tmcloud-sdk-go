package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyImageCrossRegionRequest struct {
	ImageId string `json:"image_id"`

	Body *CopyImageCrossRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageCrossRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequest", string(data)}, " ")
}

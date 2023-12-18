package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyImageInRegionRequest struct {
	ImageId string `json:"image_id"`

	Body *CopyImageInRegionRequestBody `json:"body,omitempty"`
}

func (o CopyImageInRegionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionRequest struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionRequest", string(data)}, " ")
}

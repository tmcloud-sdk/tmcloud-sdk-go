package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateImageRequest struct {
	ImageId string `json:"image_id"`

	Body *[]UpdateImageRequestBody `json:"body,omitempty"`
}

func (o UpdateImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageRequest struct{}"
	}

	return strings.Join([]string{"UpdateImageRequest", string(data)}, " ")
}

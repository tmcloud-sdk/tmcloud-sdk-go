package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RegisterImageRequest struct {
	ImageId string `json:"image_id"`

	Body *RegisterImageRequestBody `json:"body,omitempty"`
}

func (o RegisterImageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterImageRequest struct{}"
	}

	return strings.Join([]string{"RegisterImageRequest", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AddOrUpdateTagsRequestBody struct {
	ImageId string `json:"image_id"`

	Tag *string `json:"tag,omitempty"`

	ImageTag *ResourceTag `json:"image_tag,omitempty"`
}

func (o AddOrUpdateTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrUpdateTagsRequestBody struct{}"
	}

	return strings.Join([]string{"AddOrUpdateTagsRequestBody", string(data)}, " ")
}

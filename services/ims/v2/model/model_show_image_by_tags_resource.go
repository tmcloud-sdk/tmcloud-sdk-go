package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowImageByTagsResource struct {
	ResourceId string `json:"resource_id"`

	ResourceDetail *QueryImageByTagsResourceDetail `json:"resource_detail"`

	Tags []TagKeyValue `json:"tags"`

	ResourceName string `json:"resource_name"`
}

func (o ShowImageByTagsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowImageByTagsResource struct{}"
	}

	return strings.Join([]string{"ShowImageByTagsResource", string(data)}, " ")
}

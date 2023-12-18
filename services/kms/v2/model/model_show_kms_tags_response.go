package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowKmsTagsResponse struct {
	Tags *[]TagItem `json:"tags,omitempty"`

	ExistTagsNum   *int32 `json:"existTagsNum,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowKmsTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowKmsTagsResponse", string(data)}, " ")
}

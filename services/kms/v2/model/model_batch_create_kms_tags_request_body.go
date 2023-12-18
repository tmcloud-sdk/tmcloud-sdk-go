package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateKmsTagsRequestBody struct {
	Tags *[]TagItem `json:"tags,omitempty"`

	Action *string `json:"action,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o BatchCreateKmsTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateKmsTagsRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateKmsTagsRequestBody", string(data)}, " ")
}

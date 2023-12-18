package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListKmsByTagsRequestBody struct {
	Limit *string `json:"limit,omitempty"`

	Offset *string `json:"offset,omitempty"`

	Action *string `json:"action,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	Matches *[]TagItem `json:"matches,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ListKmsByTagsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsByTagsRequestBody struct{}"
	}

	return strings.Join([]string{"ListKmsByTagsRequestBody", string(data)}, " ")
}

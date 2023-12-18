package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListGrantsRequestBody struct {
	KeyId string `json:"key_id"`

	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ListGrantsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGrantsRequestBody struct{}"
	}

	return strings.Join([]string{"ListGrantsRequestBody", string(data)}, " ")
}

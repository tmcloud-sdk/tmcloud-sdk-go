package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRetirableGrantsRequestBody struct {
	Limit *string `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ListRetirableGrantsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsRequestBody struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsRequestBody", string(data)}, " ")
}

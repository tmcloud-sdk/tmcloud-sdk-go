package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListTagsOfTenantResponse struct {
	Tags           *[]Tag `json:"tags,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTagsOfTenantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTagsOfTenantResponse struct{}"
	}

	return strings.Join([]string{"ListTagsOfTenantResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListKmsByTagsResponse struct {
	Resources *[]ActionResources `json:"resources,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListKmsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListKmsByTagsResponse", string(data)}, " ")
}

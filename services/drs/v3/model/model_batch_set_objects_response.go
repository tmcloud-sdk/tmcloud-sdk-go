package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchSetObjectsResponse struct {
	AllCounts *int64 `json:"all_counts,omitempty"`

	Results        *[]DatabaseObjectResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchSetObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetObjectsResponse struct{}"
	}

	return strings.Join([]string{"BatchSetObjectsResponse", string(data)}, " ")
}

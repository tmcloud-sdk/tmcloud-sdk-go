package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowJobListResponse struct {
	TotalRecord *int32 `json:"total_record,omitempty"`

	Jobs           *[]JobInfo `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowJobListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobListResponse struct{}"
	}

	return strings.Join([]string{"ShowJobListResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateJobsResponse struct {
	Results *[]CreateJobResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchCreateJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateJobsResponse", string(data)}, " ")
}

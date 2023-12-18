package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchStopJobsResponse struct {
	Results *[]PauseJobResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchStopJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStopJobsResponse struct{}"
	}

	return strings.Join([]string{"BatchStopJobsResponse", string(data)}, " ")
}

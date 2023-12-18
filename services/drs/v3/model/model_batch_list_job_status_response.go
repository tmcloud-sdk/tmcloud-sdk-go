package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchListJobStatusResponse struct {
	Results *[]QueryJobStatusResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchListJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListJobStatusResponse struct{}"
	}

	return strings.Join([]string{"BatchListJobStatusResponse", string(data)}, " ")
}

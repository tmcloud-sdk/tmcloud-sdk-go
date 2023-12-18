package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ResizePostPaidServerResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResizePostPaidServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerResponse struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerResponse", string(data)}, " ")
}

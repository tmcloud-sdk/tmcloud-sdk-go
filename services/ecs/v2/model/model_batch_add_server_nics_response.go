package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchAddServerNicsResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddServerNicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddServerNicsResponse struct{}"
	}

	return strings.Join([]string{"BatchAddServerNicsResponse", string(data)}, " ")
}

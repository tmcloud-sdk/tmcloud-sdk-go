package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateDataImageResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataImageResponse struct{}"
	}

	return strings.Join([]string{"CreateDataImageResponse", string(data)}, " ")
}

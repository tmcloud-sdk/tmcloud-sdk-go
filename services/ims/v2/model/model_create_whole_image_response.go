package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateWholeImageResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWholeImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageResponse struct{}"
	}

	return strings.Join([]string{"CreateWholeImageResponse", string(data)}, " ")
}

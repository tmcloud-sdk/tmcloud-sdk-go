package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StartupInstanceResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartupInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartupInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartupInstanceResponse", string(data)}, " ")
}

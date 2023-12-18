package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StartInstanceRestartActionResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartInstanceRestartActionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceRestartActionResponse struct{}"
	}

	return strings.Join([]string{"StartInstanceRestartActionResponse", string(data)}, " ")
}

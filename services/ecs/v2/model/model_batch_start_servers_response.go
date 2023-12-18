package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchStartServersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchStartServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchStartServersResponse struct{}"
	}

	return strings.Join([]string{"BatchStartServersResponse", string(data)}, " ")
}

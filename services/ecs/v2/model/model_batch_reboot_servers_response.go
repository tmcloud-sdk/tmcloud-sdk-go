package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchRebootServersResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchRebootServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRebootServersResponse struct{}"
	}

	return strings.Join([]string{"BatchRebootServersResponse", string(data)}, " ")
}

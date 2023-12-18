package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MigrateServerResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateServerResponse struct{}"
	}

	return strings.Join([]string{"MigrateServerResponse", string(data)}, " ")
}

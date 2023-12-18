package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdatePostgresqlParameterValueResponse struct {
	JobId *string `json:"job_id,omitempty"`

	RestartRequired *bool `json:"restart_required,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o UpdatePostgresqlParameterValueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePostgresqlParameterValueResponse struct{}"
	}

	return strings.Join([]string{"UpdatePostgresqlParameterValueResponse", string(data)}, " ")
}

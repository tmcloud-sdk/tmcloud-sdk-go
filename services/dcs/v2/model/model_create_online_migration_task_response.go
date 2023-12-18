package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateOnlineMigrationTaskResponse struct {
	InstanceId     *string `json:"instance_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateOnlineMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskResponse", string(data)}, " ")
}

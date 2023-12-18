package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreatePostPaidServersResponse struct {
	JobId *string `json:"job_id,omitempty"`

	ServerIds      *[]string `json:"serverIds,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreatePostPaidServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePostPaidServersResponse struct{}"
	}

	return strings.Join([]string{"CreatePostPaidServersResponse", string(data)}, " ")
}

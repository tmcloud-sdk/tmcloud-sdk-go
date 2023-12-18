package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithCloudInitResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeServerOsWithCloudInitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitResponse struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitResponse", string(data)}, " ")
}

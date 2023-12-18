package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ReinstallServerWithoutCloudInitResponse struct {
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ReinstallServerWithoutCloudInitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithoutCloudInitResponse struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithoutCloudInitResponse", string(data)}, " ")
}

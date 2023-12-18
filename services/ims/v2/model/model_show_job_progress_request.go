package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowJobProgressRequest struct {
	JobId string `json:"job_id"`
}

func (o ShowJobProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobProgressRequest struct{}"
	}

	return strings.Join([]string{"ShowJobProgressRequest", string(data)}, " ")
}

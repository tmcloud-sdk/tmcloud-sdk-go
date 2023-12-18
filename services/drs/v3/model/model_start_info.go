package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type StartInfo struct {
	JobId string `json:"job_id"`

	StartTime *string `json:"start_time,omitempty"`
}

func (o StartInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInfo struct{}"
	}

	return strings.Join([]string{"StartInfo", string(data)}, " ")
}

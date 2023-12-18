package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ProgressInfo struct {
	Completed *string `json:"completed,omitempty"`

	RemainingTime *string `json:"remaining_time,omitempty"`
}

func (o ProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressInfo struct{}"
	}

	return strings.Join([]string{"ProgressInfo", string(data)}, " ")
}

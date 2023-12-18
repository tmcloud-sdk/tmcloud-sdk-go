package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SelectedSetAlarmTaskReq struct {
	JobId string `json:"job_id"`

	Status string `json:"status"`

	EngineType string `json:"engine_type"`
}

func (o SelectedSetAlarmTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedSetAlarmTaskReq struct{}"
	}

	return strings.Join([]string{"SelectedSetAlarmTaskReq", string(data)}, " ")
}

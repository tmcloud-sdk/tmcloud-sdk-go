package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchImportSmnInfoReq struct {
	Jobs []SelectedSetAlarmTaskReq `json:"jobs"`

	AlarmNotifyInfo *BatchSetAlarmNotifyInfo `json:"alarm_notify_info"`
}

func (o BatchImportSmnInfoReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchImportSmnInfoReq struct{}"
	}

	return strings.Join([]string{"BatchImportSmnInfoReq", string(data)}, " ")
}

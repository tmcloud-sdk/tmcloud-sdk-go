package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowLog struct {
	Count string `json:"count"`

	Time string `json:"time"`

	LockTime string `json:"lock_time"`

	RowsSent string `json:"rows_sent"`

	RowsExamined string `json:"rows_examined"`

	Database string `json:"database"`

	Users string `json:"users"`

	QuerySample string `json:"query_sample"`

	Type string `json:"type"`

	StartTime string `json:"start_time"`

	ClientIp string `json:"client_ip"`
}

func (o SlowLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLog struct{}"
	}

	return strings.Join([]string{"SlowLog", string(data)}, " ")
}

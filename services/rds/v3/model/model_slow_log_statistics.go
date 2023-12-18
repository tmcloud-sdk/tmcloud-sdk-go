package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowLogStatistics struct {
	Count string `json:"count"`

	Time string `json:"time"`

	LockTime string `json:"lockTime"`

	RowsSent int64 `json:"rowsSent"`

	RowsExamined int64 `json:"rowsExamined"`

	Database string `json:"database"`

	Users string `json:"users"`

	QuerySample string `json:"querySample"`

	Type string `json:"type"`

	ClientIP string `json:"clientIP"`
}

func (o SlowLogStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowLogStatistics struct{}"
	}

	return strings.Join([]string{"SlowLogStatistics", string(data)}, " ")
}

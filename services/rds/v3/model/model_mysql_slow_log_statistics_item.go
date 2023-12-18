package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MysqlSlowLogStatisticsItem struct {
	Count *string `json:"count,omitempty"`

	Time *string `json:"time,omitempty"`

	LockTime *string `json:"lock_time,omitempty"`

	RowsSent *int64 `json:"rows_sent,omitempty"`

	RowsExamined *int64 `json:"rows_examined,omitempty"`

	Database *string `json:"database,omitempty"`

	Users *string `json:"users,omitempty"`

	QuerySample *string `json:"query_sample,omitempty"`

	ClientIp *string `json:"client_ip,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o MysqlSlowLogStatisticsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlSlowLogStatisticsItem struct{}"
	}

	return strings.Join([]string{"MysqlSlowLogStatisticsItem", string(data)}, " ")
}

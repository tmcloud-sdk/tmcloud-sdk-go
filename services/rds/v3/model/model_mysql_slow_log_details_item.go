package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MysqlSlowLogDetailsItem struct {
	Count *string `json:"count,omitempty"`

	Time *string `json:"time,omitempty"`

	LockTime *string `json:"lock_time,omitempty"`

	RowsSent *string `json:"rows_sent,omitempty"`

	RowsExamined *string `json:"rows_examined,omitempty"`

	Database *string `json:"database,omitempty"`

	Users *string `json:"users,omitempty"`

	QuerySample *string `json:"query_sample,omitempty"`

	Type *string `json:"type,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	ClientIp *string `json:"client_ip,omitempty"`

	LineNum *string `json:"line_num,omitempty"`
}

func (o MysqlSlowLogDetailsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlSlowLogDetailsItem struct{}"
	}

	return strings.Join([]string{"MysqlSlowLogDetailsItem", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSlowLogStatisticsForLtsResponse struct {
	SlowLogList *[]MysqlSlowLogStatisticsItem `json:"slow_log_list,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowLogStatisticsForLtsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogStatisticsForLtsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowLogStatisticsForLtsResponse", string(data)}, " ")
}

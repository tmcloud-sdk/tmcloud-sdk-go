package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSlowlogStatisticsResponse struct {
	PageNumber *int32 `json:"pageNumber,omitempty"`

	PageRecord *int32 `json:"pageRecord,omitempty"`

	SlowLogList *[]SlowLogStatistics `json:"slowLogList,omitempty"`

	TotalRecord *int32 `json:"totalRecord,omitempty"`

	StartTime *int64 `json:"startTime,omitempty"`

	EndTime        *int64 `json:"endTime,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSlowlogStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowlogStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSlowlogStatisticsResponse", string(data)}, " ")
}

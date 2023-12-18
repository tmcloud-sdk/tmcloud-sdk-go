package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryDataGuardMonitorAndChartResp struct {
	Id string `json:"id"`

	DataGuardMinitor *QueryDataGuardMonitorResponse `json:"data_guard_minitor"`
}

func (o QueryDataGuardMonitorAndChartResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDataGuardMonitorAndChartResp struct{}"
	}

	return strings.Join([]string{"QueryDataGuardMonitorAndChartResp", string(data)}, " ")
}

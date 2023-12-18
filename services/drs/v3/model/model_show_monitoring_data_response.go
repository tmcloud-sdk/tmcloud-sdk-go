package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowMonitoringDataResponse struct {
	Results *[]QueryDataGuardMonitorAndChartResp `json:"results,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMonitoringDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMonitoringDataResponse struct{}"
	}

	return strings.Join([]string{"ShowMonitoringDataResponse", string(data)}, " ")
}

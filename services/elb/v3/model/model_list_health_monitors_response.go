package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListHealthMonitorsResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	Healthmonitors *[]HealthMonitor `json:"healthmonitors,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListHealthMonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthMonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListHealthMonitorsResponse", string(data)}, " ")
}

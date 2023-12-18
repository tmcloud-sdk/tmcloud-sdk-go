package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMsdtcHostsResponse struct {
	TotalCount *int32 `json:"total_count,omitempty"`

	Hosts          *[]DbsInstanceHostInfoResult `json:"hosts,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMsdtcHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMsdtcHostsResponse struct{}"
	}

	return strings.Join([]string{"ListMsdtcHostsResponse", string(data)}, " ")
}

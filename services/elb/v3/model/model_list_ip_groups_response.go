package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListIpGroupsResponse struct {
	Ipgroups *[]IpGroup `json:"ipgroups,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListIpGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListIpGroupsResponse", string(data)}, " ")
}

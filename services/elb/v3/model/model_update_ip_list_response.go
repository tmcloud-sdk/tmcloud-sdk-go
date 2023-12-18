package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateIpListResponse struct {
	Ipgroup *IpGroup `json:"ipgroup,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpListResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpListResponse", string(data)}, " ")
}

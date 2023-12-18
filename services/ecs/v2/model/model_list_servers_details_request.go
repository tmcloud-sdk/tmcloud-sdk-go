package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListServersDetailsRequest struct {
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Flavor *string `json:"flavor,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`

	NotTags *string `json:"not-tags,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	ReservationId *string `json:"reservation_id,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags *string `json:"tags,omitempty"`

	IpEq *string `json:"ip_eq,omitempty"`
}

func (o ListServersDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListServersDetailsRequest", string(data)}, " ")
}

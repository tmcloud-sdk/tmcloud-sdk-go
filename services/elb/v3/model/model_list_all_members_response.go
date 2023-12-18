package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAllMembersResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	Members        *[]Member `json:"members,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllMembersResponse struct{}"
	}

	return strings.Join([]string{"ListAllMembersResponse", string(data)}, " ")
}

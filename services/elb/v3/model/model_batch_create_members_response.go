package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchCreateMembersResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Members        *[]BatchMember `json:"members,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o BatchCreateMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateMembersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateMembersResponse", string(data)}, " ")
}

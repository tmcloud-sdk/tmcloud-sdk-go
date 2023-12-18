package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListL7RulesResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	Rules          *[]L7Rule `json:"rules,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListL7RulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesResponse struct{}"
	}

	return strings.Join([]string{"ListL7RulesResponse", string(data)}, " ")
}

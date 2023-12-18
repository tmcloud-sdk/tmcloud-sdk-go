package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7RuleResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7RuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateL7RuleResponse", string(data)}, " ")
}

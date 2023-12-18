package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateL7RuleResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Rule           *L7Rule `json:"rule,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateL7RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7RuleResponse struct{}"
	}

	return strings.Join([]string{"CreateL7RuleResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowL7PolicyRequest struct {
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7PolicyRequest", string(data)}, " ")
}

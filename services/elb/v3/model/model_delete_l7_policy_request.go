package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteL7PolicyRequest struct {
	L7policyId string `json:"l7policy_id"`
}

func (o DeleteL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7PolicyRequest", string(data)}, " ")
}

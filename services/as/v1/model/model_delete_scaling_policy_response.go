package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingPolicyResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7PolicyRequestBody struct {
	L7policy *UpdateL7PolicyOption `json:"l7policy"`
}

func (o UpdateL7PolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyRequestBody", string(data)}, " ")
}

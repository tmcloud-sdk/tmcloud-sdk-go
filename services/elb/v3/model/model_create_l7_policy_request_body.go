package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateL7PolicyRequestBody struct {
	L7policy *CreateL7PolicyOption `json:"l7policy"`
}

func (o CreateL7PolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyRequestBody", string(data)}, " ")
}

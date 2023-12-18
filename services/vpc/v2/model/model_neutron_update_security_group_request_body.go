package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NeutronUpdateSecurityGroupRequestBody struct {
	SecurityGroup *NeutronUpdateSecurityGroupOption `json:"security_group"`
}

func (o NeutronUpdateSecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronUpdateSecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"NeutronUpdateSecurityGroupRequestBody", string(data)}, " ")
}

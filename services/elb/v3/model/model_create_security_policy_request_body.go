package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityPolicyRequestBody struct {
	SecurityPolicy *CreateSecurityPolicyOption `json:"security_policy"`
}

func (o CreateSecurityPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyRequestBody", string(data)}, " ")
}

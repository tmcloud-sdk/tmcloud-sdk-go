package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSecurityPolicyRequest struct {
	Body *CreateSecurityPolicyRequestBody `json:"body,omitempty"`
}

func (o CreateSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"CreateSecurityPolicyRequest", string(data)}, " ")
}

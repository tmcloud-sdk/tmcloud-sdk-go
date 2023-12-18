package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSystemSecurityPoliciesResponse struct {
	SystemSecurityPolicies *[]SystemSecurityPolicy `json:"system_security_policies,omitempty"`

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSystemSecurityPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemSecurityPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSystemSecurityPoliciesResponse", string(data)}, " ")
}

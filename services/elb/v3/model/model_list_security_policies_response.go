package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityPoliciesResponse struct {
	SecurityPolicies *[]SecurityPolicy `json:"security_policies,omitempty"`

	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSecurityPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityPoliciesResponse", string(data)}, " ")
}

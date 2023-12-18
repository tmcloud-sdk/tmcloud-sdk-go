package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSecurityPoliciesRequest struct {
	Marker *string `json:"marker,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Description *[]string `json:"description,omitempty"`

	Protocols *[]string `json:"protocols,omitempty"`

	Ciphers *[]string `json:"ciphers,omitempty"`
}

func (o ListSecurityPoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityPoliciesRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityPoliciesRequest", string(data)}, " ")
}

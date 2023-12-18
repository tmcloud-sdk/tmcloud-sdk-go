package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListL7RulesRequest struct {
	L7policyId string `json:"l7policy_id"`

	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	Id *[]string `json:"id,omitempty"`

	CompareType *[]string `json:"compare_type,omitempty"`

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`

	Invert *bool `json:"invert,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Value *[]string `json:"value,omitempty"`

	Key *[]string `json:"key,omitempty"`

	Type *[]string `json:"type,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListL7RulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7RulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7RulesRequest", string(data)}, " ")
}

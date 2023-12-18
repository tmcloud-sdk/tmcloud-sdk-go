package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateL7PolicyOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	RedirectUrlConfig *UpdateRedirectUrlConfig `json:"redirect_url_config,omitempty"`

	FixedResponseConfig *UpdateFixtedResponseConfig `json:"fixed_response_config,omitempty"`

	Rules *[]CreateRuleOption `json:"rules,omitempty"`

	Priority *int32 `json:"priority,omitempty"`

	RedirectPoolsConfig *[]CreateRedirectPoolsConfig `json:"redirect_pools_config,omitempty"`
}

func (o UpdateL7PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"UpdateL7PolicyOption", string(data)}, " ")
}

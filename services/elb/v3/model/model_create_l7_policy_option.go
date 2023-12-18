package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateL7PolicyOption struct {
	Action string `json:"action"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Description *string `json:"description,omitempty"`

	ListenerId string `json:"listener_id"`

	Name *string `json:"name,omitempty"`

	Position *int32 `json:"position,omitempty"`

	Priority *int32 `json:"priority,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	RedirectListenerId *string `json:"redirect_listener_id,omitempty"`

	RedirectPoolId *string `json:"redirect_pool_id,omitempty"`

	RedirectPoolsConfig *[]CreateRedirectPoolsConfig `json:"redirect_pools_config,omitempty"`

	RedirectUrl *string `json:"redirect_url,omitempty"`

	RedirectUrlConfig *CreateRedirectUrlConfig `json:"redirect_url_config,omitempty"`

	FixedResponseConfig *CreateFixtedResponseConfig `json:"fixed_response_config,omitempty"`

	Rules *[]CreateL7PolicyRuleOption `json:"rules,omitempty"`
}

func (o CreateL7PolicyOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateL7PolicyOption struct{}"
	}

	return strings.Join([]string{"CreateL7PolicyOption", string(data)}, " ")
}

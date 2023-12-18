package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type L7Policy struct {
	Action string `json:"action"`

	AdminStateUp bool `json:"admin_state_up"`

	Description string `json:"description"`

	Id string `json:"id"`

	ListenerId string `json:"listener_id"`

	Name string `json:"name"`

	Position int32 `json:"position"`

	Priority *int32 `json:"priority,omitempty"`

	ProjectId string `json:"project_id"`

	ProvisioningStatus string `json:"provisioning_status"`

	RedirectPoolId string `json:"redirect_pool_id"`

	RedirectPoolsConfig []CreateRedirectPoolsConfig `json:"redirect_pools_config"`

	RedirectListenerId string `json:"redirect_listener_id"`

	RedirectUrl string `json:"redirect_url"`

	Rules []RuleRef `json:"rules"`

	RedirectUrlConfig *RedirectUrlConfig `json:"redirect_url_config"`

	FixedResponseConfig *FixtedResponseConfig `json:"fixed_response_config"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o L7Policy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "L7Policy struct{}"
	}

	return strings.Join([]string{"L7Policy", string(data)}, " ")
}

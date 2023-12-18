package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type HealthMonitor struct {
	AdminStateUp bool `json:"admin_state_up"`

	Delay int32 `json:"delay"`

	DomainName string `json:"domain_name"`

	ExpectedCodes string `json:"expected_codes"`

	HttpMethod string `json:"http_method"`

	Id string `json:"id"`

	MaxRetries int32 `json:"max_retries"`

	MaxRetriesDown int32 `json:"max_retries_down"`

	MonitorPort int32 `json:"monitor_port"`

	Name string `json:"name"`

	Pools []PoolRef `json:"pools"`

	ProjectId string `json:"project_id"`

	Timeout int32 `json:"timeout"`

	Type string `json:"type"`

	UrlPath string `json:"url_path"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (o HealthMonitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthMonitor struct{}"
	}

	return strings.Join([]string{"HealthMonitor", string(data)}, " ")
}

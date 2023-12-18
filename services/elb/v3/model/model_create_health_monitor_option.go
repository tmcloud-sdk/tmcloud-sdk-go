package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateHealthMonitorOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Delay int32 `json:"delay"`

	DomainName *string `json:"domain_name,omitempty"`

	ExpectedCodes *string `json:"expected_codes,omitempty"`

	HttpMethod *string `json:"http_method,omitempty"`

	MaxRetries int32 `json:"max_retries"`

	MaxRetriesDown *int32 `json:"max_retries_down,omitempty"`

	MonitorPort *int32 `json:"monitor_port,omitempty"`

	Name *string `json:"name,omitempty"`

	PoolId string `json:"pool_id"`

	ProjectId *string `json:"project_id,omitempty"`

	Timeout int32 `json:"timeout"`

	Type string `json:"type"`

	UrlPath *string `json:"url_path,omitempty"`
}

func (o CreateHealthMonitorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthMonitorOption struct{}"
	}

	return strings.Join([]string{"CreateHealthMonitorOption", string(data)}, " ")
}

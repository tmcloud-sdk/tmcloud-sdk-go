package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListListenersRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	PageReverse *bool `json:"page_reverse,omitempty"`

	ProtocolPort *[]string `json:"protocol_port,omitempty"`

	Protocol *[]string `json:"protocol,omitempty"`

	Description *[]string `json:"description,omitempty"`

	DefaultTlsContainerRef *[]string `json:"default_tls_container_ref,omitempty"`

	ClientCaTlsContainerRef *[]string `json:"client_ca_tls_container_ref,omitempty"`

	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ConnectionLimit *[]int32 `json:"connection_limit,omitempty"`

	DefaultPoolId *[]string `json:"default_pool_id,omitempty"`

	Id *[]string `json:"id,omitempty"`

	Name *[]string `json:"name,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	LoadbalancerId *[]string `json:"loadbalancer_id,omitempty"`

	TlsCiphersPolicy *[]string `json:"tls_ciphers_policy,omitempty"`

	MemberAddress *[]string `json:"member_address,omitempty"`

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	MemberTimeout *[]int32 `json:"member_timeout,omitempty"`

	ClientTimeout *[]int32 `json:"client_timeout,omitempty"`

	KeepaliveTimeout *[]int32 `json:"keepalive_timeout,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	MemberInstanceId *[]string `json:"member_instance_id,omitempty"`
}

func (o ListListenersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListListenersRequest struct{}"
	}

	return strings.Join([]string{"ListListenersRequest", string(data)}, " ")
}

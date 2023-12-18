package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	Description *string `json:"description,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`

	LoadbalancerId string `json:"loadbalancer_id"`

	Name *string `json:"name,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Protocol string `json:"protocol"`

	ProtocolPort int32 `json:"protocol_port"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	SniMatchAlgo *string `json:"sni_match_algo,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	SecurityPolicyId *string `json:"security_policy_id,omitempty"`

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	Ipgroup *CreateListenerIpGroupOption `json:"ipgroup,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	QuicConfig *CreateListenerQuicConfigOption `json:"quic_config,omitempty"`
}

func (o CreateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerOption struct{}"
	}

	return strings.Join([]string{"CreateListenerOption", string(data)}, " ")
}

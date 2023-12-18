package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateListenerOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`

	DefaultPoolId *string `json:"default_pool_id,omitempty"`

	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`

	Description *string `json:"description,omitempty"`

	Http2Enable *bool `json:"http2_enable,omitempty"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`

	Name *string `json:"name,omitempty"`

	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`

	SniMatchAlgo *string `json:"sni_match_algo,omitempty"`

	TlsCiphersPolicy *string `json:"tls_ciphers_policy,omitempty"`

	SecurityPolicyId *string `json:"security_policy_id,omitempty"`

	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`

	MemberTimeout *int32 `json:"member_timeout,omitempty"`

	ClientTimeout *int32 `json:"client_timeout,omitempty"`

	KeepaliveTimeout *int32 `json:"keepalive_timeout,omitempty"`

	Ipgroup *UpdateListenerIpGroupOption `json:"ipgroup,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	EnhanceL7policyEnable *bool `json:"enhance_l7policy_enable,omitempty"`

	QuicConfig *UpdateListenerQuicConfigOption `json:"quic_config,omitempty"`
}

func (o UpdateListenerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateListenerOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerOption", string(data)}, " ")
}

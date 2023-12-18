package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Listener struct {
	AdminStateUp bool `json:"admin_state_up"`

	ClientCaTlsContainerRef string `json:"client_ca_tls_container_ref"`

	ConnectionLimit int32 `json:"connection_limit"`

	CreatedAt string `json:"created_at"`

	DefaultPoolId string `json:"default_pool_id"`

	DefaultTlsContainerRef string `json:"default_tls_container_ref"`

	Description string `json:"description"`

	Http2Enable bool `json:"http2_enable"`

	Id string `json:"id"`

	InsertHeaders *ListenerInsertHeaders `json:"insert_headers"`

	Loadbalancers []LoadBalancerRef `json:"loadbalancers"`

	Name string `json:"name"`

	ProjectId string `json:"project_id"`

	Protocol string `json:"protocol"`

	ProtocolPort int32 `json:"protocol_port"`

	SniContainerRefs []string `json:"sni_container_refs"`

	SniMatchAlgo string `json:"sni_match_algo"`

	Tags []Tag `json:"tags"`

	UpdatedAt string `json:"updated_at"`

	TlsCiphersPolicy string `json:"tls_ciphers_policy"`

	SecurityPolicyId string `json:"security_policy_id"`

	EnableMemberRetry bool `json:"enable_member_retry"`

	KeepaliveTimeout int32 `json:"keepalive_timeout"`

	ClientTimeout int32 `json:"client_timeout"`

	MemberTimeout int32 `json:"member_timeout"`

	Ipgroup *ListenerIpGroup `json:"ipgroup"`

	TransparentClientIpEnable bool `json:"transparent_client_ip_enable"`

	EnhanceL7policyEnable bool `json:"enhance_l7policy_enable"`

	QuicConfig *ListenerQuicConfig `json:"quic_config,omitempty"`
}

func (o Listener) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Listener struct{}"
	}

	return strings.Join([]string{"Listener", string(data)}, " ")
}

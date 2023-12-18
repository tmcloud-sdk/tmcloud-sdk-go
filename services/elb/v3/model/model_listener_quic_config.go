package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListenerQuicConfig struct {
	QuicListenerId *string `json:"quic_listener_id,omitempty"`

	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

func (o ListenerQuicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListenerQuicConfig struct{}"
	}

	return strings.Join([]string{"ListenerQuicConfig", string(data)}, " ")
}

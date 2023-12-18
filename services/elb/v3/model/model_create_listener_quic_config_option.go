package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateListenerQuicConfigOption struct {
	QuicListenerId string `json:"quic_listener_id"`

	EnableQuicUpgrade *bool `json:"enable_quic_upgrade,omitempty"`
}

func (o CreateListenerQuicConfigOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateListenerQuicConfigOption struct{}"
	}

	return strings.Join([]string{"CreateListenerQuicConfigOption", string(data)}, " ")
}

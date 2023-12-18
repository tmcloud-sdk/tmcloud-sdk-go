package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LbaasListenersResult struct {
	ListenerId *string `json:"listener_id,omitempty"`

	PoolId *string `json:"pool_id,omitempty"`

	ProtocolPort *int32 `json:"protocol_port,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o LbaasListenersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbaasListenersResult struct{}"
	}

	return strings.Join([]string{"LbaasListenersResult", string(data)}, " ")
}

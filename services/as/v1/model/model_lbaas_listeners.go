package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type LbaasListeners struct {
	PoolId string `json:"pool_id"`

	ProtocolPort int32 `json:"protocol_port"`

	Weight int32 `json:"weight"`
}

func (o LbaasListeners) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LbaasListeners struct{}"
	}

	return strings.Join([]string{"LbaasListeners", string(data)}, " ")
}

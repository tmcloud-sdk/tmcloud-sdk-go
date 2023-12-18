package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceStatistic struct {
	InputKbps *string `json:"input_kbps,omitempty"`

	OutputKbps *string `json:"output_kbps,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Keys *int64 `json:"keys,omitempty"`

	UsedMemory *int64 `json:"used_memory,omitempty"`

	MaxMemory *int64 `json:"max_memory,omitempty"`

	CmdGetCount *int64 `json:"cmd_get_count,omitempty"`

	CmdSetCount *int64 `json:"cmd_set_count,omitempty"`

	UsedCpu *string `json:"used_cpu,omitempty"`
}

func (o InstanceStatistic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceStatistic struct{}"
	}

	return strings.Join([]string{"InstanceStatistic", string(data)}, " ")
}

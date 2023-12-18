package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EipSpec struct {
	Bandwidth *EipSpecBandwidth `json:"bandwidth,omitempty"`
}

func (o EipSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipSpec struct{}"
	}

	return strings.Join([]string{"EipSpec", string(data)}, " ")
}

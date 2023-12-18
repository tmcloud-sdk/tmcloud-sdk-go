package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServerEip struct {
	Iptype string `json:"iptype"`

	Bandwidth *PrePaidServerEipBandwidth `json:"bandwidth"`

	Extendparam *PrePaidServerEipExtendParam `json:"extendparam,omitempty"`
}

func (o PrePaidServerEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerEip struct{}"
	}

	return strings.Join([]string{"PrePaidServerEip", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Proxy struct {
	PoolId string `json:"pool_id"`

	Status string `json:"status"`

	Address string `json:"address"`

	ElbVip string `json:"elb_vip"`

	Eip string `json:"eip"`

	Port int32 `json:"port"`

	PoolStatus string `json:"pool_status"`

	DelayThresholdInKilobytes int32 `json:"delay_threshold_in_kilobytes"`

	Cpu string `json:"cpu"`

	Mem string `json:"mem"`

	NodeNum int32 `json:"node_num"`

	Nodes []ProxyNode `json:"nodes"`

	Mode string `json:"mode"`
}

func (o Proxy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Proxy struct{}"
	}

	return strings.Join([]string{"Proxy", string(data)}, " ")
}

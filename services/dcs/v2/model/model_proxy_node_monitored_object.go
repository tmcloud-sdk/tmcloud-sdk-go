package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ProxyNodeMonitoredObject struct {
	DcsInstanceId *string `json:"dcs_instance_id,omitempty"`

	Name *string `json:"name,omitempty"`

	DcsClusterProxyNode *string `json:"dcs_cluster_proxy_node,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ProxyNodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyNodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"ProxyNodeMonitoredObject", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Proxy2NodeMonitoredObject struct {
	DcsInstanceId *string `json:"dcs_instance_id,omitempty"`

	Name *string `json:"name,omitempty"`

	DcsClusterProxy2Node *string `json:"dcs_cluster_proxy2_node,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o Proxy2NodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Proxy2NodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"Proxy2NodeMonitoredObject", string(data)}, " ")
}

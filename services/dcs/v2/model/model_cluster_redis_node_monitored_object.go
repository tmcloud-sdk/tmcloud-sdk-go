package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterRedisNodeMonitoredObject struct {
	DcsInstanceId *string `json:"dcs_instance_id,omitempty"`

	Name *string `json:"name,omitempty"`

	DcsClusterRedisNode *string `json:"dcs_cluster_redis_node,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o ClusterRedisNodeMonitoredObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRedisNodeMonitoredObject struct{}"
	}

	return strings.Join([]string{"ClusterRedisNodeMonitoredObject", string(data)}, " ")
}

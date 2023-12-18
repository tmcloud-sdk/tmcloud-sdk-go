package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMonitoredObjectsOfInstanceResponse struct {
	Router *[]string `json:"router,omitempty"`

	Children *[]DimChild `json:"children,omitempty"`

	Instances *[]InstancesMonitoredObject `json:"instances,omitempty"`

	DcsClusterRedisNode *[]ClusterRedisNodeMonitoredObject `json:"dcs_cluster_redis_node,omitempty"`

	DcsClusterProxyNode *[]ProxyNodeMonitoredObject `json:"dcs_cluster_proxy_node,omitempty"`

	DcsClusterProxy2Node *[]Proxy2NodeMonitoredObject `json:"dcs_cluster_proxy2_node,omitempty"`

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMonitoredObjectsOfInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonitoredObjectsOfInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListMonitoredObjectsOfInstanceResponse", string(data)}, " ")
}

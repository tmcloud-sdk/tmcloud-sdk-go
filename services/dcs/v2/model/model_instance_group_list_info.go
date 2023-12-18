package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceGroupListInfo struct {
	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	ReplicationList *[]InstanceReplicationListInfo `json:"replication_list,omitempty"`
}

func (o InstanceGroupListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceGroupListInfo struct{}"
	}

	return strings.Join([]string{"InstanceGroupListInfo", string(data)}, " ")
}

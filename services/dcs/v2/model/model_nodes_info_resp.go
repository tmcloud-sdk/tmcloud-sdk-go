package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NodesInfoResp struct {
	LogicalNodeId *string `json:"logical_node_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *NodesInfoRespStatus `json:"status,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	NodeRole *NodesInfoRespNodeRole `json:"node_role,omitempty"`

	NodeType *NodesInfoRespNodeType `json:"node_type,omitempty"`

	NodeIp *string `json:"node_ip,omitempty"`

	NodePort *string `json:"node_port,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	PriorityWeight *int32 `json:"priority_weight,omitempty"`

	IsAccess *bool `json:"is_access,omitempty"`

	GroupId *string `json:"group_id,omitempty"`

	GroupName *string `json:"group_name,omitempty"`

	IsRemoveIp *bool `json:"is_remove_ip,omitempty"`

	ReplicationId *string `json:"replication_id,omitempty"`

	Dimensions *[]InstanceReplicationDimensionsInfo `json:"dimensions,omitempty"`
}

func (o NodesInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodesInfoResp struct{}"
	}

	return strings.Join([]string{"NodesInfoResp", string(data)}, " ")
}

type NodesInfoRespStatus struct {
	value string
}

type NodesInfoRespStatusEnum struct {
	CREATING     NodesInfoRespStatus
	ACTIVE       NodesInfoRespStatus
	INACTIVE     NodesInfoRespStatus
	DELETING     NodesInfoRespStatus
	ADD_SHARDING NodesInfoRespStatus
}

func GetNodesInfoRespStatusEnum() NodesInfoRespStatusEnum {
	return NodesInfoRespStatusEnum{
		CREATING: NodesInfoRespStatus{
			value: "Creating",
		},
		ACTIVE: NodesInfoRespStatus{
			value: "Active",
		},
		INACTIVE: NodesInfoRespStatus{
			value: "Inactive",
		},
		DELETING: NodesInfoRespStatus{
			value: "Deleting",
		},
		ADD_SHARDING: NodesInfoRespStatus{
			value: "AddSharding",
		},
	}
}

func (c NodesInfoRespStatus) Value() string {
	return c.value
}

func (c NodesInfoRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type NodesInfoRespNodeRole struct {
	value string
}

type NodesInfoRespNodeRoleEnum struct {
	REDIS_SERVER NodesInfoRespNodeRole
	REDIS_PROXY  NodesInfoRespNodeRole
}

func GetNodesInfoRespNodeRoleEnum() NodesInfoRespNodeRoleEnum {
	return NodesInfoRespNodeRoleEnum{
		REDIS_SERVER: NodesInfoRespNodeRole{
			value: "redis-server",
		},
		REDIS_PROXY: NodesInfoRespNodeRole{
			value: "redis-proxy",
		},
	}
}

func (c NodesInfoRespNodeRole) Value() string {
	return c.value
}

func (c NodesInfoRespNodeRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespNodeRole) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type NodesInfoRespNodeType struct {
	value string
}

type NodesInfoRespNodeTypeEnum struct {
	MASTER NodesInfoRespNodeType
	SLAVE  NodesInfoRespNodeType
	PROXY  NodesInfoRespNodeType
}

func GetNodesInfoRespNodeTypeEnum() NodesInfoRespNodeTypeEnum {
	return NodesInfoRespNodeTypeEnum{
		MASTER: NodesInfoRespNodeType{
			value: "master",
		},
		SLAVE: NodesInfoRespNodeType{
			value: "slave",
		},
		PROXY: NodesInfoRespNodeType{
			value: "proxy",
		},
	}
}

func (c NodesInfoRespNodeType) Value() string {
	return c.value
}

func (c NodesInfoRespNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodesInfoRespNodeType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

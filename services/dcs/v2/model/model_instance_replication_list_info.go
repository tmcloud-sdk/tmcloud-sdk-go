package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type InstanceReplicationListInfo struct {
	ReplicationRole *string `json:"replication_role,omitempty"`

	ReplicationIp *string `json:"replication_ip,omitempty"`

	IsReplication *bool `json:"is_replication,omitempty"`

	ReplicationId *string `json:"replication_id,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	Status *InstanceReplicationListInfoStatus `json:"status,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	Dimensions *[]InstanceReplicationDimensionsInfo `json:"dimensions,omitempty"`
}

func (o InstanceReplicationListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceReplicationListInfo struct{}"
	}

	return strings.Join([]string{"InstanceReplicationListInfo", string(data)}, " ")
}

type InstanceReplicationListInfoStatus struct {
	value string
}

type InstanceReplicationListInfoStatusEnum struct {
	ACTIVE   InstanceReplicationListInfoStatus
	INACTIVE InstanceReplicationListInfoStatus
}

func GetInstanceReplicationListInfoStatusEnum() InstanceReplicationListInfoStatusEnum {
	return InstanceReplicationListInfoStatusEnum{
		ACTIVE: InstanceReplicationListInfoStatus{
			value: "Active",
		},
		INACTIVE: InstanceReplicationListInfoStatus{
			value: "Inactive",
		},
	}
}

func (c InstanceReplicationListInfoStatus) Value() string {
	return c.value
}

func (c InstanceReplicationListInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceReplicationListInfoStatus) UnmarshalJSON(b []byte) error {
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

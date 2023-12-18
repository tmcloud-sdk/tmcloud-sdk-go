package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NodeStatus struct {
	Phase *NodeStatusPhase `json:"phase,omitempty"`

	LastProbeTime *string `json:"lastProbeTime,omitempty"`

	JobID *string `json:"jobID,omitempty"`

	ServerId *string `json:"serverId,omitempty"`

	PrivateIP *string `json:"privateIP,omitempty"`

	PrivateIPv6IP *string `json:"privateIPv6IP,omitempty"`

	PublicIP *string `json:"publicIP,omitempty"`

	DeleteStatus *DeleteStatus `json:"deleteStatus,omitempty"`
}

func (o NodeStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeStatus struct{}"
	}

	return strings.Join([]string{"NodeStatus", string(data)}, " ")
}

type NodeStatusPhase struct {
	value string
}

type NodeStatusPhaseEnum struct {
	BUILD      NodeStatusPhase
	INSTALLING NodeStatusPhase
	INSTALLED  NodeStatusPhase
	SHUT_DOWN  NodeStatusPhase
	UPGRADING  NodeStatusPhase
	ACTIVE     NodeStatusPhase
	ABNORMAL   NodeStatusPhase
	DELETING   NodeStatusPhase
	ERROR      NodeStatusPhase
}

func GetNodeStatusPhaseEnum() NodeStatusPhaseEnum {
	return NodeStatusPhaseEnum{
		BUILD: NodeStatusPhase{
			value: "Build",
		},
		INSTALLING: NodeStatusPhase{
			value: "Installing",
		},
		INSTALLED: NodeStatusPhase{
			value: "Installed",
		},
		SHUT_DOWN: NodeStatusPhase{
			value: "ShutDown",
		},
		UPGRADING: NodeStatusPhase{
			value: "Upgrading",
		},
		ACTIVE: NodeStatusPhase{
			value: "Active",
		},
		ABNORMAL: NodeStatusPhase{
			value: "Abnormal",
		},
		DELETING: NodeStatusPhase{
			value: "Deleting",
		},
		ERROR: NodeStatusPhase{
			value: "Error",
		},
	}
}

func (c NodeStatusPhase) Value() string {
	return c.value
}

func (c NodeStatusPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeStatusPhase) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NodePoolStatus struct {
	CurrentNode *int32 `json:"currentNode,omitempty"`

	CreatingNode *int32 `json:"creatingNode,omitempty"`

	DeletingNode *int32 `json:"deletingNode,omitempty"`

	Phase *NodePoolStatusPhase `json:"phase,omitempty"`

	JobId *string `json:"jobId,omitempty"`

	Conditions *[]NodePoolCondition `json:"conditions,omitempty"`
}

func (o NodePoolStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolStatus struct{}"
	}

	return strings.Join([]string{"NodePoolStatus", string(data)}, " ")
}

type NodePoolStatusPhase struct {
	value string
}

type NodePoolStatusPhaseEnum struct {
	SYNCHRONIZING NodePoolStatusPhase
	SYNCHRONIZED  NodePoolStatusPhase
	SOLD_OUT      NodePoolStatusPhase
	DELETING      NodePoolStatusPhase
	ERROR         NodePoolStatusPhase
}

func GetNodePoolStatusPhaseEnum() NodePoolStatusPhaseEnum {
	return NodePoolStatusPhaseEnum{
		SYNCHRONIZING: NodePoolStatusPhase{
			value: "Synchronizing",
		},
		SYNCHRONIZED: NodePoolStatusPhase{
			value: "Synchronized",
		},
		SOLD_OUT: NodePoolStatusPhase{
			value: "SoldOut",
		},
		DELETING: NodePoolStatusPhase{
			value: "Deleting",
		},
		ERROR: NodePoolStatusPhase{
			value: "Error",
		},
	}
}

func (c NodePoolStatusPhase) Value() string {
	return c.value
}

func (c NodePoolStatusPhase) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodePoolStatusPhase) UnmarshalJSON(b []byte) error {
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

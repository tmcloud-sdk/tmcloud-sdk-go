package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	NodepoolScaleUp *CreateNodeRequestNodepoolScaleUp `json:"nodepoolScaleUp,omitempty"`

	Body *NodeCreateRequest `json:"body,omitempty"`
}

func (o CreateNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateNodeRequest", string(data)}, " ")
}

type CreateNodeRequestNodepoolScaleUp struct {
	value string
}

type CreateNodeRequestNodepoolScaleUpEnum struct {
	NODEPOOL_SCALE_UP CreateNodeRequestNodepoolScaleUp
}

func GetCreateNodeRequestNodepoolScaleUpEnum() CreateNodeRequestNodepoolScaleUpEnum {
	return CreateNodeRequestNodepoolScaleUpEnum{
		NODEPOOL_SCALE_UP: CreateNodeRequestNodepoolScaleUp{
			value: "NodepoolScaleUp",
		},
	}
}

func (c CreateNodeRequestNodepoolScaleUp) Value() string {
	return c.value
}

func (c CreateNodeRequestNodepoolScaleUp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateNodeRequestNodepoolScaleUp) UnmarshalJSON(b []byte) error {
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

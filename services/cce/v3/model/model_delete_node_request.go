package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteNodeRequest struct {
	ClusterId string `json:"cluster_id"`

	NodeId string `json:"node_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	NodepoolScaleDown *DeleteNodeRequestNodepoolScaleDown `json:"nodepoolScaleDown,omitempty"`
}

func (o DeleteNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteNodeRequest", string(data)}, " ")
}

type DeleteNodeRequestNodepoolScaleDown struct {
	value string
}

type DeleteNodeRequestNodepoolScaleDownEnum struct {
	NO_SCALE_DOWN DeleteNodeRequestNodepoolScaleDown
}

func GetDeleteNodeRequestNodepoolScaleDownEnum() DeleteNodeRequestNodepoolScaleDownEnum {
	return DeleteNodeRequestNodepoolScaleDownEnum{
		NO_SCALE_DOWN: DeleteNodeRequestNodepoolScaleDown{
			value: "NoScaleDown",
		},
	}
}

func (c DeleteNodeRequestNodepoolScaleDown) Value() string {
	return c.value
}

func (c DeleteNodeRequestNodepoolScaleDown) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteNodeRequestNodepoolScaleDown) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type MasterEipRequestSpec struct {
	Action *MasterEipRequestSpecAction `json:"action,omitempty"`

	Spec *MasterEipRequestSpecSpec `json:"spec,omitempty"`

	Bandwidth *string `json:"bandwidth,omitempty"`

	ElasticIp *string `json:"elasticIp,omitempty"`
}

func (o MasterEipRequestSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequestSpec struct{}"
	}

	return strings.Join([]string{"MasterEipRequestSpec", string(data)}, " ")
}

type MasterEipRequestSpecAction struct {
	value string
}

type MasterEipRequestSpecActionEnum struct {
	BIND   MasterEipRequestSpecAction
	UNBIND MasterEipRequestSpecAction
}

func GetMasterEipRequestSpecActionEnum() MasterEipRequestSpecActionEnum {
	return MasterEipRequestSpecActionEnum{
		BIND: MasterEipRequestSpecAction{
			value: "bind",
		},
		UNBIND: MasterEipRequestSpecAction{
			value: "unbind",
		},
	}
}

func (c MasterEipRequestSpecAction) Value() string {
	return c.value
}

func (c MasterEipRequestSpecAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MasterEipRequestSpecAction) UnmarshalJSON(b []byte) error {
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

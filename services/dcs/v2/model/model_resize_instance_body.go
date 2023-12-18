package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ResizeInstanceBody struct {
	SpecCode string `json:"spec_code"`

	NewCapacity int32 `json:"new_capacity"`

	BssParam *BssParamEntity `json:"bss_param,omitempty"`

	ReservedIp *[]string `json:"reserved_ip,omitempty"`

	ChangeType *ResizeInstanceBodyChangeType `json:"change_type,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	NodeList *[]string `json:"node_list,omitempty"`

	ExecuteImmediately *bool `json:"execute_immediately,omitempty"`
}

func (o ResizeInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstanceBody struct{}"
	}

	return strings.Join([]string{"ResizeInstanceBody", string(data)}, " ")
}

type ResizeInstanceBodyChangeType struct {
	value string
}

type ResizeInstanceBodyChangeTypeEnum struct {
	CREATE_REPLICATION ResizeInstanceBodyChangeType
	DELETE_REPLICATION ResizeInstanceBodyChangeType
}

func GetResizeInstanceBodyChangeTypeEnum() ResizeInstanceBodyChangeTypeEnum {
	return ResizeInstanceBodyChangeTypeEnum{
		CREATE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "createReplication",
		},
		DELETE_REPLICATION: ResizeInstanceBodyChangeType{
			value: "deleteReplication",
		},
	}
}

func (c ResizeInstanceBodyChangeType) Value() string {
	return c.value
}

func (c ResizeInstanceBodyChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResizeInstanceBodyChangeType) UnmarshalJSON(b []byte) error {
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

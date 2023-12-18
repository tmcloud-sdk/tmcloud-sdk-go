package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type InstInfo struct {
	EngineType *InstInfoEngineType `json:"engine_type,omitempty"`

	InstType *InstInfoInstType `json:"inst_type,omitempty"`

	Ip *string `json:"ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	Status *InstInfoStatus `json:"status,omitempty"`

	VolumeSize *int32 `json:"volume_size,omitempty"`
}

func (o InstInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstInfo struct{}"
	}

	return strings.Join([]string{"InstInfo", string(data)}, " ")
}

type InstInfoEngineType struct {
	value string
}

type InstInfoEngineTypeEnum struct {
	MYSQL                  InstInfoEngineType
	MONGODB                InstInfoEngineType
	CLOUD_DATA_GUARD_MYSQL InstInfoEngineType
}

func GetInstInfoEngineTypeEnum() InstInfoEngineTypeEnum {
	return InstInfoEngineTypeEnum{
		MYSQL: InstInfoEngineType{
			value: "mysql",
		},
		MONGODB: InstInfoEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: InstInfoEngineType{
			value: "cloudDataGuard-mysql",
		},
	}
}

func (c InstInfoEngineType) Value() string {
	return c.value
}

func (c InstInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoEngineType) UnmarshalJSON(b []byte) error {
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

type InstInfoInstType struct {
	value string
}

type InstInfoInstTypeEnum struct {
	HIGH InstInfoInstType
}

func GetInstInfoInstTypeEnum() InstInfoInstTypeEnum {
	return InstInfoInstTypeEnum{
		HIGH: InstInfoInstType{
			value: "high",
		},
	}
}

func (c InstInfoInstType) Value() string {
	return c.value
}

func (c InstInfoInstType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoInstType) UnmarshalJSON(b []byte) error {
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

type InstInfoStatus struct {
	value string
}

type InstInfoStatusEnum struct {
	ACTIVE  InstInfoStatus
	DELETED InstInfoStatus
}

func GetInstInfoStatusEnum() InstInfoStatusEnum {
	return InstInfoStatusEnum{
		ACTIVE: InstInfoStatus{
			value: "active",
		},
		DELETED: InstInfoStatus{
			value: "deleted",
		},
	}
}

func (c InstInfoStatus) Value() string {
	return c.value
}

func (c InstInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstInfoStatus) UnmarshalJSON(b []byte) error {
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

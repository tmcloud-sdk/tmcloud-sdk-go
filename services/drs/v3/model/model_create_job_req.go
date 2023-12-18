package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateJobReq struct {
	BindEip *bool `json:"bind_eip,omitempty"`

	DbUseType CreateJobReqDbUseType `json:"db_use_type"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	EngineType CreateJobReqEngineType `json:"engine_type"`

	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	JobDirection CreateJobReqJobDirection `json:"job_direction"`

	MultiWrite *bool `json:"multi_write,omitempty"`

	NetType CreateJobReqNetType `json:"net_type"`

	NodeNum *int32 `json:"node_num,omitempty"`

	NodeType CreateJobReqNodeType `json:"node_type"`

	SourceEndpoint *Endpoint `json:"source_endpoint"`

	TargetEndpoint *Endpoint `json:"target_endpoint"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	TaskType CreateJobReqTaskType `json:"task_type"`

	CustomizeSutnetId string `json:"customize_sutnet_id"`

	ProductId *string `json:"product_id,omitempty"`

	SysTags *[]ResourceTag `json:"sys_tags,omitempty"`

	ExpiredDays *string `json:"expired_days,omitempty"`

	MasterAz *string `json:"master_az,omitempty"`

	SlaveAz *string `json:"slave_az,omitempty"`

	ChargingMode *CreateJobReqChargingMode `json:"charging_mode,omitempty"`

	PeriodOrder *PeriodOrderInfo `json:"period_order,omitempty"`
}

func (o CreateJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobReq struct{}"
	}

	return strings.Join([]string{"CreateJobReq", string(data)}, " ")
}

type CreateJobReqDbUseType struct {
	value string
}

type CreateJobReqDbUseTypeEnum struct {
	MIGRATION        CreateJobReqDbUseType
	SYNC             CreateJobReqDbUseType
	CLOUD_DATA_GUARD CreateJobReqDbUseType
}

func GetCreateJobReqDbUseTypeEnum() CreateJobReqDbUseTypeEnum {
	return CreateJobReqDbUseTypeEnum{
		MIGRATION: CreateJobReqDbUseType{
			value: "migration",
		},
		SYNC: CreateJobReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: CreateJobReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c CreateJobReqDbUseType) Value() string {
	return c.value
}

func (c CreateJobReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqDbUseType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqEngineType struct {
	value string
}

type CreateJobReqEngineTypeEnum struct {
	MYSQL                  CreateJobReqEngineType
	MONGODB                CreateJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL CreateJobReqEngineType
	GAUSSDBV5              CreateJobReqEngineType
	POSTGRESQL             CreateJobReqEngineType
}

func GetCreateJobReqEngineTypeEnum() CreateJobReqEngineTypeEnum {
	return CreateJobReqEngineTypeEnum{
		MYSQL: CreateJobReqEngineType{
			value: "mysql",
		},
		MONGODB: CreateJobReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: CreateJobReqEngineType{
			value: "cloudDataGuard-mysql",
		},
		GAUSSDBV5: CreateJobReqEngineType{
			value: "gaussdbv5",
		},
		POSTGRESQL: CreateJobReqEngineType{
			value: "postgresql",
		},
	}
}

func (c CreateJobReqEngineType) Value() string {
	return c.value
}

func (c CreateJobReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqEngineType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqJobDirection struct {
	value string
}

type CreateJobReqJobDirectionEnum struct {
	UP      CreateJobReqJobDirection
	DOWN    CreateJobReqJobDirection
	NON_DBS CreateJobReqJobDirection
}

func GetCreateJobReqJobDirectionEnum() CreateJobReqJobDirectionEnum {
	return CreateJobReqJobDirectionEnum{
		UP: CreateJobReqJobDirection{
			value: "up",
		},
		DOWN: CreateJobReqJobDirection{
			value: "down",
		},
		NON_DBS: CreateJobReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c CreateJobReqJobDirection) Value() string {
	return c.value
}

func (c CreateJobReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqJobDirection) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNetType struct {
	value string
}

type CreateJobReqNetTypeEnum struct {
	VPN CreateJobReqNetType
	VPC CreateJobReqNetType
	EIP CreateJobReqNetType
}

func GetCreateJobReqNetTypeEnum() CreateJobReqNetTypeEnum {
	return CreateJobReqNetTypeEnum{
		VPN: CreateJobReqNetType{
			value: "vpn",
		},
		VPC: CreateJobReqNetType{
			value: "vpc",
		},
		EIP: CreateJobReqNetType{
			value: "eip",
		},
	}
}

func (c CreateJobReqNetType) Value() string {
	return c.value
}

func (c CreateJobReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNetType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqNodeType struct {
	value string
}

type CreateJobReqNodeTypeEnum struct {
	HIGH CreateJobReqNodeType
}

func GetCreateJobReqNodeTypeEnum() CreateJobReqNodeTypeEnum {
	return CreateJobReqNodeTypeEnum{
		HIGH: CreateJobReqNodeType{
			value: "high",
		},
	}
}

func (c CreateJobReqNodeType) Value() string {
	return c.value
}

func (c CreateJobReqNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqNodeType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqTaskType struct {
	value string
}

type CreateJobReqTaskTypeEnum struct {
	FULL_TRANS      CreateJobReqTaskType
	FULL_INCR_TRANS CreateJobReqTaskType
	INCR_TRANS      CreateJobReqTaskType
}

func GetCreateJobReqTaskTypeEnum() CreateJobReqTaskTypeEnum {
	return CreateJobReqTaskTypeEnum{
		FULL_TRANS: CreateJobReqTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: CreateJobReqTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: CreateJobReqTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c CreateJobReqTaskType) Value() string {
	return c.value
}

func (c CreateJobReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqTaskType) UnmarshalJSON(b []byte) error {
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

type CreateJobReqChargingMode struct {
	value string
}

type CreateJobReqChargingModeEnum struct {
	PERIOD    CreateJobReqChargingMode
	ON_DEMAND CreateJobReqChargingMode
}

func GetCreateJobReqChargingModeEnum() CreateJobReqChargingModeEnum {
	return CreateJobReqChargingModeEnum{
		PERIOD: CreateJobReqChargingMode{
			value: "period",
		},
		ON_DEMAND: CreateJobReqChargingMode{
			value: "on_demand",
		},
	}
}

func (c CreateJobReqChargingMode) Value() string {
	return c.value
}

func (c CreateJobReqChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateJobReqChargingMode) UnmarshalJSON(b []byte) error {
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

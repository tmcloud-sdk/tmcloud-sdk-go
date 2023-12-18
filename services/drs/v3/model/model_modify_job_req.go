package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ModifyJobReq struct {
	JobId string `json:"job_id"`

	Description *string `json:"description,omitempty"`

	Name *string `json:"name,omitempty"`

	AlarmNotify *AlarmNotifyInfo `json:"alarm_notify,omitempty"`

	TaskType *ModifyJobReqTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	NodeType *ModifyJobReqNodeType `json:"node_type,omitempty"`

	EngineType *ModifyJobReqEngineType `json:"engine_type,omitempty"`

	NetType *ModifyJobReqNetType `json:"net_type,omitempty"`

	StoreDbInfo *bool `json:"store_db_info,omitempty"`

	IsRecreate *bool `json:"is_recreate,omitempty"`

	JobDirection *ModifyJobReqJobDirection `json:"job_direction,omitempty"`

	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	DbUseType *ModifyJobReqDbUseType `json:"db_use_type,omitempty"`

	ProductId *string `json:"product_id,omitempty"`
}

func (o ModifyJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyJobReq struct{}"
	}

	return strings.Join([]string{"ModifyJobReq", string(data)}, " ")
}

type ModifyJobReqTaskType struct {
	value string
}

type ModifyJobReqTaskTypeEnum struct {
	FULL_TRANS      ModifyJobReqTaskType
	INCR_TRANS      ModifyJobReqTaskType
	FULL_INCR_TRANS ModifyJobReqTaskType
}

func GetModifyJobReqTaskTypeEnum() ModifyJobReqTaskTypeEnum {
	return ModifyJobReqTaskTypeEnum{
		FULL_TRANS: ModifyJobReqTaskType{
			value: "FULL_TRANS",
		},
		INCR_TRANS: ModifyJobReqTaskType{
			value: "INCR_TRANS",
		},
		FULL_INCR_TRANS: ModifyJobReqTaskType{
			value: "FULL_INCR_TRANS",
		},
	}
}

func (c ModifyJobReqTaskType) Value() string {
	return c.value
}

func (c ModifyJobReqTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqTaskType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqNodeType struct {
	value string
}

type ModifyJobReqNodeTypeEnum struct {
	HIGH ModifyJobReqNodeType
}

func GetModifyJobReqNodeTypeEnum() ModifyJobReqNodeTypeEnum {
	return ModifyJobReqNodeTypeEnum{
		HIGH: ModifyJobReqNodeType{
			value: "high",
		},
	}
}

func (c ModifyJobReqNodeType) Value() string {
	return c.value
}

func (c ModifyJobReqNodeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqNodeType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqEngineType struct {
	value string
}

type ModifyJobReqEngineTypeEnum struct {
	MYSQL                  ModifyJobReqEngineType
	MONGODB                ModifyJobReqEngineType
	CLOUD_DATA_GUARD_MYSQL ModifyJobReqEngineType
}

func GetModifyJobReqEngineTypeEnum() ModifyJobReqEngineTypeEnum {
	return ModifyJobReqEngineTypeEnum{
		MYSQL: ModifyJobReqEngineType{
			value: "mysql",
		},
		MONGODB: ModifyJobReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: ModifyJobReqEngineType{
			value: "cloudDataGuard-mysql",
		},
	}
}

func (c ModifyJobReqEngineType) Value() string {
	return c.value
}

func (c ModifyJobReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqEngineType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqNetType struct {
	value string
}

type ModifyJobReqNetTypeEnum struct {
	VPC ModifyJobReqNetType
	VPN ModifyJobReqNetType
	EIP ModifyJobReqNetType
}

func GetModifyJobReqNetTypeEnum() ModifyJobReqNetTypeEnum {
	return ModifyJobReqNetTypeEnum{
		VPC: ModifyJobReqNetType{
			value: "vpc",
		},
		VPN: ModifyJobReqNetType{
			value: "vpn",
		},
		EIP: ModifyJobReqNetType{
			value: "eip",
		},
	}
}

func (c ModifyJobReqNetType) Value() string {
	return c.value
}

func (c ModifyJobReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqNetType) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqJobDirection struct {
	value string
}

type ModifyJobReqJobDirectionEnum struct {
	UP      ModifyJobReqJobDirection
	DOWN    ModifyJobReqJobDirection
	NON_DBS ModifyJobReqJobDirection
}

func GetModifyJobReqJobDirectionEnum() ModifyJobReqJobDirectionEnum {
	return ModifyJobReqJobDirectionEnum{
		UP: ModifyJobReqJobDirection{
			value: "up",
		},
		DOWN: ModifyJobReqJobDirection{
			value: "down",
		},
		NON_DBS: ModifyJobReqJobDirection{
			value: "non-dbs",
		},
	}
}

func (c ModifyJobReqJobDirection) Value() string {
	return c.value
}

func (c ModifyJobReqJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqJobDirection) UnmarshalJSON(b []byte) error {
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

type ModifyJobReqDbUseType struct {
	value string
}

type ModifyJobReqDbUseTypeEnum struct {
	MIGRATION        ModifyJobReqDbUseType
	SYNC             ModifyJobReqDbUseType
	CLOUD_DATA_GUARD ModifyJobReqDbUseType
}

func GetModifyJobReqDbUseTypeEnum() ModifyJobReqDbUseTypeEnum {
	return ModifyJobReqDbUseTypeEnum{
		MIGRATION: ModifyJobReqDbUseType{
			value: "migration",
		},
		SYNC: ModifyJobReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ModifyJobReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c ModifyJobReqDbUseType) Value() string {
	return c.value
}

func (c ModifyJobReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifyJobReqDbUseType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ChildrenJobInfo struct {
	BillingTag bool `json:"billing_tag"`

	CreateTime string `json:"create_time"`

	DbUseType ChildrenJobInfoDbUseType `json:"db_use_type"`

	Description string `json:"description"`

	EngineType ChildrenJobInfoEngineType `json:"engine_type"`

	ErrorMsg string `json:"error_msg"`

	Id string `json:"id"`

	JobDirection ChildrenJobInfoJobDirection `json:"job_direction"`

	Name string `json:"name"`

	NetType ChildrenJobInfoNetType `json:"net_type"`

	NodeNewFramework bool `json:"node_newFramework"`

	Status ChildrenJobInfoStatus `json:"status"`

	TaskType ChildrenJobInfoTaskType `json:"task_type"`

	JobAction *JobActionResp `json:"job_action,omitempty"`
}

func (o ChildrenJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChildrenJobInfo struct{}"
	}

	return strings.Join([]string{"ChildrenJobInfo", string(data)}, " ")
}

type ChildrenJobInfoDbUseType struct {
	value string
}

type ChildrenJobInfoDbUseTypeEnum struct {
	MIGRATION        ChildrenJobInfoDbUseType
	SYNC             ChildrenJobInfoDbUseType
	CLOUD_DATA_GUARD ChildrenJobInfoDbUseType
}

func GetChildrenJobInfoDbUseTypeEnum() ChildrenJobInfoDbUseTypeEnum {
	return ChildrenJobInfoDbUseTypeEnum{
		MIGRATION: ChildrenJobInfoDbUseType{
			value: "migration",
		},
		SYNC: ChildrenJobInfoDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: ChildrenJobInfoDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c ChildrenJobInfoDbUseType) Value() string {
	return c.value
}

func (c ChildrenJobInfoDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoDbUseType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoEngineType struct {
	value string
}

type ChildrenJobInfoEngineTypeEnum struct {
	CLOUD_DATA_GUARD_CASSANDRA       ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_DDM             ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_TAURUS_TO_MYSQL ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL           ChildrenJobInfoEngineType
	CLOUD_DATA_GUARD_MYSQL_TO_TAURUS ChildrenJobInfoEngineType
}

func GetChildrenJobInfoEngineTypeEnum() ChildrenJobInfoEngineTypeEnum {
	return ChildrenJobInfoEngineTypeEnum{
		CLOUD_DATA_GUARD_CASSANDRA: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-cassandra",
		},
		CLOUD_DATA_GUARD_DDM: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-ddm",
		},
		CLOUD_DATA_GUARD_TAURUS_TO_MYSQL: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-taurus-to-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-mysql",
		},
		CLOUD_DATA_GUARD_MYSQL_TO_TAURUS: ChildrenJobInfoEngineType{
			value: "cloudDataGuard-mysql-to-taurus",
		},
	}
}

func (c ChildrenJobInfoEngineType) Value() string {
	return c.value
}

func (c ChildrenJobInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoEngineType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoJobDirection struct {
	value string
}

type ChildrenJobInfoJobDirectionEnum struct {
	UP     ChildrenJobInfoJobDirection
	DOWN   ChildrenJobInfoJobDirection
	NO_DBS ChildrenJobInfoJobDirection
}

func GetChildrenJobInfoJobDirectionEnum() ChildrenJobInfoJobDirectionEnum {
	return ChildrenJobInfoJobDirectionEnum{
		UP: ChildrenJobInfoJobDirection{
			value: "up",
		},
		DOWN: ChildrenJobInfoJobDirection{
			value: "down",
		},
		NO_DBS: ChildrenJobInfoJobDirection{
			value: "no-dbs",
		},
	}
}

func (c ChildrenJobInfoJobDirection) Value() string {
	return c.value
}

func (c ChildrenJobInfoJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoJobDirection) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoNetType struct {
	value string
}

type ChildrenJobInfoNetTypeEnum struct {
	VPC ChildrenJobInfoNetType
	VPN ChildrenJobInfoNetType
	EIP ChildrenJobInfoNetType
}

func GetChildrenJobInfoNetTypeEnum() ChildrenJobInfoNetTypeEnum {
	return ChildrenJobInfoNetTypeEnum{
		VPC: ChildrenJobInfoNetType{
			value: "vpc",
		},
		VPN: ChildrenJobInfoNetType{
			value: "vpn",
		},
		EIP: ChildrenJobInfoNetType{
			value: "eip",
		},
	}
}

func (c ChildrenJobInfoNetType) Value() string {
	return c.value
}

func (c ChildrenJobInfoNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoNetType) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoStatus struct {
	value string
}

type ChildrenJobInfoStatusEnum struct {
	CREATING                        ChildrenJobInfoStatus
	CREATE_FAILED                   ChildrenJobInfoStatus
	CONFIGURATION                   ChildrenJobInfoStatus
	STARTJOBING                     ChildrenJobInfoStatus
	WAITING_FOR_START               ChildrenJobInfoStatus
	START_JOB_FAILED                ChildrenJobInfoStatus
	PAUSING                         ChildrenJobInfoStatus
	FULL_TRANSFER_STARTED           ChildrenJobInfoStatus
	FULL_TRANSFER_FAILED            ChildrenJobInfoStatus
	FULL_TRANSFER_COMPLETE          ChildrenJobInfoStatus
	INCRE_TRANSFER_STARTED          ChildrenJobInfoStatus
	INCRE_TRANSFER_FAILED           ChildrenJobInfoStatus
	RELEASE_RESOURCE_STARTED        ChildrenJobInfoStatus
	RELEASE_RESOURCE_FAILED         ChildrenJobInfoStatus
	RELEASE_RESOURCE_COMPLETE       ChildrenJobInfoStatus
	REBUILD_NODE_STARTED            ChildrenJobInfoStatus
	REBUILD_NODE_FAILED             ChildrenJobInfoStatus
	CHANGE_JOB_STARTED              ChildrenJobInfoStatus
	CHANGE_JOB_FAILED               ChildrenJobInfoStatus
	DELETED                         ChildrenJobInfoStatus
	CHILD_TRANSFER_STARTING         ChildrenJobInfoStatus
	CHILD_TRANSFER_STARTED          ChildrenJobInfoStatus
	CHILD_TRANSFER_COMPLETE         ChildrenJobInfoStatus
	CHILD_TRANSFER_FAILED           ChildrenJobInfoStatus
	RELEASE_CHILD_TRANSFER_STARTED  ChildrenJobInfoStatus
	RELEASE_CHILD_TRANSFER_COMPLETE ChildrenJobInfoStatus
	NODE_UPGRADE_START              ChildrenJobInfoStatus
	NODE_UPGRADE_COMPLETE           ChildrenJobInfoStatus
	NODE_UPGRADE_FAILED             ChildrenJobInfoStatus
}

func GetChildrenJobInfoStatusEnum() ChildrenJobInfoStatusEnum {
	return ChildrenJobInfoStatusEnum{
		CREATING: ChildrenJobInfoStatus{
			value: "CREATING",
		},
		CREATE_FAILED: ChildrenJobInfoStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: ChildrenJobInfoStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: ChildrenJobInfoStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: ChildrenJobInfoStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: ChildrenJobInfoStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: ChildrenJobInfoStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: ChildrenJobInfoStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: ChildrenJobInfoStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: ChildrenJobInfoStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: ChildrenJobInfoStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: ChildrenJobInfoStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: ChildrenJobInfoStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: ChildrenJobInfoStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: ChildrenJobInfoStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: ChildrenJobInfoStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: ChildrenJobInfoStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: ChildrenJobInfoStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: ChildrenJobInfoStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: ChildrenJobInfoStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: ChildrenJobInfoStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: ChildrenJobInfoStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: ChildrenJobInfoStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: ChildrenJobInfoStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: ChildrenJobInfoStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: ChildrenJobInfoStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: ChildrenJobInfoStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: ChildrenJobInfoStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: ChildrenJobInfoStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c ChildrenJobInfoStatus) Value() string {
	return c.value
}

func (c ChildrenJobInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoStatus) UnmarshalJSON(b []byte) error {
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

type ChildrenJobInfoTaskType struct {
	value string
}

type ChildrenJobInfoTaskTypeEnum struct {
	FULL_TRANS      ChildrenJobInfoTaskType
	FULL_INCR_TRANS ChildrenJobInfoTaskType
	INCR_TRANS      ChildrenJobInfoTaskType
}

func GetChildrenJobInfoTaskTypeEnum() ChildrenJobInfoTaskTypeEnum {
	return ChildrenJobInfoTaskTypeEnum{
		FULL_TRANS: ChildrenJobInfoTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: ChildrenJobInfoTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: ChildrenJobInfoTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c ChildrenJobInfoTaskType) Value() string {
	return c.value
}

func (c ChildrenJobInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChildrenJobInfoTaskType) UnmarshalJSON(b []byte) error {
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

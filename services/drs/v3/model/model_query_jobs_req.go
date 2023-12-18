package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryJobsReq struct {
	CurPage int32 `json:"cur_page"`

	PerPage int32 `json:"per_page"`

	DbUseType QueryJobsReqDbUseType `json:"db_use_type"`

	EngineType *QueryJobsReqEngineType `json:"engine_type,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Name *string `json:"name,omitempty"`

	NetType *QueryJobsReqNetType `json:"net_type,omitempty"`

	ServiceName *string `json:"service_name,omitempty"`

	Status *QueryJobsReqStatus `json:"status,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`
}

func (o QueryJobsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobsReq struct{}"
	}

	return strings.Join([]string{"QueryJobsReq", string(data)}, " ")
}

type QueryJobsReqDbUseType struct {
	value string
}

type QueryJobsReqDbUseTypeEnum struct {
	MIGRATION        QueryJobsReqDbUseType
	SYNC             QueryJobsReqDbUseType
	CLOUD_DATA_GUARD QueryJobsReqDbUseType
}

func GetQueryJobsReqDbUseTypeEnum() QueryJobsReqDbUseTypeEnum {
	return QueryJobsReqDbUseTypeEnum{
		MIGRATION: QueryJobsReqDbUseType{
			value: "migration",
		},
		SYNC: QueryJobsReqDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryJobsReqDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryJobsReqDbUseType) Value() string {
	return c.value
}

func (c QueryJobsReqDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqEngineType struct {
	value string
}

type QueryJobsReqEngineTypeEnum struct {
	MYSQL                  QueryJobsReqEngineType
	MONGODB                QueryJobsReqEngineType
	CLOUD_DATA_GUARD_MYSQL QueryJobsReqEngineType
}

func GetQueryJobsReqEngineTypeEnum() QueryJobsReqEngineTypeEnum {
	return QueryJobsReqEngineTypeEnum{
		MYSQL: QueryJobsReqEngineType{
			value: "mysql",
		},
		MONGODB: QueryJobsReqEngineType{
			value: "mongodb",
		},
		CLOUD_DATA_GUARD_MYSQL: QueryJobsReqEngineType{
			value: "cloudDataGuard-mysql",
		},
	}
}

func (c QueryJobsReqEngineType) Value() string {
	return c.value
}

func (c QueryJobsReqEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqEngineType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqNetType struct {
	value string
}

type QueryJobsReqNetTypeEnum struct {
	VPN QueryJobsReqNetType
	VPC QueryJobsReqNetType
	EIP QueryJobsReqNetType
}

func GetQueryJobsReqNetTypeEnum() QueryJobsReqNetTypeEnum {
	return QueryJobsReqNetTypeEnum{
		VPN: QueryJobsReqNetType{
			value: "vpn",
		},
		VPC: QueryJobsReqNetType{
			value: "vpc",
		},
		EIP: QueryJobsReqNetType{
			value: "eip",
		},
	}
}

func (c QueryJobsReqNetType) Value() string {
	return c.value
}

func (c QueryJobsReqNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqNetType) UnmarshalJSON(b []byte) error {
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

type QueryJobsReqStatus struct {
	value string
}

type QueryJobsReqStatusEnum struct {
	CREATING                        QueryJobsReqStatus
	CREATE_FAILED                   QueryJobsReqStatus
	CONFIGURATION                   QueryJobsReqStatus
	STARTJOBING                     QueryJobsReqStatus
	WAITING_FOR_START               QueryJobsReqStatus
	START_JOB_FAILED                QueryJobsReqStatus
	PAUSING                         QueryJobsReqStatus
	FULL_TRANSFER_STARTED           QueryJobsReqStatus
	FULL_TRANSFER_FAILED            QueryJobsReqStatus
	FULL_TRANSFER_COMPLETE          QueryJobsReqStatus
	INCRE_TRANSFER_STARTED          QueryJobsReqStatus
	INCRE_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_RESOURCE_STARTED        QueryJobsReqStatus
	RELEASE_RESOURCE_FAILED         QueryJobsReqStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobsReqStatus
	REBUILD_NODE_STARTED            QueryJobsReqStatus
	REBUILD_NODE_FAILED             QueryJobsReqStatus
	CHANGE_JOB_STARTED              QueryJobsReqStatus
	CHANGE_JOB_FAILED               QueryJobsReqStatus
	DELETED                         QueryJobsReqStatus
	CHILD_TRANSFER_STARTING         QueryJobsReqStatus
	CHILD_TRANSFER_STARTED          QueryJobsReqStatus
	CHILD_TRANSFER_COMPLETE         QueryJobsReqStatus
	CHILD_TRANSFER_FAILED           QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobsReqStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobsReqStatus
	NODE_UPGRADE_START              QueryJobsReqStatus
	NODE_UPGRADE_COMPLETE           QueryJobsReqStatus
	NODE_UPGRADE_FAILED             QueryJobsReqStatus
}

func GetQueryJobsReqStatusEnum() QueryJobsReqStatusEnum {
	return QueryJobsReqStatusEnum{
		CREATING: QueryJobsReqStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobsReqStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobsReqStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobsReqStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobsReqStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobsReqStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: QueryJobsReqStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: QueryJobsReqStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: QueryJobsReqStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: QueryJobsReqStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobsReqStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: QueryJobsReqStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobsReqStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobsReqStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: QueryJobsReqStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: QueryJobsReqStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: QueryJobsReqStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c QueryJobsReqStatus) Value() string {
	return c.value
}

func (c QueryJobsReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobsReqStatus) UnmarshalJSON(b []byte) error {
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

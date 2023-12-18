package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type QueryJobResp struct {
	Id *string `json:"id,omitempty"`

	ParentId *string `json:"parent_id,omitempty"`

	Name *string `json:"name,omitempty"`

	Status *QueryJobRespStatus `json:"status,omitempty"`

	Description *string `json:"description,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	TaskType *QueryJobRespTaskType `json:"task_type,omitempty"`

	SourceEndpoint *Endpoint `json:"source_endpoint,omitempty"`

	DmqEndpoint *Endpoint `json:"dmq_endpoint,omitempty"`

	SourceSharding *[]Endpoint `json:"source_sharding,omitempty"`

	TargetEndpoint *Endpoint `json:"target_endpoint,omitempty"`

	NetType *QueryJobRespNetType `json:"net_type,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	InstInfo *InstInfo `json:"inst_info,omitempty"`

	ActualStartTime *string `json:"actual_start_time,omitempty"`

	FullTransferCompleteTime *string `json:"full_transfer_complete_time,omitempty"`

	UpdateTime *string `json:"update_time,omitempty"`

	JobDirection *QueryJobRespJobDirection `json:"job_direction,omitempty"`

	DbUseType *QueryJobRespDbUseType `json:"db_use_type,omitempty"`

	NeedRestart *bool `json:"need_restart,omitempty"`

	IsTargetReadonly *bool `json:"is_target_readonly,omitempty"`

	ConflictPolicy *QueryJobRespConflictPolicy `json:"conflict_policy,omitempty"`

	FilterDdlPolicy *string `json:"filter_ddl_policy,omitempty"`

	SpeedLimit *[]SpeedLimitInfo `json:"speed_limit,omitempty"`

	SchemaType *QueryJobRespSchemaType `json:"schema_type,omitempty"`

	NodeNum *string `json:"node_num,omitempty"`

	ObjectSwitch *bool `json:"object_switch,omitempty"`

	MasterJobId *string `json:"master_job_id,omitempty"`

	FullMode *string `json:"full_mode,omitempty"`

	StructTrans *bool `json:"struct_trans,omitempty"`

	IndexTrans *bool `json:"index_trans,omitempty"`

	ReplaceDefiner *bool `json:"replace_definer,omitempty"`

	MigrateUser *bool `json:"migrate_user,omitempty"`

	SyncDatabase *bool `json:"sync_database,omitempty"`

	ErrorCode *string `json:"error_code,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`

	TargetRootDb *DefaultRootDb `json:"target_root_db,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	MultiWrite *bool `json:"multi_write,omitempty"`

	SupportIpV6 *bool `json:"support_ip_v6,omitempty"`

	InheritId *string `json:"inherit_id,omitempty"`

	Gtid *string `json:"gtid,omitempty"`

	AlarmNotify *QuerySmnInfoResp `json:"alarm_notify,omitempty"`

	IncreStartPosition *string `json:"incre_start_position,omitempty"`

	IsMultiAz *bool `json:"is_multi_az,omitempty"`

	AzName *string `json:"az_name,omitempty"`

	MasterAz *string `json:"master_az,omitempty"`

	SlaveAz *string `json:"slave_az,omitempty"`

	NodeRole *string `json:"node_role,omitempty"`

	PeriodOrder *PeriodOrderResp `json:"period_order,omitempty"`

	ObjectInfos *[]DatabaseObjectInfo `json:"object_infos,omitempty"`

	OriginalJobDirection *QueryJobRespOriginalJobDirection `json:"original_job_direction,omitempty"`

	DataTransformation *GetDataTransformationResp `json:"data_transformation,omitempty"`

	Tags *[]Tag `json:"tags,omitempty"`
}

func (o QueryJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryJobResp struct{}"
	}

	return strings.Join([]string{"QueryJobResp", string(data)}, " ")
}

type QueryJobRespStatus struct {
	value string
}

type QueryJobRespStatusEnum struct {
	CREATING                        QueryJobRespStatus
	CREATE_FAILED                   QueryJobRespStatus
	CONFIGURATION                   QueryJobRespStatus
	STARTJOBING                     QueryJobRespStatus
	WAITING_FOR_START               QueryJobRespStatus
	START_JOB_FAILED                QueryJobRespStatus
	PAUSING                         QueryJobRespStatus
	FULL_TRANSFER_STARTED           QueryJobRespStatus
	FULL_TRANSFER_FAILED            QueryJobRespStatus
	FULL_TRANSFER_COMPLETE          QueryJobRespStatus
	INCRE_TRANSFER_STARTED          QueryJobRespStatus
	INCRE_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_RESOURCE_STARTED        QueryJobRespStatus
	RELEASE_RESOURCE_FAILED         QueryJobRespStatus
	RELEASE_RESOURCE_COMPLETE       QueryJobRespStatus
	REBUILD_NODE_STARTED            QueryJobRespStatus
	REBUILD_NODE_FAILED             QueryJobRespStatus
	CHANGE_JOB_STARTED              QueryJobRespStatus
	CHANGE_JOB_FAILED               QueryJobRespStatus
	DELETED                         QueryJobRespStatus
	CHILD_TRANSFER_STARTING         QueryJobRespStatus
	CHILD_TRANSFER_STARTED          QueryJobRespStatus
	CHILD_TRANSFER_COMPLETE         QueryJobRespStatus
	CHILD_TRANSFER_FAILED           QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_STARTED  QueryJobRespStatus
	RELEASE_CHILD_TRANSFER_COMPLETE QueryJobRespStatus
	NODE_UPGRADE_START              QueryJobRespStatus
	NODE_UPGRADE_COMPLETE           QueryJobRespStatus
	NODE_UPGRADE_FAILED             QueryJobRespStatus
}

func GetQueryJobRespStatusEnum() QueryJobRespStatusEnum {
	return QueryJobRespStatusEnum{
		CREATING: QueryJobRespStatus{
			value: "CREATING",
		},
		CREATE_FAILED: QueryJobRespStatus{
			value: "CREATE_FAILED",
		},
		CONFIGURATION: QueryJobRespStatus{
			value: "CONFIGURATION",
		},
		STARTJOBING: QueryJobRespStatus{
			value: "STARTJOBING",
		},
		WAITING_FOR_START: QueryJobRespStatus{
			value: "WAITING_FOR_START",
		},
		START_JOB_FAILED: QueryJobRespStatus{
			value: "START_JOB_FAILED",
		},
		PAUSING: QueryJobRespStatus{
			value: "PAUSING",
		},
		FULL_TRANSFER_STARTED: QueryJobRespStatus{
			value: "FULL_TRANSFER_STARTED",
		},
		FULL_TRANSFER_FAILED: QueryJobRespStatus{
			value: "FULL_TRANSFER_FAILED",
		},
		FULL_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "FULL_TRANSFER_COMPLETE",
		},
		INCRE_TRANSFER_STARTED: QueryJobRespStatus{
			value: "INCRE_TRANSFER_STARTED",
		},
		INCRE_TRANSFER_FAILED: QueryJobRespStatus{
			value: "INCRE_TRANSFER_FAILED",
		},
		RELEASE_RESOURCE_STARTED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_STARTED",
		},
		RELEASE_RESOURCE_FAILED: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_FAILED",
		},
		RELEASE_RESOURCE_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_RESOURCE_COMPLETE",
		},
		REBUILD_NODE_STARTED: QueryJobRespStatus{
			value: "REBUILD_NODE_STARTED",
		},
		REBUILD_NODE_FAILED: QueryJobRespStatus{
			value: "REBUILD_NODE_FAILED",
		},
		CHANGE_JOB_STARTED: QueryJobRespStatus{
			value: "CHANGE_JOB_STARTED",
		},
		CHANGE_JOB_FAILED: QueryJobRespStatus{
			value: "CHANGE_JOB_FAILED",
		},
		DELETED: QueryJobRespStatus{
			value: "DELETED",
		},
		CHILD_TRANSFER_STARTING: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTING",
		},
		CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_STARTED",
		},
		CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "CHILD_TRANSFER_COMPLETE",
		},
		CHILD_TRANSFER_FAILED: QueryJobRespStatus{
			value: "CHILD_TRANSFER_FAILED",
		},
		RELEASE_CHILD_TRANSFER_STARTED: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_STARTED",
		},
		RELEASE_CHILD_TRANSFER_COMPLETE: QueryJobRespStatus{
			value: "RELEASE_CHILD_TRANSFER_COMPLETE",
		},
		NODE_UPGRADE_START: QueryJobRespStatus{
			value: "NODE_UPGRADE_START",
		},
		NODE_UPGRADE_COMPLETE: QueryJobRespStatus{
			value: "NODE_UPGRADE_COMPLETE",
		},
		NODE_UPGRADE_FAILED: QueryJobRespStatus{
			value: "NODE_UPGRADE_FAILED",
		},
	}
}

func (c QueryJobRespStatus) Value() string {
	return c.value
}

func (c QueryJobRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespStatus) UnmarshalJSON(b []byte) error {
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

type QueryJobRespTaskType struct {
	value string
}

type QueryJobRespTaskTypeEnum struct {
	FULL_TRANS      QueryJobRespTaskType
	FULL_INCR_TRANS QueryJobRespTaskType
	INCR_TRANS      QueryJobRespTaskType
}

func GetQueryJobRespTaskTypeEnum() QueryJobRespTaskTypeEnum {
	return QueryJobRespTaskTypeEnum{
		FULL_TRANS: QueryJobRespTaskType{
			value: "FULL_TRANS",
		},
		FULL_INCR_TRANS: QueryJobRespTaskType{
			value: "FULL_INCR_TRANS",
		},
		INCR_TRANS: QueryJobRespTaskType{
			value: "INCR_TRANS",
		},
	}
}

func (c QueryJobRespTaskType) Value() string {
	return c.value
}

func (c QueryJobRespTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespTaskType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespNetType struct {
	value string
}

type QueryJobRespNetTypeEnum struct {
	VPN QueryJobRespNetType
	VPC QueryJobRespNetType
	EIP QueryJobRespNetType
}

func GetQueryJobRespNetTypeEnum() QueryJobRespNetTypeEnum {
	return QueryJobRespNetTypeEnum{
		VPN: QueryJobRespNetType{
			value: "vpn",
		},
		VPC: QueryJobRespNetType{
			value: "vpc",
		},
		EIP: QueryJobRespNetType{
			value: "eip",
		},
	}
}

func (c QueryJobRespNetType) Value() string {
	return c.value
}

func (c QueryJobRespNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespNetType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespJobDirection struct {
	value string
}

type QueryJobRespJobDirectionEnum struct {
	UP      QueryJobRespJobDirection
	DOWN    QueryJobRespJobDirection
	NON_DBS QueryJobRespJobDirection
}

func GetQueryJobRespJobDirectionEnum() QueryJobRespJobDirectionEnum {
	return QueryJobRespJobDirectionEnum{
		UP: QueryJobRespJobDirection{
			value: "up",
		},
		DOWN: QueryJobRespJobDirection{
			value: "down",
		},
		NON_DBS: QueryJobRespJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryJobRespJobDirection) Value() string {
	return c.value
}

func (c QueryJobRespJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespJobDirection) UnmarshalJSON(b []byte) error {
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

type QueryJobRespDbUseType struct {
	value string
}

type QueryJobRespDbUseTypeEnum struct {
	MIGRATION        QueryJobRespDbUseType
	SYNC             QueryJobRespDbUseType
	CLOUD_DATA_GUARD QueryJobRespDbUseType
}

func GetQueryJobRespDbUseTypeEnum() QueryJobRespDbUseTypeEnum {
	return QueryJobRespDbUseTypeEnum{
		MIGRATION: QueryJobRespDbUseType{
			value: "migration",
		},
		SYNC: QueryJobRespDbUseType{
			value: "sync",
		},
		CLOUD_DATA_GUARD: QueryJobRespDbUseType{
			value: "cloudDataGuard",
		},
	}
}

func (c QueryJobRespDbUseType) Value() string {
	return c.value
}

func (c QueryJobRespDbUseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespDbUseType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespConflictPolicy struct {
	value string
}

type QueryJobRespConflictPolicyEnum struct {
	STOP      QueryJobRespConflictPolicy
	OVERWRITE QueryJobRespConflictPolicy
	IGNORE    QueryJobRespConflictPolicy
}

func GetQueryJobRespConflictPolicyEnum() QueryJobRespConflictPolicyEnum {
	return QueryJobRespConflictPolicyEnum{
		STOP: QueryJobRespConflictPolicy{
			value: "stop",
		},
		OVERWRITE: QueryJobRespConflictPolicy{
			value: "overwrite",
		},
		IGNORE: QueryJobRespConflictPolicy{
			value: "ignore",
		},
	}
}

func (c QueryJobRespConflictPolicy) Value() string {
	return c.value
}

func (c QueryJobRespConflictPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespConflictPolicy) UnmarshalJSON(b []byte) error {
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

type QueryJobRespSchemaType struct {
	value string
}

type QueryJobRespSchemaTypeEnum struct {
	REPLICATION    QueryJobRespSchemaType
	TUNGSTEN       QueryJobRespSchemaType
	PG_BASE_BACKUP QueryJobRespSchemaType
}

func GetQueryJobRespSchemaTypeEnum() QueryJobRespSchemaTypeEnum {
	return QueryJobRespSchemaTypeEnum{
		REPLICATION: QueryJobRespSchemaType{
			value: "Replication",
		},
		TUNGSTEN: QueryJobRespSchemaType{
			value: "Tungsten",
		},
		PG_BASE_BACKUP: QueryJobRespSchemaType{
			value: "PGBaseBackup",
		},
	}
}

func (c QueryJobRespSchemaType) Value() string {
	return c.value
}

func (c QueryJobRespSchemaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespSchemaType) UnmarshalJSON(b []byte) error {
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

type QueryJobRespOriginalJobDirection struct {
	value string
}

type QueryJobRespOriginalJobDirectionEnum struct {
	UP      QueryJobRespOriginalJobDirection
	DOWN    QueryJobRespOriginalJobDirection
	NON_DBS QueryJobRespOriginalJobDirection
}

func GetQueryJobRespOriginalJobDirectionEnum() QueryJobRespOriginalJobDirectionEnum {
	return QueryJobRespOriginalJobDirectionEnum{
		UP: QueryJobRespOriginalJobDirection{
			value: "up",
		},
		DOWN: QueryJobRespOriginalJobDirection{
			value: "down",
		},
		NON_DBS: QueryJobRespOriginalJobDirection{
			value: "non-dbs",
		},
	}
}

func (c QueryJobRespOriginalJobDirection) Value() string {
	return c.value
}

func (c QueryJobRespOriginalJobDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QueryJobRespOriginalJobDirection) UnmarshalJSON(b []byte) error {
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

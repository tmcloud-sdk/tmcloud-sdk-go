package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type StopMigrationTaskResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Status *StopMigrationTaskResponseStatus `json:"status,omitempty"`

	MigrationType *StopMigrationTaskResponseMigrationType `json:"migration_type,omitempty"`

	MigrationMethod *StopMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty"`

	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	BackupFiles *BackupFilesBody `json:"backup_files,omitempty"`

	NetworkType *StopMigrationTaskResponseNetworkType `json:"network_type,omitempty"`

	SourceInstance *SourceInstanceBody `json:"source_instance,omitempty"`

	TargetInstance *TargetInstanceBody `json:"target_instance,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"StopMigrationTaskResponse", string(data)}, " ")
}

type StopMigrationTaskResponseStatus struct {
	value string
}

type StopMigrationTaskResponseStatusEnum struct {
	SUCCESS    StopMigrationTaskResponseStatus
	FAILED     StopMigrationTaskResponseStatus
	MIGRATING  StopMigrationTaskResponseStatus
	TERMINATED StopMigrationTaskResponseStatus
}

func GetStopMigrationTaskResponseStatusEnum() StopMigrationTaskResponseStatusEnum {
	return StopMigrationTaskResponseStatusEnum{
		SUCCESS: StopMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: StopMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: StopMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: StopMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c StopMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseMigrationType struct {
	value string
}

type StopMigrationTaskResponseMigrationTypeEnum struct {
	BACKUPFILE_IMPORT StopMigrationTaskResponseMigrationType
	ONLINE_MIGRATION  StopMigrationTaskResponseMigrationType
}

func GetStopMigrationTaskResponseMigrationTypeEnum() StopMigrationTaskResponseMigrationTypeEnum {
	return StopMigrationTaskResponseMigrationTypeEnum{
		BACKUPFILE_IMPORT: StopMigrationTaskResponseMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: StopMigrationTaskResponseMigrationType{
			value: "online_migration",
		},
	}
}

func (c StopMigrationTaskResponseMigrationType) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseMigrationType) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseMigrationMethod struct {
	value string
}

type StopMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION StopMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION StopMigrationTaskResponseMigrationMethod
}

func GetStopMigrationTaskResponseMigrationMethodEnum() StopMigrationTaskResponseMigrationMethodEnum {
	return StopMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: StopMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: StopMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c StopMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type StopMigrationTaskResponseNetworkType struct {
	value string
}

type StopMigrationTaskResponseNetworkTypeEnum struct {
	VPC StopMigrationTaskResponseNetworkType
	VPN StopMigrationTaskResponseNetworkType
}

func GetStopMigrationTaskResponseNetworkTypeEnum() StopMigrationTaskResponseNetworkTypeEnum {
	return StopMigrationTaskResponseNetworkTypeEnum{
		VPC: StopMigrationTaskResponseNetworkType{
			value: "vpc",
		},
		VPN: StopMigrationTaskResponseNetworkType{
			value: "vpn",
		},
	}
}

func (c StopMigrationTaskResponseNetworkType) Value() string {
	return c.value
}

func (c StopMigrationTaskResponseNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StopMigrationTaskResponseNetworkType) UnmarshalJSON(b []byte) error {
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

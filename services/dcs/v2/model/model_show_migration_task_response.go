package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowMigrationTaskResponse struct {
	TaskId *string `json:"task_id,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	Description *string `json:"description,omitempty"`

	Status *ShowMigrationTaskResponseStatus `json:"status,omitempty"`

	MigrationType *ShowMigrationTaskResponseMigrationType `json:"migration_type,omitempty"`

	MigrationMethod *ShowMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty"`

	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	BackupFiles *BackupFilesBody `json:"backup_files,omitempty"`

	NetworkType *ShowMigrationTaskResponseNetworkType `json:"network_type,omitempty"`

	SourceInstance *SourceInstanceBody `json:"source_instance,omitempty"`

	TargetInstance *TargetInstanceBody `json:"target_instance,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskResponse", string(data)}, " ")
}

type ShowMigrationTaskResponseStatus struct {
	value string
}

type ShowMigrationTaskResponseStatusEnum struct {
	SUCCESS    ShowMigrationTaskResponseStatus
	FAILED     ShowMigrationTaskResponseStatus
	MIGRATING  ShowMigrationTaskResponseStatus
	TERMINATED ShowMigrationTaskResponseStatus
}

func GetShowMigrationTaskResponseStatusEnum() ShowMigrationTaskResponseStatusEnum {
	return ShowMigrationTaskResponseStatusEnum{
		SUCCESS: ShowMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: ShowMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: ShowMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: ShowMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c ShowMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseMigrationType struct {
	value string
}

type ShowMigrationTaskResponseMigrationTypeEnum struct {
	BACKUPFILE_IMPORT ShowMigrationTaskResponseMigrationType
	ONLINE_MIGRATION  ShowMigrationTaskResponseMigrationType
}

func GetShowMigrationTaskResponseMigrationTypeEnum() ShowMigrationTaskResponseMigrationTypeEnum {
	return ShowMigrationTaskResponseMigrationTypeEnum{
		BACKUPFILE_IMPORT: ShowMigrationTaskResponseMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: ShowMigrationTaskResponseMigrationType{
			value: "online_migration",
		},
	}
}

func (c ShowMigrationTaskResponseMigrationType) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseMigrationType) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseMigrationMethod struct {
	value string
}

type ShowMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION ShowMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION ShowMigrationTaskResponseMigrationMethod
}

func GetShowMigrationTaskResponseMigrationMethodEnum() ShowMigrationTaskResponseMigrationMethodEnum {
	return ShowMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: ShowMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: ShowMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c ShowMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type ShowMigrationTaskResponseNetworkType struct {
	value string
}

type ShowMigrationTaskResponseNetworkTypeEnum struct {
	VPC ShowMigrationTaskResponseNetworkType
	VPN ShowMigrationTaskResponseNetworkType
}

func GetShowMigrationTaskResponseNetworkTypeEnum() ShowMigrationTaskResponseNetworkTypeEnum {
	return ShowMigrationTaskResponseNetworkTypeEnum{
		VPC: ShowMigrationTaskResponseNetworkType{
			value: "vpc",
		},
		VPN: ShowMigrationTaskResponseNetworkType{
			value: "vpn",
		},
	}
}

func (c ShowMigrationTaskResponseNetworkType) Value() string {
	return c.value
}

func (c ShowMigrationTaskResponseNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMigrationTaskResponseNetworkType) UnmarshalJSON(b []byte) error {
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

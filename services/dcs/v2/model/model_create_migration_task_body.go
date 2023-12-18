package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateMigrationTaskBody struct {
	TaskName string `json:"task_name"`

	Description *string `json:"description,omitempty"`

	MigrationType CreateMigrationTaskBodyMigrationType `json:"migration_type"`

	MigrationMethod CreateMigrationTaskBodyMigrationMethod `json:"migration_method"`

	BackupFiles *BackupFilesBody `json:"backup_files,omitempty"`

	NetworkType *CreateMigrationTaskBodyNetworkType `json:"network_type,omitempty"`

	SourceInstance *SourceInstanceBody `json:"source_instance,omitempty"`

	TargetInstance *TargetInstanceBody `json:"target_instance"`
}

func (o CreateMigrationTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMigrationTaskBody struct{}"
	}

	return strings.Join([]string{"CreateMigrationTaskBody", string(data)}, " ")
}

type CreateMigrationTaskBodyMigrationType struct {
	value string
}

type CreateMigrationTaskBodyMigrationTypeEnum struct {
	BACKUPFILE_IMPORT CreateMigrationTaskBodyMigrationType
	ONLINE_MIGRATION  CreateMigrationTaskBodyMigrationType
}

func GetCreateMigrationTaskBodyMigrationTypeEnum() CreateMigrationTaskBodyMigrationTypeEnum {
	return CreateMigrationTaskBodyMigrationTypeEnum{
		BACKUPFILE_IMPORT: CreateMigrationTaskBodyMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: CreateMigrationTaskBodyMigrationType{
			value: "online_migration",
		},
	}
}

func (c CreateMigrationTaskBodyMigrationType) Value() string {
	return c.value
}

func (c CreateMigrationTaskBodyMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMigrationTaskBodyMigrationType) UnmarshalJSON(b []byte) error {
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

type CreateMigrationTaskBodyMigrationMethod struct {
	value string
}

type CreateMigrationTaskBodyMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION CreateMigrationTaskBodyMigrationMethod
	INCREMENTAL_MIGRATION CreateMigrationTaskBodyMigrationMethod
}

func GetCreateMigrationTaskBodyMigrationMethodEnum() CreateMigrationTaskBodyMigrationMethodEnum {
	return CreateMigrationTaskBodyMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: CreateMigrationTaskBodyMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: CreateMigrationTaskBodyMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c CreateMigrationTaskBodyMigrationMethod) Value() string {
	return c.value
}

func (c CreateMigrationTaskBodyMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMigrationTaskBodyMigrationMethod) UnmarshalJSON(b []byte) error {
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

type CreateMigrationTaskBodyNetworkType struct {
	value string
}

type CreateMigrationTaskBodyNetworkTypeEnum struct {
	VPC CreateMigrationTaskBodyNetworkType
	VPN CreateMigrationTaskBodyNetworkType
}

func GetCreateMigrationTaskBodyNetworkTypeEnum() CreateMigrationTaskBodyNetworkTypeEnum {
	return CreateMigrationTaskBodyNetworkTypeEnum{
		VPC: CreateMigrationTaskBodyNetworkType{
			value: "vpc",
		},
		VPN: CreateMigrationTaskBodyNetworkType{
			value: "vpn",
		},
	}
}

func (c CreateMigrationTaskBodyNetworkType) Value() string {
	return c.value
}

func (c CreateMigrationTaskBodyNetworkType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMigrationTaskBodyNetworkType) UnmarshalJSON(b []byte) error {
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

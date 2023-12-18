package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type MigrationTaskList struct {
	TaskId *string `json:"task_id,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	Status *MigrationTaskListStatus `json:"status,omitempty"`

	MigrationType *MigrationTaskListMigrationType `json:"migration_type,omitempty"`

	MigrationMethod *MigrationTaskListMigrationMethod `json:"migration_method,omitempty"`

	EcsTenantPrivateIp *string `json:"ecs_tenant_private_ip,omitempty"`

	DataSource *string `json:"data_source,omitempty"`

	SourceInstanceName *string `json:"source_instance_name,omitempty"`

	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	TargetInstanceAddrs *string `json:"target_instance_addrs,omitempty"`

	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`
}

func (o MigrationTaskList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationTaskList struct{}"
	}

	return strings.Join([]string{"MigrationTaskList", string(data)}, " ")
}

type MigrationTaskListStatus struct {
	value string
}

type MigrationTaskListStatusEnum struct {
	SUCCESS    MigrationTaskListStatus
	FAILED     MigrationTaskListStatus
	MIGRATING  MigrationTaskListStatus
	TERMINATED MigrationTaskListStatus
}

func GetMigrationTaskListStatusEnum() MigrationTaskListStatusEnum {
	return MigrationTaskListStatusEnum{
		SUCCESS: MigrationTaskListStatus{
			value: "SUCCESS",
		},
		FAILED: MigrationTaskListStatus{
			value: "FAILED",
		},
		MIGRATING: MigrationTaskListStatus{
			value: "MIGRATING",
		},
		TERMINATED: MigrationTaskListStatus{
			value: "TERMINATED",
		},
	}
}

func (c MigrationTaskListStatus) Value() string {
	return c.value
}

func (c MigrationTaskListStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListStatus) UnmarshalJSON(b []byte) error {
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

type MigrationTaskListMigrationType struct {
	value string
}

type MigrationTaskListMigrationTypeEnum struct {
	BACKUPFILE_IMPORT MigrationTaskListMigrationType
	ONLINE_MIGRATION  MigrationTaskListMigrationType
}

func GetMigrationTaskListMigrationTypeEnum() MigrationTaskListMigrationTypeEnum {
	return MigrationTaskListMigrationTypeEnum{
		BACKUPFILE_IMPORT: MigrationTaskListMigrationType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: MigrationTaskListMigrationType{
			value: "online_migration",
		},
	}
}

func (c MigrationTaskListMigrationType) Value() string {
	return c.value
}

func (c MigrationTaskListMigrationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationType) UnmarshalJSON(b []byte) error {
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

type MigrationTaskListMigrationMethod struct {
	value string
}

type MigrationTaskListMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION MigrationTaskListMigrationMethod
	INCREMENTAL_MIGRATION MigrationTaskListMigrationMethod
}

func GetMigrationTaskListMigrationMethodEnum() MigrationTaskListMigrationMethodEnum {
	return MigrationTaskListMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: MigrationTaskListMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: MigrationTaskListMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c MigrationTaskListMigrationMethod) Value() string {
	return c.value
}

func (c MigrationTaskListMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MigrationTaskListMigrationMethod) UnmarshalJSON(b []byte) error {
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

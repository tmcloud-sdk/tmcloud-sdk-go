package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListMigrationTaskResponse struct {
	Count *int32 `json:"count,omitempty"`

	MigrationTasks *[]MigrationTaskList `json:"migration_tasks,omitempty"`

	TargetInstanceAddress *string `json:"target_instance_address,omitempty"`

	MigrationMethod *ListMigrationTaskResponseMigrationMethod `json:"migration_method,omitempty"`

	TaskName *string `json:"task_name,omitempty"`

	TargetInstanceId *string `json:"target_instance_id,omitempty"`

	SourceInstanceName *string `json:"source_instance_name,omitempty"`

	TargetInstanceName *string `json:"target_instance_name,omitempty"`

	MigrateType *ListMigrationTaskResponseMigrateType `json:"migrate_type,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	SourceInstanceId *string `json:"source_instance_id,omitempty"`

	TaskId *string `json:"task_id,omitempty"`

	DataSource *string `json:"data_source,omitempty"`

	Status         *ListMigrationTaskResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListMigrationTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskResponse struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskResponse", string(data)}, " ")
}

type ListMigrationTaskResponseMigrationMethod struct {
	value string
}

type ListMigrationTaskResponseMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION ListMigrationTaskResponseMigrationMethod
	INCREMENTAL_MIGRATION ListMigrationTaskResponseMigrationMethod
}

func GetListMigrationTaskResponseMigrationMethodEnum() ListMigrationTaskResponseMigrationMethodEnum {
	return ListMigrationTaskResponseMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: ListMigrationTaskResponseMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: ListMigrationTaskResponseMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c ListMigrationTaskResponseMigrationMethod) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseMigrationMethod) UnmarshalJSON(b []byte) error {
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

type ListMigrationTaskResponseMigrateType struct {
	value string
}

type ListMigrationTaskResponseMigrateTypeEnum struct {
	BACKUPFILE_IMPORT ListMigrationTaskResponseMigrateType
	ONLINE_MIGRATION  ListMigrationTaskResponseMigrateType
}

func GetListMigrationTaskResponseMigrateTypeEnum() ListMigrationTaskResponseMigrateTypeEnum {
	return ListMigrationTaskResponseMigrateTypeEnum{
		BACKUPFILE_IMPORT: ListMigrationTaskResponseMigrateType{
			value: "backupfile_import",
		},
		ONLINE_MIGRATION: ListMigrationTaskResponseMigrateType{
			value: "online_migration",
		},
	}
}

func (c ListMigrationTaskResponseMigrateType) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseMigrateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseMigrateType) UnmarshalJSON(b []byte) error {
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

type ListMigrationTaskResponseStatus struct {
	value string
}

type ListMigrationTaskResponseStatusEnum struct {
	SUCCESS    ListMigrationTaskResponseStatus
	FAILED     ListMigrationTaskResponseStatus
	MIGRATING  ListMigrationTaskResponseStatus
	TERMINATED ListMigrationTaskResponseStatus
}

func GetListMigrationTaskResponseStatusEnum() ListMigrationTaskResponseStatusEnum {
	return ListMigrationTaskResponseStatusEnum{
		SUCCESS: ListMigrationTaskResponseStatus{
			value: "SUCCESS",
		},
		FAILED: ListMigrationTaskResponseStatus{
			value: "FAILED",
		},
		MIGRATING: ListMigrationTaskResponseStatus{
			value: "MIGRATING",
		},
		TERMINATED: ListMigrationTaskResponseStatus{
			value: "TERMINATED",
		},
	}
}

func (c ListMigrationTaskResponseStatus) Value() string {
	return c.value
}

func (c ListMigrationTaskResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListMigrationTaskResponseStatus) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type SetOnlineMigrationTaskBody struct {
	MigrationMethod SetOnlineMigrationTaskBodyMigrationMethod `json:"migration_method"`

	ResumeMode SetOnlineMigrationTaskBodyResumeMode `json:"resume_mode"`

	BandwidthLimitMb *string `json:"bandwidth_limit_mb,omitempty"`

	SourceInstance *ConfigMigrationInstanceBody `json:"source_instance"`

	TargetInstance *ConfigMigrationInstanceBody `json:"target_instance"`
}

func (o SetOnlineMigrationTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOnlineMigrationTaskBody struct{}"
	}

	return strings.Join([]string{"SetOnlineMigrationTaskBody", string(data)}, " ")
}

type SetOnlineMigrationTaskBodyMigrationMethod struct {
	value string
}

type SetOnlineMigrationTaskBodyMigrationMethodEnum struct {
	FULL_AMOUNT_MIGRATION SetOnlineMigrationTaskBodyMigrationMethod
	INCREMENTAL_MIGRATION SetOnlineMigrationTaskBodyMigrationMethod
}

func GetSetOnlineMigrationTaskBodyMigrationMethodEnum() SetOnlineMigrationTaskBodyMigrationMethodEnum {
	return SetOnlineMigrationTaskBodyMigrationMethodEnum{
		FULL_AMOUNT_MIGRATION: SetOnlineMigrationTaskBodyMigrationMethod{
			value: "full_amount_migration",
		},
		INCREMENTAL_MIGRATION: SetOnlineMigrationTaskBodyMigrationMethod{
			value: "incremental_migration",
		},
	}
}

func (c SetOnlineMigrationTaskBodyMigrationMethod) Value() string {
	return c.value
}

func (c SetOnlineMigrationTaskBodyMigrationMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskBodyMigrationMethod) UnmarshalJSON(b []byte) error {
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

type SetOnlineMigrationTaskBodyResumeMode struct {
	value string
}

type SetOnlineMigrationTaskBodyResumeModeEnum struct {
	AUTO   SetOnlineMigrationTaskBodyResumeMode
	MANUAL SetOnlineMigrationTaskBodyResumeMode
}

func GetSetOnlineMigrationTaskBodyResumeModeEnum() SetOnlineMigrationTaskBodyResumeModeEnum {
	return SetOnlineMigrationTaskBodyResumeModeEnum{
		AUTO: SetOnlineMigrationTaskBodyResumeMode{
			value: "auto",
		},
		MANUAL: SetOnlineMigrationTaskBodyResumeMode{
			value: "manual",
		},
	}
}

func (c SetOnlineMigrationTaskBodyResumeMode) Value() string {
	return c.value
}

func (c SetOnlineMigrationTaskBodyResumeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetOnlineMigrationTaskBodyResumeMode) UnmarshalJSON(b []byte) error {
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

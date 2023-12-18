package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BackupInfo struct {
	Id string `json:"id"`

	InstanceId string `json:"instance_id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Databases *[]BackupDatabase `json:"databases,omitempty"`

	BeginTime string `json:"begin_time"`

	Status BackupInfoStatus `json:"status"`

	Type BackupInfoType `json:"type"`
}

func (o BackupInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInfo struct{}"
	}

	return strings.Join([]string{"BackupInfo", string(data)}, " ")
}

type BackupInfoStatus struct {
	value string
}

type BackupInfoStatusEnum struct {
	BUILDING  BackupInfoStatus
	COMPLETED BackupInfoStatus
	FAILED    BackupInfoStatus
	DELETING  BackupInfoStatus
}

func GetBackupInfoStatusEnum() BackupInfoStatusEnum {
	return BackupInfoStatusEnum{
		BUILDING: BackupInfoStatus{
			value: "BUILDING",
		},
		COMPLETED: BackupInfoStatus{
			value: "COMPLETED",
		},
		FAILED: BackupInfoStatus{
			value: "FAILED",
		},
		DELETING: BackupInfoStatus{
			value: "DELETING",
		},
	}
}

func (c BackupInfoStatus) Value() string {
	return c.value
}

func (c BackupInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInfoStatus) UnmarshalJSON(b []byte) error {
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

type BackupInfoType struct {
	value string
}

type BackupInfoTypeEnum struct {
	AUTO        BackupInfoType
	MANUAL      BackupInfoType
	FRAGMENT    BackupInfoType
	INCREMENTAL BackupInfoType
}

func GetBackupInfoTypeEnum() BackupInfoTypeEnum {
	return BackupInfoTypeEnum{
		AUTO: BackupInfoType{
			value: "auto",
		},
		MANUAL: BackupInfoType{
			value: "manual",
		},
		FRAGMENT: BackupInfoType{
			value: "fragment",
		},
		INCREMENTAL: BackupInfoType{
			value: "incremental",
		},
	}
}

func (c BackupInfoType) Value() string {
	return c.value
}

func (c BackupInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInfoType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type VersionInfo struct {
	Id *VersionInfoId `json:"id,omitempty"`

	Links *[]Links `json:"links,omitempty"`

	MinVersion *string `json:"min_version,omitempty"`

	Status *VersionInfoStatus `json:"status,omitempty"`

	Update *sdktime.SdkTime `json:"update,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (o VersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionInfo struct{}"
	}

	return strings.Join([]string{"VersionInfo", string(data)}, " ")
}

type VersionInfoId struct {
	value string
}

type VersionInfoIdEnum struct {
	V1 VersionInfoId
	V2 VersionInfoId
}

func GetVersionInfoIdEnum() VersionInfoIdEnum {
	return VersionInfoIdEnum{
		V1: VersionInfoId{
			value: "v1",
		},
		V2: VersionInfoId{
			value: "v2",
		},
	}
}

func (c VersionInfoId) Value() string {
	return c.value
}

func (c VersionInfoId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionInfoId) UnmarshalJSON(b []byte) error {
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

type VersionInfoStatus struct {
	value string
}

type VersionInfoStatusEnum struct {
	CURRENT    VersionInfoStatus
	SUPPORT    VersionInfoStatus
	DEPRECATED VersionInfoStatus
}

func GetVersionInfoStatusEnum() VersionInfoStatusEnum {
	return VersionInfoStatusEnum{
		CURRENT: VersionInfoStatus{
			value: "CURRENT",
		},
		SUPPORT: VersionInfoStatus{
			value: "SUPPORT",
		},
		DEPRECATED: VersionInfoStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionInfoStatus) Value() string {
	return c.value
}

func (c VersionInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionInfoStatus) UnmarshalJSON(b []byte) error {
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

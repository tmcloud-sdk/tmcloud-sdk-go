package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DiskInfo struct {
	Size int32 `json:"size"`

	VolumeType DiskInfoVolumeType `json:"volume_type"`

	DiskType DiskInfoDiskType `json:"disk_type"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	DataDiskImageId *string `json:"data_disk_image_id,omitempty"`

	SnapshotId *string `json:"snapshot_id,omitempty"`

	Metadata *MetaData `json:"metadata,omitempty"`
}

func (o DiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskInfo struct{}"
	}

	return strings.Join([]string{"DiskInfo", string(data)}, " ")
}

type DiskInfoVolumeType struct {
	value string
}

type DiskInfoVolumeTypeEnum struct {
	SATA  DiskInfoVolumeType
	SAS   DiskInfoVolumeType
	SSD   DiskInfoVolumeType
	CO_PL DiskInfoVolumeType
	UH_11 DiskInfoVolumeType
	GPSSD DiskInfoVolumeType
}

func GetDiskInfoVolumeTypeEnum() DiskInfoVolumeTypeEnum {
	return DiskInfoVolumeTypeEnum{
		SATA: DiskInfoVolumeType{
			value: "SATA",
		},
		SAS: DiskInfoVolumeType{
			value: "SAS",
		},
		SSD: DiskInfoVolumeType{
			value: "SSD",
		},
		CO_PL: DiskInfoVolumeType{
			value: "co-pl",
		},
		UH_11: DiskInfoVolumeType{
			value: "uh-11",
		},
		GPSSD: DiskInfoVolumeType{
			value: "GPSSD",
		},
	}
}

func (c DiskInfoVolumeType) Value() string {
	return c.value
}

func (c DiskInfoVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskInfoVolumeType) UnmarshalJSON(b []byte) error {
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

type DiskInfoDiskType struct {
	value string
}

type DiskInfoDiskTypeEnum struct {
	SYS  DiskInfoDiskType
	DATA DiskInfoDiskType
}

func GetDiskInfoDiskTypeEnum() DiskInfoDiskTypeEnum {
	return DiskInfoDiskTypeEnum{
		SYS: DiskInfoDiskType{
			value: "SYS",
		},
		DATA: DiskInfoDiskType{
			value: "DATA",
		},
	}
}

func (c DiskInfoDiskType) Value() string {
	return c.value
}

func (c DiskInfoDiskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskInfoDiskType) UnmarshalJSON(b []byte) error {
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

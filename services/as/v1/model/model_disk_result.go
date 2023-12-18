package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DiskResult struct {
	Size *int32 `json:"size,omitempty"`

	VolumeType *DiskResultVolumeType `json:"volume_type,omitempty"`

	DiskType *DiskResultDiskType `json:"disk_type,omitempty"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	DataDiskImageId *string `json:"data_disk_image_id,omitempty"`

	SnapshotId *string `json:"snapshot_id,omitempty"`

	Metadata *MetaData `json:"metadata,omitempty"`
}

func (o DiskResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskResult struct{}"
	}

	return strings.Join([]string{"DiskResult", string(data)}, " ")
}

type DiskResultVolumeType struct {
	value string
}

type DiskResultVolumeTypeEnum struct {
	SATA  DiskResultVolumeType
	SAS   DiskResultVolumeType
	SSD   DiskResultVolumeType
	CO_PL DiskResultVolumeType
	UH_11 DiskResultVolumeType
}

func GetDiskResultVolumeTypeEnum() DiskResultVolumeTypeEnum {
	return DiskResultVolumeTypeEnum{
		SATA: DiskResultVolumeType{
			value: "SATA",
		},
		SAS: DiskResultVolumeType{
			value: "SAS",
		},
		SSD: DiskResultVolumeType{
			value: "SSD",
		},
		CO_PL: DiskResultVolumeType{
			value: "co-pl",
		},
		UH_11: DiskResultVolumeType{
			value: "uh-11",
		},
	}
}

func (c DiskResultVolumeType) Value() string {
	return c.value
}

func (c DiskResultVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskResultVolumeType) UnmarshalJSON(b []byte) error {
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

type DiskResultDiskType struct {
	value string
}

type DiskResultDiskTypeEnum struct {
	SYS  DiskResultDiskType
	DATA DiskResultDiskType
}

func GetDiskResultDiskTypeEnum() DiskResultDiskTypeEnum {
	return DiskResultDiskTypeEnum{
		SYS: DiskResultDiskType{
			value: "SYS",
		},
		DATA: DiskResultDiskType{
			value: "DATA",
		},
	}
}

func (c DiskResultDiskType) Value() string {
	return c.value
}

func (c DiskResultDiskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskResultDiskType) UnmarshalJSON(b []byte) error {
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

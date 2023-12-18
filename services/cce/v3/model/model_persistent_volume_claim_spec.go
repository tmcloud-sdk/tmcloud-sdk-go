package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PersistentVolumeClaimSpec struct {
	VolumeID string `json:"volumeID"`

	StorageType string `json:"storageType"`

	AccessModes []PersistentVolumeClaimSpecAccessModes `json:"accessModes"`

	StorageClassName *string `json:"storageClassName,omitempty"`

	VolumeName *string `json:"volumeName,omitempty"`

	Resources *ResourceRequirements `json:"resources,omitempty"`

	VolumeMode *string `json:"volumeMode,omitempty"`
}

func (o PersistentVolumeClaimSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaimSpec struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaimSpec", string(data)}, " ")
}

type PersistentVolumeClaimSpecAccessModes struct {
	value string
}

type PersistentVolumeClaimSpecAccessModesEnum struct {
	READ_ONLY_MANY  PersistentVolumeClaimSpecAccessModes
	READ_WRITE_MANY PersistentVolumeClaimSpecAccessModes
}

func GetPersistentVolumeClaimSpecAccessModesEnum() PersistentVolumeClaimSpecAccessModesEnum {
	return PersistentVolumeClaimSpecAccessModesEnum{
		READ_ONLY_MANY: PersistentVolumeClaimSpecAccessModes{
			value: "ReadOnlyMany",
		},
		READ_WRITE_MANY: PersistentVolumeClaimSpecAccessModes{
			value: "ReadWriteMany",
		},
	}
}

func (c PersistentVolumeClaimSpecAccessModes) Value() string {
	return c.value
}

func (c PersistentVolumeClaimSpecAccessModes) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PersistentVolumeClaimSpecAccessModes) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrePaidServerRootVolume struct {
	Volumetype PrePaidServerRootVolumeVolumetype `json:"volumetype"`

	Size *int32 `json:"size,omitempty"`

	Extendparam *PrePaidServerRootVolumeExtendParam `json:"extendparam,omitempty"`

	ClusterType *PrePaidServerRootVolumeClusterType `json:"cluster_type,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`
}

func (o PrePaidServerRootVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerRootVolume struct{}"
	}

	return strings.Join([]string{"PrePaidServerRootVolume", string(data)}, " ")
}

type PrePaidServerRootVolumeVolumetype struct {
	value string
}

type PrePaidServerRootVolumeVolumetypeEnum struct {
	SATA  PrePaidServerRootVolumeVolumetype
	SAS   PrePaidServerRootVolumeVolumetype
	SSD   PrePaidServerRootVolumeVolumetype
	GPSSD PrePaidServerRootVolumeVolumetype
	CO_P1 PrePaidServerRootVolumeVolumetype
	UH_L1 PrePaidServerRootVolumeVolumetype
	ESSD  PrePaidServerRootVolumeVolumetype
}

func GetPrePaidServerRootVolumeVolumetypeEnum() PrePaidServerRootVolumeVolumetypeEnum {
	return PrePaidServerRootVolumeVolumetypeEnum{
		SATA: PrePaidServerRootVolumeVolumetype{
			value: "SATA",
		},
		SAS: PrePaidServerRootVolumeVolumetype{
			value: "SAS",
		},
		SSD: PrePaidServerRootVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PrePaidServerRootVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PrePaidServerRootVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PrePaidServerRootVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PrePaidServerRootVolumeVolumetype{
			value: "ESSD",
		},
	}
}

func (c PrePaidServerRootVolumeVolumetype) Value() string {
	return c.value
}

func (c PrePaidServerRootVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerRootVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PrePaidServerRootVolumeClusterType struct {
	value string
}

type PrePaidServerRootVolumeClusterTypeEnum struct {
	DSS PrePaidServerRootVolumeClusterType
}

func GetPrePaidServerRootVolumeClusterTypeEnum() PrePaidServerRootVolumeClusterTypeEnum {
	return PrePaidServerRootVolumeClusterTypeEnum{
		DSS: PrePaidServerRootVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PrePaidServerRootVolumeClusterType) Value() string {
	return c.value
}

func (c PrePaidServerRootVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerRootVolumeClusterType) UnmarshalJSON(b []byte) error {
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

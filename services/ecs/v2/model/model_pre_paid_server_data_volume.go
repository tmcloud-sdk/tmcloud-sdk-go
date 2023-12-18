package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PrePaidServerDataVolume struct {
	Volumetype PrePaidServerDataVolumeVolumetype `json:"volumetype"`

	Size int32 `json:"size"`

	Shareable *bool `json:"shareable,omitempty"`

	Multiattach *bool `json:"multiattach,omitempty"`

	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	Extendparam *PrePaidServerDataVolumeExtendParam `json:"extendparam,omitempty"`

	ClusterType *PrePaidServerDataVolumeClusterType `json:"cluster_type,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Metadata *PrePaidServerDataVolumeMetadata `json:"metadata,omitempty"`

	DataImageId *string `json:"data_image_id,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o PrePaidServerDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServerDataVolume struct{}"
	}

	return strings.Join([]string{"PrePaidServerDataVolume", string(data)}, " ")
}

type PrePaidServerDataVolumeVolumetype struct {
	value string
}

type PrePaidServerDataVolumeVolumetypeEnum struct {
	SATA  PrePaidServerDataVolumeVolumetype
	SAS   PrePaidServerDataVolumeVolumetype
	SSD   PrePaidServerDataVolumeVolumetype
	GPSSD PrePaidServerDataVolumeVolumetype
	CO_P1 PrePaidServerDataVolumeVolumetype
	UH_L1 PrePaidServerDataVolumeVolumetype
	ESSD  PrePaidServerDataVolumeVolumetype
}

func GetPrePaidServerDataVolumeVolumetypeEnum() PrePaidServerDataVolumeVolumetypeEnum {
	return PrePaidServerDataVolumeVolumetypeEnum{
		SATA: PrePaidServerDataVolumeVolumetype{
			value: "SATA",
		},
		SAS: PrePaidServerDataVolumeVolumetype{
			value: "SAS",
		},
		SSD: PrePaidServerDataVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PrePaidServerDataVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PrePaidServerDataVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PrePaidServerDataVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PrePaidServerDataVolumeVolumetype{
			value: "ESSD",
		},
	}
}

func (c PrePaidServerDataVolumeVolumetype) Value() string {
	return c.value
}

func (c PrePaidServerDataVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerDataVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PrePaidServerDataVolumeClusterType struct {
	value string
}

type PrePaidServerDataVolumeClusterTypeEnum struct {
	DSS PrePaidServerDataVolumeClusterType
}

func GetPrePaidServerDataVolumeClusterTypeEnum() PrePaidServerDataVolumeClusterTypeEnum {
	return PrePaidServerDataVolumeClusterTypeEnum{
		DSS: PrePaidServerDataVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PrePaidServerDataVolumeClusterType) Value() string {
	return c.value
}

func (c PrePaidServerDataVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrePaidServerDataVolumeClusterType) UnmarshalJSON(b []byte) error {
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

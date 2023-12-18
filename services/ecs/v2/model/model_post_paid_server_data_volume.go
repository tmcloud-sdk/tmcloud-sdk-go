package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PostPaidServerDataVolume struct {
	Volumetype PostPaidServerDataVolumeVolumetype `json:"volumetype"`

	Size int32 `json:"size"`

	Shareable *bool `json:"shareable,omitempty"`

	Multiattach *bool `json:"multiattach,omitempty"`

	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	Extendparam *PostPaidServerDataVolumeExtendParam `json:"extendparam,omitempty"`

	ClusterType *PostPaidServerDataVolumeClusterType `json:"cluster_type,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Metadata *PostPaidServerDataVolumeMetadata `json:"metadata,omitempty"`

	DataImageId *string `json:"data_image_id,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`
}

func (o PostPaidServerDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerDataVolume struct{}"
	}

	return strings.Join([]string{"PostPaidServerDataVolume", string(data)}, " ")
}

type PostPaidServerDataVolumeVolumetype struct {
	value string
}

type PostPaidServerDataVolumeVolumetypeEnum struct {
	SATA  PostPaidServerDataVolumeVolumetype
	SAS   PostPaidServerDataVolumeVolumetype
	SSD   PostPaidServerDataVolumeVolumetype
	GPSSD PostPaidServerDataVolumeVolumetype
	CO_P1 PostPaidServerDataVolumeVolumetype
	UH_L1 PostPaidServerDataVolumeVolumetype
	ESSD  PostPaidServerDataVolumeVolumetype
}

func GetPostPaidServerDataVolumeVolumetypeEnum() PostPaidServerDataVolumeVolumetypeEnum {
	return PostPaidServerDataVolumeVolumetypeEnum{
		SATA: PostPaidServerDataVolumeVolumetype{
			value: "SATA",
		},
		SAS: PostPaidServerDataVolumeVolumetype{
			value: "SAS",
		},
		SSD: PostPaidServerDataVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PostPaidServerDataVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PostPaidServerDataVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PostPaidServerDataVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PostPaidServerDataVolumeVolumetype{
			value: "ESSD",
		},
	}
}

func (c PostPaidServerDataVolumeVolumetype) Value() string {
	return c.value
}

func (c PostPaidServerDataVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerDataVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PostPaidServerDataVolumeClusterType struct {
	value string
}

type PostPaidServerDataVolumeClusterTypeEnum struct {
	DSS PostPaidServerDataVolumeClusterType
}

func GetPostPaidServerDataVolumeClusterTypeEnum() PostPaidServerDataVolumeClusterTypeEnum {
	return PostPaidServerDataVolumeClusterTypeEnum{
		DSS: PostPaidServerDataVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PostPaidServerDataVolumeClusterType) Value() string {
	return c.value
}

func (c PostPaidServerDataVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerDataVolumeClusterType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PostPaidServerRootVolume struct {
	Volumetype PostPaidServerRootVolumeVolumetype `json:"volumetype"`

	Size *int32 `json:"size,omitempty"`

	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	ClusterType *PostPaidServerRootVolumeClusterType `json:"cluster_type,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	Extendparam *PostPaidServerRootVolumeExtendParam `json:"extendparam,omitempty"`
}

func (o PostPaidServerRootVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerRootVolume struct{}"
	}

	return strings.Join([]string{"PostPaidServerRootVolume", string(data)}, " ")
}

type PostPaidServerRootVolumeVolumetype struct {
	value string
}

type PostPaidServerRootVolumeVolumetypeEnum struct {
	SATA  PostPaidServerRootVolumeVolumetype
	SAS   PostPaidServerRootVolumeVolumetype
	SSD   PostPaidServerRootVolumeVolumetype
	GPSSD PostPaidServerRootVolumeVolumetype
	CO_P1 PostPaidServerRootVolumeVolumetype
	UH_L1 PostPaidServerRootVolumeVolumetype
	ESSD  PostPaidServerRootVolumeVolumetype
}

func GetPostPaidServerRootVolumeVolumetypeEnum() PostPaidServerRootVolumeVolumetypeEnum {
	return PostPaidServerRootVolumeVolumetypeEnum{
		SATA: PostPaidServerRootVolumeVolumetype{
			value: "SATA",
		},
		SAS: PostPaidServerRootVolumeVolumetype{
			value: "SAS",
		},
		SSD: PostPaidServerRootVolumeVolumetype{
			value: "SSD",
		},
		GPSSD: PostPaidServerRootVolumeVolumetype{
			value: "GPSSD",
		},
		CO_P1: PostPaidServerRootVolumeVolumetype{
			value: "co-p1",
		},
		UH_L1: PostPaidServerRootVolumeVolumetype{
			value: "uh-l1",
		},
		ESSD: PostPaidServerRootVolumeVolumetype{
			value: "ESSD",
		},
	}
}

func (c PostPaidServerRootVolumeVolumetype) Value() string {
	return c.value
}

func (c PostPaidServerRootVolumeVolumetype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerRootVolumeVolumetype) UnmarshalJSON(b []byte) error {
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

type PostPaidServerRootVolumeClusterType struct {
	value string
}

type PostPaidServerRootVolumeClusterTypeEnum struct {
	DSS PostPaidServerRootVolumeClusterType
}

func GetPostPaidServerRootVolumeClusterTypeEnum() PostPaidServerRootVolumeClusterTypeEnum {
	return PostPaidServerRootVolumeClusterTypeEnum{
		DSS: PostPaidServerRootVolumeClusterType{
			value: "DSS",
		},
	}
}

func (c PostPaidServerRootVolumeClusterType) Value() string {
	return c.value
}

func (c PostPaidServerRootVolumeClusterType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerRootVolumeClusterType) UnmarshalJSON(b []byte) error {
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

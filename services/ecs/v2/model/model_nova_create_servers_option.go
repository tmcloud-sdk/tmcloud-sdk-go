package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaCreateServersOption struct {
	AutoTerminateTime *string `json:"auto_terminate_time,omitempty"`

	ImageRef *string `json:"imageRef,omitempty"`

	FlavorRef string `json:"flavorRef"`

	Name string `json:"name"`

	Metadata map[string]string `json:"metadata,omitempty"`

	AdminPass *string `json:"adminPass,omitempty"`

	BlockDeviceMappingV2 *[]NovaServerBlockDeviceMapping `json:"block_device_mapping_v2,omitempty"`

	ConfigDrive *string `json:"config_drive,omitempty"`

	SecurityGroups *[]NovaServerSecurityGroup `json:"security_groups,omitempty"`

	Networks []NovaServerNetwork `json:"networks"`

	KeyName *string `json:"key_name,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	ReturnReservationId *bool `json:"return_reservation_id,omitempty"`

	MinCount *int32 `json:"min_count,omitempty"`

	MaxCount *int32 `json:"max_count,omitempty"`

	OSDCFdiskConfig *NovaCreateServersOptionOSDCFdiskConfig `json:"OS-DCF:diskConfig,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o NovaCreateServersOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersOption struct{}"
	}

	return strings.Join([]string{"NovaCreateServersOption", string(data)}, " ")
}

type NovaCreateServersOptionOSDCFdiskConfig struct {
	value string
}

type NovaCreateServersOptionOSDCFdiskConfigEnum struct {
	AUTO   NovaCreateServersOptionOSDCFdiskConfig
	MANUAL NovaCreateServersOptionOSDCFdiskConfig
}

func GetNovaCreateServersOptionOSDCFdiskConfigEnum() NovaCreateServersOptionOSDCFdiskConfigEnum {
	return NovaCreateServersOptionOSDCFdiskConfigEnum{
		AUTO: NovaCreateServersOptionOSDCFdiskConfig{
			value: "AUTO",
		},
		MANUAL: NovaCreateServersOptionOSDCFdiskConfig{
			value: "MANUAL",
		},
	}
}

func (c NovaCreateServersOptionOSDCFdiskConfig) Value() string {
	return c.value
}

func (c NovaCreateServersOptionOSDCFdiskConfig) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaCreateServersOptionOSDCFdiskConfig) UnmarshalJSON(b []byte) error {
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

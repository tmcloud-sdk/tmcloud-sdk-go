package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaCreateServersResult struct {
	Id string `json:"id"`

	Links []NovaLink `json:"links"`

	SecurityGroups []NovaServerSecurityGroup `json:"security_groups"`

	OSDCFdiskConfig NovaCreateServersResultOSDCFdiskConfig `json:"OS-DCF:diskConfig"`

	ReservationId *string `json:"reservation_id,omitempty"`

	AdminPass string `json:"adminPass"`
}

func (o NovaCreateServersResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaCreateServersResult struct{}"
	}

	return strings.Join([]string{"NovaCreateServersResult", string(data)}, " ")
}

type NovaCreateServersResultOSDCFdiskConfig struct {
	value string
}

type NovaCreateServersResultOSDCFdiskConfigEnum struct {
	MANUAL NovaCreateServersResultOSDCFdiskConfig
	AUTO   NovaCreateServersResultOSDCFdiskConfig
}

func GetNovaCreateServersResultOSDCFdiskConfigEnum() NovaCreateServersResultOSDCFdiskConfigEnum {
	return NovaCreateServersResultOSDCFdiskConfigEnum{
		MANUAL: NovaCreateServersResultOSDCFdiskConfig{
			value: "MANUAL",
		},
		AUTO: NovaCreateServersResultOSDCFdiskConfig{
			value: "AUTO",
		},
	}
}

func (c NovaCreateServersResultOSDCFdiskConfig) Value() string {
	return c.value
}

func (c NovaCreateServersResultOSDCFdiskConfig) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaCreateServersResultOSDCFdiskConfig) UnmarshalJSON(b []byte) error {
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

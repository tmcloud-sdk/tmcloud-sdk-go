package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Ha struct {
	Mode HaMode `json:"mode"`

	ReplicationMode HaReplicationMode `json:"replication_mode"`
}

func (o Ha) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Ha struct{}"
	}

	return strings.Join([]string{"Ha", string(data)}, " ")
}

type HaMode struct {
	value string
}

type HaModeEnum struct {
	HA     HaMode
	SINGLE HaMode
}

func GetHaModeEnum() HaModeEnum {
	return HaModeEnum{
		HA: HaMode{
			value: "Ha",
		},
		SINGLE: HaMode{
			value: "Single",
		},
	}
}

func (c HaMode) Value() string {
	return c.value
}

func (c HaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HaMode) UnmarshalJSON(b []byte) error {
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

type HaReplicationMode struct {
	value string
}

type HaReplicationModeEnum struct {
	ASYNC    HaReplicationMode
	SEMISYNC HaReplicationMode
	SYNC     HaReplicationMode
}

func GetHaReplicationModeEnum() HaReplicationModeEnum {
	return HaReplicationModeEnum{
		ASYNC: HaReplicationMode{
			value: "async",
		},
		SEMISYNC: HaReplicationMode{
			value: "semisync",
		},
		SYNC: HaReplicationMode{
			value: "sync",
		},
	}
}

func (c HaReplicationMode) Value() string {
	return c.value
}

func (c HaReplicationMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HaReplicationMode) UnmarshalJSON(b []byte) error {
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

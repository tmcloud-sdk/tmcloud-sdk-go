package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ServerAddress struct {
	Version string `json:"version"`

	Addr string `json:"addr"`

	OSEXTIPStype *ServerAddressOSEXTIPStype `json:"OS-EXT-IPS:type,omitempty"`

	OSEXTIPSMACmacAddr *string `json:"OS-EXT-IPS-MAC:mac_addr,omitempty"`

	OSEXTIPSportId *string `json:"OS-EXT-IPS:port_id,omitempty"`
}

func (o ServerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerAddress struct{}"
	}

	return strings.Join([]string{"ServerAddress", string(data)}, " ")
}

type ServerAddressOSEXTIPStype struct {
	value string
}

type ServerAddressOSEXTIPStypeEnum struct {
	FIXED    ServerAddressOSEXTIPStype
	FLOATING ServerAddressOSEXTIPStype
}

func GetServerAddressOSEXTIPStypeEnum() ServerAddressOSEXTIPStypeEnum {
	return ServerAddressOSEXTIPStypeEnum{
		FIXED: ServerAddressOSEXTIPStype{
			value: "fixed",
		},
		FLOATING: ServerAddressOSEXTIPStype{
			value: "floating",
		},
	}
}

func (c ServerAddressOSEXTIPStype) Value() string {
	return c.value
}

func (c ServerAddressOSEXTIPStype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ServerAddressOSEXTIPStype) UnmarshalJSON(b []byte) error {
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

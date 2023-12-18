package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type EipInfo struct {
	IpType EipInfoIpType `json:"ip_type"`

	Bandwidth *BandwidthInfo `json:"bandwidth"`
}

func (o EipInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipInfo struct{}"
	}

	return strings.Join([]string{"EipInfo", string(data)}, " ")
}

type EipInfoIpType struct {
	value string
}

type EipInfoIpTypeEnum struct {
	E_5_BGP    EipInfoIpType
	E_5_SBGP   EipInfoIpType
	E_5_TELCOM EipInfoIpType
	E_5_UNION  EipInfoIpType
}

func GetEipInfoIpTypeEnum() EipInfoIpTypeEnum {
	return EipInfoIpTypeEnum{
		E_5_BGP: EipInfoIpType{
			value: "5_bgp",
		},
		E_5_SBGP: EipInfoIpType{
			value: "5_sbgp",
		},
		E_5_TELCOM: EipInfoIpType{
			value: "5_telcom",
		},
		E_5_UNION: EipInfoIpType{
			value: "5_union",
		},
	}
}

func (c EipInfoIpType) Value() string {
	return c.value
}

func (c EipInfoIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EipInfoIpType) UnmarshalJSON(b []byte) error {
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

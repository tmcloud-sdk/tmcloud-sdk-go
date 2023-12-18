package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type InstanceConfig struct {
	InstanceId *string `json:"instance_id,omitempty"`

	FlavorRef *string `json:"flavorRef,omitempty"`

	ImageRef *string `json:"imageRef,omitempty"`

	Disk *[]DiskInfo `json:"disk,omitempty"`

	KeyName *string `json:"key_name,omitempty"`

	Personality *[]PersonalityInfo `json:"personality,omitempty"`

	PublicIp *PublicIp `json:"public_ip,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Metadata *VmMetaData `json:"metadata,omitempty"`

	SecurityGroups *[]SecurityGroups `json:"security_groups,omitempty"`

	ServerGroupId *string `json:"server_group_id,omitempty"`

	Tenancy *InstanceConfigTenancy `json:"tenancy,omitempty"`

	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	MultiFlavorPriorityPolicy *InstanceConfigMultiFlavorPriorityPolicy `json:"multi_flavor_priority_policy,omitempty"`

	MarketType *InstanceConfigMarketType `json:"market_type,omitempty"`
}

func (o InstanceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceConfig struct{}"
	}

	return strings.Join([]string{"InstanceConfig", string(data)}, " ")
}

type InstanceConfigTenancy struct {
	value string
}

type InstanceConfigTenancyEnum struct {
	DEDICATED InstanceConfigTenancy
}

func GetInstanceConfigTenancyEnum() InstanceConfigTenancyEnum {
	return InstanceConfigTenancyEnum{
		DEDICATED: InstanceConfigTenancy{
			value: "dedicated",
		},
	}
}

func (c InstanceConfigTenancy) Value() string {
	return c.value
}

func (c InstanceConfigTenancy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceConfigTenancy) UnmarshalJSON(b []byte) error {
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

type InstanceConfigMultiFlavorPriorityPolicy struct {
	value string
}

type InstanceConfigMultiFlavorPriorityPolicyEnum struct {
	PICK_FIRST InstanceConfigMultiFlavorPriorityPolicy
	COST_FIRST InstanceConfigMultiFlavorPriorityPolicy
}

func GetInstanceConfigMultiFlavorPriorityPolicyEnum() InstanceConfigMultiFlavorPriorityPolicyEnum {
	return InstanceConfigMultiFlavorPriorityPolicyEnum{
		PICK_FIRST: InstanceConfigMultiFlavorPriorityPolicy{
			value: "PICK_FIRST",
		},
		COST_FIRST: InstanceConfigMultiFlavorPriorityPolicy{
			value: "COST_FIRST",
		},
	}
}

func (c InstanceConfigMultiFlavorPriorityPolicy) Value() string {
	return c.value
}

func (c InstanceConfigMultiFlavorPriorityPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceConfigMultiFlavorPriorityPolicy) UnmarshalJSON(b []byte) error {
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

type InstanceConfigMarketType struct {
	value string
}

type InstanceConfigMarketTypeEnum struct {
	SPOT InstanceConfigMarketType
}

func GetInstanceConfigMarketTypeEnum() InstanceConfigMarketTypeEnum {
	return InstanceConfigMarketTypeEnum{
		SPOT: InstanceConfigMarketType{
			value: "spot",
		},
	}
}

func (c InstanceConfigMarketType) Value() string {
	return c.value
}

func (c InstanceConfigMarketType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceConfigMarketType) UnmarshalJSON(b []byte) error {
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

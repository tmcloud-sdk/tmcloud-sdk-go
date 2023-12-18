package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type CreateScalingGroupOption struct {
	ScalingGroupName string `json:"scaling_group_name"`

	ScalingConfigurationId string `json:"scaling_configuration_id"`

	DesireInstanceNumber *int32 `json:"desire_instance_number,omitempty"`

	MinInstanceNumber *int32 `json:"min_instance_number,omitempty"`

	MaxInstanceNumber *int32 `json:"max_instance_number,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	LbListenerId *string `json:"lb_listener_id,omitempty"`

	LbaasListeners *[]LbaasListeners `json:"lbaas_listeners,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	Networks []Networks `json:"networks"`

	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	VpcId string `json:"vpc_id"`

	HealthPeriodicAuditMethod *CreateScalingGroupOptionHealthPeriodicAuditMethod `json:"health_periodic_audit_method,omitempty"`

	HealthPeriodicAuditTime *int32 `json:"health_periodic_audit_time,omitempty"`

	HealthPeriodicAuditGracePeriod *int32 `json:"health_periodic_audit_grace_period,omitempty"`

	InstanceTerminatePolicy *CreateScalingGroupOptionInstanceTerminatePolicy `json:"instance_terminate_policy,omitempty"`

	Notifications *[]string `json:"notifications,omitempty"`

	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	DeleteVolume *bool `json:"delete_volume,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	MultiAzPriorityPolicy *CreateScalingGroupOptionMultiAzPriorityPolicy `json:"multi_az_priority_policy,omitempty"`

	IamAgencyName *string `json:"iam_agency_name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o CreateScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingGroupOption struct{}"
	}

	return strings.Join([]string{"CreateScalingGroupOption", string(data)}, " ")
}

type CreateScalingGroupOptionHealthPeriodicAuditMethod struct {
	value string
}

type CreateScalingGroupOptionHealthPeriodicAuditMethodEnum struct {
	ELB_AUDIT  CreateScalingGroupOptionHealthPeriodicAuditMethod
	NOVA_AUDIT CreateScalingGroupOptionHealthPeriodicAuditMethod
}

func GetCreateScalingGroupOptionHealthPeriodicAuditMethodEnum() CreateScalingGroupOptionHealthPeriodicAuditMethodEnum {
	return CreateScalingGroupOptionHealthPeriodicAuditMethodEnum{
		ELB_AUDIT: CreateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "ELB_AUDIT",
		},
		NOVA_AUDIT: CreateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "NOVA_AUDIT",
		},
	}
}

func (c CreateScalingGroupOptionHealthPeriodicAuditMethod) Value() string {
	return c.value
}

func (c CreateScalingGroupOptionHealthPeriodicAuditMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionHealthPeriodicAuditMethod) UnmarshalJSON(b []byte) error {
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

type CreateScalingGroupOptionInstanceTerminatePolicy struct {
	value string
}

type CreateScalingGroupOptionInstanceTerminatePolicyEnum struct {
	OLD_CONFIG_OLD_INSTANCE CreateScalingGroupOptionInstanceTerminatePolicy
	OLD_CONFIG_NEW_INSTANCE CreateScalingGroupOptionInstanceTerminatePolicy
	OLD_INSTANCE            CreateScalingGroupOptionInstanceTerminatePolicy
	NEW_INSTANCE            CreateScalingGroupOptionInstanceTerminatePolicy
}

func GetCreateScalingGroupOptionInstanceTerminatePolicyEnum() CreateScalingGroupOptionInstanceTerminatePolicyEnum {
	return CreateScalingGroupOptionInstanceTerminatePolicyEnum{
		OLD_CONFIG_OLD_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_OLD_INSTANCE",
		},
		OLD_CONFIG_NEW_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_NEW_INSTANCE",
		},
		OLD_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_INSTANCE",
		},
		NEW_INSTANCE: CreateScalingGroupOptionInstanceTerminatePolicy{
			value: "NEW_INSTANCE",
		},
	}
}

func (c CreateScalingGroupOptionInstanceTerminatePolicy) Value() string {
	return c.value
}

func (c CreateScalingGroupOptionInstanceTerminatePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionInstanceTerminatePolicy) UnmarshalJSON(b []byte) error {
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

type CreateScalingGroupOptionMultiAzPriorityPolicy struct {
	value string
}

type CreateScalingGroupOptionMultiAzPriorityPolicyEnum struct {
	EQUILIBRIUM_DISTRIBUTE CreateScalingGroupOptionMultiAzPriorityPolicy
	PICK_FIRST             CreateScalingGroupOptionMultiAzPriorityPolicy
}

func GetCreateScalingGroupOptionMultiAzPriorityPolicyEnum() CreateScalingGroupOptionMultiAzPriorityPolicyEnum {
	return CreateScalingGroupOptionMultiAzPriorityPolicyEnum{
		EQUILIBRIUM_DISTRIBUTE: CreateScalingGroupOptionMultiAzPriorityPolicy{
			value: "EQUILIBRIUM_DISTRIBUTE",
		},
		PICK_FIRST: CreateScalingGroupOptionMultiAzPriorityPolicy{
			value: "PICK_FIRST",
		},
	}
}

func (c CreateScalingGroupOptionMultiAzPriorityPolicy) Value() string {
	return c.value
}

func (c CreateScalingGroupOptionMultiAzPriorityPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateScalingGroupOptionMultiAzPriorityPolicy) UnmarshalJSON(b []byte) error {
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

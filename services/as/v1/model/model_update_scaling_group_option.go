package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateScalingGroupOption struct {
	ScalingGroupName *string `json:"scaling_group_name,omitempty"`

	DesireInstanceNumber *int32 `json:"desire_instance_number,omitempty"`

	MinInstanceNumber *int32 `json:"min_instance_number,omitempty"`

	MaxInstanceNumber *int32 `json:"max_instance_number,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	Networks *[]Networks `json:"networks,omitempty"`

	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	LbListenerId *string `json:"lb_listener_id,omitempty"`

	LbaasListeners *[]LbaasListeners `json:"lbaas_listeners,omitempty"`

	HealthPeriodicAuditMethod *UpdateScalingGroupOptionHealthPeriodicAuditMethod `json:"health_periodic_audit_method,omitempty"`

	HealthPeriodicAuditTime *int32 `json:"health_periodic_audit_time,omitempty"`

	HealthPeriodicAuditGracePeriod *int32 `json:"health_periodic_audit_grace_period,omitempty"`

	InstanceTerminatePolicy *UpdateScalingGroupOptionInstanceTerminatePolicy `json:"instance_terminate_policy,omitempty"`

	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`

	Notifications *[]string `json:"notifications,omitempty"`

	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	DeleteVolume *bool `json:"delete_volume,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	MultiAzPriorityPolicy *string `json:"multi_az_priority_policy,omitempty"`

	IamAgencyName *string `json:"iam_agency_name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o UpdateScalingGroupOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScalingGroupOption struct{}"
	}

	return strings.Join([]string{"UpdateScalingGroupOption", string(data)}, " ")
}

type UpdateScalingGroupOptionHealthPeriodicAuditMethod struct {
	value string
}

type UpdateScalingGroupOptionHealthPeriodicAuditMethodEnum struct {
	ELB_AUDIT  UpdateScalingGroupOptionHealthPeriodicAuditMethod
	NOVA_AUDIT UpdateScalingGroupOptionHealthPeriodicAuditMethod
}

func GetUpdateScalingGroupOptionHealthPeriodicAuditMethodEnum() UpdateScalingGroupOptionHealthPeriodicAuditMethodEnum {
	return UpdateScalingGroupOptionHealthPeriodicAuditMethodEnum{
		ELB_AUDIT: UpdateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "ELB_AUDIT",
		},
		NOVA_AUDIT: UpdateScalingGroupOptionHealthPeriodicAuditMethod{
			value: "NOVA_AUDIT",
		},
	}
}

func (c UpdateScalingGroupOptionHealthPeriodicAuditMethod) Value() string {
	return c.value
}

func (c UpdateScalingGroupOptionHealthPeriodicAuditMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingGroupOptionHealthPeriodicAuditMethod) UnmarshalJSON(b []byte) error {
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

type UpdateScalingGroupOptionInstanceTerminatePolicy struct {
	value string
}

type UpdateScalingGroupOptionInstanceTerminatePolicyEnum struct {
	OLD_CONFIG_OLD_INSTANCE UpdateScalingGroupOptionInstanceTerminatePolicy
	OLD_CONFIG_NEW_INSTANCE UpdateScalingGroupOptionInstanceTerminatePolicy
	OLD_INSTANCE            UpdateScalingGroupOptionInstanceTerminatePolicy
	NEW_INSTANCE            UpdateScalingGroupOptionInstanceTerminatePolicy
}

func GetUpdateScalingGroupOptionInstanceTerminatePolicyEnum() UpdateScalingGroupOptionInstanceTerminatePolicyEnum {
	return UpdateScalingGroupOptionInstanceTerminatePolicyEnum{
		OLD_CONFIG_OLD_INSTANCE: UpdateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_OLD_INSTANCE",
		},
		OLD_CONFIG_NEW_INSTANCE: UpdateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_CONFIG_NEW_INSTANCE",
		},
		OLD_INSTANCE: UpdateScalingGroupOptionInstanceTerminatePolicy{
			value: "OLD_INSTANCE",
		},
		NEW_INSTANCE: UpdateScalingGroupOptionInstanceTerminatePolicy{
			value: "NEW_INSTANCE",
		},
	}
}

func (c UpdateScalingGroupOptionInstanceTerminatePolicy) Value() string {
	return c.value
}

func (c UpdateScalingGroupOptionInstanceTerminatePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateScalingGroupOptionInstanceTerminatePolicy) UnmarshalJSON(b []byte) error {
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

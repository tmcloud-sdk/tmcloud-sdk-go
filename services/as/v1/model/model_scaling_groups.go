package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type ScalingGroups struct {
	ScalingGroupName *string `json:"scaling_group_name,omitempty"`

	ScalingGroupId *string `json:"scaling_group_id,omitempty"`

	ScalingGroupStatus *ScalingGroupsScalingGroupStatus `json:"scaling_group_status,omitempty"`

	ScalingConfigurationId *string `json:"scaling_configuration_id,omitempty"`

	ScalingConfigurationName *string `json:"scaling_configuration_name,omitempty"`

	CurrentInstanceNumber *int32 `json:"current_instance_number,omitempty"`

	DesireInstanceNumber *int32 `json:"desire_instance_number,omitempty"`

	MinInstanceNumber *int32 `json:"min_instance_number,omitempty"`

	MaxInstanceNumber *int32 `json:"max_instance_number,omitempty"`

	CoolDownTime *int32 `json:"cool_down_time,omitempty"`

	LbListenerId *string `json:"lb_listener_id,omitempty"`

	LbaasListeners *[]LbaasListenersResult `json:"lbaas_listeners,omitempty"`

	AvailableZones *[]string `json:"available_zones,omitempty"`

	Networks *[]NetworksResult `json:"networks,omitempty"`

	SecurityGroups *[]SecurityGroupsResult `json:"security_groups,omitempty"`

	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	Detail *string `json:"detail,omitempty"`

	IsScaling *bool `json:"is_scaling,omitempty"`

	HealthPeriodicAuditMethod *ScalingGroupsHealthPeriodicAuditMethod `json:"health_periodic_audit_method,omitempty"`

	HealthPeriodicAuditTime *int32 `json:"health_periodic_audit_time,omitempty"`

	HealthPeriodicAuditGracePeriod *int32 `json:"health_periodic_audit_grace_period,omitempty"`

	InstanceTerminatePolicy *ScalingGroupsInstanceTerminatePolicy `json:"instance_terminate_policy,omitempty"`

	Notifications *[]string `json:"notifications,omitempty"`

	DeletePublicip *bool `json:"delete_publicip,omitempty"`

	DeleteVolume *bool `json:"delete_volume,omitempty"`

	CloudLocationId *string `json:"cloud_location_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ActivityType *string `json:"activity_type,omitempty"`

	MultiAzPriorityPolicy *string `json:"multi_az_priority_policy,omitempty"`

	IamAgencyName *string `json:"iam_agency_name,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o ScalingGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingGroups struct{}"
	}

	return strings.Join([]string{"ScalingGroups", string(data)}, " ")
}

type ScalingGroupsScalingGroupStatus struct {
	value string
}

type ScalingGroupsScalingGroupStatusEnum struct {
	INSERVICE ScalingGroupsScalingGroupStatus
	PAUSED    ScalingGroupsScalingGroupStatus
	ERROR     ScalingGroupsScalingGroupStatus
	DELETING  ScalingGroupsScalingGroupStatus
	FREEZED   ScalingGroupsScalingGroupStatus
}

func GetScalingGroupsScalingGroupStatusEnum() ScalingGroupsScalingGroupStatusEnum {
	return ScalingGroupsScalingGroupStatusEnum{
		INSERVICE: ScalingGroupsScalingGroupStatus{
			value: "INSERVICE",
		},
		PAUSED: ScalingGroupsScalingGroupStatus{
			value: "PAUSED",
		},
		ERROR: ScalingGroupsScalingGroupStatus{
			value: "ERROR",
		},
		DELETING: ScalingGroupsScalingGroupStatus{
			value: "DELETING",
		},
		FREEZED: ScalingGroupsScalingGroupStatus{
			value: "FREEZED",
		},
	}
}

func (c ScalingGroupsScalingGroupStatus) Value() string {
	return c.value
}

func (c ScalingGroupsScalingGroupStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingGroupsScalingGroupStatus) UnmarshalJSON(b []byte) error {
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

type ScalingGroupsHealthPeriodicAuditMethod struct {
	value string
}

type ScalingGroupsHealthPeriodicAuditMethodEnum struct {
	ELB_AUDIT  ScalingGroupsHealthPeriodicAuditMethod
	NOVA_AUDIT ScalingGroupsHealthPeriodicAuditMethod
}

func GetScalingGroupsHealthPeriodicAuditMethodEnum() ScalingGroupsHealthPeriodicAuditMethodEnum {
	return ScalingGroupsHealthPeriodicAuditMethodEnum{
		ELB_AUDIT: ScalingGroupsHealthPeriodicAuditMethod{
			value: "ELB_AUDIT",
		},
		NOVA_AUDIT: ScalingGroupsHealthPeriodicAuditMethod{
			value: "NOVA_AUDIT",
		},
	}
}

func (c ScalingGroupsHealthPeriodicAuditMethod) Value() string {
	return c.value
}

func (c ScalingGroupsHealthPeriodicAuditMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingGroupsHealthPeriodicAuditMethod) UnmarshalJSON(b []byte) error {
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

type ScalingGroupsInstanceTerminatePolicy struct {
	value string
}

type ScalingGroupsInstanceTerminatePolicyEnum struct {
	OLD_CONFIG_OLD_INSTANCE ScalingGroupsInstanceTerminatePolicy
	OLD_CONFIG_NEW_INSTANCE ScalingGroupsInstanceTerminatePolicy
	OLD_INSTANCE            ScalingGroupsInstanceTerminatePolicy
	NEW_INSTANCE            ScalingGroupsInstanceTerminatePolicy
}

func GetScalingGroupsInstanceTerminatePolicyEnum() ScalingGroupsInstanceTerminatePolicyEnum {
	return ScalingGroupsInstanceTerminatePolicyEnum{
		OLD_CONFIG_OLD_INSTANCE: ScalingGroupsInstanceTerminatePolicy{
			value: "OLD_CONFIG_OLD_INSTANCE",
		},
		OLD_CONFIG_NEW_INSTANCE: ScalingGroupsInstanceTerminatePolicy{
			value: "OLD_CONFIG_NEW_INSTANCE",
		},
		OLD_INSTANCE: ScalingGroupsInstanceTerminatePolicy{
			value: "OLD_INSTANCE",
		},
		NEW_INSTANCE: ScalingGroupsInstanceTerminatePolicy{
			value: "NEW_INSTANCE",
		},
	}
}

func (c ScalingGroupsInstanceTerminatePolicy) Value() string {
	return c.value
}

func (c ScalingGroupsInstanceTerminatePolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScalingGroupsInstanceTerminatePolicy) UnmarshalJSON(b []byte) error {
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

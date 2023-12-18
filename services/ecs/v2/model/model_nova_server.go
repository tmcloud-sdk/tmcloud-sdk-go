package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaServer struct {
	Name string `json:"name"`

	Id string `json:"id"`

	Status NovaServerStatus `json:"status"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	Flavor *NovaServerFlavor `json:"flavor"`

	Image *NovaServerImage `json:"image"`

	TenantId string `json:"tenant_id"`

	KeyName string `json:"key_name"`

	UserId string `json:"user_id"`

	Metadata map[string]string `json:"metadata"`

	HostId string `json:"hostId"`

	Addresses map[string][]NovaNetwork `json:"addresses"`

	SecurityGroups []NovaServerSecurityGroup `json:"security_groups"`

	Links []NovaLink `json:"links"`

	OSDCFdiskConfig NovaServerOSDCFdiskConfig `json:"OS-DCF:diskConfig"`

	OSEXTAZavailabilityZone string `json:"OS-EXT-AZ:availability_zone"`

	OSEXTSRVATTRhost string `json:"OS-EXT-SRV-ATTR:host"`

	OSEXTSRVATTRhypervisorHostname string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`

	OSEXTSRVATTRinstanceName string `json:"OS-EXT-SRV-ATTR:instance_name"`

	OSEXTSTSpowerState int32 `json:"OS-EXT-STS:power_state"`

	OSEXTSTStaskState NovaServerOSEXTSTStaskState `json:"OS-EXT-STS:task_state"`

	OSEXTSTSvmState NovaServerOSEXTSTSvmState `json:"OS-EXT-STS:vm_state"`

	OSSRVUSGlaunchedAt string `json:"OS-SRV-USG:launched_at"`

	OSSRVUSGterminatedAt string `json:"OS-SRV-USG:terminated_at"`

	OsExtendedVolumesvolumesAttached []NovaServerVolume `json:"os-extended-volumes:volumes_attached"`

	Fault *NovaServerFault `json:"fault,omitempty"`

	Description *string `json:"description,omitempty"`

	HostStatus NovaServerHostStatus `json:"host_status"`

	OSEXTSRVATTRhostname *string `json:"OS-EXT-SRV-ATTR:hostname,omitempty"`

	OSEXTSRVATTRreservationId *string `json:"OS-EXT-SRV-ATTR:reservation_id,omitempty"`

	OSEXTSRVATTRlaunchIndex *int32 `json:"OS-EXT-SRV-ATTR:launch_index,omitempty"`

	OSEXTSRVATTRkernelId *string `json:"OS-EXT-SRV-ATTR:kernel_id,omitempty"`

	OSEXTSRVATTRramdiskId *string `json:"OS-EXT-SRV-ATTR:ramdisk_id,omitempty"`

	OSEXTSRVATTRrootDeviceName *string `json:"OS-EXT-SRV-ATTR:root_device_name,omitempty"`

	OSEXTSRVATTRuserData *string `json:"OS-EXT-SRV-ATTR:user_data,omitempty"`

	Tags []string `json:"tags"`

	Locked *bool `json:"locked,omitempty"`

	AccessIPv4 string `json:"accessIPv4"`

	AccessIPv6 string `json:"accessIPv6"`

	ConfigDrive string `json:"config_drive"`

	Progress int32 `json:"progress"`

	OsschedulerHints *NovaServerSchedulerHints `json:"os:scheduler_hints,omitempty"`
}

func (o NovaServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaServer struct{}"
	}

	return strings.Join([]string{"NovaServer", string(data)}, " ")
}

type NovaServerStatus struct {
	value string
}

type NovaServerStatusEnum struct {
	ACTIVE            NovaServerStatus
	BUILD             NovaServerStatus
	DELETED           NovaServerStatus
	ERROR             NovaServerStatus
	HARD_REBOOT       NovaServerStatus
	MIGRATING         NovaServerStatus
	REBOOT            NovaServerStatus
	RESIZE            NovaServerStatus
	REVERT_RESIZE     NovaServerStatus
	SHELVED           NovaServerStatus
	SHELVED_OFFLOADED NovaServerStatus
	SHUTOFF           NovaServerStatus
	UNKNOWN           NovaServerStatus
	VERIFY_RESIZE     NovaServerStatus
}

func GetNovaServerStatusEnum() NovaServerStatusEnum {
	return NovaServerStatusEnum{
		ACTIVE: NovaServerStatus{
			value: "ACTIVE",
		},
		BUILD: NovaServerStatus{
			value: " BUILD",
		},
		DELETED: NovaServerStatus{
			value: "DELETED",
		},
		ERROR: NovaServerStatus{
			value: "ERROR",
		},
		HARD_REBOOT: NovaServerStatus{
			value: "HARD_REBOOT",
		},
		MIGRATING: NovaServerStatus{
			value: "MIGRATING",
		},
		REBOOT: NovaServerStatus{
			value: "REBOOT",
		},
		RESIZE: NovaServerStatus{
			value: "RESIZE",
		},
		REVERT_RESIZE: NovaServerStatus{
			value: "REVERT_RESIZE",
		},
		SHELVED: NovaServerStatus{
			value: "SHELVED",
		},
		SHELVED_OFFLOADED: NovaServerStatus{
			value: "SHELVED_OFFLOADED",
		},
		SHUTOFF: NovaServerStatus{
			value: "SHUTOFF",
		},
		UNKNOWN: NovaServerStatus{
			value: "UNKNOWN",
		},
		VERIFY_RESIZE: NovaServerStatus{
			value: "VERIFY_RESIZE",
		},
	}
}

func (c NovaServerStatus) Value() string {
	return c.value
}

func (c NovaServerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerStatus) UnmarshalJSON(b []byte) error {
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

type NovaServerOSDCFdiskConfig struct {
	value string
}

type NovaServerOSDCFdiskConfigEnum struct {
	AUTO   NovaServerOSDCFdiskConfig
	MANUAL NovaServerOSDCFdiskConfig
}

func GetNovaServerOSDCFdiskConfigEnum() NovaServerOSDCFdiskConfigEnum {
	return NovaServerOSDCFdiskConfigEnum{
		AUTO: NovaServerOSDCFdiskConfig{
			value: "AUTO",
		},
		MANUAL: NovaServerOSDCFdiskConfig{
			value: "MANUAL",
		},
	}
}

func (c NovaServerOSDCFdiskConfig) Value() string {
	return c.value
}

func (c NovaServerOSDCFdiskConfig) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerOSDCFdiskConfig) UnmarshalJSON(b []byte) error {
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

type NovaServerOSEXTSTStaskState struct {
	value string
}

type NovaServerOSEXTSTStaskStateEnum struct {
	SHOUTOFF          NovaServerOSEXTSTStaskState
	RESIZE            NovaServerOSEXTSTStaskState
	REBUILD           NovaServerOSEXTSTStaskState
	VERIFY_RESIZE     NovaServerOSEXTSTStaskState
	REVERT_RESIZE     NovaServerOSEXTSTStaskState
	PAUSED            NovaServerOSEXTSTStaskState
	MIGRATING         NovaServerOSEXTSTStaskState
	SUSPENDED         NovaServerOSEXTSTStaskState
	RESCUE            NovaServerOSEXTSTStaskState
	ERROR             NovaServerOSEXTSTStaskState
	DELETED           NovaServerOSEXTSTStaskState
	SOFT_DELETED      NovaServerOSEXTSTStaskState
	SHELVED           NovaServerOSEXTSTStaskState
	SHELVED_OFFLOADED NovaServerOSEXTSTStaskState
}

func GetNovaServerOSEXTSTStaskStateEnum() NovaServerOSEXTSTStaskStateEnum {
	return NovaServerOSEXTSTStaskStateEnum{
		SHOUTOFF: NovaServerOSEXTSTStaskState{
			value: "SHOUTOFF",
		},
		RESIZE: NovaServerOSEXTSTStaskState{
			value: " RESIZE",
		},
		REBUILD: NovaServerOSEXTSTStaskState{
			value: " REBUILD",
		},
		VERIFY_RESIZE: NovaServerOSEXTSTStaskState{
			value: " VERIFY_RESIZE",
		},
		REVERT_RESIZE: NovaServerOSEXTSTStaskState{
			value: " REVERT_RESIZE",
		},
		PAUSED: NovaServerOSEXTSTStaskState{
			value: " PAUSED",
		},
		MIGRATING: NovaServerOSEXTSTStaskState{
			value: " MIGRATING",
		},
		SUSPENDED: NovaServerOSEXTSTStaskState{
			value: " SUSPENDED",
		},
		RESCUE: NovaServerOSEXTSTStaskState{
			value: " RESCUE",
		},
		ERROR: NovaServerOSEXTSTStaskState{
			value: " ERROR",
		},
		DELETED: NovaServerOSEXTSTStaskState{
			value: " DELETED",
		},
		SOFT_DELETED: NovaServerOSEXTSTStaskState{
			value: "SOFT_DELETED",
		},
		SHELVED: NovaServerOSEXTSTStaskState{
			value: "SHELVED",
		},
		SHELVED_OFFLOADED: NovaServerOSEXTSTStaskState{
			value: "SHELVED_OFFLOADED",
		},
	}
}

func (c NovaServerOSEXTSTStaskState) Value() string {
	return c.value
}

func (c NovaServerOSEXTSTStaskState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerOSEXTSTStaskState) UnmarshalJSON(b []byte) error {
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

type NovaServerOSEXTSTSvmState struct {
	value string
}

type NovaServerOSEXTSTSvmStateEnum struct {
	ACTIVE            NovaServerOSEXTSTSvmState
	BUILDING          NovaServerOSEXTSTSvmState
	STOPPED           NovaServerOSEXTSTSvmState
	RESIZED           NovaServerOSEXTSTSvmState
	PAUSED            NovaServerOSEXTSTSvmState
	SUSPENDED         NovaServerOSEXTSTSvmState
	RESCUED           NovaServerOSEXTSTSvmState
	ERROR             NovaServerOSEXTSTSvmState
	DELETED           NovaServerOSEXTSTSvmState
	SOFT_DELETED      NovaServerOSEXTSTSvmState
	SHELVED           NovaServerOSEXTSTSvmState
	SHELVED_OFFLOADED NovaServerOSEXTSTSvmState
}

func GetNovaServerOSEXTSTSvmStateEnum() NovaServerOSEXTSTSvmStateEnum {
	return NovaServerOSEXTSTSvmStateEnum{
		ACTIVE: NovaServerOSEXTSTSvmState{
			value: "ACTIVE",
		},
		BUILDING: NovaServerOSEXTSTSvmState{
			value: "BUILDING",
		},
		STOPPED: NovaServerOSEXTSTSvmState{
			value: "STOPPED",
		},
		RESIZED: NovaServerOSEXTSTSvmState{
			value: "RESIZED",
		},
		PAUSED: NovaServerOSEXTSTSvmState{
			value: "PAUSED",
		},
		SUSPENDED: NovaServerOSEXTSTSvmState{
			value: "SUSPENDED",
		},
		RESCUED: NovaServerOSEXTSTSvmState{
			value: "RESCUED",
		},
		ERROR: NovaServerOSEXTSTSvmState{
			value: "ERROR",
		},
		DELETED: NovaServerOSEXTSTSvmState{
			value: "DELETED",
		},
		SOFT_DELETED: NovaServerOSEXTSTSvmState{
			value: "SOFT_DELETED",
		},
		SHELVED: NovaServerOSEXTSTSvmState{
			value: "SHELVED",
		},
		SHELVED_OFFLOADED: NovaServerOSEXTSTSvmState{
			value: "SHELVED_OFFLOADED",
		},
	}
}

func (c NovaServerOSEXTSTSvmState) Value() string {
	return c.value
}

func (c NovaServerOSEXTSTSvmState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerOSEXTSTSvmState) UnmarshalJSON(b []byte) error {
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

type NovaServerHostStatus struct {
	value string
}

type NovaServerHostStatusEnum struct {
	UP          NovaServerHostStatus
	UNKNOWN     NovaServerHostStatus
	DOWN        NovaServerHostStatus
	MAINTENANCE NovaServerHostStatus
}

func GetNovaServerHostStatusEnum() NovaServerHostStatusEnum {
	return NovaServerHostStatusEnum{
		UP: NovaServerHostStatus{
			value: "UP",
		},
		UNKNOWN: NovaServerHostStatus{
			value: "UNKNOWN",
		},
		DOWN: NovaServerHostStatus{
			value: "DOWN",
		},
		MAINTENANCE: NovaServerHostStatus{
			value: "MAINTENANCE",
		},
	}
}

func (c NovaServerHostStatus) Value() string {
	return c.value
}

func (c NovaServerHostStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaServerHostStatus) UnmarshalJSON(b []byte) error {
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

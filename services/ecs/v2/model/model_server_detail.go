package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ServerDetail struct {
	Status string `json:"status"`

	Updated string `json:"updated"`

	AutoTerminateTime string `json:"auto_terminate_time"`

	HostId string `json:"hostId"`

	OSEXTSRVATTRhost string `json:"OS-EXT-SRV-ATTR:host"`

	Addresses map[string][]ServerAddress `json:"addresses"`

	KeyName string `json:"key_name"`

	Image *ServerImage `json:"image"`

	OSEXTSTStaskState string `json:"OS-EXT-STS:task_state"`

	OSEXTSTSvmState string `json:"OS-EXT-STS:vm_state"`

	OSEXTSRVATTRinstanceName string `json:"OS-EXT-SRV-ATTR:instance_name"`

	OSEXTSRVATTRhypervisorHostname string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname"`

	Flavor *ServerFlavor `json:"flavor"`

	Id string `json:"id"`

	SecurityGroups []ServerSecurityGroup `json:"security_groups"`

	OSEXTAZavailabilityZone string `json:"OS-EXT-AZ:availability_zone"`

	UserId string `json:"user_id"`

	Name string `json:"name"`

	Created string `json:"created"`

	TenantId string `json:"tenant_id"`

	OSDCFdiskConfig *string `json:"OS-DCF:diskConfig,omitempty"`

	AccessIPv4 string `json:"accessIPv4"`

	AccessIPv6 string `json:"accessIPv6"`

	Fault *ServerFault `json:"fault,omitempty"`

	Progress *int32 `json:"progress,omitempty"`

	OSEXTSTSpowerState int32 `json:"OS-EXT-STS:power_state"`

	ConfigDrive string `json:"config_drive"`

	Metadata map[string]string `json:"metadata"`

	OSSRVUSGlaunchedAt string `json:"OS-SRV-USG:launched_at"`

	OSSRVUSGterminatedAt string `json:"OS-SRV-USG:terminated_at"`

	OsExtendedVolumesvolumesAttached []ServerExtendVolumeAttachment `json:"os-extended-volumes:volumes_attached"`

	Description *string `json:"description,omitempty"`

	HostStatus string `json:"host_status"`

	OSEXTSRVATTRhostname string `json:"OS-EXT-SRV-ATTR:hostname"`

	OSEXTSRVATTRreservationId *string `json:"OS-EXT-SRV-ATTR:reservation_id,omitempty"`

	OSEXTSRVATTRlaunchIndex int32 `json:"OS-EXT-SRV-ATTR:launch_index"`

	OSEXTSRVATTRkernelId string `json:"OS-EXT-SRV-ATTR:kernel_id"`

	OSEXTSRVATTRramdiskId string `json:"OS-EXT-SRV-ATTR:ramdisk_id"`

	OSEXTSRVATTRrootDeviceName string `json:"OS-EXT-SRV-ATTR:root_device_name"`

	OSEXTSRVATTRuserData *string `json:"OS-EXT-SRV-ATTR:user_data,omitempty"`

	Locked bool `json:"locked"`

	Tags *[]string `json:"tags,omitempty"`

	OsschedulerHints *ServerSchedulerHints `json:"os:scheduler_hints,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	SysTags *[]ServerSystemTag `json:"sys_tags,omitempty"`

	CpuOptions *CpuOptions `json:"cpu_options,omitempty"`

	Hypervisor *Hypervisor `json:"hypervisor,omitempty"`
}

func (o ServerDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerDetail struct{}"
	}

	return strings.Join([]string{"ServerDetail", string(data)}, " ")
}

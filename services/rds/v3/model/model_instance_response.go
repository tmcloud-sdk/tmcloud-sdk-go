package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceResponse struct {
	Id string `json:"id"`

	Status string `json:"status"`

	EnableSsl bool `json:"enable_ssl"`

	PrivateIps []string `json:"private_ips"`

	PrivateDnsNames *[]string `json:"private_dns_names,omitempty"`

	PublicIps []string `json:"public_ips"`

	Type string `json:"type"`

	Created string `json:"created"`

	Updated string `json:"updated"`

	DbUserName string `json:"db_user_name"`

	SwitchStrategy string `json:"switch_strategy"`

	ReadOnlyByUser *bool `json:"read_only_by_user,omitempty"`

	MaintenanceWindow string `json:"maintenance_window"`

	Nodes []NodeResponse `json:"nodes"`

	RelatedInstance []RelatedInstance `json:"related_instance"`

	Name string `json:"name"`

	Datastore *Datastore `json:"datastore"`

	Ha *HaResponse `json:"ha,omitempty"`

	Port int32 `json:"port"`

	BackupStrategy *BackupStrategyForResponse `json:"backup_strategy"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	DiskEncryptionId string `json:"disk_encryption_id"`

	FlavorRef string `json:"flavor_ref"`

	Cpu *string `json:"cpu,omitempty"`

	Mem *string `json:"mem,omitempty"`

	Volume *Volume `json:"volume"`

	Region string `json:"region"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	ChargeInfo *ChargeInfoResponse `json:"charge_info"`

	TimeZone string `json:"time_zone"`

	Tags []TagResponse `json:"tags"`

	BackupUsedSpace *float64 `json:"backup_used_space,omitempty"`

	StorageUsedSpace *float64 `json:"storage_used_space,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	AssociatedWithDdm *bool `json:"associated_with_ddm,omitempty"`

	Alias *string `json:"alias,omitempty"`

	MaxIops *int64 `json:"max_iops,omitempty"`

	ExpirationTime *string `json:"expiration_time,omitempty"`
}

func (o InstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceResponse struct{}"
	}

	return strings.Join([]string{"InstanceResponse", string(data)}, " ")
}

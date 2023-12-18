package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRestoreInstanceRequestBody struct {
	Name string `json:"name"`

	Datastore *Datastore `json:"datastore"`

	Ha *Ha `json:"ha,omitempty"`

	ConfigurationId *string `json:"configuration_id,omitempty"`

	Port *string `json:"port,omitempty"`

	Password *string `json:"password,omitempty"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	FlavorRef string `json:"flavor_ref"`

	Volume *Volume `json:"volume"`

	Region string `json:"region"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	DataVip *string `json:"data_vip,omitempty"`

	SecurityGroupId string `json:"security_group_id"`

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`

	DsspoolId *string `json:"dsspool_id,omitempty"`

	ReplicaOfId *string `json:"replica_of_id,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point,omitempty"`

	Collation *string `json:"collation,omitempty"`

	Tags *[]TagWithKeyValue `json:"tags,omitempty"`

	UnchangeableParam *UnchangeableParam `json:"unchangeable_param,omitempty"`

	DryRun *bool `json:"dry_run,omitempty"`
}

func (o CreateRestoreInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRestoreInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateRestoreInstanceRequestBody", string(data)}, " ")
}

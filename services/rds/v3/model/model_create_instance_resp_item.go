package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceRespItem struct {
	Id *string `json:"id,omitempty"`

	Name string `json:"name"`

	Status *string `json:"status,omitempty"`

	Datastore *Datastore `json:"datastore"`

	Ha *Ha `json:"ha,omitempty"`

	ConfigurationId *string `json:"configuration_id,omitempty"`

	Port *string `json:"port,omitempty"`

	BackupStrategy *BackupStrategy `json:"backup_strategy,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	DiskEncryptionId *string `json:"disk_encryption_id,omitempty"`

	FlavorRef string `json:"flavor_ref"`

	Volume *Volume `json:"volume"`

	Region string `json:"region"`

	AvailabilityZone string `json:"availability_zone"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`

	ChargeInfo *ChargeInfo `json:"charge_info,omitempty"`

	Collation *string `json:"collation,omitempty"`

	RestorePoint *RestorePoint `json:"restore_point,omitempty"`
}

func (o CreateInstanceRespItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceRespItem struct{}"
	}

	return strings.Join([]string{"CreateInstanceRespItem", string(data)}, " ")
}

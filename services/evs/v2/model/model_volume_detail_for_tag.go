package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type VolumeDetailForTag struct {
	Id string `json:"id"`

	Links []Link `json:"links"`

	Name string `json:"name"`

	Status string `json:"status"`

	Attachments []Attachment `json:"attachments"`

	AvailabilityZone string `json:"availability_zone"`

	OsVolHostAttrhost string `json:"os-vol-host-attr:host"`

	SourceVolid *string `json:"source_volid,omitempty"`

	SnapshotId string `json:"snapshot_id"`

	Description string `json:"description"`

	CreatedAt string `json:"created_at"`

	OsVolTenantAttrtenantId string `json:"os-vol-tenant-attr:tenant_id"`

	VolumeImageMetadata map[string]interface{} `json:"volume_image_metadata"`

	VolumeType string `json:"volume_type"`

	Size int32 `json:"size"`

	ConsistencygroupId *string `json:"consistencygroup_id,omitempty"`

	Bootable string `json:"bootable"`

	Metadata *VolumeMetadata `json:"metadata"`

	UpdatedAt string `json:"updated_at"`

	Encrypted *bool `json:"encrypted,omitempty"`

	ReplicationStatus string `json:"replication_status"`

	OsVolumeReplicationextendedStatus string `json:"os-volume-replication:extended_status"`

	OsVolMigStatusAttrmigstat string `json:"os-vol-mig-status-attr:migstat"`

	OsVolMigStatusAttrnameId string `json:"os-vol-mig-status-attr:name_id"`

	Shareable bool `json:"shareable"`

	UserId string `json:"user_id"`

	ServiceType string `json:"service_type"`

	Multiattach bool `json:"multiattach"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	Tags map[string]string `json:"tags"`

	Wwn *string `json:"wwn,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o VolumeDetailForTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDetailForTag struct{}"
	}

	return strings.Join([]string{"VolumeDetailForTag", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListVolumesRequest struct {
	Marker *string `json:"marker,omitempty"`

	Name *string `json:"name,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	SortKey *string `json:"sort_key,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	SortDir *string `json:"sort_dir,omitempty"`

	Status *string `json:"status,omitempty"`

	Metadata *string `json:"metadata,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	Multiattach *bool `json:"multiattach,omitempty"`

	ServiceType *string `json:"service_type,omitempty"`

	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	VolumeTypeId *string `json:"volume_type_id,omitempty"`

	Id *string `json:"id,omitempty"`

	Ids *string `json:"ids,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ServerId *string `json:"server_id,omitempty"`
}

func (o ListVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesRequest", string(data)}, " ")
}

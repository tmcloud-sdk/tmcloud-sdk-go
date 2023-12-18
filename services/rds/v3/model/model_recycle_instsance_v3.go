package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RecycleInstsanceV3 struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	HaMode *string `json:"ha_mode,omitempty"`

	EngineName *string `json:"engine_name,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	PayModel *string `json:"pay_model,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	DeletedAt *string `json:"deleted_at,omitempty"`

	VolumeType *string `json:"volume_type,omitempty"`

	VolumeSize *int32 `json:"volume_size,omitempty"`

	DataVip *string `json:"data_vip,omitempty"`

	DataVipV6 *string `json:"data_vip_v6,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	RetainedUntil *string `json:"retained_until,omitempty"`

	RecycleBackupId *string `json:"recycle_backup_id,omitempty"`

	RecycleStatus *string `json:"recycle_status,omitempty"`
}

func (o RecycleInstsanceV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleInstsanceV3 struct{}"
	}

	return strings.Join([]string{"RecycleInstsanceV3", string(data)}, " ")
}

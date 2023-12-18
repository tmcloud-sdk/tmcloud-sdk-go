package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type OffsiteBackupInstance struct {
	Id string `json:"id"`

	Name *string `json:"name,omitempty"`

	SourceRegion *string `json:"source_region,omitempty"`

	SourceProjectId *string `json:"source_project_id,omitempty"`

	Datastore *ParaGroupDatastore `json:"datastore,omitempty"`

	DestinationRegion *string `json:"destination_region,omitempty"`

	DestinationProjectId *string `json:"destination_project_id,omitempty"`

	KeepDays *int64 `json:"keep_days,omitempty"`
}

func (o OffsiteBackupInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OffsiteBackupInstance struct{}"
	}

	return strings.Join([]string{"OffsiteBackupInstance", string(data)}, " ")
}

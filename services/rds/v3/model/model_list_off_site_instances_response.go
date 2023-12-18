package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListOffSiteInstancesResponse struct {
	OffsiteBackupInstances *[]OffsiteBackupInstance `json:"offsite_backup_instances,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOffSiteInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListOffSiteInstancesResponse", string(data)}, " ")
}

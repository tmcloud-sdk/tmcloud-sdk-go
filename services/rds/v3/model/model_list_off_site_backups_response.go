package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListOffSiteBackupsResponse struct {
	Backups *[]OffSiteBackupForList `json:"backups,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOffSiteBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListOffSiteBackupsResponse", string(data)}, " ")
}

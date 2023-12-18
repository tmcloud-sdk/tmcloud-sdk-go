package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListOffSiteRestoreTimesResponse struct {
	RestoreTime    *[]GetRestoreTimeResponseRestoreTime `json:"restore_time,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListOffSiteRestoreTimesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOffSiteRestoreTimesResponse struct{}"
	}

	return strings.Join([]string{"ListOffSiteRestoreTimesResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMaintenanceWindowsResponse struct {
	MaintainWindows *[]MaintainWindowsEntity `json:"maintain_windows,omitempty"`
	HttpStatusCode  int                      `json:"-"`
}

func (o ListMaintenanceWindowsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMaintenanceWindowsResponse struct{}"
	}

	return strings.Join([]string{"ListMaintenanceWindowsResponse", string(data)}, " ")
}

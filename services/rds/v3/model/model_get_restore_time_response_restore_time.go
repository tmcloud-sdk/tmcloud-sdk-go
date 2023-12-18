package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GetRestoreTimeResponseRestoreTime struct {
	StartTime int64 `json:"start_time"`

	EndTime int64 `json:"end_time"`
}

func (o GetRestoreTimeResponseRestoreTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetRestoreTimeResponseRestoreTime struct{}"
	}

	return strings.Join([]string{"GetRestoreTimeResponseRestoreTime", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListRestoreRecordsResponse struct {
	RestoreRecordResponse *[]InstanceRestoreInfo `json:"restore_record_response,omitempty"`

	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRestoreRecordsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreRecordsResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreRecordsResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ScheduleKeyDeletionRequestBody struct {
	KeyId string `json:"key_id"`

	PendingDays string `json:"pending_days"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o ScheduleKeyDeletionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleKeyDeletionRequestBody struct{}"
	}

	return strings.Join([]string{"ScheduleKeyDeletionRequestBody", string(data)}, " ")
}

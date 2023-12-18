package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteKeyRequest struct {
	Body *ScheduleKeyDeletionRequestBody `json:"body,omitempty"`
}

func (o DeleteKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteKeyRequest", string(data)}, " ")
}

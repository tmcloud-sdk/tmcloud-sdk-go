package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowReplicationStatusResponse struct {
	ReplicationStatus *string `json:"replication_status,omitempty"`

	AbnormalReason *string `json:"abnormal_reason,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowReplicationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowReplicationStatusResponse", string(data)}, " ")
}

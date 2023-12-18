package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowDrReplicaStatusResponse struct {
	ReplicaState *string `json:"replica_state,omitempty"`

	WalWriteReceiveDelayInMb *string `json:"wal_write_receive_delay_in_mb,omitempty"`

	WalWriteReplayDelayInMb *string `json:"wal_write_replay_delay_in_mb,omitempty"`

	WalReceiveReplayDelayInMs *string `json:"wal_receive_replay_delay_in_ms,omitempty"`
	HttpStatusCode            int     `json:"-"`
}

func (o ShowDrReplicaStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrReplicaStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowDrReplicaStatusResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterStatus struct {
	Phase *string `json:"phase,omitempty"`

	JobID *string `json:"jobID,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Message *string `json:"message,omitempty"`

	Endpoints *[]ClusterEndpoints `json:"endpoints,omitempty"`

	IsLocked *bool `json:"isLocked,omitempty"`

	LockScene *string `json:"lockScene,omitempty"`

	LockSource *string `json:"lockSource,omitempty"`

	LockSourceId *string `json:"lockSourceId,omitempty"`

	DeleteOption *interface{} `json:"deleteOption,omitempty"`

	DeleteStatus *interface{} `json:"deleteStatus,omitempty"`
}

func (o ClusterStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterStatus struct{}"
	}

	return strings.Join([]string{"ClusterStatus", string(data)}, " ")
}

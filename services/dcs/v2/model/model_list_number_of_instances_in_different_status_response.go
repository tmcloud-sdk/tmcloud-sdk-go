package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListNumberOfInstancesInDifferentStatusResponse struct {
	Redis *StatusStatistic `json:"redis,omitempty"`

	Memcached *StatusStatistic `json:"memcached,omitempty"`

	PayingCount *int32 `json:"paying_count,omitempty"`

	FreezingCount *int32 `json:"freezing_count,omitempty"`

	MigratingCount *int32 `json:"migrating_count,omitempty"`

	FlushingCount *int32 `json:"flushing_count,omitempty"`

	UpgradingCount *int32 `json:"upgrading_count,omitempty"`

	RestoringCount *int32 `json:"restoring_count,omitempty"`

	ExtendingCount *int32 `json:"extending_count,omitempty"`

	CreatingCount *int32 `json:"creating_count,omitempty"`

	RunningCount *int32 `json:"running_count,omitempty"`

	ErrorCount *int32 `json:"error_count,omitempty"`

	FrozenCount *int32 `json:"frozen_count,omitempty"`

	CreatefailedCount *int32 `json:"createfailed_count,omitempty"`

	RestartingCount *int32 `json:"restarting_count,omitempty"`
	HttpStatusCode  int    `json:"-"`
}

func (o ListNumberOfInstancesInDifferentStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNumberOfInstancesInDifferentStatusResponse struct{}"
	}

	return strings.Join([]string{"ListNumberOfInstancesInDifferentStatusResponse", string(data)}, " ")
}

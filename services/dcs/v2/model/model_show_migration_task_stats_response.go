package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowMigrationTaskStatsResponse struct {
	FullMigrationProgress *string `json:"full_migration_progress,omitempty"`

	Offset *string `json:"offset,omitempty"`

	SourceDbsize *string `json:"source_dbsize,omitempty"`

	TargetDbsize *string `json:"target_dbsize,omitempty"`

	TargetInputKbps *string `json:"target_input_kbps,omitempty"`

	TargetOps *string `json:"target_ops,omitempty"`

	IsMigrating    *bool `json:"is_migrating,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowMigrationTaskStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationTaskStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowMigrationTaskStatsResponse", string(data)}, " ")
}

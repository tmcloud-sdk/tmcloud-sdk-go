package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListMigrationTaskRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Name *string `json:"name,omitempty"`
}

func (o ListMigrationTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMigrationTaskRequest struct{}"
	}

	return strings.Join([]string{"ListMigrationTaskRequest", string(data)}, " ")
}

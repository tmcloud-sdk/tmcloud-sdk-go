package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SingleBackgroundTask struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Details *DetailsBody `json:"details,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	Params *string `json:"params,omitempty"`

	Status *string `json:"status,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	EnableShow *bool `json:"enable_show,omitempty"`
}

func (o SingleBackgroundTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SingleBackgroundTask struct{}"
	}

	return strings.Join([]string{"SingleBackgroundTask", string(data)}, " ")
}

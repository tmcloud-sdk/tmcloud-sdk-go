package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListBackgroundTaskResponse struct {
	TaskCount *string `json:"task_count,omitempty"`

	Tasks *[]SingleBackgroundTask `json:"tasks,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBackgroundTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackgroundTaskResponse struct{}"
	}

	return strings.Join([]string{"ListBackgroundTaskResponse", string(data)}, " ")
}

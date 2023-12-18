package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type JobEntitiesResult struct {
	ImageId *string `json:"image_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	Status *string `json:"status,omitempty"`
}

func (o JobEntitiesResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntitiesResult struct{}"
	}

	return strings.Join([]string{"JobEntitiesResult", string(data)}, " ")
}

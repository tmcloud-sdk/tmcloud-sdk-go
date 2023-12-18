package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type JobEntities struct {
	VolumeType *string `json:"volume_type,omitempty"`

	Size *int32 `json:"size,omitempty"`

	VolumeId *string `json:"volume_id,omitempty"`

	Name *string `json:"name,omitempty"`

	SubJobs *[]SubJob `json:"sub_jobs,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}

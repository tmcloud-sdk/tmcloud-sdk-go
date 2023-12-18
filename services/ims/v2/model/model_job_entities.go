package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type JobEntities struct {
	ImageId *string `json:"image_id,omitempty"`

	CurrentTask *string `json:"current_task,omitempty"`

	ImageName *string `json:"image_name,omitempty"`

	ProcessPercent *float64 `json:"process_percent,omitempty"`

	Results *[]JobEntitiesResult `json:"results,omitempty"`

	SubJobsResult *[]SubJobResult `json:"sub_jobs_result,omitempty"`

	SubJobsList *[]string `json:"sub_jobs_list,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}

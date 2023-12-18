package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type JobSpec struct {
	Type *string `json:"type,omitempty"`

	ClusterUID *string `json:"clusterUID,omitempty"`

	ResourceID *string `json:"resourceID,omitempty"`

	ResourceName *string `json:"resourceName,omitempty"`

	ExtendParam map[string]string `json:"extendParam,omitempty"`

	SubJobs *[]Job `json:"subJobs,omitempty"`
}

func (o JobSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSpec struct{}"
	}

	return strings.Join([]string{"JobSpec", string(data)}, " ")
}

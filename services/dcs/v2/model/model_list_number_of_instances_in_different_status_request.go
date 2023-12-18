package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListNumberOfInstancesInDifferentStatusRequest struct {
	IncludeFailure *string `json:"include_failure,omitempty"`
}

func (o ListNumberOfInstancesInDifferentStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNumberOfInstancesInDifferentStatusRequest struct{}"
	}

	return strings.Join([]string{"ListNumberOfInstancesInDifferentStatusRequest", string(data)}, " ")
}

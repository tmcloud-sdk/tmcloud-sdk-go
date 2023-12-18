package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListSlowLogFileRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSlowLogFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogFileRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogFileRequest", string(data)}, " ")
}

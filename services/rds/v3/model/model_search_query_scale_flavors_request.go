package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SearchQueryScaleFlavorsRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`
}

func (o SearchQueryScaleFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryScaleFlavorsRequest struct{}"
	}

	return strings.Join([]string{"SearchQueryScaleFlavorsRequest", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchUpdateServersNameRequestBody struct {
	Name string `json:"name"`

	DryRun *bool `json:"dry_run,omitempty"`

	Servers []ServerId `json:"servers"`
}

func (o BatchUpdateServersNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateServersNameRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateServersNameRequestBody", string(data)}, " ")
}

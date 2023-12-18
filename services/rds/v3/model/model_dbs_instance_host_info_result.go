package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DbsInstanceHostInfoResult struct {
	Id *string `json:"id,omitempty"`

	Host *string `json:"host,omitempty"`

	HostName *string `json:"host_name,omitempty"`
}

func (o DbsInstanceHostInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DbsInstanceHostInfoResult struct{}"
	}

	return strings.Join([]string{"DbsInstanceHostInfoResult", string(data)}, " ")
}

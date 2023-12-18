package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerAddress struct {
	Version int32 `json:"version"`

	Addr string `json:"addr"`
}

func (o UpdateServerAddress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerAddress struct{}"
	}

	return strings.Join([]string{"UpdateServerAddress", string(data)}, " ")
}

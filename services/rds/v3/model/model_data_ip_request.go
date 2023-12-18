package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DataIpRequest struct {
	NewIp string `json:"new_ip"`
}

func (o DataIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataIpRequest struct{}"
	}

	return strings.Join([]string{"DataIpRequest", string(data)}, " ")
}

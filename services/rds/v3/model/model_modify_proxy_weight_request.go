package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyProxyWeightRequest struct {
	MasterWeight string `json:"master_weight"`

	ReadonlyInstances []ProxyReadonlyInstances `json:"readonly_instances"`
}

func (o ModifyProxyWeightRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyProxyWeightRequest struct{}"
	}

	return strings.Join([]string{"ModifyProxyWeightRequest", string(data)}, " ")
}

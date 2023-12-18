package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAddonInstancesRequest struct {
	AddonTemplateName *string `json:"addon_template_name,omitempty"`

	ClusterId string `json:"cluster_id"`
}

func (o ListAddonInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAddonInstancesRequest", string(data)}, " ")
}

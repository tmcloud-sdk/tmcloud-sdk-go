package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceSpec struct {
	ClusterID string `json:"clusterID"`

	Version string `json:"version"`

	AddonTemplateName string `json:"addonTemplateName"`

	AddonTemplateType string `json:"addonTemplateType"`

	AddonTemplateLogo *string `json:"addonTemplateLogo,omitempty"`

	AddonTemplateLabels *[]string `json:"addonTemplateLabels,omitempty"`

	Description string `json:"description"`

	Values map[string]interface{} `json:"values"`
}

func (o InstanceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceSpec struct{}"
	}

	return strings.Join([]string{"InstanceSpec", string(data)}, " ")
}

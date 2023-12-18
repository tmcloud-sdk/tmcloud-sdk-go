package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyImageCrossRegionRequestBody struct {
	AgencyName string `json:"agency_name"`

	Description *string `json:"description,omitempty"`

	Name string `json:"name"`

	ProjectName string `json:"project_name"`

	Region string `json:"region"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o CopyImageCrossRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageCrossRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageCrossRegionRequestBody", string(data)}, " ")
}

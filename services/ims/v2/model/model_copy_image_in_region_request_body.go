package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyImageInRegionRequestBody struct {
	CmkId *string `json:"cmk_id,omitempty"`

	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Name string `json:"name"`
}

func (o CopyImageInRegionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyImageInRegionRequestBody struct{}"
	}

	return strings.Join([]string{"CopyImageInRegionRequestBody", string(data)}, " ")
}

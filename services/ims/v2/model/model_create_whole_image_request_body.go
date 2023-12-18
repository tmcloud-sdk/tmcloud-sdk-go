package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateWholeImageRequestBody struct {
	Description *string `json:"description,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	ImageTags *[]TagKeyValue `json:"image_tags,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Name string `json:"name"`

	Tags *[]string `json:"tags,omitempty"`

	BackupId *string `json:"backup_id,omitempty"`

	WholeImageType *string `json:"whole_image_type,omitempty"`

	MaxRam *int32 `json:"max_ram,omitempty"`

	MinRam *int32 `json:"min_ram,omitempty"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o CreateWholeImageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWholeImageRequestBody struct{}"
	}

	return strings.Join([]string{"CreateWholeImageRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateInstanceBody struct {
	Name string `json:"name"`

	Engine string `json:"engine"`

	EngineVersion *string `json:"engine_version,omitempty"`

	Capacity float32 `json:"capacity"`

	SpecCode string `json:"spec_code"`

	AzCodes []string `json:"az_codes"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	Description *string `json:"description,omitempty"`

	EnableSsl *bool `json:"enable_ssl,omitempty"`

	PrivateIp *string `json:"private_ip,omitempty"`

	InstanceNum *int32 `json:"instance_num,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	Password *string `json:"password,omitempty"`

	NoPasswordAccess *bool `json:"no_password_access,omitempty"`

	BssParam *BssParam `json:"bss_param,omitempty"`

	InstanceBackupPolicy *BackupPolicy `json:"instance_backup_policy,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	AccessUser *string `json:"access_user,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	Port *int32 `json:"port,omitempty"`

	RenameCommands *interface{} `json:"rename_commands,omitempty"`
}

func (o CreateInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceBody struct{}"
	}

	return strings.Join([]string{"CreateInstanceBody", string(data)}, " ")
}

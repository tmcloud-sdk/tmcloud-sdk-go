package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type InstanceListInfo struct {
	PublicipId *string `json:"publicip_id,omitempty"`

	VpcName *string `json:"vpc_name,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	EnableSsl *bool `json:"enable_ssl,omitempty"`

	MaxMemory *int32 `json:"max_memory,omitempty"`

	UsedMemory *int32 `json:"used_memory,omitempty"`

	PublicipAddress *string `json:"publicip_address,omitempty"`

	Capacity *int32 `json:"capacity,omitempty"`

	CapacityMinor *string `json:"capacity_minor,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	Engine *string `json:"engine,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	ServiceUpgrade *bool `json:"service_upgrade,omitempty"`

	NoPasswordAccess *string `json:"no_password_access,omitempty"`

	ServiceTaskId *string `json:"service_task_id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	AccessUser *string `json:"access_user,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	Port *int32 `json:"port,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	ReadonlyDomainName *string `json:"readonly_domain_name,omitempty"`

	Name *string `json:"name,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	Status *string `json:"status,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Description *string `json:"description,omitempty"`

	CpuType *string `json:"cpu_type,omitempty"`

	AzCodes *[]string `json:"az_codes,omitempty"`

	Features *Features `json:"features,omitempty"`

	SubStatus *string `json:"sub_status,omitempty"`
}

func (o InstanceListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceListInfo struct{}"
	}

	return strings.Join([]string{"InstanceListInfo", string(data)}, " ")
}

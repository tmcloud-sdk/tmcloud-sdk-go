package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ShowInstanceResponse struct {
	VpcName *string `json:"vpc_name,omitempty"`

	ChargingMode *int32 `json:"charging_mode,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	UserName *string `json:"user_name,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	Description *string `json:"description,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	SecurityGroupName *string `json:"security_group_name,omitempty"`

	MaxMemory *int32 `json:"max_memory,omitempty"`

	UsedMemory *int32 `json:"used_memory,omitempty"`

	Capacity *int32 `json:"capacity,omitempty"`

	CapacityMinor *string `json:"capacity_minor,omitempty"`

	MaintainBegin *string `json:"maintain_begin,omitempty"`

	MaintainEnd *string `json:"maintain_end,omitempty"`

	Engine *string `json:"engine,omitempty"`

	NoPasswordAccess *string `json:"no_password_access,omitempty"`

	Ip *string `json:"ip,omitempty"`

	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty"`

	AzCodes *[]string `json:"az_codes,omitempty"`

	AccessUser *string `json:"access_user,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	Port *int32 `json:"port,omitempty"`

	UserId *string `json:"user_id,omitempty"`

	Name *string `json:"name,omitempty"`

	SpecCode *string `json:"spec_code,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SubnetName *string `json:"subnet_name,omitempty"`

	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	EngineVersion *string `json:"engine_version,omitempty"`

	OrderId *string `json:"order_id,omitempty"`

	Status *string `json:"status,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	ReadonlyDomainName *string `json:"readonly_domain_name,omitempty"`

	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	PublicipId *string `json:"publicip_id,omitempty"`

	PublicipAddress *string `json:"publicip_address,omitempty"`

	EnableSsl *bool `json:"enable_ssl,omitempty"`

	ServiceUpgrade *bool `json:"service_upgrade,omitempty"`

	ServiceTaskId *string `json:"service_task_id,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	BackendAddrs *string `json:"backend_addrs,omitempty"`

	Features *Features `json:"features,omitempty"`

	DomainNameInfo *DomainNameInfo `json:"domain_name_info,omitempty"`

	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	SubStatus *string `json:"sub_status,omitempty"`

	Tags *[]ResourceTag `json:"tags,omitempty"`

	CpuType *string `json:"cpu_type,omitempty"`

	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	UpdateAt *string `json:"update_at,omitempty"`

	ProductType *ShowInstanceResponseProductType `json:"product_type,omitempty"`

	StorageType *ShowInstanceResponseStorageType `json:"storage_type,omitempty"`

	LaunchedAt *string `json:"launched_at,omitempty"`

	CacheMode *ShowInstanceResponseCacheMode `json:"cache_mode,omitempty"`

	SupportSlowLogFlag *string `json:"support_slow_log_flag,omitempty"`

	DbNumber *int32 `json:"db_number,omitempty"`

	ReplicaCount *int32 `json:"replica_count,omitempty"`

	ShardingCount *int32 `json:"sharding_count,omitempty"`

	BandwidthInfo  *BandwidthInfo `json:"bandwidth_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}

type ShowInstanceResponseProductType struct {
	value string
}

type ShowInstanceResponseProductTypeEnum struct {
	GENERIC    ShowInstanceResponseProductType
	ENTERPRISE ShowInstanceResponseProductType
}

func GetShowInstanceResponseProductTypeEnum() ShowInstanceResponseProductTypeEnum {
	return ShowInstanceResponseProductTypeEnum{
		GENERIC: ShowInstanceResponseProductType{
			value: "generic",
		},
		ENTERPRISE: ShowInstanceResponseProductType{
			value: "enterprise",
		},
	}
}

func (c ShowInstanceResponseProductType) Value() string {
	return c.value
}

func (c ShowInstanceResponseProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseProductType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowInstanceResponseStorageType struct {
	value string
}

type ShowInstanceResponseStorageTypeEnum struct {
	DRAM ShowInstanceResponseStorageType
}

func GetShowInstanceResponseStorageTypeEnum() ShowInstanceResponseStorageTypeEnum {
	return ShowInstanceResponseStorageTypeEnum{
		DRAM: ShowInstanceResponseStorageType{
			value: "DRAM",
		},
	}
}

func (c ShowInstanceResponseStorageType) Value() string {
	return c.value
}

func (c ShowInstanceResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseStorageType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowInstanceResponseCacheMode struct {
	value string
}

type ShowInstanceResponseCacheModeEnum struct {
	SINGLE      ShowInstanceResponseCacheMode
	HA          ShowInstanceResponseCacheMode
	HA_RW_SPLIT ShowInstanceResponseCacheMode
	PROXY       ShowInstanceResponseCacheMode
	CLUSTER     ShowInstanceResponseCacheMode
}

func GetShowInstanceResponseCacheModeEnum() ShowInstanceResponseCacheModeEnum {
	return ShowInstanceResponseCacheModeEnum{
		SINGLE: ShowInstanceResponseCacheMode{
			value: "single",
		},
		HA: ShowInstanceResponseCacheMode{
			value: "ha",
		},
		HA_RW_SPLIT: ShowInstanceResponseCacheMode{
			value: "ha_rw_split",
		},
		PROXY: ShowInstanceResponseCacheMode{
			value: "proxy",
		},
		CLUSTER: ShowInstanceResponseCacheMode{
			value: "cluster",
		},
	}
}

func (c ShowInstanceResponseCacheMode) Value() string {
	return c.value
}

func (c ShowInstanceResponseCacheMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseCacheMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

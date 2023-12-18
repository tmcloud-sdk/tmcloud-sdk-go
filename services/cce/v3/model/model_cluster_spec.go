package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ClusterSpec struct {
	Category *ClusterSpecCategory `json:"category,omitempty"`

	Type *ClusterSpecType `json:"type,omitempty"`

	Flavor string `json:"flavor"`

	Version *string `json:"version,omitempty"`

	PlatformVersion *string `json:"platformVersion,omitempty"`

	Description *string `json:"description,omitempty"`

	CustomSan *[]string `json:"customSan,omitempty"`

	Ipv6enable *bool `json:"ipv6enable,omitempty"`

	OffloadCluster *bool `json:"offloadCluster,omitempty"`

	HostNetwork *HostNetwork `json:"hostNetwork"`

	ContainerNetwork *ContainerNetwork `json:"containerNetwork"`

	EniNetwork *EniNetwork `json:"eniNetwork,omitempty"`

	Authentication *Authentication `json:"authentication,omitempty"`

	BillingMode *int32 `json:"billingMode,omitempty"`

	Masters *[]MasterSpec `json:"masters,omitempty"`

	KubernetesSvcIpRange *string `json:"kubernetesSvcIpRange,omitempty"`

	ClusterTags *[]ResourceTag `json:"clusterTags,omitempty"`

	KubeProxyMode *ClusterSpecKubeProxyMode `json:"kubeProxyMode,omitempty"`

	Az *string `json:"az,omitempty"`

	ExtendParam *ClusterExtendParam `json:"extendParam,omitempty"`

	SupportIstio *bool `json:"supportIstio,omitempty"`

	ConfigurationsOverride *[]PackageConfiguration `json:"configurationsOverride,omitempty"`
}

func (o ClusterSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterSpec struct{}"
	}

	return strings.Join([]string{"ClusterSpec", string(data)}, " ")
}

type ClusterSpecCategory struct {
	value string
}

type ClusterSpecCategoryEnum struct {
	CCE   ClusterSpecCategory
	TURBO ClusterSpecCategory
}

func GetClusterSpecCategoryEnum() ClusterSpecCategoryEnum {
	return ClusterSpecCategoryEnum{
		CCE: ClusterSpecCategory{
			value: "CCE",
		},
		TURBO: ClusterSpecCategory{
			value: "Turbo",
		},
	}
}

func (c ClusterSpecCategory) Value() string {
	return c.value
}

func (c ClusterSpecCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSpecCategory) UnmarshalJSON(b []byte) error {
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

type ClusterSpecType struct {
	value string
}

type ClusterSpecTypeEnum struct {
	VIRTUAL_MACHINE ClusterSpecType
	ARM64           ClusterSpecType
}

func GetClusterSpecTypeEnum() ClusterSpecTypeEnum {
	return ClusterSpecTypeEnum{
		VIRTUAL_MACHINE: ClusterSpecType{
			value: "VirtualMachine",
		},
		ARM64: ClusterSpecType{
			value: "ARM64",
		},
	}
}

func (c ClusterSpecType) Value() string {
	return c.value
}

func (c ClusterSpecType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSpecType) UnmarshalJSON(b []byte) error {
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

type ClusterSpecKubeProxyMode struct {
	value string
}

type ClusterSpecKubeProxyModeEnum struct {
	IPTABLES ClusterSpecKubeProxyMode
	IPVS     ClusterSpecKubeProxyMode
}

func GetClusterSpecKubeProxyModeEnum() ClusterSpecKubeProxyModeEnum {
	return ClusterSpecKubeProxyModeEnum{
		IPTABLES: ClusterSpecKubeProxyMode{
			value: "iptables",
		},
		IPVS: ClusterSpecKubeProxyMode{
			value: "ipvs",
		},
	}
}

func (c ClusterSpecKubeProxyMode) Value() string {
	return c.value
}

func (c ClusterSpecKubeProxyMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterSpecKubeProxyMode) UnmarshalJSON(b []byte) error {
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

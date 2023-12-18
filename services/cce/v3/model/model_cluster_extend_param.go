package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterExtendParam struct {
	ClusterAZ *string `json:"clusterAZ,omitempty"`

	DssMasterVolumes *string `json:"dssMasterVolumes,omitempty"`

	EnterpriseProjectId *string `json:"enterpriseProjectId,omitempty"`

	KubeProxyMode *string `json:"kubeProxyMode,omitempty"`

	ClusterExternalIP *string `json:"clusterExternalIP,omitempty"`

	AlphaCceFixPoolMask *string `json:"alpha.cce/fixPoolMask,omitempty"`

	DecMasterFlavor *string `json:"decMasterFlavor,omitempty"`

	DockerUmaskMode *string `json:"dockerUmaskMode,omitempty"`

	KubernetesIoCpuManagerPolicy *string `json:"kubernetes.io/cpuManagerPolicy,omitempty"`

	OrderID *string `json:"orderID,omitempty"`

	PeriodType *string `json:"periodType,omitempty"`

	PeriodNum *int32 `json:"periodNum,omitempty"`

	IsAutoRenew *string `json:"isAutoRenew,omitempty"`

	IsAutoPay *string `json:"isAutoPay,omitempty"`

	Upgradefrom *string `json:"upgradefrom,omitempty"`
}

func (o ClusterExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterExtendParam struct{}"
	}

	return strings.Join([]string{"ClusterExtendParam", string(data)}, " ")
}

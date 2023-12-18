package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodeExtendParam struct {
	Ecsperformancetype *string `json:"ecs:performancetype,omitempty"`

	OrderID *string `json:"orderID,omitempty"`

	ProductID *string `json:"productID,omitempty"`

	MaxPods *int32 `json:"maxPods,omitempty"`

	PeriodType *string `json:"periodType,omitempty"`

	PeriodNum *int32 `json:"periodNum,omitempty"`

	IsAutoRenew *string `json:"isAutoRenew,omitempty"`

	IsAutoPay *string `json:"isAutoPay,omitempty"`

	DockerLVMConfigOverride *string `json:"DockerLVMConfigOverride,omitempty"`

	DockerBaseSize *int32 `json:"dockerBaseSize,omitempty"`

	OffloadNode *string `json:"offloadNode,omitempty"`

	PublicKey *string `json:"publicKey,omitempty"`

	AlphaCcePreInstall *string `json:"alpha.cce/preInstall,omitempty"`

	AlphaCcePostInstall *string `json:"alpha.cce/postInstall,omitempty"`

	AlphaCceNodeImageID *string `json:"alpha.cce/NodeImageID,omitempty"`

	NicMultiqueue *string `json:"nicMultiqueue,omitempty"`

	NicThreshold *string `json:"nicThreshold,omitempty"`

	ChargingMode *int32 `json:"chargingMode,omitempty"`

	AgencyName *string `json:"agency_name,omitempty"`

	KubeReservedMem *int32 `json:"kube-reserved-mem,omitempty"`

	SystemReservedMem *int32 `json:"system-reserved-mem,omitempty"`
}

func (o NodeExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExtendParam struct{}"
	}

	return strings.Join([]string{"NodeExtendParam", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodeSpec struct {
	Flavor string `json:"flavor"`

	Az string `json:"az"`

	Os *string `json:"os,omitempty"`

	Login *Login `json:"login"`

	RootVolume *Volume `json:"rootVolume"`

	DataVolumes []Volume `json:"dataVolumes"`

	Storage *Storage `json:"storage,omitempty"`

	PublicIP *NodePublicIp `json:"publicIP,omitempty"`

	NodeNicSpec *NodeNicSpec `json:"nodeNicSpec,omitempty"`

	Count *int32 `json:"count,omitempty"`

	BillingMode *int32 `json:"billingMode,omitempty"`

	Taints *[]Taint `json:"taints,omitempty"`

	K8sTags map[string]string `json:"k8sTags,omitempty"`

	EcsGroupId *string `json:"ecsGroupId,omitempty"`

	FaultDomain *string `json:"faultDomain,omitempty"`

	DedicatedHostId *string `json:"dedicatedHostId,omitempty"`

	OffloadNode *bool `json:"offloadNode,omitempty"`

	IsStatic *bool `json:"isStatic,omitempty"`

	UserTags *[]UserTag `json:"userTags,omitempty"`

	Runtime *Runtime `json:"runtime,omitempty"`

	InitializedConditions *[]string `json:"initializedConditions,omitempty"`

	ExtendParam *NodeExtendParam `json:"extendParam,omitempty"`
}

func (o NodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpec struct{}"
	}

	return strings.Join([]string{"NodeSpec", string(data)}, " ")
}

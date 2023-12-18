package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type FlavorExtraSpec struct {
	Ecsperformancetype *string `json:"ecs:performancetype,omitempty"`

	HwnumaNodes *string `json:"hw:numa_nodes,omitempty"`

	ResourceType *string `json:"resource_type,omitempty"`

	HpetSupport *string `json:"hpet_support,omitempty"`

	InstanceVnictype *string `json:"instance_vnic:type,omitempty"`

	InstanceVnicinstanceBandwidth *string `json:"instance_vnic:instance_bandwidth,omitempty"`

	InstanceVnicmaxCount *string `json:"instance_vnic:max_count,omitempty"`

	QuotalocalDisk *string `json:"quota:local_disk,omitempty"`

	QuotanvmeSsd *string `json:"quota:nvme_ssd,omitempty"`

	ExtraSpeciopersistentGrant *string `json:"extra_spec:io:persistent_grant,omitempty"`

	Ecsgeneration *string `json:"ecs:generation,omitempty"`

	EcsvirtualizationEnvTypes *string `json:"ecs:virtualization_env_types,omitempty"`

	PciPassthroughenableGpu *string `json:"pci_passthrough:enable_gpu,omitempty"`

	PciPassthroughgpuSpecs *string `json:"pci_passthrough:gpu_specs,omitempty"`

	PciPassthroughalias *string `json:"pci_passthrough:alias,omitempty"`

	Condoperationstatus *string `json:"cond:operation:status,omitempty"`

	Condoperationaz *string `json:"cond:operation:az,omitempty"`

	QuotamaxRate *string `json:"quota:max_rate,omitempty"`

	QuotaminRate *string `json:"quota:min_rate,omitempty"`

	QuotamaxPps *string `json:"quota:max_pps,omitempty"`

	Condoperationcharge *string `json:"cond:operation:charge,omitempty"`

	Condoperationchargestop *string `json:"cond:operation:charge:stop,omitempty"`

	Condspotoperationaz *string `json:"cond:spot:operation:az,omitempty"`

	Condoperationroles *string `json:"cond:operation:roles,omitempty"`

	Condspotoperationstatus *string `json:"cond:spot:operation:status,omitempty"`

	Condnetwork *string `json:"cond:network,omitempty"`

	Condstorage *string `json:"cond:storage,omitempty"`

	CondcomputeliveResizable *string `json:"cond:compute:live_resizable,omitempty"`

	Condcompute *string `json:"cond:compute,omitempty"`

	Infogpuname *string `json:"info:gpu:name,omitempty"`

	Infocpuname *string `json:"info:cpu:name,omitempty"`

	Quotagpu *string `json:"quota:gpu,omitempty"`

	EcsinstanceArchitecture *string `json:"ecs:instance_architecture,omitempty"`
}

func (o FlavorExtraSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorExtraSpec struct{}"
	}

	return strings.Join([]string{"FlavorExtraSpec", string(data)}, " ")
}

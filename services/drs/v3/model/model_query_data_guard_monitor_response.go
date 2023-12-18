package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QueryDataGuardMonitorResponse struct {
	Bandwidth *string `json:"bandwidth,omitempty"`

	CpuUsedPercent *string `json:"cpuUsed_percent,omitempty"`

	DstDelay *int64 `json:"dst_delay,omitempty"`

	DstIo *string `json:"dst_io,omitempty"`

	DstNormal *bool `json:"dst_normal,omitempty"`

	DstOffset *string `json:"dst_offset,omitempty"`

	DstRps *string `json:"dst_rps,omitempty"`

	MemUsedInMB *string `json:"mem_used_inMB,omitempty"`

	NodeMemInMB *int64 `json:"node_mem_inMB,omitempty"`

	NodeOffset *string `json:"node_offset,omitempty"`

	NodeVolumeInGB *int64 `json:"node_volume_inGB,omitempty"`

	SrDelay *int64 `json:"sr_delay,omitempty"`

	SrOffset *string `json:"sr_offset,omitempty"`

	SrcIo *string `json:"src_io,omitempty"`

	SrcNormal *bool `json:"src_normal,omitempty"`

	SrcRps *string `json:"src_rps,omitempty"`

	TransInMB *string `json:"trans_inMB,omitempty"`

	TransLines *string `json:"trans_lines,omitempty"`

	VolumeUsedInGB *string `json:"volume_used_inGB,omitempty"`

	MigrationBytesPerSecond *int64 `json:"migration_bytes_per_second,omitempty"`
}

func (o QueryDataGuardMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryDataGuardMonitorResponse struct{}"
	}

	return strings.Join([]string{"QueryDataGuardMonitorResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PrePaidServer struct {
	AutoTerminateTime *string `json:"auto_terminate_time,omitempty"`

	ImageRef string `json:"imageRef"`

	FlavorRef string `json:"flavorRef"`

	Name string `json:"name"`

	UserData *string `json:"user_data,omitempty"`

	AdminPass *string `json:"adminPass,omitempty"`

	KeyName *string `json:"key_name,omitempty"`

	Vpcid string `json:"vpcid"`

	Nics []PrePaidServerNic `json:"nics"`

	Publicip *PrePaidServerPublicip `json:"publicip,omitempty"`

	Count *int32 `json:"count,omitempty"`

	IsAutoRename *bool `json:"isAutoRename,omitempty"`

	RootVolume *PrePaidServerRootVolume `json:"root_volume"`

	DataVolumes *[]PrePaidServerDataVolume `json:"data_volumes,omitempty"`

	SecurityGroups *[]PrePaidServerSecurityGroup `json:"security_groups,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	BatchCreateInMultiAz *bool `json:"batch_create_in_multi_az,omitempty"`

	Extendparam *PrePaidServerExtendParam `json:"extendparam,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	OsschedulerHints *PrePaidServerSchedulerHints `json:"os:scheduler_hints,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	ServerTags *[]PrePaidServerTag `json:"server_tags,omitempty"`

	Description *string `json:"description,omitempty"`
}

func (o PrePaidServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrePaidServer struct{}"
	}

	return strings.Join([]string{"PrePaidServer", string(data)}, " ")
}

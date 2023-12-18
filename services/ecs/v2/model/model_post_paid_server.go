package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PostPaidServer struct {
	AutoTerminateTime *string `json:"auto_terminate_time,omitempty"`

	AdminPass *string `json:"adminPass,omitempty"`

	AvailabilityZone *string `json:"availability_zone,omitempty"`

	BatchCreateInMultiAz *bool `json:"batch_create_in_multi_az,omitempty"`

	Count *int32 `json:"count,omitempty"`

	DataVolumes *[]PostPaidServerDataVolume `json:"data_volumes,omitempty"`

	Extendparam *PostPaidServerExtendParam `json:"extendparam,omitempty"`

	FlavorRef string `json:"flavorRef"`

	ImageRef string `json:"imageRef"`

	IsAutoRename *bool `json:"isAutoRename,omitempty"`

	KeyName *string `json:"key_name,omitempty"`

	Metadata map[string]string `json:"metadata,omitempty"`

	Name string `json:"name"`

	Nics []PostPaidServerNic `json:"nics"`

	OsschedulerHints *PostPaidServerSchedulerHints `json:"os:scheduler_hints,omitempty"`

	Publicip *PostPaidServerPublicip `json:"publicip,omitempty"`

	RootVolume *PostPaidServerRootVolume `json:"root_volume"`

	SecurityGroups *[]PostPaidServerSecurityGroup `json:"security_groups,omitempty"`

	ServerTags *[]PostPaidServerTag `json:"server_tags,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	UserData *string `json:"user_data,omitempty"`

	Vpcid string `json:"vpcid"`

	Description *string `json:"description,omitempty"`
}

func (o PostPaidServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServer struct{}"
	}

	return strings.Join([]string{"PostPaidServer", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CopyConfigurationResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DatastoreVersionName *string `json:"datastore_version_name,omitempty"`

	DatastoreName *string `json:"datastore_name,omitempty"`

	CreateTime *string `json:"create_time,omitempty"`

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CopyConfigurationResponse", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowInformationAboutDatabaseProxyResponse struct {
	Proxy *Proxy `json:"proxy,omitempty"`

	MasterInstance *MasterInstance `json:"master_instance,omitempty"`

	ReadonlyInstances *[]ReadonlyInstances `json:"readonly_instances,omitempty"`
	HttpStatusCode    int                  `json:"-"`
}

func (o ShowInformationAboutDatabaseProxyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInformationAboutDatabaseProxyResponse struct{}"
	}

	return strings.Join([]string{"ShowInformationAboutDatabaseProxyResponse", string(data)}, " ")
}

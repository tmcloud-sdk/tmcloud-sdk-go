package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListAddonInstancesResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Items          *[]AddonInstance `json:"items,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAddonInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAddonInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListAddonInstancesResponse", string(data)}, " ")
}

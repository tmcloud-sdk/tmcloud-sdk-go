package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowFlavorResponse struct {
	RequestId *string `json:"request_id,omitempty"`

	Flavor         *Flavor `json:"flavor,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowFlavorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFlavorResponse struct{}"
	}

	return strings.Join([]string{"ShowFlavorResponse", string(data)}, " ")
}

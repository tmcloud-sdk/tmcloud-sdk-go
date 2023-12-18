package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MasterEipResponseSpecSpec struct {
	Id *string `json:"id,omitempty"`

	Eip *EipSpec `json:"eip,omitempty"`

	IsDynamic *bool `json:"IsDynamic,omitempty"`
}

func (o MasterEipResponseSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipResponseSpecSpec", string(data)}, " ")
}

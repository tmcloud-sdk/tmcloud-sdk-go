package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MasterEipRequestSpecSpec struct {
	Id *string `json:"id,omitempty"`
}

func (o MasterEipRequestSpecSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequestSpecSpec struct{}"
	}

	return strings.Join([]string{"MasterEipRequestSpecSpec", string(data)}, " ")
}

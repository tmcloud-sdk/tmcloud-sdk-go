package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MasterEipRequest struct {
	Spec *MasterEipRequestSpec `json:"spec"`
}

func (o MasterEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipRequest struct{}"
	}

	return strings.Join([]string{"MasterEipRequest", string(data)}, " ")
}

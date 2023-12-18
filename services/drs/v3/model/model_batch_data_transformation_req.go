package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchDataTransformationReq struct {
	Jobs []CheckDataTransformationReq `json:"jobs"`
}

func (o BatchDataTransformationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDataTransformationReq struct{}"
	}

	return strings.Join([]string{"BatchDataTransformationReq", string(data)}, " ")
}

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CheckDataTransformationReq struct {
	JobId *string `json:"job_id,omitempty"`

	ObjectInfo *[]DatabaseObjectVo `json:"object_info,omitempty"`

	TransformationInfo *TransformationInfo `json:"transformation_info"`

	ConfigTransformation *ConfigTransformationVo `json:"config_transformation,omitempty"`
}

func (o CheckDataTransformationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckDataTransformationReq struct{}"
	}

	return strings.Join([]string{"CheckDataTransformationReq", string(data)}, " ")
}

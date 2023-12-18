package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DataTransformationInfo struct {
	TransformationInfo *TransformationInfo `json:"transformation_info,omitempty"`

	ConfigTransformation *ConfigTransformationVo `json:"config_transformation,omitempty"`

	DataTransformationObjectInfos *[]DataTransformationObjectVo `json:"data_transformation_object_infos,omitempty"`
}

func (o DataTransformationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataTransformationInfo struct{}"
	}

	return strings.Join([]string{"DataTransformationInfo", string(data)}, " ")
}

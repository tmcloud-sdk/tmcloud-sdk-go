package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type GlanceUpdateImageMemberRequestBody struct {
	Status GlanceUpdateImageMemberRequestBodyStatus `json:"status"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o GlanceUpdateImageMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberRequestBody struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberRequestBody", string(data)}, " ")
}

type GlanceUpdateImageMemberRequestBodyStatus struct {
	value string
}

type GlanceUpdateImageMemberRequestBodyStatusEnum struct {
	ACCEPTED GlanceUpdateImageMemberRequestBodyStatus
	REJECTED GlanceUpdateImageMemberRequestBodyStatus
}

func GetGlanceUpdateImageMemberRequestBodyStatusEnum() GlanceUpdateImageMemberRequestBodyStatusEnum {
	return GlanceUpdateImageMemberRequestBodyStatusEnum{
		ACCEPTED: GlanceUpdateImageMemberRequestBodyStatus{
			value: "accepted",
		},
		REJECTED: GlanceUpdateImageMemberRequestBodyStatus{
			value: "rejected",
		},
	}
}

func (c GlanceUpdateImageMemberRequestBodyStatus) Value() string {
	return c.value
}

func (c GlanceUpdateImageMemberRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GlanceUpdateImageMemberRequestBodyStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

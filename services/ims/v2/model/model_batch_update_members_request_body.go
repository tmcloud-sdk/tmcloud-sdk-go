package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type BatchUpdateMembersRequestBody struct {
	Images []string `json:"images"`

	ProjectId string `json:"project_id"`

	Status BatchUpdateMembersRequestBodyStatus `json:"status"`

	VaultId *string `json:"vault_id,omitempty"`
}

func (o BatchUpdateMembersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateMembersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateMembersRequestBody", string(data)}, " ")
}

type BatchUpdateMembersRequestBodyStatus struct {
	value string
}

type BatchUpdateMembersRequestBodyStatusEnum struct {
	ACCEPTED BatchUpdateMembersRequestBodyStatus
	REJECTED BatchUpdateMembersRequestBodyStatus
}

func GetBatchUpdateMembersRequestBodyStatusEnum() BatchUpdateMembersRequestBodyStatusEnum {
	return BatchUpdateMembersRequestBodyStatusEnum{
		ACCEPTED: BatchUpdateMembersRequestBodyStatus{
			value: "accepted",
		},
		REJECTED: BatchUpdateMembersRequestBodyStatus{
			value: "rejected",
		},
	}
}

func (c BatchUpdateMembersRequestBodyStatus) Value() string {
	return c.value
}

func (c BatchUpdateMembersRequestBodyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateMembersRequestBodyStatus) UnmarshalJSON(b []byte) error {
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

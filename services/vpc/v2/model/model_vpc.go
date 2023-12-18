package model

import (
	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdktime"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"
	"strings"
)

type Vpc struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Cidr string `json:"cidr"`

	Description string `json:"description"`

	Routes []Route `json:"routes"`

	Status VpcStatus `json:"status"`

	EnterpriseProjectId string `json:"enterprise_project_id"`

	TenantId string `json:"tenant_id"`

	CreatedAt *sdktime.SdkTime `json:"created_at"`

	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Vpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Vpc struct{}"
	}

	return strings.Join([]string{"Vpc", string(data)}, " ")
}

type VpcStatus struct {
	value string
}

type VpcStatusEnum struct {
	CREATING VpcStatus
	OK       VpcStatus
	ERROR    VpcStatus
}

func GetVpcStatusEnum() VpcStatusEnum {
	return VpcStatusEnum{
		CREATING: VpcStatus{
			value: "CREATING",
		},
		OK: VpcStatus{
			value: "OK",
		},
		ERROR: VpcStatus{
			value: "ERROR",
		},
	}
}

func (c VpcStatus) Value() string {
	return c.value
}

func (c VpcStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VpcStatus) UnmarshalJSON(b []byte) error {
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

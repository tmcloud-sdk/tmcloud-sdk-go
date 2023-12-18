package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type PostPaidServerExtendParam struct {
	ChargingMode *int32 `json:"chargingMode,omitempty"`

	RegionID *string `json:"regionID,omitempty"`

	SupportAutoRecovery *bool `json:"support_auto_recovery,omitempty"`

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	MarketType *string `json:"marketType,omitempty"`

	SpotPrice *string `json:"spotPrice,omitempty"`

	DiskPrior *string `json:"diskPrior,omitempty"`

	SpotDurationHours *int32 `json:"spot_duration_hours,omitempty"`

	InterruptionPolicy *PostPaidServerExtendParamInterruptionPolicy `json:"interruption_policy,omitempty"`

	SpotDurationCount *int32 `json:"spot_duration_count,omitempty"`

	CbCsbsBackup *string `json:"CB_CSBS_BACKUP,omitempty"`
}

func (o PostPaidServerExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostPaidServerExtendParam struct{}"
	}

	return strings.Join([]string{"PostPaidServerExtendParam", string(data)}, " ")
}

type PostPaidServerExtendParamInterruptionPolicy struct {
	value string
}

type PostPaidServerExtendParamInterruptionPolicyEnum struct {
	IMMEDIATE PostPaidServerExtendParamInterruptionPolicy
}

func GetPostPaidServerExtendParamInterruptionPolicyEnum() PostPaidServerExtendParamInterruptionPolicyEnum {
	return PostPaidServerExtendParamInterruptionPolicyEnum{
		IMMEDIATE: PostPaidServerExtendParamInterruptionPolicy{
			value: "immediate",
		},
	}
}

func (c PostPaidServerExtendParamInterruptionPolicy) Value() string {
	return c.value
}

func (c PostPaidServerExtendParamInterruptionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PostPaidServerExtendParamInterruptionPolicy) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type AddonInstanceStatus struct {
	Status AddonInstanceStatusStatus `json:"status"`

	Reason string `json:"Reason"`

	Message string `json:"message"`

	TargetVersions *[]string `json:"targetVersions,omitempty"`

	CurrentVersion *Versions `json:"currentVersion"`

	IsRollbackable *bool `json:"isRollbackable,omitempty"`

	PreviousVersion *string `json:"previousVersion,omitempty"`
}

func (o AddonInstanceStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstanceStatus struct{}"
	}

	return strings.Join([]string{"AddonInstanceStatus", string(data)}, " ")
}

type AddonInstanceStatusStatus struct {
	value string
}

type AddonInstanceStatusStatusEnum struct {
	RUNNING         AddonInstanceStatusStatus
	ABNORMAL        AddonInstanceStatusStatus
	INSTALLING      AddonInstanceStatusStatus
	INSTALL_FAILED  AddonInstanceStatusStatus
	UPGRADING       AddonInstanceStatusStatus
	UPGRADE_FAILED  AddonInstanceStatusStatus
	DELETING        AddonInstanceStatusStatus
	DELETE_SUCCESS  AddonInstanceStatusStatus
	DELETE_FAILED   AddonInstanceStatusStatus
	AVAILABLE       AddonInstanceStatusStatus
	ROLLBACKING     AddonInstanceStatusStatus
	ROLLBACK_FAILED AddonInstanceStatusStatus
}

func GetAddonInstanceStatusStatusEnum() AddonInstanceStatusStatusEnum {
	return AddonInstanceStatusStatusEnum{
		RUNNING: AddonInstanceStatusStatus{
			value: "running",
		},
		ABNORMAL: AddonInstanceStatusStatus{
			value: "abnormal",
		},
		INSTALLING: AddonInstanceStatusStatus{
			value: "installing",
		},
		INSTALL_FAILED: AddonInstanceStatusStatus{
			value: "installFailed",
		},
		UPGRADING: AddonInstanceStatusStatus{
			value: "upgrading",
		},
		UPGRADE_FAILED: AddonInstanceStatusStatus{
			value: "upgradeFailed",
		},
		DELETING: AddonInstanceStatusStatus{
			value: "deleting",
		},
		DELETE_SUCCESS: AddonInstanceStatusStatus{
			value: "deleteSuccess",
		},
		DELETE_FAILED: AddonInstanceStatusStatus{
			value: "deleteFailed",
		},
		AVAILABLE: AddonInstanceStatusStatus{
			value: "available",
		},
		ROLLBACKING: AddonInstanceStatusStatus{
			value: "rollbacking",
		},
		ROLLBACK_FAILED: AddonInstanceStatusStatus{
			value: "rollbackFailed",
		},
	}
}

func (c AddonInstanceStatusStatus) Value() string {
	return c.value
}

func (c AddonInstanceStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddonInstanceStatusStatus) UnmarshalJSON(b []byte) error {
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

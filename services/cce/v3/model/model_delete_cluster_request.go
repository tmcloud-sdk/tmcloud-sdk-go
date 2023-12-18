package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type DeleteClusterRequest struct {
	ClusterId string `json:"cluster_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	DeleteEfs *DeleteClusterRequestDeleteEfs `json:"delete_efs,omitempty"`

	DeleteEni *DeleteClusterRequestDeleteEni `json:"delete_eni,omitempty"`

	DeleteEvs *DeleteClusterRequestDeleteEvs `json:"delete_evs,omitempty"`

	DeleteNet *DeleteClusterRequestDeleteNet `json:"delete_net,omitempty"`

	DeleteObs *DeleteClusterRequestDeleteObs `json:"delete_obs,omitempty"`

	DeleteSfs *DeleteClusterRequestDeleteSfs `json:"delete_sfs,omitempty"`

	DeleteSfs30 *DeleteClusterRequestDeleteSfs30 `json:"delete_sfs30,omitempty"`

	Tobedeleted *DeleteClusterRequestTobedeleted `json:"tobedeleted,omitempty"`
}

func (o DeleteClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteClusterRequest", string(data)}, " ")
}

type DeleteClusterRequestDeleteEfs struct {
	value string
}

type DeleteClusterRequestDeleteEfsEnum struct {
	TRUE  DeleteClusterRequestDeleteEfs
	BLOCK DeleteClusterRequestDeleteEfs
	TRY   DeleteClusterRequestDeleteEfs
	FALSE DeleteClusterRequestDeleteEfs
	SKIP  DeleteClusterRequestDeleteEfs
}

func GetDeleteClusterRequestDeleteEfsEnum() DeleteClusterRequestDeleteEfsEnum {
	return DeleteClusterRequestDeleteEfsEnum{
		TRUE: DeleteClusterRequestDeleteEfs{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteEfs{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteEfs{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteEfs{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteEfs{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteEfs) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteEfs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteEfs) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteEni struct {
	value string
}

type DeleteClusterRequestDeleteEniEnum struct {
	TRUE  DeleteClusterRequestDeleteEni
	BLOCK DeleteClusterRequestDeleteEni
	TRY   DeleteClusterRequestDeleteEni
	FALSE DeleteClusterRequestDeleteEni
	SKIP  DeleteClusterRequestDeleteEni
}

func GetDeleteClusterRequestDeleteEniEnum() DeleteClusterRequestDeleteEniEnum {
	return DeleteClusterRequestDeleteEniEnum{
		TRUE: DeleteClusterRequestDeleteEni{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteEni{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteEni{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteEni{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteEni{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteEni) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteEni) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteEni) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteEvs struct {
	value string
}

type DeleteClusterRequestDeleteEvsEnum struct {
	TRUE  DeleteClusterRequestDeleteEvs
	BLOCK DeleteClusterRequestDeleteEvs
	TRY   DeleteClusterRequestDeleteEvs
	FALSE DeleteClusterRequestDeleteEvs
	SKIP  DeleteClusterRequestDeleteEvs
}

func GetDeleteClusterRequestDeleteEvsEnum() DeleteClusterRequestDeleteEvsEnum {
	return DeleteClusterRequestDeleteEvsEnum{
		TRUE: DeleteClusterRequestDeleteEvs{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteEvs{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteEvs{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteEvs{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteEvs{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteEvs) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteEvs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteEvs) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteNet struct {
	value string
}

type DeleteClusterRequestDeleteNetEnum struct {
	TRUE  DeleteClusterRequestDeleteNet
	BLOCK DeleteClusterRequestDeleteNet
	TRY   DeleteClusterRequestDeleteNet
	FALSE DeleteClusterRequestDeleteNet
	SKIP  DeleteClusterRequestDeleteNet
}

func GetDeleteClusterRequestDeleteNetEnum() DeleteClusterRequestDeleteNetEnum {
	return DeleteClusterRequestDeleteNetEnum{
		TRUE: DeleteClusterRequestDeleteNet{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteNet{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteNet{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteNet{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteNet{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteNet) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteNet) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteNet) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteObs struct {
	value string
}

type DeleteClusterRequestDeleteObsEnum struct {
	TRUE  DeleteClusterRequestDeleteObs
	BLOCK DeleteClusterRequestDeleteObs
	TRY   DeleteClusterRequestDeleteObs
	FALSE DeleteClusterRequestDeleteObs
	SKIP  DeleteClusterRequestDeleteObs
}

func GetDeleteClusterRequestDeleteObsEnum() DeleteClusterRequestDeleteObsEnum {
	return DeleteClusterRequestDeleteObsEnum{
		TRUE: DeleteClusterRequestDeleteObs{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteObs{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteObs{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteObs{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteObs{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteObs) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteObs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteObs) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteSfs struct {
	value string
}

type DeleteClusterRequestDeleteSfsEnum struct {
	TRUE  DeleteClusterRequestDeleteSfs
	BLOCK DeleteClusterRequestDeleteSfs
	TRY   DeleteClusterRequestDeleteSfs
	FALSE DeleteClusterRequestDeleteSfs
	SKIP  DeleteClusterRequestDeleteSfs
}

func GetDeleteClusterRequestDeleteSfsEnum() DeleteClusterRequestDeleteSfsEnum {
	return DeleteClusterRequestDeleteSfsEnum{
		TRUE: DeleteClusterRequestDeleteSfs{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteSfs{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteSfs{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteSfs{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteSfs{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteSfs) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteSfs) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteSfs) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestDeleteSfs30 struct {
	value string
}

type DeleteClusterRequestDeleteSfs30Enum struct {
	TRUE  DeleteClusterRequestDeleteSfs30
	BLOCK DeleteClusterRequestDeleteSfs30
	TRY   DeleteClusterRequestDeleteSfs30
	FALSE DeleteClusterRequestDeleteSfs30
	SKIP  DeleteClusterRequestDeleteSfs30
}

func GetDeleteClusterRequestDeleteSfs30Enum() DeleteClusterRequestDeleteSfs30Enum {
	return DeleteClusterRequestDeleteSfs30Enum{
		TRUE: DeleteClusterRequestDeleteSfs30{
			value: "true",
		},
		BLOCK: DeleteClusterRequestDeleteSfs30{
			value: "block",
		},
		TRY: DeleteClusterRequestDeleteSfs30{
			value: "try",
		},
		FALSE: DeleteClusterRequestDeleteSfs30{
			value: "false",
		},
		SKIP: DeleteClusterRequestDeleteSfs30{
			value: "skip",
		},
	}
}

func (c DeleteClusterRequestDeleteSfs30) Value() string {
	return c.value
}

func (c DeleteClusterRequestDeleteSfs30) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestDeleteSfs30) UnmarshalJSON(b []byte) error {
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

type DeleteClusterRequestTobedeleted struct {
	value string
}

type DeleteClusterRequestTobedeletedEnum struct {
	TRUE DeleteClusterRequestTobedeleted
}

func GetDeleteClusterRequestTobedeletedEnum() DeleteClusterRequestTobedeletedEnum {
	return DeleteClusterRequestTobedeletedEnum{
		TRUE: DeleteClusterRequestTobedeleted{
			value: "true",
		},
	}
}

func (c DeleteClusterRequestTobedeleted) Value() string {
	return c.value
}

func (c DeleteClusterRequestTobedeleted) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteClusterRequestTobedeleted) UnmarshalJSON(b []byte) error {
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

/*
 * HCS API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package hcsschema

type Chipset struct {
	Uefi *Uefi `json:"Uefi,omitempty"`

	IsNumLockDisabled bool `json:"IsNumLockDisabled,omitempty"`

	BaseBoardSerialNumber string `json:"BaseBoardSerialNumber,omitempty"`

	ChassisSerialNumber string `json:"ChassisSerialNumber,omitempty"`

	ChassisAssetTag string `json:"ChassisAssetTag,omitempty"`

	UseUtc bool `json:"UseUtc,omitempty"`

	// LinuxKernelDirect - Added in v2.2 Builds >=181117
	LinuxKernelDirect *LinuxKernelDirect `json:"LinuxKernelDirect,omitempty"`

	FirmwareFile *FirmwareFile `json:"FirmwareFile,omitempty"`
}

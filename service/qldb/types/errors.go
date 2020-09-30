// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// One or more parameters in the request aren't valid.
type InvalidParameterException struct {
	Message *string

	ParameterName *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}
func (e *InvalidParameterException) GetParameterName() string {
	return ptr.ToString(e.ParameterName)
}
func (e *InvalidParameterException) HasParameterName() bool {
	return e.ParameterName != nil
}

// You have reached the limit on the maximum number of resources allowed.
type LimitExceededException struct {
	Message *string

	ResourceType *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *LimitExceededException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string

	ResourceName *string
	ResourceType *string
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceAlreadyExistsException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *ResourceAlreadyExistsException) HasResourceName() bool {
	return e.ResourceName != nil
}
func (e *ResourceAlreadyExistsException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourceAlreadyExistsException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *ResourceAlreadyExistsException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceAlreadyExistsException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource can't be modified at this time.
type ResourceInUseException struct {
	Message *string

	ResourceType *string
	ResourceName *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceInUseException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourceInUseException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *ResourceInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceInUseException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceInUseException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *ResourceInUseException) HasResourceName() bool {
	return e.ResourceName != nil
}

// The specified resource doesn't exist.
type ResourceNotFoundException struct {
	Message *string

	ResourceType *string
	ResourceName *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourceNotFoundException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *ResourceNotFoundException) HasResourceName() bool {
	return e.ResourceName != nil
}

// The operation failed because a condition wasn't satisfied in advance.
type ResourcePreconditionNotMetException struct {
	Message *string

	ResourceName *string
	ResourceType *string
}

func (e *ResourcePreconditionNotMetException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourcePreconditionNotMetException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourcePreconditionNotMetException) ErrorCode() string {
	return "ResourcePreconditionNotMetException"
}
func (e *ResourcePreconditionNotMetException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *ResourcePreconditionNotMetException) GetResourceName() string {
	return ptr.ToString(e.ResourceName)
}
func (e *ResourcePreconditionNotMetException) HasResourceName() bool {
	return e.ResourceName != nil
}
func (e *ResourcePreconditionNotMetException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *ResourcePreconditionNotMetException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *ResourcePreconditionNotMetException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourcePreconditionNotMetException) HasMessage() bool {
	return e.Message != nil
}
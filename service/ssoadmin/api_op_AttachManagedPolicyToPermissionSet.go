// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches an IAM managed policy ARN to a permission set. If the permission set is
// already referenced by one or more account assignments, you will need to call
// ProvisionPermissionSet after this action to apply the corresponding IAM policy
// updates to all assigned accounts.
func (c *Client) AttachManagedPolicyToPermissionSet(ctx context.Context, params *AttachManagedPolicyToPermissionSetInput, optFns ...func(*Options)) (*AttachManagedPolicyToPermissionSetOutput, error) {
	if params == nil {
		params = &AttachManagedPolicyToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachManagedPolicyToPermissionSet", params, optFns, addOperationAttachManagedPolicyToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachManagedPolicyToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachManagedPolicyToPermissionSetInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The IAM managed policy ARN to be attached to a permission set.
	//
	// This member is required.
	ManagedPolicyArn *string

	// The ARN of the PermissionSet that the managed policy should be attached to.
	//
	// This member is required.
	PermissionSetArn *string
}

type AttachManagedPolicyToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachManagedPolicyToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachManagedPolicyToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachManagedPolicyToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpAttachManagedPolicyToPermissionSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachManagedPolicyToPermissionSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAttachManagedPolicyToPermissionSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "AttachManagedPolicyToPermissionSet",
	}
}
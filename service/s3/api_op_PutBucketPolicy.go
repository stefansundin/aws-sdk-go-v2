// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Applies an Amazon S3 bucket policy to an Amazon S3 bucket. If you are using an
// identity other than the root user of the AWS account that owns the bucket, the
// calling identity must have the PutBucketPolicy permissions on the specified
// bucket and belong to the bucket owner's account in order to use this operation.
// <p>If you don't have <code>PutBucketPolicy</code> permissions, Amazon S3 returns
// a <code>403 Access Denied</code> error. If you have the correct permissions, but
// you're not using an identity that belongs to the bucket owner's account, Amazon
// S3 returns a <code>405 Method Not Allowed</code> error.</p> <important> <p> As a
// security precaution, the root user of the AWS account that owns a bucket can
// always use this operation, even if the policy explicitly denies the root user
// the ability to perform this action. </p> </important> <p>For more information
// about bucket policies, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html">Using
// Bucket Policies and User Policies</a>.</p> <p>The following operations are
// related to <code>PutBucketPolicy</code>:</p> <ul> <li> <p> <a>CreateBucket</a>
// </p> </li> <li> <p> <a>DeleteBucket</a> </p> </li> </ul>
func (c *Client) PutBucketPolicy(ctx context.Context, params *PutBucketPolicyInput, optFns ...func(*Options)) (*PutBucketPolicyOutput, error) {
	stack := middleware.NewStack("PutBucketPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketPolicyMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpPutBucketPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketPolicy(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	addMetadataRetrieverMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "PutBucketPolicy",
			Err:           err,
		}
	}
	out := result.(*PutBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketPolicyInput struct {
	// The name of the bucket.
	Bucket *string
	// The bucket policy as a JSON document.
	Policy *string
	// Set this parameter to true to confirm that you want to remove your permissions
	// to change this bucket policy in the future.
	ConfirmRemoveSelfBucketAccess *bool
	// The MD5 hash of the request body.
	ContentMD5 *string
}

type PutBucketPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketPolicy",
	}
}

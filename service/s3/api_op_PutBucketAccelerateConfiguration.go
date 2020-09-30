// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the accelerate configuration of an existing bucket. Amazon S3 Transfer
// Acceleration is a bucket-level feature that enables you to perform faster data
// transfers to Amazon S3.  <p> To use this operation, you must have permission to
// perform the s3:PutAccelerateConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p> The Transfer
// Acceleration state of a bucket can be set to one of the following two
// values:</p> <ul> <li> <p> Enabled – Enables accelerated data transfers to the
// bucket.</p> </li> <li> <p> Suspended – Disables accelerated data transfers to
// the bucket.</p> </li> </ul> <p>The <a>GetBucketAccelerateConfiguration</a>
// operation returns the transfer acceleration state of a bucket.</p> <p>After
// setting the Transfer Acceleration state of a bucket to Enabled, it might take up
// to thirty minutes before the data transfer rates to the bucket increase.</p> <p>
// The name of the bucket used for Transfer Acceleration must be DNS-compliant and
// must not contain periods (".").</p> <p> For more information about transfer
// acceleration, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html">Transfer
// Acceleration</a>.</p> <p>The following operations are related to
// <code>PutBucketAccelerateConfiguration</code>:</p> <ul> <li> <p>
// <a>GetBucketAccelerateConfiguration</a> </p> </li> <li> <p> <a>CreateBucket</a>
// </p> </li> </ul>
func (c *Client) PutBucketAccelerateConfiguration(ctx context.Context, params *PutBucketAccelerateConfigurationInput, optFns ...func(*Options)) (*PutBucketAccelerateConfigurationOutput, error) {
	stack := middleware.NewStack("PutBucketAccelerateConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketAccelerateConfigurationMiddlewares(stack)
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
	addOpPutBucketAccelerateConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAccelerateConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutBucketAccelerateConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutBucketAccelerateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAccelerateConfigurationInput struct {
	// Name of the bucket for which the accelerate configuration is set.
	Bucket *string
	// Container for setting the transfer acceleration state.
	AccelerateConfiguration *types.AccelerateConfiguration
}

type PutBucketAccelerateConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketAccelerateConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAccelerateConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAccelerateConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketAccelerateConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAccelerateConfiguration",
	}
}

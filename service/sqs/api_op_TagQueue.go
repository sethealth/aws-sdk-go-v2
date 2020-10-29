// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Add cost allocation tags to the specified Amazon SQS queue. For an overview, see
// Tagging Your Amazon SQS Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide. When you use queue tags,
// keep the following guidelines in mind:
//
// * Adding more than 50 tags to a queue
// isn't recommended.
//
// * Tags don't have any semantic meaning. Amazon SQS
// interprets tags as character strings.
//
// * Tags are case-sensitive.
//
// * A new tag
// with a key identical to that of an existing tag overwrites the existing
// tag.
//
// For a full list of tag restrictions, see Limits Related to Queues
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues)
// in the Amazon Simple Queue Service Developer Guide. Cross-account permissions
// don't apply to this action. For more information, see Grant Cross-Account
// Permissions to a Role and a User Name
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
func (c *Client) TagQueue(ctx context.Context, params *TagQueueInput, optFns ...func(*Options)) (*TagQueueOutput, error) {
	if params == nil {
		params = &TagQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagQueue", params, optFns, addOperationTagQueueMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagQueueInput struct {

	// The URL of the queue.
	//
	// This member is required.
	QueueUrl *string

	// The list of tags to be added to the specified queue.
	//
	// This member is required.
	Tags map[string]*string
}

type TagQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagQueueMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagQueue{}, middleware.After)
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
	addOpTagQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagQueue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTagQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "TagQueue",
	}
}

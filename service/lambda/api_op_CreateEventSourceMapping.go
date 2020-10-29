// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a mapping between an event source and an AWS Lambda function. Lambda
// reads items from the event source and triggers the function. For details about
// each event source type, see the following topics.
//
// * Using AWS Lambda with
// Amazon DynamoDB (https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html)
//
// *
// Using AWS Lambda with Amazon Kinesis
// (https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html)
//
// * Using AWS
// Lambda with Amazon SQS
// (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html)
//
// * Using AWS Lambda
// with Amazon MSK
// (https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html)
//
// The following
// error handling options are only available for stream sources (DynamoDB and
// Kinesis):
//
// * BisectBatchOnFunctionError - If the function returns an error,
// split the batch in two and retry.
//
// * DestinationConfig - Send discarded records
// to an Amazon SQS queue or Amazon SNS topic.
//
// * MaximumRecordAgeInSeconds -
// Discard records older than the specified age. Default -1 (infinite). Minimum 60.
// Maximum 604800.
//
// * MaximumRetryAttempts - Discard records after the specified
// number of retries. Default -1 (infinite). Minimum 0. Maximum 10000. When
// infinite, failed records will be retried until the record expires.
//
// *
// ParallelizationFactor - Process multiple batches from each shard concurrently.
func (c *Client) CreateEventSourceMapping(ctx context.Context, params *CreateEventSourceMappingInput, optFns ...func(*Options)) (*CreateEventSourceMappingOutput, error) {
	if params == nil {
		params = &CreateEventSourceMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventSourceMapping", params, optFns, addOperationCreateEventSourceMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventSourceMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventSourceMappingInput struct {

	// The Amazon Resource Name (ARN) of the event source.
	//
	// * Amazon Kinesis - The ARN
	// of the data stream or a stream consumer.
	//
	// * Amazon DynamoDB Streams - The ARN of
	// the stream.
	//
	// * Amazon Simple Queue Service - The ARN of the queue.
	//
	// * Amazon
	// Managed Streaming for Apache Kafka - The ARN of the cluster.
	//
	// This member is required.
	EventSourceArn *string

	// The name of the Lambda function. Name formats
	//
	// * Function name - MyFunction.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	// *
	// Version or Alias ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.
	//
	// * Partial ARN -
	// 123456789012:function:MyFunction.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it's limited to 64 characters
	// in length.
	//
	// This member is required.
	FunctionName *string

	// The maximum number of items to retrieve in a single batch.
	//
	// * Amazon Kinesis -
	// Default 100. Max 10,000.
	//
	// * Amazon DynamoDB Streams - Default 100. Max 1,000.
	//
	// *
	// Amazon Simple Queue Service - Default 10. Max 10.
	//
	// * Amazon Managed Streaming
	// for Apache Kafka - Default 100. Max 10,000.
	BatchSize *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

	// If true, the event source mapping is active. Set to false to pause polling and
	// invocation.
	Enabled *bool

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int32

	// (Streams) Discard records older than the specified age. The default value is
	// infinite (-1).
	MaximumRecordAgeInSeconds *int32

	// (Streams) Discard records after the specified number of retries. The default
	// value is infinite (-1). When set to infinite (-1), failed records will be
	// retried until the record expires.
	MaximumRetryAttempts *int32

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int32

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is only
	// supported for Amazon Kinesis streams.
	StartingPosition types.EventSourcePosition

	// With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.
	StartingPositionTimestamp *time.Time

	// (MSK) The name of the Kafka topic.
	Topics []*string
}

// A mapping between an AWS resource and an AWS Lambda function. See
// CreateEventSourceMapping for details.
type CreateEventSourceMappingOutput struct {

	// The maximum number of items to retrieve in a single batch.
	BatchSize *int32

	// (Streams) If the function returns an error, split the batch in two and retry.
	BisectBatchOnFunctionError *bool

	// (Streams) An Amazon SQS queue or Amazon SNS topic destination for discarded
	// records.
	DestinationConfig *types.DestinationConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// The ARN of the Lambda function.
	FunctionArn *string

	// The date that the event source mapping was last updated, or its state changed.
	LastModified *time.Time

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string

	// (Streams) The maximum amount of time to gather records before invoking the
	// function, in seconds.
	MaximumBatchingWindowInSeconds *int32

	// (Streams) The maximum age of a record that Lambda sends to a function for
	// processing.
	MaximumRecordAgeInSeconds *int32

	// (Streams) The maximum number of times to retry when the function returns an
	// error.
	MaximumRetryAttempts *int32

	// (Streams) The number of batches to process from each shard concurrently.
	ParallelizationFactor *int32

	// The state of the event source mapping. It can be one of the following: Creating,
	// Enabling, Enabled, Disabling, Disabled, Updating, or Deleting.
	State *string

	// Indicates whether the last change to the event source mapping was made by a
	// user, or by the Lambda service.
	StateTransitionReason *string

	// (MSK) The name of the Kafka topic.
	Topics []*string

	// The identifier of the event source mapping.
	UUID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEventSourceMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEventSourceMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEventSourceMapping{}, middleware.After)
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
	addOpCreateEventSourceMappingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventSourceMapping(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEventSourceMapping(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "CreateEventSourceMapping",
	}
}

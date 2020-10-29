// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Kinesis data stream. A stream captures and transports data records
// that are continuously emitted from different data sources or producers.
// Scale-out within a stream is explicitly supported by means of shards, which are
// uniquely identified groups of data records in a stream. You specify and control
// the number of shards that a stream is composed of. Each shard can support reads
// up to five transactions per second, up to a maximum data read total of 2 MiB per
// second. Each shard can support writes up to 1,000 records per second, up to a
// maximum data write total of 1 MiB per second. If the amount of data input
// increases or decreases, you can add or remove shards. The stream name identifies
// the stream. The name is scoped to the AWS account used by the application. It is
// also scoped by AWS Region. That is, two streams in two different accounts can
// have the same name, and two streams in the same account, but in two different
// Regions, can have the same name. CreateStream is an asynchronous operation. Upon
// receiving a CreateStream request, Kinesis Data Streams immediately returns and
// sets the stream status to CREATING. After the stream is created, Kinesis Data
// Streams sets the stream status to ACTIVE. You should perform read and write
// operations only on an ACTIVE stream. You receive a LimitExceededException when
// making a CreateStream request when you try to do one of the following:
//
// * Have
// more than five streams in the CREATING state at any point in time.
//
// * Create
// more shards than are authorized for your account.
//
// For the default shard limit
// for an AWS account, see Amazon Kinesis Data Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide. To increase this limit,
// contact AWS Support
// (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html). You can
// use DescribeStream to check the stream status, which is returned in
// StreamStatus. CreateStream has a limit of five transactions per second per
// account.
func (c *Client) CreateStream(ctx context.Context, params *CreateStreamInput, optFns ...func(*Options)) (*CreateStreamOutput, error) {
	if params == nil {
		params = &CreateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStream", params, optFns, addOperationCreateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for CreateStream.
type CreateStreamInput struct {

	// The number of shards that the stream will use. The throughput of the stream is a
	// function of the number of shards; more shards are required for greater
	// provisioned throughput.
	//
	// This member is required.
	ShardCount *int32

	// A name to identify the stream. The stream name is scoped to the AWS account used
	// by the application that creates the stream. It is also scoped by AWS Region.
	// That is, two streams in two different AWS accounts can have the same name. Two
	// streams in the same AWS account but in two different Regions can also have the
	// same name.
	//
	// This member is required.
	StreamName *string
}

type CreateStreamOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateStream{}, middleware.After)
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
	addOpCreateStreamValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStream(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateStream(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "CreateStream",
	}
}

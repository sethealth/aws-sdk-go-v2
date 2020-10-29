// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Schedules delivery of a configuration snapshot to the Amazon S3 bucket in the
// specified delivery channel. After the delivery has started, AWS Config sends the
// following notifications using an Amazon SNS topic that you have specified.
//
// *
// Notification of the start of the delivery.
//
// * Notification of the completion of
// the delivery, if the delivery was successfully completed.
//
// * Notification of
// delivery failure, if the delivery failed.
func (c *Client) DeliverConfigSnapshot(ctx context.Context, params *DeliverConfigSnapshotInput, optFns ...func(*Options)) (*DeliverConfigSnapshotOutput, error) {
	if params == nil {
		params = &DeliverConfigSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeliverConfigSnapshot", params, optFns, addOperationDeliverConfigSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeliverConfigSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeliverConfigSnapshot action.
type DeliverConfigSnapshotInput struct {

	// The name of the delivery channel through which the snapshot is delivered.
	//
	// This member is required.
	DeliveryChannelName *string
}

// The output for the DeliverConfigSnapshot action, in JSON format.
type DeliverConfigSnapshotOutput struct {

	// The ID of the snapshot that is being created.
	ConfigSnapshotId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeliverConfigSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeliverConfigSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeliverConfigSnapshot{}, middleware.After)
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
	addOpDeliverConfigSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeliverConfigSnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeliverConfigSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeliverConfigSnapshot",
	}
}

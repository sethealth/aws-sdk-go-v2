// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an event tracker that you use when sending event data to the specified
// dataset group using the PutEvents
// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.
// When Amazon Personalize creates an event tracker, it also creates an
// event-interactions dataset in the dataset group associated with the event
// tracker. The event-interactions dataset stores the event data from the PutEvents
// call. The contents of this dataset are not available to the user. Only one event
// tracker can be associated with a dataset group. You will get an error if you
// call CreateEventTracker using the same dataset group as an existing event
// tracker. When you send event data you include your tracking ID. The tracking ID
// identifies the customer and authorizes the customer to send the data. The event
// tracker can be in one of the following states:
//
// * CREATE PENDING > CREATE
// IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
// * DELETE PENDING > DELETE
// IN_PROGRESS
//
// To get the status of the event tracker, call DescribeEventTracker.
// The event tracker must be in the ACTIVE state before using the tracking ID.
// Related APIs
//
// * ListEventTrackers
//
// * DescribeEventTracker
//
// * DeleteEventTracker
func (c *Client) CreateEventTracker(ctx context.Context, params *CreateEventTrackerInput, optFns ...func(*Options)) (*CreateEventTrackerOutput, error) {
	if params == nil {
		params = &CreateEventTrackerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventTracker", params, optFns, addOperationCreateEventTrackerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventTrackerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventTrackerInput struct {

	// The Amazon Resource Name (ARN) of the dataset group that receives the event
	// data.
	//
	// This member is required.
	DatasetGroupArn *string

	// The name for the event tracker.
	//
	// This member is required.
	Name *string
}

type CreateEventTrackerOutput struct {

	// The ARN of the event tracker.
	EventTrackerArn *string

	// The ID of the event tracker. Include this ID in requests to the PutEvents
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.
	TrackingId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateEventTrackerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEventTracker{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEventTracker{}, middleware.After)
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
	addOpCreateEventTrackerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventTracker(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateEventTracker(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateEventTracker",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a contact flow to start a new chat for the customer. Response of this
// API provides a token required to obtain credentials from the
// CreateParticipantConnection
// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// API in the Amazon Connect Participant Service. When a new chat contact is
// successfully created, clients must subscribe to the participant’s connection for
// the created chat within 5 minutes. This is achieved by invoking
// CreateParticipantConnection
// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html)
// with WEBSOCKET and CONNECTION_CREDENTIALS. A 429 error occurs in two
// situations:
//
// * API rate limit is exceeded. API TPS throttling returns a
// TooManyRequests exception from the API Gateway.
//
// * The quota for concurrent
// active chats
// (https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html)
// is exceeded. Active chat throttling returns a LimitExceededException.
//
// For more
// information about chat, see Chat
// (https://docs.aws.amazon.com/connect/latest/adminguide/chat.html) in the Amazon
// Connect Administrator Guide.
func (c *Client) StartChatContact(ctx context.Context, params *StartChatContactInput, optFns ...func(*Options)) (*StartChatContactOutput, error) {
	if params == nil {
		params = &StartChatContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartChatContact", params, optFns, addOperationStartChatContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartChatContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartChatContactInput struct {

	// The identifier of the contact flow for initiating the chat. To see the
	// ContactFlowId in the Amazon Connect console user interface, on the navigation
	// menu go to Routing, Contact Flows. Choose the contact flow. On the contact flow
	// page, under the name of the contact flow, choose Show additional flow
	// information. The ContactFlowId is the last part of the ARN, shown here in bold:
	// arn:aws:connect:us-west-2:xxxxxxxxxxxx:instance/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/contact-flow/846ec553-a005-41c0-8341-xxxxxxxxxxxx
	//
	// This member is required.
	ContactFlowId *string

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// Information identifying the participant.
	//
	// This member is required.
	ParticipantDetails *types.ParticipantDetails

	// A custom key-value pair using an attribute map. The attributes are standard
	// Amazon Connect attributes. They can be accessed in contact flows just like any
	// other contact attributes. There can be up to 32,768 UTF-8 bytes across all
	// key-value pairs per contact. Attribute keys can include only alphanumeric, dash,
	// and underscore characters.
	Attributes map[string]string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// The initial message to be sent to the newly created chat.
	InitialMessage *types.ChatMessage
}

type StartChatContactOutput struct {

	// The identifier of this contact within the Amazon Connect instance.
	ContactId *string

	// The identifier for a chat participant. The participantId for a chat participant
	// is the same throughout the chat lifecycle.
	ParticipantId *string

	// The token used by the chat participant to call CreateParticipantConnection
	// (https://docs.aws.amazon.com/connect-participant/latest/APIReference/API_CreateParticipantConnection.html).
	// The participant token is valid for the lifetime of a chat participant.
	ParticipantToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartChatContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartChatContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartChatContact{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opStartChatContactMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartChatContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartChatContact(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpStartChatContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartChatContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartChatContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartChatContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartChatContactInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartChatContactMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartChatContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartChatContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "StartChatContact",
	}
}

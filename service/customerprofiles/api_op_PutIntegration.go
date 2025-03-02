// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Adds an integration between the service and a third-party service, which
// includes Amazon AppFlow and Amazon Connect. An integration can belong to only
// one domain.
func (c *Client) PutIntegration(ctx context.Context, params *PutIntegrationInput, optFns ...func(*Options)) (*PutIntegrationOutput, error) {
	if params == nil {
		params = &PutIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutIntegration", params, optFns, addOperationPutIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutIntegrationInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	FlowDefinition *types.FlowDefinition

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// The URI of the S3 bucket or any other type of data source.
	Uri *string
}

type PutIntegrationOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The name of the profile object type.
	//
	// This member is required.
	ObjectTypeName *string

	// The URI of the S3 bucket or any other type of data source.
	//
	// This member is required.
	Uri *string

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutIntegration{}, middleware.After)
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
	if err = addOpPutIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutIntegration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "PutIntegration",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the provisioned concurrency configuration for a function's alias or
// version.
func (c *Client) GetProvisionedConcurrencyConfig(ctx context.Context, params *GetProvisionedConcurrencyConfigInput, optFns ...func(*Options)) (*GetProvisionedConcurrencyConfigOutput, error) {
	if params == nil {
		params = &GetProvisionedConcurrencyConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProvisionedConcurrencyConfig", params, optFns, addOperationGetProvisionedConcurrencyConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProvisionedConcurrencyConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProvisionedConcurrencyConfigInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - my-function.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// The version number or alias name.
	//
	// This member is required.
	Qualifier *string
}

type GetProvisionedConcurrencyConfigOutput struct {

	// The amount of provisioned concurrency allocated.
	AllocatedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// The date and time that a user last updated the configuration, in ISO 8601 format
	// (https://www.iso.org/iso-8601-date-and-time-format.html).
	LastModified *string

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The status of the allocation process.
	Status types.ProvisionedConcurrencyStatusEnum

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetProvisionedConcurrencyConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProvisionedConcurrencyConfig{}, middleware.After)
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
	addOpGetProvisionedConcurrencyConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetProvisionedConcurrencyConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetProvisionedConcurrencyConfig",
	}
}

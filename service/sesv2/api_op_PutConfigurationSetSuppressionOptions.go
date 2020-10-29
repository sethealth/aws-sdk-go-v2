// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specify the account suppression list preferences for a configuration set.
func (c *Client) PutConfigurationSetSuppressionOptions(ctx context.Context, params *PutConfigurationSetSuppressionOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetSuppressionOptionsOutput, error) {
	if params == nil {
		params = &PutConfigurationSetSuppressionOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationSetSuppressionOptions", params, optFns, addOperationPutConfigurationSetSuppressionOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationSetSuppressionOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to change the account suppression list preferences for a specific
// configuration set.
type PutConfigurationSetSuppressionOptionsInput struct {

	// The name of the configuration set that you want to change the suppression list
	// preferences for.
	//
	// This member is required.
	ConfigurationSetName *string

	// A list that contains the reasons that email addresses are automatically added to
	// the suppression list for your account. This list can contain any or all of the
	// following:
	//
	// * COMPLAINT – Amazon SES adds an email address to the suppression
	// list for your account when a message sent to that address results in a
	// complaint.
	//
	// * BOUNCE – Amazon SES adds an email address to the suppression list
	// for your account when a message sent to that address results in a hard bounce.
	SuppressedReasons []types.SuppressionListReason
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetSuppressionOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigurationSetSuppressionOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetSuppressionOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetSuppressionOptions{}, middleware.After)
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
	addOpPutConfigurationSetSuppressionOptionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetSuppressionOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutConfigurationSetSuppressionOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetSuppressionOptions",
	}
}

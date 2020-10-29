// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the resource-based permission policy attached to the secret. Minimum
// permissions To run this command, you must have the following permissions:
//
// *
// secretsmanager:DeleteResourcePolicy
//
// Related operations
//
// * To attach a resource
// policy to a secret, use PutResourcePolicy.
//
// * To retrieve the current
// resource-based policy that's attached to a secret, use GetResourcePolicy.
//
// * To
// list all of the currently available secrets, use ListSecrets.
func (c *Client) DeleteResourcePolicy(ctx context.Context, params *DeleteResourcePolicyInput, optFns ...func(*Options)) (*DeleteResourcePolicyOutput, error) {
	if params == nil {
		params = &DeleteResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResourcePolicy", params, optFns, addOperationDeleteResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResourcePolicyInput struct {

	// Specifies the secret that you want to delete the attached resource-based policy
	// for. You can specify either the Amazon Resource Name (ARN) or the friendly name
	// of the secret. If you specify an ARN, we generally recommend that you specify a
	// complete ARN. You can specify a partial ARN too—for example, if you don’t
	// include the final hyphen and six random characters that Secrets Manager adds at
	// the end of the ARN when you created the secret. A partial ARN match can work as
	// long as it uniquely matches only one secret. However, if your secret has a name
	// that ends in a hyphen followed by six characters (before Secrets Manager adds
	// the hyphen and six characters to the ARN) and you try to use that as a partial
	// ARN, then those characters cause Secrets Manager to assume that you’re
	// specifying a complete ARN. This confusion can cause unexpected results. To avoid
	// this situation, we recommend that you don’t create secret names ending with a
	// hyphen followed by six characters. If you specify an incomplete ARN without the
	// random suffix, and instead provide the 'friendly name', you must not include the
	// random suffix. If you do include the random suffix added by Secrets Manager, you
	// receive either a ResourceNotFoundException or an AccessDeniedException error,
	// depending on your permissions.
	//
	// This member is required.
	SecretId *string
}

type DeleteResourcePolicyOutput struct {

	// The ARN of the secret that the resource-based policy was deleted for.
	ARN *string

	// The friendly name of the secret that the resource-based policy was deleted for.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteResourcePolicy{}, middleware.After)
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
	addOpDeleteResourcePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResourcePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteResourcePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "DeleteResourcePolicy",
	}
}

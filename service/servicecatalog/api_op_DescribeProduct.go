// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about the specified product.
func (c *Client) DescribeProduct(ctx context.Context, params *DescribeProductInput, optFns ...func(*Options)) (*DescribeProductOutput, error) {
	if params == nil {
		params = &DescribeProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeProduct", params, optFns, addOperationDescribeProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeProductInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The product identifier.
	Id *string

	// The product name.
	Name *string
}

type DescribeProductOutput struct {

	// Information about the associated budgets.
	Budgets []*types.BudgetDetail

	// Information about the associated launch paths.
	LaunchPaths []*types.LaunchPath

	// Summary information about the product view.
	ProductViewSummary *types.ProductViewSummary

	// Information about the provisioning artifacts for the specified product.
	ProvisioningArtifacts []*types.ProvisioningArtifact

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeProduct{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeProduct(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "DescribeProduct",
	}
}

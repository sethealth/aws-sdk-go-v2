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

// Gets information about the products to which the caller has access.
func (c *Client) SearchProducts(ctx context.Context, params *SearchProductsInput, optFns ...func(*Options)) (*SearchProductsOutput, error) {
	if params == nil {
		params = &SearchProductsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProducts", params, optFns, addOperationSearchProductsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProductsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProductsInput struct {

	// The language code.
	//
	// * en - English (default)
	//
	// * jp - Japanese
	//
	// * zh - Chinese
	AcceptLanguage *string

	// The search filters. If no search filters are specified, the output includes all
	// products to which the caller has access.
	Filters map[string][]*string

	// The maximum number of items to return with this call.
	PageSize *int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	// The sort field. If no value is specified, the results are not sorted.
	SortBy types.ProductViewSortBy

	// The sort order. If no value is specified, the results are not sorted.
	SortOrder types.SortOrder
}

type SearchProductsOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// The product view aggregations.
	ProductViewAggregations map[string][]*types.ProductViewAggregationValue

	// Information about the product views.
	ProductViewSummaries []*types.ProductViewSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchProductsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProducts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProducts{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProducts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opSearchProducts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "SearchProducts",
	}
}

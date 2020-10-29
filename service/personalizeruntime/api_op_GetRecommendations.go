// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeruntime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of recommended items. The required input depends on the recipe
// type used to create the solution backing the campaign, as follows:
//
// *
// RELATED_ITEMS - itemId required, userId not used
//
// * USER_PERSONALIZATION -
// itemId optional, userId required
//
// Campaigns that are backed by a solution
// created using a recipe of type PERSONALIZED_RANKING use the API.
func (c *Client) GetRecommendations(ctx context.Context, params *GetRecommendationsInput, optFns ...func(*Options)) (*GetRecommendationsOutput, error) {
	if params == nil {
		params = &GetRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecommendations", params, optFns, addOperationGetRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecommendationsInput struct {

	// The Amazon Resource Name (ARN) of the campaign to use for getting
	// recommendations.
	//
	// This member is required.
	CampaignArn *string

	// The contextual metadata to use when getting recommendations. Contextual metadata
	// includes any interaction information that might be relevant when getting a
	// user's recommendations, such as the user's current location or device type.
	Context map[string]*string

	// The ARN of the filter to apply to the returned recommendations. For more
	// information, see Using Filters with Amazon Personalize
	// (https://docs.aws.amazon.com/personalize/latest/dg/filters.html). When using
	// this parameter, be sure the filter resource is ACTIVE.
	FilterArn *string

	// The item ID to provide recommendations for. Required for RELATED_ITEMS recipe
	// type.
	ItemId *string

	// The number of results to return. The default is 25. The maximum is 500.
	NumResults *int32

	// The user ID to provide recommendations for. Required for USER_PERSONALIZATION
	// recipe type.
	UserId *string
}

type GetRecommendationsOutput struct {

	// A list of recommendations sorted in ascending order by prediction score. There
	// can be a maximum of 500 items in the list.
	ItemList []*types.PredictedItem

	// The ID of the recommendation.
	RecommendationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecommendations{}, middleware.After)
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
	addOpGetRecommendationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecommendations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetRecommendations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "GetRecommendations",
	}
}

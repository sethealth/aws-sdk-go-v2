// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing matchmaking rule set. To delete the rule set, provide the
// rule set name. Rule sets cannot be deleted if they are currently being used by a
// matchmaking configuration. Learn more
//
// * Build a Rule Set
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-rulesets.html)
//
// Related
// operations
//
// * CreateMatchmakingConfiguration
//
// *
// DescribeMatchmakingConfigurations
//
// * UpdateMatchmakingConfiguration
//
// *
// DeleteMatchmakingConfiguration
//
// * CreateMatchmakingRuleSet
//
// *
// DescribeMatchmakingRuleSets
//
// * ValidateMatchmakingRuleSet
//
// *
// DeleteMatchmakingRuleSet
func (c *Client) DeleteMatchmakingRuleSet(ctx context.Context, params *DeleteMatchmakingRuleSetInput, optFns ...func(*Options)) (*DeleteMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &DeleteMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMatchmakingRuleSet", params, optFns, addOperationDeleteMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteMatchmakingRuleSetInput struct {

	// A unique identifier for a matchmaking rule set to be deleted. (Note: The rule
	// set name is different from the optional "name" field in the rule set body.) You
	// can use either the rule set name or ARN value.
	//
	// This member is required.
	Name *string
}

// Represents the returned data in response to a request operation.
type DeleteMatchmakingRuleSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteMatchmakingRuleSet{}, middleware.After)
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
	addOpDeleteMatchmakingRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMatchmakingRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteMatchmakingRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteMatchmakingRuleSet",
	}
}

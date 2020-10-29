// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Add tags to the specified Amazon SNS topic. For an overview, see Amazon SNS Tags
// (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html) in the Amazon SNS
// Developer Guide. When you use topic tags, keep the following guidelines in
// mind:
//
// * Adding more than 50 tags to a topic isn't recommended.
//
// * Tags don't
// have any semantic meaning. Amazon SNS interprets tags as character strings.
//
// *
// Tags are case-sensitive.
//
// * A new tag with a key identical to that of an
// existing tag overwrites the existing tag.
//
// * Tagging actions are limited to 10
// TPS per AWS account, per AWS region. If your application requires a higher
// throughput, file a technical support request
// (https://console.aws.amazon.com/support/home#/case/create?issueType=technical).
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	if params == nil {
		params = &TagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResource", params, optFns, addOperationTagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {

	// The ARN of the topic to which to add tags.
	//
	// This member is required.
	ResourceArn *string

	// The tags to be added to the specified topic. A tag consists of a required key
	// and an optional value.
	//
	// This member is required.
	Tags []*types.Tag
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTagResource{}, middleware.After)
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
	addOpTagResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTagResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "TagResource",
	}
}

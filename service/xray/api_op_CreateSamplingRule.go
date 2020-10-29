// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a rule to control sampling behavior for instrumented applications.
// Services retrieve rules with GetSamplingRules, and evaluate each rule in
// ascending order of priority for each request. If a rule matches, the service
// records a trace, borrowing it from the reservoir size. After 10 seconds, the
// service reports back to X-Ray with GetSamplingTargets to get updated versions of
// each in-use rule. The updated rule contains a trace quota that the service can
// use instead of borrowing from the reservoir.
func (c *Client) CreateSamplingRule(ctx context.Context, params *CreateSamplingRuleInput, optFns ...func(*Options)) (*CreateSamplingRuleOutput, error) {
	if params == nil {
		params = &CreateSamplingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSamplingRule", params, optFns, addOperationCreateSamplingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSamplingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSamplingRuleInput struct {

	// The rule definition.
	//
	// This member is required.
	SamplingRule *types.SamplingRule

	// A map that contains one or more tag keys and tag values to attach to an X-Ray
	// sampling rule. For more information about ways to use tags, see Tagging AWS
	// resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in
	// the AWS General Reference. The following restrictions apply to tags:
	//
	// * Maximum
	// number of user-applied tags per resource: 50
	//
	// * Maximum tag key length: 128
	// Unicode characters
	//
	// * Maximum tag value length: 256 Unicode characters
	//
	// * Valid
	// values for key and value: a-z, A-Z, 0-9, space, and the following characters: _
	// . : / = + - and @
	//
	// * Tag keys and values are case sensitive.
	//
	// * Don't use aws:
	// as a prefix for keys; it's reserved for AWS use.
	Tags []*types.Tag
}

type CreateSamplingRuleOutput struct {

	// The saved rule definition and metadata.
	SamplingRuleRecord *types.SamplingRuleRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSamplingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSamplingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSamplingRule{}, middleware.After)
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
	addOpCreateSamplingRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSamplingRule(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSamplingRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "CreateSamplingRule",
	}
}

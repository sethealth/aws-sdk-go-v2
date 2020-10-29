// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS CloudFormation WAFV2 template for the specified web ACL in the
// specified Amazon S3 bucket. Then, in CloudFormation, you create a stack from the
// template, to create the web ACL and its resources in AWS WAFV2. Use this to
// migrate your AWS WAF Classic web ACL to the latest version of AWS WAF. This is
// part of a larger migration procedure for web ACLs from AWS WAF Classic to the
// latest version of AWS WAF. For the full procedure, including caveats and manual
// steps to complete the migration and switch over to the new web ACL, see
// Migrating your AWS WAF Classic resources to AWS WAF
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-migrating-from-classic.html)
// in the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
func (c *Client) CreateWebACLMigrationStack(ctx context.Context, params *CreateWebACLMigrationStackInput, optFns ...func(*Options)) (*CreateWebACLMigrationStackOutput, error) {
	if params == nil {
		params = &CreateWebACLMigrationStackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebACLMigrationStack", params, optFns, addOperationCreateWebACLMigrationStackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebACLMigrationStackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebACLMigrationStackInput struct {

	// Indicates whether to exclude entities that can't be migrated or to stop the
	// migration. Set this to true to ignore unsupported entities in the web ACL during
	// the migration. Otherwise, if AWS WAF encounters unsupported entities, it stops
	// the process and throws an exception.
	//
	// This member is required.
	IgnoreUnsupportedType *bool

	// The name of the Amazon S3 bucket to store the CloudFormation template in. The S3
	// bucket must be configured as follows for the migration:
	//
	// * The bucket name must
	// start with aws-waf-migration-. For example, aws-waf-migration-my-web-acl.
	//
	// * The
	// bucket must be in the Region where you are deploying the template. For example,
	// for a web ACL in us-west-2, you must use an Amazon S3 bucket in us-west-2 and
	// you must deploy the template stack to us-west-2.
	//
	// * The bucket policies must
	// permit the migration process to write data. For listings of the bucket policies,
	// see the Examples section.
	//
	// This member is required.
	S3BucketName *string

	// The UUID of the WAF Classic web ACL that you want to migrate to WAF v2.
	//
	// This member is required.
	WebACLId *string
}

type CreateWebACLMigrationStackOutput struct {

	// The URL of the template created in Amazon S3.
	//
	// This member is required.
	S3ObjectUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateWebACLMigrationStackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWebACLMigrationStack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWebACLMigrationStack{}, middleware.After)
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
	addOpCreateWebACLMigrationStackValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebACLMigrationStack(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateWebACLMigrationStack(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateWebACLMigrationStack",
	}
}

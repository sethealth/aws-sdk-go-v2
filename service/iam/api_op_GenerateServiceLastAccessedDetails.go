// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Generates a report that includes details about when an IAM resource (user,
// group, role, or policy) was last used in an attempt to access AWS services.
// Recent activity usually appears within four hours. IAM reports activity for the
// last 365 days, or less if your Region began supporting this feature within the
// last year. For more information, see Regions Where Data Is Tracked
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html#access-advisor_tracking-period).
// The service last accessed data includes all attempts to access an AWS API, not
// just the successful ones. This includes all attempts that were made using the
// AWS Management Console, the AWS API through any of the SDKs, or any of the
// command line tools. An unexpected entry in the service last accessed data does
// not mean that your account has been compromised, because the request might have
// been denied. Refer to your CloudTrail logs as the authoritative source for
// information about all API calls and whether they were successful or denied
// access. For more information, see Logging IAM Events with CloudTrail
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/cloudtrail-integration.html)
// in the IAM User Guide. The GenerateServiceLastAccessedDetails operation returns
// a JobId. Use this parameter in the following operations to retrieve the
// following details from your report:
//
// * GetServiceLastAccessedDetails – Use this
// operation for users, groups, roles, or policies to list every AWS service that
// the resource could access using permissions policies. For each service, the
// response includes information about the most recent access attempt. The JobId
// returned by GenerateServiceLastAccessedDetail must be used by the same role
// within a session, or by the same user when used to call
// GetServiceLastAccessedDetail.
//
// * GetServiceLastAccessedDetailsWithEntities – Use
// this operation for groups and policies to list information about the associated
// entities (users or roles) that attempted to access a specific AWS service.
//
// To
// check the status of the GenerateServiceLastAccessedDetails request, use the
// JobId parameter in the same operations and test the JobStatus response
// parameter. For additional information about the permissions policies that allow
// an identity (user, group, or role) to access specific services, use the
// ListPoliciesGrantingServiceAccess operation. Service last accessed data does not
// use other policy types when determining whether a resource could access a
// service. These other policy types include resource-based policies, access
// control lists, AWS Organizations policies, IAM permissions boundaries, and AWS
// STS assume role policies. It only applies permissions policy logic. For more
// about the evaluation of policy types, see Evaluating Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_evaluation-logic.html#policy-eval-basics)
// in the IAM User Guide. For more information about service and action last
// accessed data, see Reducing Permissions Using Service Last Accessed Data
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_access-advisor.html)
// in the IAM User Guide.
func (c *Client) GenerateServiceLastAccessedDetails(ctx context.Context, params *GenerateServiceLastAccessedDetailsInput, optFns ...func(*Options)) (*GenerateServiceLastAccessedDetailsOutput, error) {
	if params == nil {
		params = &GenerateServiceLastAccessedDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GenerateServiceLastAccessedDetails", params, optFns, addOperationGenerateServiceLastAccessedDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GenerateServiceLastAccessedDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GenerateServiceLastAccessedDetailsInput struct {

	// The ARN of the IAM resource (user, group, role, or managed policy) used to
	// generate information about when the resource was last used in an attempt to
	// access an AWS service.
	//
	// This member is required.
	Arn *string

	// The level of detail that you want to generate. You can specify whether you want
	// to generate information about the last attempt to access services or actions. If
	// you specify service-level granularity, this operation generates only service
	// data. If you specify action-level granularity, it generates service and action
	// data. If you don't include this optional parameter, the operation generates
	// service data.
	Granularity types.AccessAdvisorUsageGranularityType
}

type GenerateServiceLastAccessedDetailsOutput struct {

	// The JobId that you can use in the GetServiceLastAccessedDetails or
	// GetServiceLastAccessedDetailsWithEntities operations. The JobId returned by
	// GenerateServiceLastAccessedDetail must be used by the same role within a
	// session, or by the same user when used to call GetServiceLastAccessedDetail.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGenerateServiceLastAccessedDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGenerateServiceLastAccessedDetails{}, middleware.After)
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
	addOpGenerateServiceLastAccessedDetailsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGenerateServiceLastAccessedDetails(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GenerateServiceLastAccessedDetails",
	}
}

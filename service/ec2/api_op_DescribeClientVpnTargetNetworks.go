// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the target networks associated with the specified Client VPN endpoint.
func (c *Client) DescribeClientVpnTargetNetworks(ctx context.Context, params *DescribeClientVpnTargetNetworksInput, optFns ...func(*Options)) (*DescribeClientVpnTargetNetworksOutput, error) {
	if params == nil {
		params = &DescribeClientVpnTargetNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClientVpnTargetNetworks", params, optFns, addOperationDescribeClientVpnTargetNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClientVpnTargetNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnTargetNetworksInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IDs of the target network associations.
	AssociationIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	// *
	// association-id - The ID of the association.
	//
	// * target-network-id - The ID of the
	// subnet specified as the target network.
	//
	// * vpc-id - The ID of the VPC in which
	// the target network is located.
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeClientVpnTargetNetworksOutput struct {

	// Information about the associated target networks.
	ClientVpnTargetNetworks []*types.TargetNetwork

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClientVpnTargetNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnTargetNetworks{}, middleware.After)
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
	addOpDescribeClientVpnTargetNetworksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeClientVpnTargetNetworks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClientVpnTargetNetworks",
	}
}

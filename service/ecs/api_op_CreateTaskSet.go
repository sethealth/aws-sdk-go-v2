// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a task set in the specified cluster and service. This is used when a
// service uses the EXTERNAL deployment controller type. For more information, see
// Amazon ECS Deployment Types
// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/deployment-types.html)
// in the Amazon Elastic Container Service Developer Guide.
func (c *Client) CreateTaskSet(ctx context.Context, params *CreateTaskSetInput, optFns ...func(*Options)) (*CreateTaskSetOutput, error) {
	if params == nil {
		params = &CreateTaskSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTaskSet", params, optFns, addOperationCreateTaskSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTaskSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTaskSetInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// service to create the task set in.
	//
	// This member is required.
	Cluster *string

	// The short name or full Amazon Resource Name (ARN) of the service to create the
	// task set in.
	//
	// This member is required.
	Service *string

	// The task definition for the tasks in the task set to use.
	//
	// This member is required.
	TaskDefinition *string

	// The capacity provider strategy to use for the task set. A capacity provider
	// strategy consists of one or more capacity providers along with the base and
	// weight to assign to them. A capacity provider must be associated with the
	// cluster to be used in a capacity provider strategy. The
	// PutClusterCapacityProviders API is used to associate a capacity provider with a
	// cluster. Only capacity providers with an ACTIVE or UPDATING status can be used.
	// If a capacityProviderStrategy is specified, the launchType parameter must be
	// omitted. If no capacityProviderStrategy or launchType is specified, the
	// defaultCapacityProviderStrategy for the cluster is used. If specifying a
	// capacity provider that uses an Auto Scaling group, the capacity provider must
	// already be created. New capacity providers can be created with the
	// CreateCapacityProvider API operation. To use a AWS Fargate capacity provider,
	// specify either the FARGATE or FARGATE_SPOT capacity providers. The AWS Fargate
	// capacity providers are available to all accounts and only need to be associated
	// with a cluster to be used. The PutClusterCapacityProviders API operation is used
	// to update the list of available capacity providers for a cluster after the
	// cluster is created.
	CapacityProviderStrategy []*types.CapacityProviderStrategyItem

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. Up to 32 ASCII characters are allowed.
	ClientToken *string

	// An optional non-unique tag that identifies this task set in external systems. If
	// the task set is associated with a service discovery registry, the tasks in this
	// task set will have the ECS_TASK_SET_EXTERNAL_ID AWS Cloud Map attribute set to
	// the provided value.
	ExternalId *string

	// The launch type that new tasks in the task set will use. For more information,
	// see Amazon ECS Launch Types
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html)
	// in the Amazon Elastic Container Service Developer Guide. If a launchType is
	// specified, the capacityProviderStrategy parameter must be omitted.
	LaunchType types.LaunchType

	// A load balancer object representing the load balancer to use with the task set.
	// The supported load balancer types are either an Application Load Balancer or a
	// Network Load Balancer.
	LoadBalancers []*types.LoadBalancer

	// An object representing the network configuration for a task or service.
	NetworkConfiguration *types.NetworkConfiguration

	// The platform version that the tasks in the task set should use. A platform
	// version is specified only for tasks using the Fargate launch type. If one isn't
	// specified, the LATEST platform version is used by default.
	PlatformVersion *string

	// A floating-point percentage of the desired number of tasks to place and keep
	// running in the task set.
	Scale *types.Scale

	// The details of the service discovery registries to assign to this task set. For
	// more information, see Service Discovery
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-discovery.html).
	ServiceRegistries []*types.ServiceRegistry

	// The metadata that you apply to the task set to help you categorize and organize
	// them. Each tag consists of a key and an optional value, both of which you
	// define. When a service is deleted, the tags are deleted as well. The following
	// basic restrictions apply to tags:
	//
	// * Maximum number of tags per resource - 50
	//
	// *
	// For each resource, each tag key must be unique, and each tag key can have only
	// one value.
	//
	// * Maximum key length - 128 Unicode characters in UTF-8
	//
	// * Maximum
	// value length - 256 Unicode characters in UTF-8
	//
	// * If your tagging schema is used
	// across multiple services and resources, remember that other services may have
	// restrictions on allowed characters. Generally allowed characters are: letters,
	// numbers, and spaces representable in UTF-8, and the following characters: + - =
	// . _ : / @.
	//
	// * Tag keys and values are case-sensitive.
	//
	// * Do not use aws:, AWS:,
	// or any upper or lowercase combination of such as a prefix for either keys or
	// values as it is reserved for AWS use. You cannot edit or delete tag keys or
	// values with this prefix. Tags with this prefix do not count against your tags
	// per resource limit.
	Tags []*types.Tag
}

type CreateTaskSetOutput struct {

	// Information about a set of Amazon ECS tasks in either an AWS CodeDeploy or an
	// EXTERNAL deployment. An Amazon ECS task set includes details such as the desired
	// number of tasks, how many tasks are running, and whether the task set serves
	// production traffic.
	TaskSet *types.TaskSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateTaskSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTaskSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTaskSet{}, middleware.After)
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
	addOpCreateTaskSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTaskSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateTaskSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "CreateTaskSet",
	}
}

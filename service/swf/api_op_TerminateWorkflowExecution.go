// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Records a WorkflowExecutionTerminated event and forces closure of the workflow
// execution identified by the given domain, runId, and workflowId. The child
// policy, registered with the workflow type or specified when starting this
// execution, is applied to any open child workflow executions of this workflow
// execution. If the identified workflow execution was in progress, it is
// terminated immediately. If a runId isn't specified, then the
// WorkflowExecutionTerminated event is recorded in the history of the current open
// workflow with the matching workflowId in the domain. You should consider using
// RequestCancelWorkflowExecution action instead because it allows the workflow to
// gracefully close while TerminateWorkflowExecution doesn't. Access Control You
// can use IAM policies to control this action's access to Amazon SWF resources as
// follows:
//
// * Use a Resource element with the domain name to limit the action to
// only specified domains.
//
// * Use an Action element to allow or deny permission to
// call this action.
//
// * You cannot use an IAM policy to constrain this action's
// parameters.
//
// If the caller doesn't have sufficient permissions to invoke the
// action, or the parameter values fall outside the specified constraints, the
// action fails. The associated event attribute's cause parameter is set to
// OPERATION_NOT_PERMITTED. For details and example IAM policies, see Using IAM to
// Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) TerminateWorkflowExecution(ctx context.Context, params *TerminateWorkflowExecutionInput, optFns ...func(*Options)) (*TerminateWorkflowExecutionOutput, error) {
	if params == nil {
		params = &TerminateWorkflowExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TerminateWorkflowExecution", params, optFns, addOperationTerminateWorkflowExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TerminateWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TerminateWorkflowExecutionInput struct {

	// The domain of the workflow execution to terminate.
	//
	// This member is required.
	Domain *string

	// The workflowId of the workflow execution to terminate.
	//
	// This member is required.
	WorkflowId *string

	// If set, specifies the policy to use for the child workflow executions of the
	// workflow execution being terminated. This policy overrides the child policy
	// specified for the workflow execution at registration time or when starting the
	// execution. The supported child policies are:
	//
	// * TERMINATE – The child executions
	// are terminated.
	//
	// * REQUEST_CANCEL – A request to cancel is attempted for each
	// child execution by recording a WorkflowExecutionCancelRequested event in its
	// history. It is up to the decider to take appropriate actions when it receives an
	// execution history with this event.
	//
	// * ABANDON – No action is taken. The child
	// executions continue to run.
	//
	// A child policy for this workflow execution must be
	// specified either as a default for the workflow type or through this parameter.
	// If neither this parameter is set nor a default child policy was specified at
	// registration time then a fault is returned.
	ChildPolicy types.ChildPolicy

	// Details for terminating the workflow execution.
	Details *string

	// A descriptive reason for terminating the workflow execution.
	Reason *string

	// The runId of the workflow execution to terminate.
	RunId *string
}

type TerminateWorkflowExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTerminateWorkflowExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpTerminateWorkflowExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpTerminateWorkflowExecution{}, middleware.After)
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
	addOpTerminateWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateWorkflowExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTerminateWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "TerminateWorkflowExecution",
	}
}

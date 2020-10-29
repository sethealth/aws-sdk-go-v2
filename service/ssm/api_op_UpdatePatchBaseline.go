// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Modifies an existing patch baseline. Fields not specified in the request are
// left unchanged. For information about valid key and value pairs in PatchFilters
// for each supported operating system type, see PatchFilter
// (http://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html).
func (c *Client) UpdatePatchBaseline(ctx context.Context, params *UpdatePatchBaselineInput, optFns ...func(*Options)) (*UpdatePatchBaselineOutput, error) {
	if params == nil {
		params = &UpdatePatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePatchBaseline", params, optFns, addOperationUpdatePatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePatchBaselineInput struct {

	// The ID of the patch baseline to update.
	//
	// This member is required.
	BaselineId *string

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline. For information about
	// accepted formats for lists of approved patches and rejected patches, see About
	// package name formats for approved and rejected patch lists
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	ApprovedPatches []*string

	// Assigns a new compliance severity level to an existing patch baseline.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is 'false'. Applies
	// to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to include patches in the baseline.
	GlobalFilters *types.PatchFilterGroup

	// The name of the patch baseline.
	Name *string

	// A list of explicitly rejected patches for the baseline. For information about
	// accepted formats for lists of approved patches and rejected patches, see About
	// package name formats for approved and rejected patch lists
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/patch-manager-approved-rejected-package-name-formats.html)
	// in the AWS Systems Manager User Guide.
	RejectedPatches []*string

	// The action for Patch Manager to take on patches included in the RejectedPackages
	// list.
	//
	// * ALLOW_AS_DEPENDENCY: A package in the Rejected patches list is
	// installed only if it is a dependency of another package. It is considered
	// compliant with the patch baseline, and its status is reported as InstalledOther.
	// This is the default action if no option is specified.
	//
	// * BLOCK: Packages in the
	// RejectedPatches list, and packages that include them as dependencies, are not
	// installed under any circumstances. If a package was installed before it was
	// added to the Rejected patches list, it is considered non-compliant with the
	// patch baseline, and its status is reported as InstalledRejected.
	RejectedPatchesAction types.PatchAction

	// If True, then all fields that are required by the CreatePatchBaseline action are
	// also required for this API request. Optional fields that are not specified are
	// set to null.
	Replace *bool

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []*types.PatchSource
}

type UpdatePatchBaselineOutput struct {

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []*string

	// The compliance severity level assigned to the patch baseline after the update
	// completed.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the instances. The default value is 'false'. Applies
	// to Linux instances only.
	ApprovedPatchesEnableNonSecurity *bool

	// The ID of the deleted patch baseline.
	BaselineId *string

	// The date when the patch baseline was created.
	CreatedDate *time.Time

	// A description of the Patch Baseline.
	Description *string

	// A set of global filters used to exclude patches from the baseline.
	GlobalFilters *types.PatchFilterGroup

	// The date when the patch baseline was last modified.
	ModifiedDate *time.Time

	// The name of the patch baseline.
	Name *string

	// The operating system rule used by the updated patch baseline.
	OperatingSystem types.OperatingSystem

	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []*string

	// The action specified to take on patches included in the RejectedPatches list. A
	// patch can be allowed only if it is a dependency of another package, or blocked
	// entirely along with packages that include it as a dependency.
	RejectedPatchesAction types.PatchAction

	// Information about the patches to use to update the instances, including target
	// operating systems and source repositories. Applies to Linux instances only.
	Sources []*types.PatchSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdatePatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePatchBaseline{}, middleware.After)
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
	addOpUpdatePatchBaselineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePatchBaseline(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdatePatchBaseline(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdatePatchBaseline",
	}
}

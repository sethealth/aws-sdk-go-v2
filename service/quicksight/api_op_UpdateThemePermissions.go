// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the resource permissions for a theme. Permissions apply to the action to
// grant or revoke permissions on, for example "quicksight:DescribeTheme". Theme
// permissions apply in groupings. Valid groupings include the following for the
// three levels of permissions, which are user, owner, or no permissions:
//
// *
// User
//
// * "quicksight:DescribeTheme"
//
// * "quicksight:DescribeThemeAlias"
//
// *
// "quicksight:ListThemeAliases"
//
// * "quicksight:ListThemeVersions"
//
// * Owner
//
// *
// "quicksight:DescribeTheme"
//
// * "quicksight:DescribeThemeAlias"
//
// *
// "quicksight:ListThemeAliases"
//
// * "quicksight:ListThemeVersions"
//
// *
// "quicksight:DeleteTheme"
//
// * "quicksight:UpdateTheme"
//
// *
// "quicksight:CreateThemeAlias"
//
// * "quicksight:DeleteThemeAlias"
//
// *
// "quicksight:UpdateThemeAlias"
//
// * "quicksight:UpdateThemePermissions"
//
// *
// "quicksight:DescribeThemePermissions"
//
// * To specify no permissions, omit the
// permissions list.
func (c *Client) UpdateThemePermissions(ctx context.Context, params *UpdateThemePermissionsInput, optFns ...func(*Options)) (*UpdateThemePermissionsOutput, error) {
	if params == nil {
		params = &UpdateThemePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThemePermissions", params, optFns, addOperationUpdateThemePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThemePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateThemePermissionsInput struct {

	// The ID of the AWS account that contains the theme.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the theme.
	//
	// This member is required.
	ThemeId *string

	// A list of resource permissions to be granted for the theme.
	GrantPermissions []*types.ResourcePermission

	// A list of resource permissions to be revoked from the theme.
	RevokePermissions []*types.ResourcePermission
}

type UpdateThemePermissionsOutput struct {

	// The resulting list of resource permissions for the theme.
	Permissions []*types.ResourcePermission

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// The Amazon Resource Name (ARN) of the theme.
	ThemeArn *string

	// The ID for the theme.
	ThemeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateThemePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThemePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThemePermissions{}, middleware.After)
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
	addOpUpdateThemePermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThemePermissions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateThemePermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateThemePermissions",
	}
}

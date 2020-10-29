// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified IAM user. Unlike the AWS Management Console, when you
// delete a user programmatically, you must delete the items attached to the user
// manually, or the deletion fails. For more information, see Deleting an IAM User
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_manage.html#id_users_deleting_cli).
// Before attempting to delete a user, remove the following items:
//
// * Password
// (DeleteLoginProfile)
//
// * Access keys (DeleteAccessKey)
//
// * Signing certificate
// (DeleteSigningCertificate)
//
// * SSH public key (DeleteSSHPublicKey)
//
// * Git
// credentials (DeleteServiceSpecificCredential)
//
// * Multi-factor authentication
// (MFA) device (DeactivateMFADevice, DeleteVirtualMFADevice)
//
// * Inline policies
// (DeleteUserPolicy)
//
// * Attached managed policies (DetachUserPolicy)
//
// * Group
// memberships (RemoveUserFromGroup)
func (c *Client) DeleteUser(ctx context.Context, params *DeleteUserInput, optFns ...func(*Options)) (*DeleteUserOutput, error) {
	if params == nil {
		params = &DeleteUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUser", params, optFns, addOperationDeleteUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserInput struct {

	// The name of the user to delete. This parameter allows (through its regex pattern
	// (http://wikipedia.org/wiki/regex)) a string of characters consisting of upper
	// and lowercase alphanumeric characters with no spaces. You can also include any
	// of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

type DeleteUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteUser{}, middleware.After)
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
	addOpDeleteUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteUser",
	}
}

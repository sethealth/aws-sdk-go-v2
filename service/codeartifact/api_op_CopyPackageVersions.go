// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Copies package versions from one repository to another repository in the same
// domain. You must specify versions or versionRevisions. You cannot specify both.
func (c *Client) CopyPackageVersions(ctx context.Context, params *CopyPackageVersionsInput, optFns ...func(*Options)) (*CopyPackageVersionsOutput, error) {
	if params == nil {
		params = &CopyPackageVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopyPackageVersions", params, optFns, addOperationCopyPackageVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopyPackageVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopyPackageVersionsInput struct {

	// The name of the repository into which package versions are copied.
	//
	// This member is required.
	DestinationRepository *string

	// The name of the domain that contains the source and destination repositories.
	//
	// This member is required.
	Domain *string

	// The format of the package that is copied. The valid package types are:
	//
	// * npm: A
	// Node Package Manager (npm) package.
	//
	// * pypi: A Python Package Index (PyPI)
	// package.
	//
	// * maven: A Maven package that contains compiled code in a
	// distributable format, such as a JAR file.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package that is copied.
	//
	// This member is required.
	Package *string

	// The name of the repository that contains the package versions to copy.
	//
	// This member is required.
	SourceRepository *string

	// Set to true to overwrite a package version that already exists in the
	// destination repository. If set to false and the package version already exists
	// in the destination repository, the package version is returned in the
	// failedVersions field of the response with an ALREADY_EXISTS error code.
	AllowOverwrite *bool

	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string

	// Set to true to copy packages from repositories that are upstream from the source
	// repository to the destination repository. The default setting is false. For more
	// information, see Working with upstream repositories
	// (https://docs.aws.amazon.com/codeartifact/latest/ug/repos-upstream.html).
	IncludeFromUpstream *bool

	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	// * The namespace of a Maven package is its
	// groupId.
	//
	// * The namespace of an npm package is its scope.
	//
	// * A Python package
	// does not contain a corresponding component, so Python packages do not have a
	// namespace.
	Namespace *string

	// A list of key-value pairs. The keys are package versions and the values are
	// package version revisions. A CopyPackageVersion operation succeeds if the
	// specified versions in the source repository match the specified package version
	// revision. You must specify versions or versionRevisions. You cannot specify
	// both.
	VersionRevisions map[string]*string

	// The versions of the package to copy. You must specify versions or
	// versionRevisions. You cannot specify both.
	Versions []*string
}

type CopyPackageVersionsOutput struct {

	// A map of package versions that failed to copy and their error codes. The
	// possible error codes are in the PackageVersionError data type. They are:
	//
	// *
	// ALREADY_EXISTS
	//
	// * MISMATCHED_REVISION
	//
	// * MISMATCHED_STATUS
	//
	// * NOT_ALLOWED
	//
	// *
	// NOT_FOUND
	//
	// * SKIPPED
	FailedVersions map[string]*types.PackageVersionError

	// A list of the package versions that were successfully copied to your repository.
	SuccessfulVersions map[string]*types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopyPackageVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCopyPackageVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCopyPackageVersions{}, middleware.After)
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
	addOpCopyPackageVersionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopyPackageVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCopyPackageVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "CopyPackageVersions",
	}
}

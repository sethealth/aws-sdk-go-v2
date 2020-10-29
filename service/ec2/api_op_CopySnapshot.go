// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	presignedurlcust "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"net/http"
)

// Copies a point-in-time snapshot of an EBS volume and stores it in Amazon S3. You
// can copy the snapshot within the same Region or from one Region to another. You
// can use the snapshot to create EBS volumes or Amazon Machine Images (AMIs).
// Copies of encrypted EBS snapshots remain encrypted. Copies of unencrypted
// snapshots remain unencrypted, unless you enable encryption for the snapshot copy
// operation. By default, encrypted snapshot copies use the default AWS Key
// Management Service (AWS KMS) customer master key (CMK); however, you can specify
// a different CMK. To copy an encrypted snapshot that has been shared from another
// account, you must have permissions for the CMK used to encrypt the snapshot.
// Snapshots created by copying another snapshot have an arbitrary volume ID that
// should not be used for any purpose. For more information, see Copying an Amazon
// EBS snapshot
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-copy-snapshot.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) CopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (*CopySnapshotOutput, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CopySnapshot", params, optFns, addOperationCopySnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CopySnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CopySnapshotInput struct {

	// The ID of the Region that contains the snapshot to be copied.
	//
	// This member is required.
	SourceRegion *string

	// The ID of the EBS snapshot to copy.
	//
	// This member is required.
	SourceSnapshotId *string

	// A description for the EBS snapshot.
	Description *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// To encrypt a copy of an unencrypted snapshot if encryption by default is not
	// enabled, enable encryption using this parameter. Otherwise, omit this parameter.
	// Encrypted snapshots are encrypted, even if you omit this parameter and
	// encryption by default is not enabled. You cannot set this parameter to false.
	// For more information, see Amazon EBS Encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// The identifier of the AWS Key Management Service (AWS KMS) customer master key
	// (CMK) to use for Amazon EBS encryption. If this parameter is not specified, your
	// AWS managed CMK for EBS is used. If KmsKeyId is specified, the encrypted state
	// must be true. You can specify the CMK using any of the following:
	//
	// * Key ID. For
	// example, key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Key alias. For example,
	// alias/ExampleAlias.
	//
	// * Key ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:key/abcd1234-a123-456a-a12b-a123b4cd56ef.
	//
	// *
	// Alias ARN. For example,
	// arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.
	//
	// AWS authenticates the
	// CMK asynchronously. Therefore, if you specify an ID, alias, or ARN that is not
	// valid, the action can appear to complete, but eventually fails.
	KmsKeyId *string

	// When you copy an encrypted source snapshot using the Amazon EC2 Query API, you
	// must supply a pre-signed URL. This parameter is optional for unencrypted
	// snapshots. For more information, see Query Requests
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/Query-Requests.html).
	// The PresignedUrl should use the snapshot source endpoint, the CopySnapshot
	// action, and include the SourceRegion, SourceSnapshotId, and DestinationRegion
	// parameters. The PresignedUrl must be signed using AWS Signature Version 4.
	// Because EBS snapshots are stored in Amazon S3, the signing algorithm for this
	// parameter uses the same logic that is described in Authenticating Requests by
	// Using Query Parameters (AWS Signature Version 4)
	// (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)
	// in the Amazon Simple Storage Service API Reference. An invalid or improperly
	// signed PresignedUrl will cause the copy operation to fail asynchronously, and
	// the snapshot will move to an error state.
	PresignedUrl *string

	// The tags to apply to the new snapshot.
	TagSpecifications []*types.TagSpecification

	// Used by the SDK's PresignURL autofill customization to specify the region the of
	// the client's request.
	destinationRegion *string
}

type CopySnapshotOutput struct {

	// The ID of the new snapshot.
	SnapshotId *string

	// Any tags applied to the new snapshot.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCopySnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCopySnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCopySnapshot{}, middleware.After)
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
	addCopySnapshotPresignURLMiddleware(stack, options)
	addOpCopySnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCopySnapshot(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func copyCopySnapshotInputForPresign(params interface{}) (interface{}, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return nil, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCopySnapshotPresignedUrl(params interface{}) (string, bool, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	if input.PresignedUrl == nil || len(*input.PresignedUrl) == 0 {
		return ``, false, nil
	}
	return *input.PresignedUrl, true, nil
}
func getCopySnapshotSourceRegion(params interface{}) (string, bool, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return ``, false, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	if input.SourceRegion == nil || len(*input.SourceRegion) == 0 {
		return ``, false, nil
	}
	return *input.SourceRegion, true, nil
}
func setCopySnapshotPresignedUrl(params interface{}, value string) error {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	input.PresignedUrl = &value
	return nil
}
func setCopySnapshotdestinationRegion(params interface{}, value string) error {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	input.destinationRegion = &value
	return nil
}

type copySnapshotHTTPPresignURLClient struct {
	client    *Client
	presigner *v4.Signer
}

func newCopySnapshotHTTPPresignURLClient(options Options, optFns ...func(*Options)) *copySnapshotHTTPPresignURLClient {
	return &copySnapshotHTTPPresignURLClient{
		client:    New(options, optFns...),
		presigner: v4.NewSigner(),
	}
}
func (c *copySnapshotHTTPPresignURLClient) PresignCopySnapshot(ctx context.Context, params *CopySnapshotInput, optFns ...func(*Options)) (string, http.Header, error) {
	if params == nil {
		params = &CopySnapshotInput{}
	}

	optFns = append(optFns, func(o *Options) {
		o.HTTPClient = &smithyhttp.NopClient{}
	})

	ctx = presignedurlcust.WithIsPresigning(ctx)
	result, _, err := c.client.invokeOperation(ctx, "CopySnapshot", params, optFns,
		addOperationCopySnapshotMiddlewares,
		c.convertToPresignMiddleware,
	)
	if err != nil {
		return ``, nil, err
	}

	out := result.(*v4.PresignedHTTPRequest)
	return out.URL, out.SignedHeader, nil
}
func (c *copySnapshotHTTPPresignURLClient) convertToPresignMiddleware(stack *middleware.Stack, options Options) (err error) {
	stack.Finalize.Clear()
	stack.Deserialize.Clear()
	stack.Build.Remove(awsmiddleware.RequestInvocationIDMiddleware{}.ID())
	err = stack.Finalize.Add(v4.NewPresignHTTPRequestMiddleware(options.Credentials, c.presigner), middleware.After)
	if err != nil {
		return err
	}
	err = query.AddAsGetRequestMiddleware(stack)
	if err != nil {
		return err
	}
	return nil
}
func addCopySnapshotPresignURLMiddleware(stack *middleware.Stack, options Options) error {
	return presignedurlcust.AddMiddleware(stack, presignedurlcust.Options{
		Accessor: presignedurlcust.ParameterAccessor{
			GetPresignedURL:      getCopySnapshotPresignedUrl,
			GetSourceRegion:      getCopySnapshotSourceRegion,
			CopyInput:            copyCopySnapshotInputForPresign,
			SetDestinationRegion: setCopySnapshotdestinationRegion,
			SetPresignedURL:      setCopySnapshotPresignedUrl,
		},
		Presigner: &presignAutoFillCopySnapshotClient{client: newCopySnapshotHTTPPresignURLClient(options)},
	})
}

type presignAutoFillCopySnapshotClient struct {
	client *copySnapshotHTTPPresignURLClient
}

func (c *presignAutoFillCopySnapshotClient) PresignURL(ctx context.Context, region string, params interface{}) (string, http.Header, error) {
	input, ok := params.(*CopySnapshotInput)
	if !ok {
		return ``, nil, fmt.Errorf("expect *CopySnapshotInput type, got %T", params)
	}
	optFn := func(o *Options) {
		o.Region = region
		o.APIOptions = append(o.APIOptions, presignedurlcust.RemoveMiddleware)
	}
	return c.client.PresignCopySnapshot(ctx, input, optFn)
}

func newServiceMetadataMiddleware_opCopySnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CopySnapshot",
	}
}

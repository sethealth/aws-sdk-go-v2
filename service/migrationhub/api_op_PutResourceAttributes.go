// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides identifying details of the resource being migrated so that it can be
// associated in the Application Discovery Service repository. This association
// occurs asynchronously after PutResourceAttributes returns.
//
// * Keep in mind that
// subsequent calls to PutResourceAttributes will override previously stored
// attributes. For example, if it is first called with a MAC address, but later, it
// is desired to add an IP address, it will then be required to call it with both
// the IP and MAC addresses to prevent overriding the MAC address.
//
// * Note the
// instructions regarding the special use case of the ResourceAttributeList
// (https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#migrationhub-PutResourceAttributes-request-ResourceAttributeList)
// parameter when specifying any "VM" related value.
//
// Because this is an
// asynchronous call, it will always return 200, whether an association occurs or
// not. To confirm if an association was found based on the provided details, call
// ListDiscoveredResources.
func (c *Client) PutResourceAttributes(ctx context.Context, params *PutResourceAttributesInput, optFns ...func(*Options)) (*PutResourceAttributesOutput, error) {
	if params == nil {
		params = &PutResourceAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourceAttributes", params, optFns, addOperationPutResourceAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourceAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourceAttributesInput struct {

	// Unique identifier that references the migration task. Do not store personal data
	// in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// Information about the resource that is being migrated. This data will be used to
	// map the task to a resource in the Application Discovery Service repository.
	// Takes the object array of ResourceAttribute where the Type field is reserved for
	// the following values: IPV4_ADDRESS | IPV6_ADDRESS | MAC_ADDRESS | FQDN |
	// VM_MANAGER_ID | VM_MANAGED_OBJECT_REFERENCE | VM_NAME | VM_PATH | BIOS_ID |
	// MOTHERBOARD_SERIAL_NUMBER where the identifying value can be a string up to 256
	// characters.
	//
	// * If any "VM" related value is set for a ResourceAttribute object,
	// it is required that VM_MANAGER_ID, as a minimum, is always set. If VM_MANAGER_ID
	// is not set, then all "VM" fields will be discarded and "VM" fields will not be
	// used for matching the migration task to a server in Application Discovery
	// Service repository. See the Example
	// (https://docs.aws.amazon.com/migrationhub/latest/ug/API_PutResourceAttributes.html#API_PutResourceAttributes_Examples)
	// section below for a use case of specifying "VM" related values.
	//
	// * If a server
	// you are trying to match has multiple IP or MAC addresses, you should provide as
	// many as you know in separate type/value pairs passed to the
	// ResourceAttributeList parameter to maximize the chances of matching.
	//
	// This member is required.
	ResourceAttributeList []*types.ResourceAttribute

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type PutResourceAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutResourceAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourceAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourceAttributes{}, middleware.After)
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
	addOpPutResourceAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourceAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutResourceAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "PutResourceAttributes",
	}
}

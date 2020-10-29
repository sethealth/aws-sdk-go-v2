// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Creates an uninitialized HSM instance. There is an upfront fee charged for each
// HSM instance that you create with the CreateHsm operation. If you accidentally
// provision an HSM and want to request a refund, delete the instance using the
// DeleteHsm operation, go to the AWS Support Center
// (https://console.aws.amazon.com/support/home), create a new case, and select
// Account and Billing Support. It can take up to 20 minutes to create and
// provision an HSM. You can monitor the status of the HSM with the DescribeHsm
// operation. The HSM is ready to be initialized when the status changes to
// RUNNING.
func (c *Client) CreateHsm(ctx context.Context, params *CreateHsmInput, optFns ...func(*Options)) (*CreateHsmOutput, error) {
	if params == nil {
		params = &CreateHsmInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateHsm", params, optFns, addOperationCreateHsmMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the CreateHsm operation.
type CreateHsmInput struct {

	// The ARN of an IAM role to enable the AWS CloudHSM service to allocate an ENI on
	// your behalf.
	//
	// This member is required.
	IamRoleArn *string

	// The SSH public key to install on the HSM.
	//
	// This member is required.
	SshKey *string

	// The identifier of the subnet in your VPC in which to place the HSM.
	//
	// This member is required.
	SubnetId *string

	// Specifies the type of subscription for the HSM.
	//
	// * PRODUCTION - The HSM is being
	// used in a production environment.
	//
	// * TRIAL - The HSM is being used in a product
	// trial.
	//
	// This member is required.
	SubscriptionType types.SubscriptionType

	// A user-defined token to ensure idempotence. Subsequent calls to this operation
	// with the same token will be ignored.
	ClientToken *string

	// The IP address to assign to the HSM's ENI. If an IP address is not specified, an
	// IP address will be randomly chosen from the CIDR range of the subnet.
	EniIp *string

	// The external ID from IamRoleArn, if present.
	ExternalId *string

	// The IP address for the syslog monitoring server. The AWS CloudHSM service only
	// supports one syslog monitoring server.
	SyslogIp *string
}

// Contains the output of the CreateHsm operation.
type CreateHsmOutput struct {

	// The ARN of the HSM.
	HsmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateHsmMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateHsm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateHsm{}, middleware.After)
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
	addOpCreateHsmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateHsm(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateHsm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "CreateHsm",
	}
}

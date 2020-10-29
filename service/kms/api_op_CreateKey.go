// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a unique customer managed customer master key
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master-keys)
// (CMK) in your AWS account and Region. You cannot use this operation to create a
// CMK in a different AWS account. You can use the CreateKey operation to create
// symmetric or asymmetric CMKs.
//
// * Symmetric CMKs contain a 256-bit symmetric key
// that never leaves AWS KMS unencrypted. To use the CMK, you must call AWS KMS.
// You can use a symmetric CMK to encrypt and decrypt small amounts of data, but
// they are typically used to generate data keys
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys)
// and data keys pairs
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-key-pairs).
// For details, see GenerateDataKey and GenerateDataKeyPair.
//
// * Asymmetric CMKs can
// contain an RSA key pair or an Elliptic Curve (ECC) key pair. The private key in
// an asymmetric CMK never leaves AWS KMS unencrypted. However, you can use the
// GetPublicKey operation to download the public key so it can be used outside of
// AWS KMS. CMKs with RSA key pairs can be used to encrypt or decrypt data or sign
// and verify messages (but not both). CMKs with ECC key pairs can be used only to
// sign and verify messages.
//
// For information about symmetric and asymmetric CMKs,
// see Using Symmetric and Asymmetric CMKs
// (https://docs.aws.amazon.com/kms/latest/developerguide/symmetric-asymmetric.html)
// in the AWS Key Management Service Developer Guide. To create different types of
// CMKs, use the following guidance: Asymmetric CMKs To create an asymmetric CMK,
// use the CustomerMasterKeySpec parameter to specify the type of key material in
// the CMK. Then, use the KeyUsage parameter to determine whether the CMK will be
// used to encrypt and decrypt or sign and verify. You can't change these
// properties after the CMK is created. Symmetric CMKs When creating a symmetric
// CMK, you don't need to specify the CustomerMasterKeySpec or KeyUsage parameters.
// The default value for CustomerMasterKeySpec, SYMMETRIC_DEFAULT, and the default
// value for KeyUsage, ENCRYPT_DECRYPT, are the only valid values for symmetric
// CMKs. Imported Key Material To import your own key material, begin by creating a
// symmetric CMK with no key material. To do this, use the Origin parameter of
// CreateKey with a value of EXTERNAL. Next, use GetParametersForImport operation
// to get a public key and import token, and use the public key to encrypt your key
// material. Then, use ImportKeyMaterial with your import token to import the key
// material. For step-by-step instructions, see Importing Key Material
// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
// the AWS Key Management Service Developer Guide . You cannot import the key
// material into an asymmetric CMK. Custom Key Stores To create a symmetric CMK in
// a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html),
// use the CustomKeyStoreId parameter to specify the custom key store. You must
// also use the Origin parameter with a value of AWS_CLOUDHSM. The AWS CloudHSM
// cluster that is associated with the custom key store must have at least two
// active HSMs in different Availability Zones in the AWS Region. You cannot create
// an asymmetric CMK in a custom key store. For information about custom key stores
// in AWS KMS see Using Custom Key Stores
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
// in the AWS Key Management Service Developer Guide .
func (c *Client) CreateKey(ctx context.Context, params *CreateKeyInput, optFns ...func(*Options)) (*CreateKeyOutput, error) {
	if params == nil {
		params = &CreateKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKey", params, optFns, addOperationCreateKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKeyInput struct {

	// A flag to indicate whether to bypass the key policy lockout safety check.
	// Setting this value to true increases the risk that the CMK becomes unmanageable.
	// Do not set this value to true indiscriminately. For more information, refer to
	// the scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section in the AWS Key Management Service Developer Guide . Use this parameter
	// only when you include a policy in the request and you intend to prevent the
	// principal that is making the request from making a subsequent PutKeyPolicy
	// request on the CMK. The default value is false.
	BypassPolicyLockoutSafetyCheck *bool

	// Creates the CMK in the specified custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and the key material in its associated AWS CloudHSM cluster. To create a CMK in
	// a custom key store, you must also specify the Origin parameter with a value of
	// AWS_CLOUDHSM. The AWS CloudHSM cluster that is associated with the custom key
	// store must have at least two active HSMs, each in a different Availability Zone
	// in the Region. This parameter is valid only for symmetric CMKs. You cannot
	// create an asymmetric CMK in a custom key store. To find the ID of a custom key
	// store, use the DescribeCustomKeyStores operation. The response includes the
	// custom key store ID and the ID of the AWS CloudHSM cluster. This operation is
	// part of the Custom Key Store feature
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// feature in AWS KMS, which combines the convenience and extensive integration of
	// AWS KMS with the isolation and control of a single-tenant key store.
	CustomKeyStoreId *string

	// Specifies the type of CMK to create. The default value, SYMMETRIC_DEFAULT,
	// creates a CMK with a 256-bit symmetric key for encryption and decryption. For
	// help choosing a key spec for your CMK, see How to Choose Your CMK Configuration
	// (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-choose.html)
	// in the AWS Key Management Service Developer Guide. The CustomerMasterKeySpec
	// determines whether the CMK contains a symmetric key or an asymmetric key pair.
	// It also determines the encryption algorithms or signing algorithms that the CMK
	// supports. You can't change the CustomerMasterKeySpec after the CMK is created.
	// To further restrict the algorithms that can be used with the CMK, use a
	// condition key in its key policy or IAM policy. For more information, see
	// kms:EncryptionAlgorithm
	// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-encryption-algorithm)
	// or kms:Signing Algorithm
	// (https://docs.aws.amazon.com/kms/latest/developerguide/policy-conditions.html#conditions-kms-signing-algorithm)
	// in the AWS Key Management Service Developer Guide. AWS services that are
	// integrated with AWS KMS
	// (http://aws.amazon.com/kms/features/#AWS_Service_Integration) use symmetric CMKs
	// to protect your data. These services do not support asymmetric CMKs. For help
	// determining whether a CMK is symmetric or asymmetric, see Identifying Symmetric
	// and Asymmetric CMKs
	// (https://docs.aws.amazon.com/kms/latest/developerguide/find-symm-asymm.html) in
	// the AWS Key Management Service Developer Guide. AWS KMS supports the following
	// key specs for CMKs:
	//
	// * Symmetric key (default)
	//
	// * SYMMETRIC_DEFAULT
	// (AES-256-GCM)
	//
	// * Asymmetric RSA key pairs
	//
	// * RSA_2048
	//
	// * RSA_3072
	//
	// * RSA_4096
	//
	// *
	// Asymmetric NIST-recommended elliptic curve key pairs
	//
	// * ECC_NIST_P256
	// (secp256r1)
	//
	// * ECC_NIST_P384 (secp384r1)
	//
	// * ECC_NIST_P521 (secp521r1)
	//
	// * Other
	// asymmetric elliptic curve key pairs
	//
	// * ECC_SECG_P256K1 (secp256k1), commonly
	// used for cryptocurrencies.
	CustomerMasterKeySpec types.CustomerMasterKeySpec

	// A description of the CMK. Use a description that helps you decide whether the
	// CMK is appropriate for a task.
	Description *string

	// Determines the cryptographic operations
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This
	// parameter is required only for asymmetric CMKs. You can't change the KeyUsage
	// value after the CMK is created. Select only one valid value.
	//
	// * For symmetric
	// CMKs, omit the parameter or specify ENCRYPT_DECRYPT.
	//
	// * For asymmetric CMKs with
	// RSA key material, specify ENCRYPT_DECRYPT or SIGN_VERIFY.
	//
	// * For asymmetric CMKs
	// with ECC key material, specify SIGN_VERIFY.
	KeyUsage types.KeyUsageType

	// The source of the key material for the CMK. You cannot change the origin after
	// you create the CMK. The default is AWS_KMS, which means AWS KMS creates the key
	// material. When the parameter value is EXTERNAL, AWS KMS creates a CMK without
	// key material so that you can import key material from your existing key
	// management infrastructure. For more information about importing key material
	// into AWS KMS, see Importing Key Material
	// (https://docs.aws.amazon.com/kms/latest/developerguide/importing-keys.html) in
	// the AWS Key Management Service Developer Guide. This value is valid only for
	// symmetric CMKs. When the parameter value is AWS_CLOUDHSM, AWS KMS creates the
	// CMK in an AWS KMS custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html)
	// and creates its key material in the associated AWS CloudHSM cluster. You must
	// also use the CustomKeyStoreId parameter to identify the custom key store. This
	// value is valid only for symmetric CMKs.
	Origin types.OriginType

	// The key policy to attach to the CMK. If you provide a key policy, it must meet
	// the following criteria:
	//
	// * If you don't set BypassPolicyLockoutSafetyCheck to
	// true, the key policy must allow the principal that is making the CreateKey
	// request to make a subsequent PutKeyPolicy request on the CMK. This reduces the
	// risk that the CMK becomes unmanageable. For more information, refer to the
	// scenario in the Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default-allow-root-enable-iam)
	// section of the AWS Key Management Service Developer Guide .
	//
	// * Each statement in
	// the key policy must contain one or more principals. The principals in the key
	// policy must exist and be visible to AWS KMS. When you create a new AWS principal
	// (for example, an IAM user or role), you might need to enforce a delay before
	// including the new principal in a key policy because the new principal might not
	// be immediately visible to AWS KMS. For more information, see Changes that I make
	// are not always immediately visible
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency)
	// in the AWS Identity and Access Management User Guide.
	//
	// If you do not provide a
	// key policy, AWS KMS attaches a default key policy to the CMK. For more
	// information, see Default Key Policy
	// (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html#key-policy-default)
	// in the AWS Key Management Service Developer Guide. The key policy size quota is
	// 32 kilobytes (32768 bytes).
	Policy *string

	// One or more tags. Each tag consists of a tag key and a tag value. Both the tag
	// key and the tag value are required, but the tag value can be an empty (null)
	// string. When you add tags to an AWS resource, AWS generates a cost allocation
	// report with usage and costs aggregated by tags. For information about adding,
	// changing, deleting and listing tags for CMKs, see Tagging Keys
	// (https://docs.aws.amazon.com/kms/latest/developerguide/tagging-keys.html). Use
	// this parameter to tag the CMK when it is created. To add tags to an existing
	// CMK, use the TagResource operation.
	Tags []*types.Tag
}

type CreateKeyOutput struct {

	// Metadata associated with the CMK.
	KeyMetadata *types.KeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateKey{}, middleware.After)
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
	addOpCreateKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "CreateKey",
	}
}

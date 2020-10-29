// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Requests an ACM certificate for use with other AWS services. To request an ACM
// certificate, you must specify a fully qualified domain name (FQDN) in the
// DomainName parameter. You can also specify additional FQDNs in the
// SubjectAlternativeNames parameter. If you are requesting a private certificate,
// domain validation is not required. If you are requesting a public certificate,
// each domain name that you specify must be validated to verify that you own or
// control the domain. You can use DNS validation
// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or
// email validation
// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html).
// We recommend that you use DNS validation. ACM issues public certificates after
// receiving approval from the domain owner.
func (c *Client) RequestCertificate(ctx context.Context, params *RequestCertificateInput, optFns ...func(*Options)) (*RequestCertificateOutput, error) {
	if params == nil {
		params = &RequestCertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestCertificate", params, optFns, addOperationRequestCertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestCertificateInput struct {

	// Fully qualified domain name (FQDN), such as www.example.com, that you want to
	// secure with an ACM certificate. Use an asterisk (*) to create a wildcard
	// certificate that protects several sites in the same domain. For example,
	// *.example.com protects www.example.com, site.example.com, and
	// images.example.com. The first domain name you enter cannot exceed 64 octets,
	// including periods. Each subsequent Subject Alternative Name (SAN), however, can
	// be up to 253 octets in length.
	//
	// This member is required.
	DomainName *string

	// The Amazon Resource Name (ARN) of the private certificate authority (CA) that
	// will be used to issue the certificate. If you do not provide an ARN and you are
	// trying to request a private certificate, ACM will attempt to issue a public
	// certificate. For more information about private CAs, see the AWS Certificate
	// Manager Private Certificate Authority (PCA)
	// (https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaWelcome.html) user
	// guide. The ARN must have the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	CertificateAuthorityArn *string

	// The domain name that you want ACM to use to send you emails so that you can
	// validate domain ownership.
	DomainValidationOptions []*types.DomainValidationOption

	// Customer chosen string that can be used to distinguish between calls to
	// RequestCertificate. Idempotency tokens time out after one hour. Therefore, if
	// you call RequestCertificate multiple times with the same idempotency token
	// within one hour, ACM recognizes that you are requesting only one certificate and
	// will issue only one. If you change the idempotency token for each call, ACM
	// recognizes that you are requesting multiple certificates.
	IdempotencyToken *string

	// Currently, you can use this parameter to specify whether to add the certificate
	// to a certificate transparency log. Certificate transparency makes it possible to
	// detect SSL/TLS certificates that have been mistakenly or maliciously issued.
	// Certificates that have not been logged typically produce an error message in a
	// browser. For more information, see Opting Out of Certificate Transparency
	// Logging
	// (https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency).
	Options *types.CertificateOptions

	// Additional FQDNs to be included in the Subject Alternative Name extension of the
	// ACM certificate. For example, add the name www.example.net to a certificate for
	// which the DomainName field is www.example.com if users can reach your site by
	// using either name. The maximum number of domain names that you can add to an ACM
	// certificate is 100. However, the initial quota is 10 domain names. If you need
	// more than 10 names, you must request a quota increase. For more information, see
	// Quotas (https://docs.aws.amazon.com/acm/latest/userguide/acm-limits.html). The
	// maximum length of a SAN DNS name is 253 octets. The name is made up of multiple
	// labels separated by periods. No label can be longer than 63 octets. Consider the
	// following examples:
	//
	// * (63 octets).(63 octets).(63 octets).(61 octets) is legal
	// because the total length is 253 octets (63+1+63+1+63+1+61) and no label exceeds
	// 63 octets.
	//
	// * (64 octets).(63 octets).(63 octets).(61 octets) is not legal
	// because the total length exceeds 253 octets (64+1+63+1+63+1+61) and the first
	// label exceeds 63 octets.
	//
	// * (63 octets).(63 octets).(63 octets).(62 octets) is
	// not legal because the total length of the DNS name (63+1+63+1+63+1+62) exceeds
	// 253 octets.
	SubjectAlternativeNames []*string

	// One or more resource tags to associate with the certificate.
	Tags []*types.Tag

	// The method you want to use if you are requesting a public certificate to
	// validate that you own or control domain. You can validate with DNS
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or
	// validate with email
	// (https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html).
	// We recommend that you use DNS validation.
	ValidationMethod types.ValidationMethod
}

type RequestCertificateOutput struct {

	// String that contains the ARN of the issued certificate. This must be of the
	// form:
	// arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
	CertificateArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRequestCertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRequestCertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRequestCertificate{}, middleware.After)
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
	addOpRequestCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRequestCertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRequestCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "RequestCertificate",
	}
}

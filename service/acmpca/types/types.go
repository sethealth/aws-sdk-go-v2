// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Contains information about the certificate subject. The certificate can be one
// issued by your private certificate authority (CA) or it can be your private CA
// certificate. The Subject field in the certificate identifies the entity that
// owns or controls the public key in the certificate. The entity can be a user,
// computer, device, or service. The Subject must contain an X.500 distinguished
// name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs
// are separated by commas in the certificate. The DN must be unique for each
// entity, but your private CA can issue more than one certificate with the same DN
// to the same entity.
type ASN1Subject struct {

	// Fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string

	// Two-digit code that specifies the country in which the certificate subject
	// located.
	Country *string

	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string

	// Typically a qualifier appended to the name of an individual. Examples include
	// Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string

	// First name.
	GivenName *string

	// Concatenation that typically contains the first letter of the GivenName, the
	// first letter of the middle name if one exists, and the first letter of the
	// SurName.
	Initials *string

	// The locality (such as a city or town) in which the certificate subject is
	// located.
	Locality *string

	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string

	// A subdivision or unit of the organization (such as sales or finance) with which
	// the certificate subject is affiliated.
	OrganizationalUnit *string

	// Typically a shortened version of a longer GivenName. For example, Jonathan is
	// often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string

	// The certificate serial number.
	SerialNumber *string

	// State in which the subject of the certificate is located.
	State *string

	// Family name. In the US and the UK, for example, the surname of an individual is
	// ordered last. In Asian cultures the surname is typically ordered first.
	Surname *string

	// A title such as Mr. or Ms., which is pre-pended to the name to refer formally to
	// the certificate subject.
	Title *string
}

// Contains information about your private certificate authority (CA). Your private
// CA can issue and revoke X.509 digital certificates. Digital certificates verify
// that the entity named in the certificate Subject field owns or controls the
// public key contained in the Subject Public Key Info field. Call the
// CreateCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
// action to create your private CA. You must then call the
// GetCertificateAuthorityCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_GetCertificateAuthorityCertificate.html)
// action to retrieve a private CA certificate signing request (CSR). Sign the CSR
// with your ACM Private CA-hosted or on-premises root or subordinate CA
// certificate. Call the ImportCertificateAuthorityCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html)
// action to import the signed certificate into AWS Certificate Manager (ACM).
type CertificateAuthority struct {

	// Amazon Resource Name (ARN) for your private certificate authority (CA). The
	// format is  12345678-1234-1234-1234-123456789012 .
	Arn *string

	// Your private CA configuration.
	CertificateAuthorityConfiguration *CertificateAuthorityConfiguration

	// Date and time at which your private CA was created.
	CreatedAt *time.Time

	// Reason the request to create your private CA failed.
	FailureReason FailureReason

	// Date and time at which your private CA was last updated.
	LastStateChangeAt *time.Time

	// Date and time after which your private CA certificate is not valid.
	NotAfter *time.Time

	// Date and time before which your private CA certificate is not valid.
	NotBefore *time.Time

	// The AWS account ID that owns the certificate authority.
	OwnerAccount *string

	// The period during which a deleted CA can be restored. For more information, see
	// the PermanentDeletionTimeInDays parameter of the
	// DeleteCertificateAuthorityRequest
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_DeleteCertificateAuthorityRequest.html)
	// action.
	RestorableUntil *time.Time

	// Information about the certificate revocation list (CRL) created and maintained
	// by your private CA.
	RevocationConfiguration *RevocationConfiguration

	// Serial number of your private CA.
	Serial *string

	// Status of your private CA.
	Status CertificateAuthorityStatus

	// Type of your private CA.
	Type CertificateAuthorityType
}

// Contains configuration information for your private certificate authority (CA).
// This includes information about the class of public key algorithm and the key
// pair that your private CA creates when it issues a certificate. It also includes
// the signature algorithm that it uses when issuing certificates, and its X.500
// distinguished name. You must specify this information when you call the
// CreateCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
// action.
type CertificateAuthorityConfiguration struct {

	// Type of the public key algorithm and size, in bits, of the key pair that your CA
	// creates when it issues a certificate. When you create a subordinate CA, you must
	// use a key algorithm supported by the parent CA.
	//
	// This member is required.
	KeyAlgorithm KeyAlgorithm

	// Name of the algorithm your private CA uses to sign certificate requests. This
	// parameter should not be confused with the SigningAlgorithm parameter used to
	// sign certificates when they are issued.
	//
	// This member is required.
	SigningAlgorithm SigningAlgorithm

	// Structure that contains X.500 distinguished name information for your private
	// CA.
	//
	// This member is required.
	Subject *ASN1Subject
}

// Contains configuration information for a certificate revocation list (CRL). Your
// private certificate authority (CA) creates base CRLs. Delta CRLs are not
// supported. You can enable CRLs for your new or an existing private CA by setting
// the Enabled parameter to true. Your private CA writes CRLs to an S3 bucket that
// you specify in the S3BucketName parameter. You can hide the name of your bucket
// by specifying a value for the CustomCname parameter. Your private CA copies the
// CNAME or the S3 bucket name to the CRL Distribution Points extension of each
// certificate it issues. Your S3 bucket policy must give write permission to ACM
// Private CA. ACM Private CAA assets that are stored in Amazon S3 can be protected
// with encryption. For more information, see Encrypting Your CRLs
// (https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaCreateCa.html#crl-encryption).
// Your private CA uses the value in the ExpirationInDays parameter to calculate
// the nextUpdate field in the CRL. The CRL is refreshed at 1/2 the age of next
// update or when a certificate is revoked. When a certificate is revoked, it is
// recorded in the next CRL that is generated and in the next audit report. Only
// time valid certificates are listed in the CRL. Expired certificates are not
// included. CRLs contain the following fields:
//
// * Version: The current version
// number defined in RFC 5280 is V2. The integer value is 0x1.
//
// * Signature
// Algorithm: The name of the algorithm used to sign the CRL.
//
// * Issuer: The X.500
// distinguished name of your private CA that issued the CRL.
//
// * Last Update: The
// issue date and time of this CRL.
//
// * Next Update: The day and time by which the
// next CRL will be issued.
//
// * Revoked Certificates: List of revoked certificates.
// Each list item contains the following information.
//
// * Serial Number: The serial
// number, in hexadecimal format, of the revoked certificate.
//
// * Revocation Date:
// Date and time the certificate was revoked.
//
// * CRL Entry Extensions: Optional
// extensions for the CRL entry.
//
// * X509v3 CRL Reason Code: Reason the certificate
// was revoked.
//
// * CRL Extensions: Optional extensions for the CRL.
//
// * X509v3
// Authority Key Identifier: Identifies the public key associated with the private
// key used to sign the certificate.
//
// * X509v3 CRL Number:: Decimal sequence number
// for the CRL.
//
// * Signature Algorithm: Algorithm used by your private CA to sign
// the CRL.
//
// * Signature Value: Signature computed over the CRL.
//
// Certificate
// revocation lists created by ACM Private CA are DER-encoded. You can use the
// following OpenSSL command to list a CRL. openssl crl -inform DER -text -in
// crl_path -noout
type CrlConfiguration struct {

	// Boolean value that specifies whether certificate revocation lists (CRLs) are
	// enabled. You can use this value to enable certificate revocation for a new CA
	// when you call the CreateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
	// action or for an existing CA when you call the UpdateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UpdateCertificateAuthority.html)
	// action.
	//
	// This member is required.
	Enabled *bool

	// Name inserted into the certificate CRL Distribution Points extension that
	// enables the use of an alias for the CRL distribution point. Use this value if
	// you don't want the name of your S3 bucket to be public.
	CustomCname *string

	// Number of days until a certificate expires.
	ExpirationInDays *int32

	// Name of the S3 bucket that contains the CRL. If you do not provide a value for
	// the CustomCname argument, the name of your S3 bucket is placed into the CRL
	// Distribution Points extension of the issued certificate. You can change the name
	// of your bucket by calling the UpdateCertificateAuthority
	// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UpdateCertificateAuthority.html)
	// action. You must specify a bucket policy that allows ACM Private CA to write the
	// CRL to your bucket.
	S3BucketName *string
}

// Permissions designate which private CA actions can be performed by an AWS
// service or entity. In order for ACM to automatically renew private certificates,
// you must give the ACM service principal all available permissions
// (IssueCertificate, GetCertificate, and ListPermissions). Permissions can be
// assigned with the CreatePermission
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreatePermission.html)
// action, removed with the DeletePermission
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_DeletePermission.html)
// action, and listed with the ListPermissions
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_ListPermissions.html)
// action.
type Permission struct {

	// The private CA actions that can be performed by the designated AWS service.
	Actions []ActionType

	// The Amazon Resource Number (ARN) of the private CA from which the permission was
	// issued.
	CertificateAuthorityArn *string

	// The time at which the permission was created.
	CreatedAt *time.Time

	// The name of the policy that is associated with the permission.
	Policy *string

	// The AWS service or entity that holds the permission. At this time, the only
	// valid principal is acm.amazonaws.com.
	Principal *string

	// The ID of the account that assigned the permission.
	SourceAccount *string
}

// Certificate revocation information used by the CreateCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_CreateCertificateAuthority.html)
// and UpdateCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UpdateCertificateAuthority.html)
// actions. Your private certificate authority (CA) can create and maintain a
// certificate revocation list (CRL). A CRL contains information about certificates
// revoked by your CA. For more information, see RevokeCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_RevokeCertificate.html).
type RevocationConfiguration struct {

	// Configuration of the certificate revocation list (CRL), if any, maintained by
	// your private CA.
	CrlConfiguration *CrlConfiguration
}

// Tags are labels that you can use to identify and organize your private CAs. Each
// tag consists of a key and an optional value. You can associate up to 50 tags
// with a private CA. To add one or more tags to a private CA, call the
// TagCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_TagCertificateAuthority.html)
// action. To remove a tag, call the UntagCertificateAuthority
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_UntagCertificateAuthority.html)
// action.
type Tag struct {

	// Key (name) of the tag.
	//
	// This member is required.
	Key *string

	// Value of the tag.
	Value *string
}

// Validity specifies the period of time during which a certificate is valid.
// Validity can be expressed as an explicit date and time when the certificate
// expires, or as a span of time after issuance, stated in days, months, or years.
// For more information, see Validity
// (https://tools.ietf.org/html/rfc5280#section-4.1.2.5) in RFC 5280. You can issue
// a certificate by calling the IssueCertificate
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_IssueCertificate.html)
// action.
type Validity struct {

	// Determines how ACM Private CA interprets the Value parameter, an integer.
	// Supported validity types include those listed below. Type definitions with
	// values include a sample input value and the resulting output. END_DATE: The
	// specific date and time when the certificate will expire, expressed using UTCTime
	// (YYMMDDHHMMSS) or GeneralizedTime (YYYYMMDDHHMMSS) format. When UTCTime is used,
	// if the year field (YY) is greater than or equal to 50, the year is interpreted
	// as 19YY. If the year field is less than 50, the year is interpreted as 20YY.
	//
	// *
	// Sample input value: 491231235959 (UTCTime format)
	//
	// * Output expiration
	// date/time: 12/31/2049 23:59:59
	//
	// ABSOLUTE: The specific date and time when the
	// certificate will expire, expressed in seconds since the Unix Epoch.
	//
	// * Sample
	// input value: 2524608000
	//
	// * Output expiration date/time: 01/01/2050
	// 00:00:00
	//
	// DAYS, MONTHS, YEARS: The relative time from the moment of issuance
	// until the certificate will expire, expressed in days, months, or years. Example
	// if DAYS, issued on 10/12/2020 at 12:34:54 UTC:
	//
	// * Sample input value: 90
	//
	// *
	// Output expiration date: 01/10/2020 12:34:54 UTC
	//
	// This member is required.
	Type ValidityPeriodType

	// A long integer interpreted according to the value of Type, below.
	//
	// This member is required.
	Value *int64
}

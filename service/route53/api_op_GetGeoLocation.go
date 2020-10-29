// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about whether a specified geographic location is supported for
// Amazon Route 53 geolocation resource record sets. Use the following syntax to
// determine whether a continent is supported for geolocation: GET
// /2013-04-01/geolocation?continentcode=two-letter abbreviation for a continent
// Use the following syntax to determine whether a country is supported for
// geolocation: GET /2013-04-01/geolocation?countrycode=two-character country code
// Use the following syntax to determine whether a subdivision of a country is
// supported for geolocation: GET /2013-04-01/geolocation?countrycode=two-character
// country code&subdivisioncode=subdivision code
func (c *Client) GetGeoLocation(ctx context.Context, params *GetGeoLocationInput, optFns ...func(*Options)) (*GetGeoLocationOutput, error) {
	if params == nil {
		params = &GetGeoLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGeoLocation", params, optFns, addOperationGetGeoLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGeoLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for information about whether a specified geographic location is
// supported for Amazon Route 53 geolocation resource record sets.
type GetGeoLocationInput struct {

	// For geolocation resource record sets, a two-letter abbreviation that identifies
	// a continent. Amazon Route 53 supports the following continent codes:
	//
	// * AF:
	// Africa
	//
	// * AN: Antarctica
	//
	// * AS: Asia
	//
	// * EU: Europe
	//
	// * OC: Oceania
	//
	// * NA: North
	// America
	//
	// * SA: South America
	ContinentCode *string

	// Amazon Route 53 uses the two-letter country codes that are specified in ISO
	// standard 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string

	// For SubdivisionCode, Amazon Route 53 supports only states of the United States.
	// For a list of state abbreviations, see Appendix B: Two–Letter State and
	// Possession Abbreviations (https://pe.usps.com/text/pub28/28apb.htm) on the
	// United States Postal Service website. If you specify subdivisioncode, you must
	// also specify US for CountryCode.
	SubdivisionCode *string
}

// A complex type that contains the response information for the specified
// geolocation code.
type GetGeoLocationOutput struct {

	// A complex type that contains the codes and full continent, country, and
	// subdivision names for the specified geolocation code.
	//
	// This member is required.
	GeoLocationDetails *types.GeoLocationDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetGeoLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetGeoLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetGeoLocation{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeoLocation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetGeoLocation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetGeoLocation",
	}
}

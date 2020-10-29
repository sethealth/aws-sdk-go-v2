// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// You can use the GetMetricWidgetImage API to retrieve a snapshot graph of one or
// more Amazon CloudWatch metrics as a bitmap image. You can then embed this image
// into your services and products, such as wiki pages, reports, and documents. You
// could also retrieve images regularly, such as every minute, and create your own
// custom live dashboard. The graph you retrieve can include all CloudWatch metric
// graph features, including metric math and horizontal and vertical annotations.
// There is a limit of 20 transactions per second for this API. Each
// GetMetricWidgetImage action has the following limits:
//
// * As many as 100 metrics
// in the graph.
//
// * Up to 100 KB uncompressed payload.
func (c *Client) GetMetricWidgetImage(ctx context.Context, params *GetMetricWidgetImageInput, optFns ...func(*Options)) (*GetMetricWidgetImageOutput, error) {
	if params == nil {
		params = &GetMetricWidgetImageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMetricWidgetImage", params, optFns, addOperationGetMetricWidgetImageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMetricWidgetImageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMetricWidgetImageInput struct {

	// A JSON string that defines the bitmap graph to be retrieved. The string includes
	// the metrics to include in the graph, statistics, annotations, title, axis
	// limits, and so on. You can include only one MetricWidget parameter in each
	// GetMetricWidgetImage call. For more information about the syntax of MetricWidget
	// see GetMetricWidgetImage: Metric Widget Structure and Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/CloudWatch-Metric-Widget-Structure.html).
	// If any metric on the graph could not load all the requested data points, an
	// orange triangle with an exclamation point appears next to the graph legend.
	//
	// This member is required.
	MetricWidget *string

	// The format of the resulting image. Only PNG images are supported. The default is
	// png. If you specify png, the API returns an HTTP response with the content-type
	// set to text/xml. The image data is in a MetricWidgetImage field. For example:
	// <GetMetricWidgetImageResponse xmlns=>
	//
	//
	// iVBORw0KGgoAAAANSUhEUgAAAlgAAAGQEAYAAAAip...
	//
	//
	// 6f0d4192-4d42-11e8-82c1-f539a07e0e3b
	//
	// The image/png setting is intended only for
	// custom HTTP requests. For most use cases, and all actions using an AWS SDK, you
	// should use png. If you specify image/png, the HTTP response has a content-type
	// set to image/png, and the body of the response is a PNG image.
	OutputFormat *string
}

type GetMetricWidgetImageOutput struct {

	// The image of the graph, in the output format specified. The output is
	// base64-encoded.
	MetricWidgetImage []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMetricWidgetImageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetMetricWidgetImage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetMetricWidgetImage{}, middleware.After)
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
	addOpGetMetricWidgetImageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMetricWidgetImage(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMetricWidgetImage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "GetMetricWidgetImage",
	}
}

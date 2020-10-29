// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the tables in a database. If neither SchemaPattern nor TablePattern are
// specified, then all tables in the database are returned. A token is returned to
// page through the table list. Depending on the authorization method, use one of
// the following combinations of request parameters:
//
// * AWS Secrets Manager -
// specify the Amazon Resource Name (ARN) of the secret and the cluster identifier
// that matches the cluster in the secret.
//
// * Temporary credentials - specify the
// cluster identifier, the database name, and the database user name. Permission to
// call the redshift:GetClusterCredentials operation is required to use this
// method.
func (c *Client) ListTables(ctx context.Context, params *ListTablesInput, optFns ...func(*Options)) (*ListTablesOutput, error) {
	if params == nil {
		params = &ListTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTables", params, optFns, addOperationListTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTablesInput struct {

	// The cluster identifier. This parameter is required when authenticating using
	// either AWS Secrets Manager or temporary credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database. This parameter is required when authenticating using
	// temporary credentials.
	//
	// This member is required.
	Database *string

	// The database user name. This parameter is required when authenticating using
	// temporary credentials.
	DbUser *string

	// The maximum number of tables to return in the response. If more tables exist
	// than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults *int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// A pattern to filter results by schema name. Within a schema pattern, "%" means
	// match any substring of 0 or more characters and "_" means match any one
	// character. Only schema name entries matching the search pattern are returned. If
	// SchemaPattern is not specified, then all tables that match TablePattern are
	// returned. If neither SchemaPattern or TablePattern are specified, then all
	// tables are returned.
	SchemaPattern *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using AWS Secrets Manager.
	SecretArn *string

	// A pattern to filter results by table name. Within a table pattern, "%" means
	// match any substring of 0 or more characters and "_" means match any one
	// character. Only table name entries matching the search pattern are returned. If
	// TablePattern is not specified, then all tables that match SchemaPatternare
	// returned. If neither SchemaPattern or TablePattern are specified, then all
	// tables are returned.
	TablePattern *string
}

type ListTablesOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The tables that match the request pattern.
	Tables []*types.TableMember

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTables{}, middleware.After)
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
	addOpListTablesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTables(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ListTables",
	}
}

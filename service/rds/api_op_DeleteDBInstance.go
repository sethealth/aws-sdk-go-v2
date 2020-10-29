// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The DeleteDBInstance action deletes a previously provisioned DB instance. When
// you delete a DB instance, all automated backups for that instance are deleted
// and can't be recovered. Manual DB snapshots of the DB instance to be deleted by
// DeleteDBInstance are not deleted. If you request a final DB snapshot the status
// of the Amazon RDS DB instance is deleting until the DB snapshot is created. The
// API action DescribeDBInstance is used to monitor the status of this operation.
// The action can't be canceled or reverted once submitted. When a DB instance is
// in a failure state and has a status of failed, incompatible-restore, or
// incompatible-network, you can only delete it when you skip creation of the final
// snapshot with the SkipFinalSnapshot parameter. If the specified DB instance is
// part of an Amazon Aurora DB cluster, you can't delete the DB instance if both of
// the following conditions are true:
//
// * The DB cluster is a read replica of
// another Amazon Aurora DB cluster.
//
// * The DB instance is the only instance in the
// DB cluster.
//
// To delete a DB instance in this case, first call the
// PromoteReadReplicaDBCluster API action to promote the DB cluster so it's no
// longer a read replica. After the promotion completes, then call the
// DeleteDBInstance API action to delete the final instance in the DB cluster.
func (c *Client) DeleteDBInstance(ctx context.Context, params *DeleteDBInstanceInput, optFns ...func(*Options)) (*DeleteDBInstanceOutput, error) {
	if params == nil {
		params = &DeleteDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBInstance", params, optFns, addOperationDeleteDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteDBInstanceInput struct {

	// The DB instance identifier for the DB instance to be deleted. This parameter
	// isn't case-sensitive. Constraints:
	//
	// * Must match the name of an existing DB
	// instance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// A value that indicates whether to remove automated backups immediately after the
	// DB instance is deleted. This parameter isn't case-sensitive. The default is to
	// remove automated backups immediately after the DB instance is deleted.
	DeleteAutomatedBackups *bool

	// The DBSnapshotIdentifier of the new DBSnapshot created when the
	// SkipFinalSnapshot parameter is disabled. Specifying this parameter and also
	// specifying to skip final DB snapshot creation in SkipFinalShapshot results in an
	// error. Constraints:
	//
	// * Must be 1 to 255 letters or numbers.
	//
	// * First character
	// must be a letter.
	//
	// * Can't end with a hyphen or contain two consecutive
	// hyphens.
	//
	// * Can't be specified when deleting a read replica.
	FinalDBSnapshotIdentifier *string

	// A value that indicates whether to skip the creation of a final DB snapshot
	// before the DB instance is deleted. If skip is specified, no DB snapshot is
	// created. If skip isn't specified, a DB snapshot is created before the DB
	// instance is deleted. By default, skip isn't specified, and the DB snapshot is
	// created. When a DB instance is in a failure state and has a status of 'failed',
	// 'incompatible-restore', or 'incompatible-network', it can only be deleted when
	// skip is specified. Specify skip when deleting a read replica. The
	// FinalDBSnapshotIdentifier parameter must be specified if skip isn't specified.
	SkipFinalSnapshot *bool
}

type DeleteDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the DescribeDBInstances action.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBInstance{}, middleware.After)
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
	addOpDeleteDBInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteDBInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBInstance",
	}
}

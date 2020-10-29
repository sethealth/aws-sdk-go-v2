// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The access level to use to filter results.
type AccessLevelFilter struct {

	// The access level.
	//
	// * Account - Filter results based on the account.
	//
	// * Role -
	// Filter results based on the federated role of the specified user.
	//
	// * User -
	// Filter results based on the specified user.
	Key AccessLevelFilterKey

	// The user to which the access level applies. The only supported value is Self.
	Value *string
}

// Information about a budget.
type BudgetDetail struct {

	// Name of the associated budget.
	BudgetName *string
}

// Information about a CloudWatch dashboard.
type CloudWatchDashboard struct {

	// The name of the CloudWatch dashboard.
	Name *string
}

// Information about a constraint.
type ConstraintDetail struct {

	// The identifier of the constraint.
	ConstraintId *string

	// The description of the constraint.
	Description *string

	// The owner of the constraint.
	Owner *string

	// The identifier of the portfolio the product resides in. The constraint applies
	// only to the instance of the product that lives within this portfolio.
	PortfolioId *string

	// The identifier of the product the constraint applies to. Note that a constraint
	// applies to a specific instance of a product within a certain portfolio.
	ProductId *string

	// The type of constraint.
	//
	// * LAUNCH
	//
	// * NOTIFICATION
	//
	// * STACKSET
	//
	// * TEMPLATE
	Type *string
}

// Summary information about a constraint.
type ConstraintSummary struct {

	// The description of the constraint.
	Description *string

	// The type of constraint.
	//
	// * LAUNCH
	//
	// * NOTIFICATION
	//
	// * STACKSET
	//
	// * TEMPLATE
	Type *string
}

// Details of an execution parameter value that is passed to a self-service action
// when executed on a provisioned product.
type ExecutionParameter struct {

	// The default values for the execution parameter.
	DefaultValues []*string

	// The name of the execution parameter.
	Name *string

	// The execution parameter type.
	Type *string
}

// An object containing information about the error, along with identifying
// information about the self-service action and its associations.
type FailedServiceActionAssociation struct {

	// The error code. Valid values are listed below.
	ErrorCode ServiceActionAssociationErrorCode

	// A text description of the error.
	ErrorMessage *string

	// The product identifier. For example, prod-abcdzk7xy33qa.
	ProductId *string

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	ProvisioningArtifactId *string

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	ServiceActionId *string
}

// A launch path object.
type LaunchPath struct {

	// The identifier of the launch path.
	Id *string

	// The name of the launch path.
	Name *string
}

// Summary information about a product path for a user.
type LaunchPathSummary struct {

	// The constraints on the portfolio-product relationship.
	ConstraintSummaries []*ConstraintSummary

	// The identifier of the product path.
	Id *string

	// The name of the portfolio to which the user was assigned.
	Name *string

	// The tags associated with this product path.
	Tags []*Tag
}

// The search filter to use when listing history records.
type ListRecordHistorySearchFilter struct {

	// The filter key.
	//
	// * product - Filter results based on the specified product
	// identifier.
	//
	// * provisionedproduct - Filter results based on the provisioned
	// product identifier.
	Key *string

	// The filter value.
	Value *string
}

// Filters to use when listing TagOptions.
type ListTagOptionsFilters struct {

	// The active state.
	Active *bool

	// The TagOption key.
	Key *string

	// The TagOption value.
	Value *string
}

// Information about the organization node.
type OrganizationNode struct {

	// The organization node type.
	Type OrganizationNodeType

	// The identifier of the organization node.
	Value *string
}

// The constraints that the administrator has put on the parameter.
type ParameterConstraints struct {

	// The values that the administrator has allowed for the parameter.
	AllowedValues []*string
}

// Information about a portfolio.
type PortfolioDetail struct {

	// The ARN assigned to the portfolio.
	ARN *string

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The description of the portfolio.
	Description *string

	// The name to use for display purposes.
	DisplayName *string

	// The portfolio identifier.
	Id *string

	// The name of the portfolio provider.
	ProviderName *string
}

// Information about a principal.
type Principal struct {

	// The ARN of the principal (IAM user, role, or group).
	PrincipalARN *string

	// The principal type. The supported value is IAM.
	PrincipalType PrincipalType
}

// A single product view aggregation value/count pair, containing metadata about
// each product to which the calling user has access.
type ProductViewAggregationValue struct {

	// An approximate count of the products that match the value.
	ApproximateCount *int32

	// The value of the product view aggregation.
	Value *string
}

// Information about a product view.
type ProductViewDetail struct {

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The ARN of the product.
	ProductARN *string

	// Summary information about the product view.
	ProductViewSummary *ProductViewSummary

	// The status of the product.
	//
	// * AVAILABLE - The product is ready for use.
	//
	// *
	// CREATING - Product creation has started; the product is not ready for use.
	//
	// *
	// FAILED - An action failed.
	Status Status
}

// Summary information about a product view.
type ProductViewSummary struct {

	// The distributor of the product. Contact the product administrator for the
	// significance of this value.
	Distributor *string

	// Indicates whether the product has a default path. If the product does not have a
	// default path, call ListLaunchPaths to disambiguate between paths. Otherwise,
	// ListLaunchPaths is not required, and the output of ProductViewSummary can be
	// used directly with DescribeProvisioningParameters.
	HasDefaultPath *bool

	// The product view identifier.
	Id *string

	// The name of the product.
	Name *string

	// The owner of the product. Contact the product administrator for the significance
	// of this value.
	Owner *string

	// The product identifier.
	ProductId *string

	// Short description of the product.
	ShortDescription *string

	// The description of the support for this Product.
	SupportDescription *string

	// The email contact information to obtain support for this Product.
	SupportEmail *string

	// The URL information to obtain support for this Product.
	SupportUrl *string

	// The product type. Contact the product administrator for the significance of this
	// value. If this value is MARKETPLACE, the product was created by AWS Marketplace.
	Type ProductType
}

// Information about a provisioned product.
type ProvisionedProductAttribute struct {

	// The ARN of the provisioned product.
	Arn *string

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The identifier of the provisioned product.
	Id *string

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	IdempotencyToken *string

	// The record identifier of the last request performed on this provisioned product
	// of the following types:
	//
	// * ProvisionedProduct
	//
	// * UpdateProvisionedProduct
	//
	// *
	// ExecuteProvisionedProductPlan
	//
	// * TerminateProvisionedProduct
	LastProvisioningRecordId *string

	// The record identifier of the last request performed on this provisioned product.
	LastRecordId *string

	// The record identifier of the last successful request performed on this
	// provisioned product of the following types:
	//
	// * ProvisionedProduct
	//
	// *
	// UpdateProvisionedProduct
	//
	// * ExecuteProvisionedProductPlan
	//
	// *
	// TerminateProvisionedProduct
	LastSuccessfulProvisioningRecordId *string

	// The user-friendly name of the provisioned product.
	Name *string

	// The assigned identifier for the resource, such as an EC2 instance ID or an S3
	// bucket name.
	PhysicalId *string

	// The product identifier.
	ProductId *string

	// The name of the product.
	ProductName *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string

	// The name of the provisioning artifact.
	ProvisioningArtifactName *string

	// The current status of the provisioned product.
	//
	// * AVAILABLE - Stable state,
	// ready to perform any operation. The most recent operation succeeded and
	// completed.
	//
	// * UNDER_CHANGE - Transitive state. Operations performed might not
	// have valid results. Wait for an AVAILABLE status before performing
	// operations.
	//
	// * TAINTED - Stable state, ready to perform any operation. The stack
	// has completed the requested operation but is not exactly what was requested. For
	// example, a request to update to a new version failed and the stack rolled back
	// to the current version.
	//
	// * ERROR - An unexpected error occurred. The provisioned
	// product exists but the stack is not running. For example, CloudFormation
	// received a parameter value that was not valid and could not launch the stack.
	//
	// *
	// PLAN_IN_PROGRESS - Transitive state. The plan operations were performed to
	// provision a new product, but resources have not yet been created. After
	// reviewing the list of resources to be created, execute the plan. Wait for an
	// AVAILABLE status before performing operations.
	Status ProvisionedProductStatus

	// The current status message of the provisioned product.
	StatusMessage *string

	// One or more tags.
	Tags []*Tag

	// The type of provisioned product. The supported values are CFN_STACK and
	// CFN_STACKSET.
	Type *string

	// The Amazon Resource Name (ARN) of the IAM user.
	UserArn *string

	// The ARN of the IAM user in the session. This ARN might contain a session ID.
	UserArnSession *string
}

// Information about a provisioned product.
type ProvisionedProductDetail struct {

	// The ARN of the provisioned product.
	Arn *string

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The identifier of the provisioned product.
	Id *string

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	IdempotencyToken *string

	// The record identifier of the last request performed on this provisioned product
	// of the following types:
	//
	// * ProvisionedProduct
	//
	// * UpdateProvisionedProduct
	//
	// *
	// ExecuteProvisionedProductPlan
	//
	// * TerminateProvisionedProduct
	LastProvisioningRecordId *string

	// The record identifier of the last request performed on this provisioned product.
	LastRecordId *string

	// The record identifier of the last successful request performed on this
	// provisioned product of the following types:
	//
	// * ProvisionedProduct
	//
	// *
	// UpdateProvisionedProduct
	//
	// * ExecuteProvisionedProductPlan
	//
	// *
	// TerminateProvisionedProduct
	LastSuccessfulProvisioningRecordId *string

	// The ARN of the launch role associated with the provisioned product.
	LaunchRoleArn *string

	// The user-friendly name of the provisioned product.
	Name *string

	// The product identifier. For example, prod-abcdzk7xy33qa.
	ProductId *string

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	ProvisioningArtifactId *string

	// The current status of the provisioned product.
	//
	// * AVAILABLE - Stable state,
	// ready to perform any operation. The most recent operation succeeded and
	// completed.
	//
	// * UNDER_CHANGE - Transitive state. Operations performed might not
	// have valid results. Wait for an AVAILABLE status before performing
	// operations.
	//
	// * TAINTED - Stable state, ready to perform any operation. The stack
	// has completed the requested operation but is not exactly what was requested. For
	// example, a request to update to a new version failed and the stack rolled back
	// to the current version.
	//
	// * ERROR - An unexpected error occurred. The provisioned
	// product exists but the stack is not running. For example, CloudFormation
	// received a parameter value that was not valid and could not launch the stack.
	//
	// *
	// PLAN_IN_PROGRESS - Transitive state. The plan operations were performed to
	// provision a new product, but resources have not yet been created. After
	// reviewing the list of resources to be created, execute the plan. Wait for an
	// AVAILABLE status before performing operations.
	Status ProvisionedProductStatus

	// The current status message of the provisioned product.
	StatusMessage *string

	// The type of provisioned product. The supported values are CFN_STACK and
	// CFN_STACKSET.
	Type *string
}

// Information about a plan.
type ProvisionedProductPlanDetails struct {

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// Passed to CloudFormation. The SNS topic ARNs to which to publish stack-related
	// events.
	NotificationArns []*string

	// The path identifier of the product. This value is optional if the product has a
	// default path, and required if the product has more than one path. To list the
	// paths for a product, use ListLaunchPaths.
	PathId *string

	// The plan identifier.
	PlanId *string

	// The name of the plan.
	PlanName *string

	// The plan type.
	PlanType ProvisionedProductPlanType

	// The product identifier.
	ProductId *string

	// The product identifier.
	ProvisionProductId *string

	// The user-friendly name of the provisioned product.
	ProvisionProductName *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string

	// Parameters specified by the administrator that are required for provisioning the
	// product.
	ProvisioningParameters []*UpdateProvisioningParameter

	// The status.
	Status ProvisionedProductPlanStatus

	// The status message.
	StatusMessage *string

	// One or more tags.
	Tags []*Tag

	// The time when the plan was last updated.
	UpdatedTime *time.Time
}

// Summary information about a plan.
type ProvisionedProductPlanSummary struct {

	// The plan identifier.
	PlanId *string

	// The name of the plan.
	PlanName *string

	// The plan type.
	PlanType ProvisionedProductPlanType

	// The product identifier.
	ProvisionProductId *string

	// The user-friendly name of the provisioned product.
	ProvisionProductName *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string
}

// Information about a provisioning artifact. A provisioning artifact is also known
// as a product version.
type ProvisioningArtifact struct {

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The description of the provisioning artifact.
	Description *string

	// Information set by the administrator to provide guidance to end users about
	// which provisioning artifacts to use.
	Guidance ProvisioningArtifactGuidance

	// The identifier of the provisioning artifact.
	Id *string

	// The name of the provisioning artifact.
	Name *string
}

// Information about a provisioning artifact (also known as a version) for a
// product.
type ProvisioningArtifactDetail struct {

	// Indicates whether the product version is active.
	Active *bool

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The description of the provisioning artifact.
	Description *string

	// Information set by the administrator to provide guidance to end users about
	// which provisioning artifacts to use.
	Guidance ProvisioningArtifactGuidance

	// The identifier of the provisioning artifact.
	Id *string

	// The name of the provisioning artifact.
	Name *string

	// The type of provisioning artifact.
	//
	// * CLOUD_FORMATION_TEMPLATE - AWS
	// CloudFormation template
	//
	// * MARKETPLACE_AMI - AWS Marketplace AMI
	//
	// *
	// MARKETPLACE_CAR - AWS Marketplace Clusters and AWS Resources
	Type ProvisioningArtifactType
}

// Provisioning artifact output.
type ProvisioningArtifactOutput struct {

	// Description of the provisioning artifact output key.
	Description *string

	// The provisioning artifact output key.
	Key *string
}

// Information about a parameter used to provision a product.
type ProvisioningArtifactParameter struct {

	// The default value.
	DefaultValue *string

	// The description of the parameter.
	Description *string

	// If this value is true, the value for this parameter is obfuscated from view when
	// the parameter is retrieved. This parameter is used to hide sensitive
	// information.
	IsNoEcho *bool

	// Constraints that the administrator has put on a parameter.
	ParameterConstraints *ParameterConstraints

	// The parameter key.
	ParameterKey *string

	// The parameter type.
	ParameterType *string
}

// The user-defined preferences that will be applied during product provisioning,
// unless overridden by ProvisioningPreferences or UpdateProvisioningPreferences.
// For more information on maximum concurrent accounts and failure tolerance, see
// Stack set operation options
// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-concepts.html#stackset-ops-options)
// in the AWS CloudFormation User Guide.
type ProvisioningArtifactPreferences struct {

	// One or more AWS accounts where stack instances are deployed from the stack set.
	// These accounts can be scoped in ProvisioningPreferences$StackSetAccounts and
	// UpdateProvisioningPreferences$StackSetAccounts. Applicable only to a
	// CFN_STACKSET provisioned product type.
	StackSetAccounts []*string

	// One or more AWS Regions where stack instances are deployed from the stack set.
	// These regions can be scoped in ProvisioningPreferences$StackSetRegions and
	// UpdateProvisioningPreferences$StackSetRegions. Applicable only to a CFN_STACKSET
	// provisioned product type.
	StackSetRegions []*string
}

// Information about a provisioning artifact (also known as a version) for a
// product.
type ProvisioningArtifactProperties struct {

	// The URL of the CloudFormation template in Amazon S3. Specify the URL in JSON
	// format as follows: "LoadTemplateFromURL":
	// "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/..."
	//
	// This member is required.
	Info map[string]*string

	// The description of the provisioning artifact, including how it differs from the
	// previous provisioning artifact.
	Description *string

	// If set to true, AWS Service Catalog stops validating the specified provisioning
	// artifact even if it is invalid.
	DisableTemplateValidation *bool

	// The name of the provisioning artifact (for example, v1 v2beta). No spaces are
	// allowed.
	Name *string

	// The type of provisioning artifact.
	//
	// * CLOUD_FORMATION_TEMPLATE - AWS
	// CloudFormation template
	//
	// * MARKETPLACE_AMI - AWS Marketplace AMI
	//
	// *
	// MARKETPLACE_CAR - AWS Marketplace Clusters and AWS Resources
	Type ProvisioningArtifactType
}

// Summary information about a provisioning artifact (also known as a version) for
// a product.
type ProvisioningArtifactSummary struct {

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The description of the provisioning artifact.
	Description *string

	// The identifier of the provisioning artifact.
	Id *string

	// The name of the provisioning artifact.
	Name *string

	// The metadata for the provisioning artifact. This is used with AWS Marketplace
	// products.
	ProvisioningArtifactMetadata map[string]*string
}

// An object that contains summary information about a product view and a
// provisioning artifact.
type ProvisioningArtifactView struct {

	// Summary information about a product view.
	ProductViewSummary *ProductViewSummary

	// Information about a provisioning artifact. A provisioning artifact is also known
	// as a product version.
	ProvisioningArtifact *ProvisioningArtifact
}

// Information about a parameter used to provision a product.
type ProvisioningParameter struct {

	// The parameter key.
	Key *string

	// The parameter value.
	Value *string
}

// The user-defined preferences that will be applied when updating a provisioned
// product. Not all preferences are applicable to all provisioned product types.
type ProvisioningPreferences struct {

	// One or more AWS accounts that will have access to the provisioned product.
	// Applicable only to a CFN_STACKSET provisioned product type. The AWS accounts
	// specified should be within the list of accounts in the STACKSET constraint. To
	// get the list of accounts in the STACKSET constraint, use the
	// DescribeProvisioningParameters operation. If no values are specified, the
	// default value is all accounts from the STACKSET constraint.
	StackSetAccounts []*string

	// The number of accounts, per region, for which this operation can fail before AWS
	// Service Catalog stops the operation in that region. If the operation is stopped
	// in a region, AWS Service Catalog doesn't attempt the operation in any subsequent
	// regions. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetFailureToleranceCount or
	// StackSetFailureTolerancePercentage, but not both. The default value is 0 if no
	// value is specified.
	StackSetFailureToleranceCount *int32

	// The percentage of accounts, per region, for which this stack operation can fail
	// before AWS Service Catalog stops the operation in that region. If the operation
	// is stopped in a region, AWS Service Catalog doesn't attempt the operation in any
	// subsequent regions. When calculating the number of accounts based on the
	// specified percentage, AWS Service Catalog rounds down to the next whole number.
	// Applicable only to a CFN_STACKSET provisioned product type. Conditional: You
	// must specify either StackSetFailureToleranceCount or
	// StackSetFailureTolerancePercentage, but not both.
	StackSetFailureTolerancePercentage *int32

	// The maximum number of accounts in which to perform this operation at one time.
	// This is dependent on the value of StackSetFailureToleranceCount.
	// StackSetMaxConcurrentCount is at most one more than the
	// StackSetFailureToleranceCount. Note that this setting lets you specify the
	// maximum for operations. For large deployments, under certain circumstances the
	// actual number of accounts acted upon concurrently may be lower due to service
	// throttling. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetMaxConcurrentCount or
	// StackSetMaxConcurrentPercentage, but not both.
	StackSetMaxConcurrencyCount *int32

	// The maximum percentage of accounts in which to perform this operation at one
	// time. When calculating the number of accounts based on the specified percentage,
	// AWS Service Catalog rounds down to the next whole number. This is true except in
	// cases where rounding down would result is zero. In this case, AWS Service
	// Catalog sets the number as 1 instead. Note that this setting lets you specify
	// the maximum for operations. For large deployments, under certain circumstances
	// the actual number of accounts acted upon concurrently may be lower due to
	// service throttling. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetMaxConcurrentCount or
	// StackSetMaxConcurrentPercentage, but not both.
	StackSetMaxConcurrencyPercentage *int32

	// One or more AWS Regions where the provisioned product will be available.
	// Applicable only to a CFN_STACKSET provisioned product type. The specified
	// regions should be within the list of regions from the STACKSET constraint. To
	// get the list of regions in the STACKSET constraint, use the
	// DescribeProvisioningParameters operation. If no values are specified, the
	// default value is all regions from the STACKSET constraint.
	StackSetRegions []*string
}

// Information about a request operation.
type RecordDetail struct {

	// The UTC time stamp of the creation time.
	CreatedTime *time.Time

	// The ARN of the launch role associated with the provisioned product.
	LaunchRoleArn *string

	// The path identifier.
	PathId *string

	// The product identifier.
	ProductId *string

	// The identifier of the provisioned product.
	ProvisionedProductId *string

	// The user-friendly name of the provisioned product.
	ProvisionedProductName *string

	// The type of provisioned product. The supported values are CFN_STACK and
	// CFN_STACKSET.
	ProvisionedProductType *string

	// The identifier of the provisioning artifact.
	ProvisioningArtifactId *string

	// The errors that occurred.
	RecordErrors []*RecordError

	// The identifier of the record.
	RecordId *string

	// One or more tags.
	RecordTags []*RecordTag

	// The record type.
	//
	// * PROVISION_PRODUCT
	//
	// * UPDATE_PROVISIONED_PRODUCT
	//
	// *
	// TERMINATE_PROVISIONED_PRODUCT
	RecordType *string

	// The status of the provisioned product.
	//
	// * CREATED - The request was created but
	// the operation has not started.
	//
	// * IN_PROGRESS - The requested operation is in
	// progress.
	//
	// * IN_PROGRESS_IN_ERROR - The provisioned product is under change but
	// the requested operation failed and some remediation is occurring. For example, a
	// rollback.
	//
	// * SUCCEEDED - The requested operation has successfully completed.
	//
	// *
	// FAILED - The requested operation has unsuccessfully completed. Investigate using
	// the error messages returned.
	Status RecordStatus

	// The time when the record was last updated.
	UpdatedTime *time.Time
}

// The error code and description resulting from an operation.
type RecordError struct {

	// The numeric value of the error.
	Code *string

	// The description of the error.
	Description *string
}

// The output for the product created as the result of a request. For example, the
// output for a CloudFormation-backed product that creates an S3 bucket would
// include the S3 bucket URL.
type RecordOutput struct {

	// The description of the output.
	Description *string

	// The output key.
	OutputKey *string

	// The output value.
	OutputValue *string
}

// Information about a tag, which is a key-value pair.
type RecordTag struct {

	// The key for this tag.
	Key *string

	// The value for this tag.
	Value *string
}

// Information about a resource change that will occur when a plan is executed.
type ResourceChange struct {

	// The change action.
	Action ChangeAction

	// Information about the resource changes.
	Details []*ResourceChangeDetail

	// The ID of the resource, as defined in the CloudFormation template.
	LogicalResourceId *string

	// The ID of the resource, if it was already created.
	PhysicalResourceId *string

	// If the change type is Modify, indicates whether the existing resource is deleted
	// and replaced with a new one.
	Replacement Replacement

	// The type of resource.
	ResourceType *string

	// The change scope.
	Scope []ResourceAttribute
}

// Information about a change to a resource attribute.
type ResourceChangeDetail struct {

	// The ID of the entity that caused the change.
	CausingEntity *string

	// For static evaluations, the value of the resource attribute will change and the
	// new value is known. For dynamic evaluations, the value might change, and any new
	// value will be determined when the plan is updated.
	Evaluation EvaluationType

	// Information about the resource attribute to be modified.
	Target *ResourceTargetDefinition
}

// Information about a resource.
type ResourceDetail struct {

	// The ARN of the resource.
	ARN *string

	// The creation time of the resource.
	CreatedTime *time.Time

	// The description of the resource.
	Description *string

	// The identifier of the resource.
	Id *string

	// The name of the resource.
	Name *string
}

// Information about a change to a resource attribute.
type ResourceTargetDefinition struct {

	// The attribute to be changed.
	Attribute ResourceAttribute

	// If the attribute is Properties, the value is the name of the property.
	// Otherwise, the value is null.
	Name *string

	// If the attribute is Properties, indicates whether a change to this property
	// causes the resource to be re-created.
	RequiresRecreation RequiresRecreation
}

// A self-service action association consisting of the Action ID, the Product ID,
// and the Provisioning Artifact ID.
type ServiceActionAssociation struct {

	// The product identifier. For example, prod-abcdzk7xy33qa.
	//
	// This member is required.
	ProductId *string

	// The identifier of the provisioning artifact. For example, pa-4abcdjnxjj6ne.
	//
	// This member is required.
	ProvisioningArtifactId *string

	// The self-service action identifier. For example, act-fs7abcd89wxyz.
	//
	// This member is required.
	ServiceActionId *string
}

// An object containing detailed information about the self-service action.
type ServiceActionDetail struct {

	// A map that defines the self-service action.
	Definition map[string]*string

	// Summary information about the self-service action.
	ServiceActionSummary *ServiceActionSummary
}

// Detailed information about the self-service action.
type ServiceActionSummary struct {

	// The self-service action definition type. For example, SSM_AUTOMATION.
	DefinitionType ServiceActionDefinitionType

	// The self-service action description.
	Description *string

	// The self-service action identifier.
	Id *string

	// The self-service action name.
	Name *string
}

// Information about the portfolio share operation.
type ShareDetails struct {

	// List of errors.
	ShareErrors []*ShareError

	// List of accounts for whom the operation succeeded.
	SuccessfulShares []*string
}

// Errors that occurred during the portfolio share operation.
type ShareError struct {

	// List of accounts impacted by the error.
	Accounts []*string

	// Error type that happened when processing the operation.
	Error *string

	// Information about the error.
	Message *string
}

// An AWS CloudFormation stack, in a specific account and region, that's part of a
// stack set operation. A stack instance is a reference to an attempted or actual
// stack in a given account within a given region. A stack instance can exist
// without a stack—for example, if the stack couldn't be created for some reason. A
// stack instance is associated with only one stack set. Each stack instance
// contains the ID of its associated stack set, as well as the ID of the actual
// stack and the stack status.
type StackInstance struct {

	// The name of the AWS account that the stack instance is associated with.
	Account *string

	// The name of the AWS region that the stack instance is associated with.
	Region *string

	// The status of the stack instance, in terms of its synchronization with its
	// associated stack set.
	//
	// * INOPERABLE: A DeleteStackInstances operation has failed
	// and left the stack in an unstable state. Stacks in this state are excluded from
	// further UpdateStackSet operations. You might need to perform a
	// DeleteStackInstances operation, with RetainStacks set to true, to delete the
	// stack instance, and then delete the stack manually.
	//
	// * OUTDATED: The stack isn't
	// currently up to date with the stack set because either the associated stack
	// failed during a CreateStackSet or UpdateStackSet operation, or the stack was
	// part of a CreateStackSet or UpdateStackSet operation that failed or was stopped
	// before the stack was created or updated.
	//
	// * CURRENT: The stack is currently up
	// to date with the stack set.
	StackInstanceStatus StackInstanceStatus
}

// Information about a tag. A tag is a key-value pair. Tags are propagated to the
// resources created when provisioning a product.
type Tag struct {

	// The tag key.
	//
	// This member is required.
	Key *string

	// The value for this key.
	//
	// This member is required.
	Value *string
}

// Information about a TagOption.
type TagOptionDetail struct {

	// The TagOption active state.
	Active *bool

	// The TagOption identifier.
	Id *string

	// The TagOption key.
	Key *string

	// The TagOption value.
	Value *string
}

// Summary information about a TagOption.
type TagOptionSummary struct {

	// The TagOption key.
	Key *string

	// The TagOption value.
	Values []*string
}

// The parameter key-value pair used to update a provisioned product.
type UpdateProvisioningParameter struct {

	// The parameter key.
	Key *string

	// If set to true, Value is ignored and the previous parameter value is kept.
	UsePreviousValue *bool

	// The parameter value.
	Value *string
}

// The user-defined preferences that will be applied when updating a provisioned
// product. Not all preferences are applicable to all provisioned product types.
type UpdateProvisioningPreferences struct {

	// One or more AWS accounts that will have access to the provisioned product.
	// Applicable only to a CFN_STACKSET provisioned product type. The AWS accounts
	// specified should be within the list of accounts in the STACKSET constraint. To
	// get the list of accounts in the STACKSET constraint, use the
	// DescribeProvisioningParameters operation. If no values are specified, the
	// default value is all accounts from the STACKSET constraint.
	StackSetAccounts []*string

	// The number of accounts, per region, for which this operation can fail before AWS
	// Service Catalog stops the operation in that region. If the operation is stopped
	// in a region, AWS Service Catalog doesn't attempt the operation in any subsequent
	// regions. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetFailureToleranceCount or
	// StackSetFailureTolerancePercentage, but not both. The default value is 0 if no
	// value is specified.
	StackSetFailureToleranceCount *int32

	// The percentage of accounts, per region, for which this stack operation can fail
	// before AWS Service Catalog stops the operation in that region. If the operation
	// is stopped in a region, AWS Service Catalog doesn't attempt the operation in any
	// subsequent regions. When calculating the number of accounts based on the
	// specified percentage, AWS Service Catalog rounds down to the next whole number.
	// Applicable only to a CFN_STACKSET provisioned product type. Conditional: You
	// must specify either StackSetFailureToleranceCount or
	// StackSetFailureTolerancePercentage, but not both.
	StackSetFailureTolerancePercentage *int32

	// The maximum number of accounts in which to perform this operation at one time.
	// This is dependent on the value of StackSetFailureToleranceCount.
	// StackSetMaxConcurrentCount is at most one more than the
	// StackSetFailureToleranceCount. Note that this setting lets you specify the
	// maximum for operations. For large deployments, under certain circumstances the
	// actual number of accounts acted upon concurrently may be lower due to service
	// throttling. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetMaxConcurrentCount or
	// StackSetMaxConcurrentPercentage, but not both.
	StackSetMaxConcurrencyCount *int32

	// The maximum percentage of accounts in which to perform this operation at one
	// time. When calculating the number of accounts based on the specified percentage,
	// AWS Service Catalog rounds down to the next whole number. This is true except in
	// cases where rounding down would result is zero. In this case, AWS Service
	// Catalog sets the number as 1 instead. Note that this setting lets you specify
	// the maximum for operations. For large deployments, under certain circumstances
	// the actual number of accounts acted upon concurrently may be lower due to
	// service throttling. Applicable only to a CFN_STACKSET provisioned product type.
	// Conditional: You must specify either StackSetMaxConcurrentCount or
	// StackSetMaxConcurrentPercentage, but not both.
	StackSetMaxConcurrencyPercentage *int32

	// Determines what action AWS Service Catalog performs to a stack set or a stack
	// instance represented by the provisioned product. The default value is UPDATE if
	// nothing is specified. Applicable only to a CFN_STACKSET provisioned product
	// type. CREATE Creates a new stack instance in the stack set represented by the
	// provisioned product. In this case, only new stack instances are created based on
	// accounts and regions; if new ProductId or ProvisioningArtifactID are passed,
	// they will be ignored. UPDATE Updates the stack set represented by the
	// provisioned product and also its stack instances. DELETE Deletes a stack
	// instance in the stack set represented by the provisioned product.
	StackSetOperationType StackSetOperationType

	// One or more AWS Regions where the provisioned product will be available.
	// Applicable only to a CFN_STACKSET provisioned product type. The specified
	// regions should be within the list of regions from the STACKSET constraint. To
	// get the list of regions in the STACKSET constraint, use the
	// DescribeProvisioningParameters operation. If no values are specified, the
	// default value is all regions from the STACKSET constraint.
	StackSetRegions []*string
}

// Additional information provided by the administrator.
type UsageInstruction struct {

	// The usage instruction type for the value.
	Type *string

	// The usage instruction value for this type.
	Value *string
}

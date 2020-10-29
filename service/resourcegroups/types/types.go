// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A resource that failed to be added to or removed from a group.
type FailedResource struct {

	// The error code associated with the failure.
	ErrorCode *string

	// The error message text associated with the failure.
	ErrorMessage *string

	// The ARN of the resource that failed to be added or removed.
	ResourceArn *string
}

// A resource group that contains AWS resources. You can assign resources to the
// group by associating either of the following elements with the group:
//
// *
// ResourceQuery - Use a resource query to specify a set of tag keys and values.
// All resources in the same AWS Region and AWS account that have those keys with
// the same values are included in the group. You can add a resource query when you
// create the group.
//
// * GroupConfiguration - Use a service configuration to
// associate the group with an AWS service. The configuration specifies which
// resource types can be included in the group.
type Group struct {

	// The ARN of the resource group.
	//
	// This member is required.
	GroupArn *string

	// The name of the resource group.
	//
	// This member is required.
	Name *string

	// The description of the resource group.
	Description *string
}

// A service configuration associated with a resource group. The configuration
// options are determined by the AWS service that defines the Type, and specifies
// which resources can be included in the group. You can add a service
// configuration when you create the group.
type GroupConfiguration struct {

	// The configuration currently associated with the group and in effect.
	Configuration []*GroupConfigurationItem

	// If present, the reason why a request to update the group configuration failed.
	FailureReason *string

	// If present, the new configuration that is in the process of being applied to the
	// group.
	ProposedConfiguration []*GroupConfigurationItem

	// The current status of an attempt to update the group configuration.
	Status GroupConfigurationStatus
}

// An item in a group configuration. A group configuration can have one or more
// items.
type GroupConfigurationItem struct {

	// Specifies the type of group configuration item. Each item must have a unique
	// value for type. You can specify the following string values:
	//
	// *
	// AWS::EC2::CapacityReservationPool For more information about EC2 capacity
	// reservation groups, see Working with capacity reservation groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html#create-cr-group)
	// in the EC2 Users Guide.
	//
	// * AWS::ResourceGroups::Generic - Supports parameters
	// that configure the behavior of resource groups of any type.
	//
	// This member is required.
	Type *string

	// A collection of parameters for this group configuration item.
	Parameters []*GroupConfigurationParameter
}

// A parameter for a group configuration item.
type GroupConfigurationParameter struct {

	// The name of the group configuration parameter. You can specify the following
	// string values:
	//
	// * For configuration item type AWS::ResourceGroups::Generic:
	//
	// *
	// allowed-resource-types Specifies the types of resources that you can add to this
	// group by using the GroupResources operation.
	//
	// * For configuration item type
	// AWS::EC2::CapacityReservationPool:
	//
	// * None - This configuration item type
	// doesn't support any parameters.
	//
	// For more information about EC2 capacity
	// reservation groups, see Working with capacity reservation groups
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html#create-cr-group)
	// in the EC2 Users Guide.
	//
	// This member is required.
	Name *string

	// The values of for this parameter. You can specify the following string value:
	//
	// *
	// For item type allowed-resource-types: the only supported parameter value is
	// AWS::EC2::CapacityReservation.
	Values []*string
}

// A filter collection that you can use to restrict the results from a List
// operation to only those you want to include.
type GroupFilter struct {

	// The name of the filter. Filter names are case-sensitive.
	//
	// This member is required.
	Name GroupFilterName

	// One or more filter values. Allowed filter values vary by group filter name, and
	// are case-sensitive.
	//
	// This member is required.
	Values []*string
}

// The unique identifiers for a resource group.
type GroupIdentifier struct {

	// The ARN of the resource group.
	GroupArn *string

	// The name of the resource group.
	GroupName *string
}

// A mapping of a query attached to a resource group that determines the AWS
// resources that are members of the group.
type GroupQuery struct {

	// The name of the resource group that is associated with the specified resource
	// query.
	//
	// This member is required.
	GroupName *string

	// The resource query that determines which AWS resources are members of the
	// associated resource group.
	//
	// This member is required.
	ResourceQuery *ResourceQuery
}

// A two-part error structure that can occur in ListGroupResources or
// SearchResources operations on CloudFormation stack-based queries. The error
// occurs if the CloudFormation stack on which the query is based either does not
// exist, or has a status that renders the stack inactive. A QueryError occurrence
// does not necessarily mean that AWS Resource Groups could not complete the
// operation, but the resulting group might have no member resources.
type QueryError struct {

	// Possible values are CLOUDFORMATION_STACK_INACTIVE and
	// CLOUDFORMATION_STACK_NOT_EXISTING.
	ErrorCode QueryErrorCode

	// A message that explains the ErrorCode value. Messages might state that the
	// specified CloudFormation stack does not exist (or no longer exists). For
	// CLOUDFORMATION_STACK_INACTIVE, the message typically states that the
	// CloudFormation stack has a status that is not (or no longer) active, such as
	// CREATE_FAILED.
	Message *string
}

// A filter name and value pair that is used to obtain more specific results from a
// list of resources.
type ResourceFilter struct {

	// The name of the filter. Filter names are case-sensitive.
	//
	// This member is required.
	Name ResourceFilterName

	// One or more filter values. Allowed filter values vary by resource filter name,
	// and are case-sensitive.
	//
	// This member is required.
	Values []*string
}

// The ARN of a resource, and its resource type.
type ResourceIdentifier struct {

	// The ARN of a resource.
	ResourceArn *string

	// The resource type of a resource, such as AWS::EC2::Instance.
	ResourceType *string
}

// The query that is used to define a resource group or a search for resources. A
// query specifies both a query type and a query string as a JSON object. See the
// examples section for example JSON strings. The examples that follow are shown as
// standard JSON strings. If you include such a string as a parameter to the AWS
// CLI or an SDK API, you might need to 'escape' the string into a single line. For
// example, see the Quoting strings
// (https://docs.aws.amazon.com/cli/latest/userguide/cli-usage-parameters-quoting-strings.html)
// in the AWS CLI User Guide. Example 1 The following generic example shows a
// resource query JSON string that includes only resources that meet the following
// criteria:
//
// * The resource type must be either resource_type1 or
// resource_type2.
//
// * The resource must have a tag Key1 with a value of either
// ValueA or ValueB.
//
// * The resource must have a tag Key2 with a value of either
// ValueC or ValueD.
//
// { "Type": "TAG_FILTERS_1_0", "Query": {
// "ResourceTypeFilters": [ "resource_type1", "resource_type2"], "TagFilters": [ {
// "Key": "Key1", "Values": ["ValueA","ValueB"] }, { "Key":"Key2",
// "Values":["ValueC","ValueD"] } ] } } This has the equivalent "shortcut" syntax
// of the following: { "Type": "TAG_FILTERS_1_0", "Query": { "ResourceTypeFilters":
// [ "resource_type1", "resource_type2"], "TagFilters": [ { "Key1":
// ["ValueA","ValueB"] }, { "Key2": ["ValueC","ValueD"] } ] } } Example 2 The
// following example shows a resource query JSON string that includes only Amazon
// EC2 instances that are tagged Stage with a value of Test. { "Type":
// "TAG_FILTERS_1_0", "Query": "{ "ResourceTypeFilters": "AWS::EC2::Instance",
// "TagFilters": { "Stage": "Test" } } } Example 3 The following example shows a
// resource query JSON string that includes resource of any supported type as long
// as it is tagged Stage with a value of Prod. { "Type": "TAG_FILTERS_1_0",
// "Query": { "ResourceTypeFilters": "AWS::AllSupported", "TagFilters": { "Stage":
// "Prod" } } } Example 4 The following example shows a resource query JSON string
// that includes only Amazon EC2 instances and Amazon S3 buckets that are part of
// the specified AWS CloudFormation stack. { "Type": "CLOUDFORMATION_STACK_1_0",
// "Query": { "ResourceTypeFilters": [ "AWS::EC2::Instance", "AWS::S3::Bucket" ],
// "StackIdentifier":
// "arn:aws:cloudformation:us-west-2:123456789012:stack/AWStestuseraccount/fb0d5000-aba8-00e8-aa9e-50d5cEXAMPLE"
// } }
type ResourceQuery struct {

	// The query that defines a group or a search.
	//
	// This member is required.
	Query *string

	// The type of the query. You can use the following values:
	//
	// *
	// CLOUDFORMATION_STACK_1_0: Specifies that the Query contains an ARN for a
	// CloudFormation stack.
	//
	// * TAG_FILTERS_1_0: Specifies that the Query parameter
	// contains a JSON string that represents a collection of simple tag filters for
	// resource types and tags. The JSON string uses a syntax similar to the
	// GetResources
	// (https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_GetResources.html)
	// operation, but uses only the  ResourceTypeFilters
	// (https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_GetResources.html#resourcegrouptagging-GetResources-request-ResourceTypeFilters)
	// and TagFilters
	// (https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/API_GetResources.html#resourcegrouptagging-GetResources-request-TagFiltersTagFilters)
	// fields. If you specify more than one tag key, only resources that match all tag
	// keys, and at least one value of each specified tag key, are returned in your
	// query. If you specify more than one value for a tag key, a resource matches the
	// filter if it has a tag key value that matches any of the specified values. For
	// example, consider the following sample query for resources that have two tags,
	// Stage and Version, with two values each:
	// [{"Stage":["Test","Deploy"]},{"Version":["1","2"]}] The results of this query
	// could include the following.
	//
	// * An EC2 instance that has the following two tags:
	// {"Stage":"Deploy"}, and {"Version":"2"}
	//
	// * An S3 bucket that has the following
	// two tags: {"Stage":"Test"}, and {"Version":"1"}
	//
	// The query would not include the
	// following items in the results, however.
	//
	// * An EC2 instance that has only the
	// following tag: {"Stage":"Deploy"}. The instance does not have all of the tag
	// keys specified in the filter, so it is excluded from the results.
	//
	// * An RDS
	// database that has the following two tags: {"Stage":"Archived"} and
	// {"Version":"4"} The database has all of the tag keys, but none of those keys has
	// an associated value that matches at least one of the specified values in the
	// filter.
	//
	// This member is required.
	Type QueryType
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// A key-value pair that describes a property of a pipeline object. The value is
// specified as either a string value (StringValue) or a reference to another
// object (RefValue) but not as both.
type Field struct {

	// The field identifier.
	//
	// This member is required.
	Key *string

	// The field value, expressed as the identifier of another object.
	RefValue *string

	// The field value, expressed as a String.
	StringValue *string
}

// Identity information for the EC2 instance that is hosting the task runner. You
// can get this value by calling a metadata URI from the EC2 instance. For more
// information, see Instance Metadata
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AESDG-chapter-instancedata.html)
// in the Amazon Elastic Compute Cloud User Guide. Passing in this value proves
// that your task runner is running on an EC2 instance, and ensures the proper AWS
// Data Pipeline service charges are applied to your pipeline.
type InstanceIdentity struct {

	// A description of an EC2 instance that is generated when the instance is launched
	// and exposed to the instance via the instance metadata service in the form of a
	// JSON representation of an object.
	Document *string

	// A signature which can be used to verify the accuracy and authenticity of the
	// information provided in the instance identity document.
	Signature *string
}

// Contains a logical operation for comparing the value of a field with a specified
// value.
type Operator struct {

	// The logical operation to be performed: equal (EQ), equal reference (REF_EQ),
	// less than or equal (LE), greater than or equal (GE), or between (BETWEEN). Equal
	// reference (REF_EQ) can be used only with reference fields. The other comparison
	// types can be used only with String fields. The comparison types you can use
	// apply only to certain object fields, as detailed below. The comparison operators
	// EQ and REF_EQ act on the following fields:
	//
	// * name
	//
	// * @sphere
	//
	// * parent
	//
	// *
	// @componentParent
	//
	// * @instanceParent
	//
	// * @status
	//
	// * @scheduledStartTime
	//
	// *
	// @scheduledEndTime
	//
	// * @actualStartTime
	//
	// * @actualEndTime
	//
	// The comparison
	// operators GE, LE, and BETWEEN act on the following fields:
	//
	// *
	// @scheduledStartTime
	//
	// * @scheduledEndTime
	//
	// * @actualStartTime
	//
	// *
	// @actualEndTime
	//
	// Note that fields beginning with the at sign (@) are read-only
	// and set by the web service. When you name fields, you should choose names
	// containing only alpha-numeric values, as symbols may be reserved by AWS Data
	// Pipeline. User-defined fields that you add to a pipeline should prefix their
	// name with the string "my".
	Type OperatorType

	// The value that the actual field value will be compared with.
	Values []*string
}

// The attributes allowed or specified with a parameter object.
type ParameterAttribute struct {

	// The field identifier.
	//
	// This member is required.
	Key *string

	// The field value, expressed as a String.
	//
	// This member is required.
	StringValue *string
}

// Contains information about a parameter object.
type ParameterObject struct {

	// The attributes of the parameter object.
	//
	// This member is required.
	Attributes []*ParameterAttribute

	// The ID of the parameter object.
	//
	// This member is required.
	Id *string
}

// A value or list of parameter values.
type ParameterValue struct {

	// The ID of the parameter value.
	//
	// This member is required.
	Id *string

	// The field value, expressed as a String.
	//
	// This member is required.
	StringValue *string
}

// Contains pipeline metadata.
type PipelineDescription struct {

	// A list of read-only fields that contain metadata about the pipeline: @userId,
	// @accountId, and @pipelineState.
	//
	// This member is required.
	Fields []*Field

	// The name of the pipeline.
	//
	// This member is required.
	Name *string

	// The pipeline identifier that was assigned by AWS Data Pipeline. This is a string
	// of the form df-297EG78HU43EEXAMPLE.
	//
	// This member is required.
	PipelineId *string

	// Description of the pipeline.
	Description *string

	// A list of tags to associated with a pipeline. Tags let you control access to
	// pipelines. For more information, see Controlling User Access to Pipelines
	// (https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	Tags []*Tag
}

// Contains the name and identifier of a pipeline.
type PipelineIdName struct {

	// The ID of the pipeline that was assigned by AWS Data Pipeline. This is a string
	// of the form df-297EG78HU43EEXAMPLE.
	Id *string

	// The name of the pipeline.
	Name *string
}

// Contains information about a pipeline object. This can be a logical, physical,
// or physical attempt pipeline object. The complete set of components of a
// pipeline defines the pipeline.
type PipelineObject struct {

	// Key-value pairs that define the properties of the object.
	//
	// This member is required.
	Fields []*Field

	// The ID of the object.
	//
	// This member is required.
	Id *string

	// The name of the object.
	//
	// This member is required.
	Name *string
}

// Defines the query to run against an object.
type Query struct {

	// List of selectors that define the query. An object must satisfy all of the
	// selectors to match the query.
	Selectors []*Selector
}

// A comparison that is used to determine whether a query should return this
// object.
type Selector struct {

	// The name of the field that the operator will be applied to. The field name is
	// the "key" portion of the field definition in the pipeline definition syntax that
	// is used by the AWS Data Pipeline API. If the field is not set on the object, the
	// condition fails.
	FieldName *string

	// Contains a logical operation for comparing the value of a field with a specified
	// value.
	Operator *Operator
}

// Tags are key/value pairs defined by a user and associated with a pipeline to
// control access. AWS Data Pipeline allows you to associate ten tags per pipeline.
// For more information, see Controlling User Access to Pipelines
// (https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
// in the AWS Data Pipeline Developer Guide.
type Tag struct {

	// The key name of a tag defined by a user. For more information, see Controlling
	// User Access to Pipelines
	// (https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	//
	// This member is required.
	Key *string

	// The optional value portion of a tag defined by a user. For more information, see
	// Controlling User Access to Pipelines
	// (https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-control-access.html)
	// in the AWS Data Pipeline Developer Guide.
	//
	// This member is required.
	Value *string
}

// Contains information about a pipeline task that is assigned to a task runner.
type TaskObject struct {

	// The ID of the pipeline task attempt object. AWS Data Pipeline uses this value to
	// track how many times a task is attempted.
	AttemptId *string

	// Connection information for the location where the task runner will publish the
	// output of the task.
	Objects map[string]*PipelineObject

	// The ID of the pipeline that provided the task.
	PipelineId *string

	// An internal identifier for the task. This ID is passed to the SetTaskStatus and
	// ReportTaskProgress actions.
	TaskId *string
}

// Defines a validation error. Validation errors prevent pipeline activation. The
// set of validation errors that can be returned are defined by AWS Data Pipeline.
type ValidationError struct {

	// A description of the validation error.
	Errors []*string

	// The identifier of the object that contains the validation error.
	Id *string
}

// Defines a validation warning. Validation warnings do not prevent pipeline
// activation. The set of validation warnings that can be returned are defined by
// AWS Data Pipeline.
type ValidationWarning struct {

	// The identifier of the object that contains the validation warning.
	Id *string

	// A description of the validation warning.
	Warnings []*string
}

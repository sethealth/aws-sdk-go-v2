// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about the AccessLog attribute.
type AccessLog struct {

	// Specifies whether access logs are enabled for the load balancer.
	//
	// This member is required.
	Enabled *bool

	// The interval for publishing the access logs. You can specify an interval of
	// either 5 minutes or 60 minutes. Default: 60 minutes
	EmitInterval *int32

	// The name of the Amazon S3 bucket where the access logs are stored.
	S3BucketName *string

	// The logical hierarchy you created for your Amazon S3 bucket, for example
	// my-bucket-prefix/prod. If the prefix is not provided, the log is placed at the
	// root level of the bucket.
	S3BucketPrefix *string
}

// Information about additional load balancer attributes.
type AdditionalAttribute struct {

	// The name of the attribute. The following attribute is supported.
	//
	// *
	// elb.http.desyncmitigationmode - Determines how the load balancer handles
	// requests that might pose a security risk to your application. The possible
	// values are monitor, defensive, and strictest. The default is defensive.
	Key *string

	// This value of the attribute.
	Value *string
}

// Information about a policy for application-controlled session stickiness.
type AppCookieStickinessPolicy struct {

	// The name of the application cookie used for stickiness.
	CookieName *string

	// The mnemonic name for the policy being created. The name must be unique within a
	// set of policies for this load balancer.
	PolicyName *string
}

// Information about the configuration of an EC2 instance.
type BackendServerDescription struct {

	// The port on which the EC2 instance is listening.
	InstancePort *int32

	// The names of the policies enabled for the EC2 instance.
	PolicyNames []*string
}

// Information about the ConnectionDraining attribute.
type ConnectionDraining struct {

	// Specifies whether connection draining is enabled for the load balancer.
	//
	// This member is required.
	Enabled *bool

	// The maximum time, in seconds, to keep the existing connections open before
	// deregistering the instances.
	Timeout *int32
}

// Information about the ConnectionSettings attribute.
type ConnectionSettings struct {

	// The time, in seconds, that the connection is allowed to be idle (no data has
	// been sent over the connection) before it is closed by the load balancer.
	//
	// This member is required.
	IdleTimeout *int32
}

// Information about the CrossZoneLoadBalancing attribute.
type CrossZoneLoadBalancing struct {

	// Specifies whether cross-zone load balancing is enabled for the load balancer.
	//
	// This member is required.
	Enabled *bool
}

// Information about a health check.
type HealthCheck struct {

	// The number of consecutive health checks successes required before moving the
	// instance to the Healthy state.
	//
	// This member is required.
	HealthyThreshold *int32

	// The approximate interval, in seconds, between health checks of an individual
	// instance.
	//
	// This member is required.
	Interval *int32

	// The instance being checked. The protocol is either TCP, HTTP, HTTPS, or SSL. The
	// range of valid ports is one (1) through 65535. TCP is the default, specified as
	// a TCP: port pair, for example "TCP:5000". In this case, a health check simply
	// attempts to open a TCP connection to the instance on the specified port. Failure
	// to connect within the configured timeout is considered unhealthy. SSL is also
	// specified as SSL: port pair, for example, SSL:5000. For HTTP/HTTPS, you must
	// include a ping path in the string. HTTP is specified as a
	// HTTP:port;/;PathToPing; grouping, for example "HTTP:80/weather/us/wa/seattle".
	// In this case, a HTTP GET request is issued to the instance on the given port and
	// path. Any answer other than "200 OK" within the timeout period is considered
	// unhealthy. The total length of the HTTP ping target must be 1024 16-bit Unicode
	// characters or less.
	//
	// This member is required.
	Target *string

	// The amount of time, in seconds, during which no response means a failed health
	// check. This value must be less than the Interval value.
	//
	// This member is required.
	Timeout *int32

	// The number of consecutive health check failures required before moving the
	// instance to the Unhealthy state.
	//
	// This member is required.
	UnhealthyThreshold *int32
}

// The ID of an EC2 instance.
type Instance struct {

	// The instance ID.
	InstanceId *string
}

// Information about the state of an EC2 instance.
type InstanceState struct {

	// A description of the instance state. This string can contain one or more of the
	// following messages.
	//
	// * N/A
	//
	// * A transient error occurred. Please try again
	// later.
	//
	// * Instance has failed at least the UnhealthyThreshold number of health
	// checks consecutively.
	//
	// * Instance has not passed the configured HealthyThreshold
	// number of health checks consecutively.
	//
	// * Instance registration is still in
	// progress.
	//
	// * Instance is in the EC2 Availability Zone for which LoadBalancer is
	// not configured to route traffic to.
	//
	// * Instance is not currently registered with
	// the LoadBalancer.
	//
	// * Instance deregistration currently in progress.
	//
	// * Disable
	// Availability Zone is currently in progress.
	//
	// * Instance is in pending state.
	//
	// *
	// Instance is in stopped state.
	//
	// * Instance is in terminated state.
	Description *string

	// The ID of the instance.
	InstanceId *string

	// Information about the cause of OutOfService instances. Specifically, whether the
	// cause is Elastic Load Balancing or the instance. Valid values: ELB | Instance |
	// N/A
	ReasonCode *string

	// The current state of the instance. Valid values: InService | OutOfService |
	// Unknown
	State *string
}

// Information about a policy for duration-based session stickiness.
type LBCookieStickinessPolicy struct {

	// The time period, in seconds, after which the cookie should be considered stale.
	// If this parameter is not specified, the stickiness session lasts for the
	// duration of the browser session.
	CookieExpirationPeriod *int64

	// The name of the policy. This name must be unique within the set of policies for
	// this load balancer.
	PolicyName *string
}

// Information about an Elastic Load Balancing resource limit for your AWS account.
type Limit struct {

	// The maximum value of the limit.
	Max *string

	// The name of the limit. The possible values are:
	//
	// * classic-listeners
	//
	// *
	// classic-load-balancers
	//
	// * classic-registered-instances
	Name *string
}

// Information about a listener. For information about the protocols and the ports
// supported by Elastic Load Balancing, see Listeners for Your Classic Load
// Balancer
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
// in the Classic Load Balancers Guide.
type Listener struct {

	// The port on which the instance is listening.
	//
	// This member is required.
	InstancePort *int32

	// The port on which the load balancer is listening. On EC2-VPC, you can specify
	// any port from the range 1-65535. On EC2-Classic, you can specify any port from
	// the following list: 25, 80, 443, 465, 587, 1024-65535.
	//
	// This member is required.
	LoadBalancerPort *int32

	// The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or
	// SSL.
	//
	// This member is required.
	Protocol *string

	// The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.
	// If the front-end protocol is TCP or SSL, the back-end protocol must be TCP or
	// SSL. If the front-end protocol is HTTP or HTTPS, the back-end protocol must be
	// HTTP or HTTPS. If there is another listener with the same InstancePort whose
	// InstanceProtocol is secure, (HTTPS or SSL), the listener's InstanceProtocol must
	// also be secure. If there is another listener with the same InstancePort whose
	// InstanceProtocol is HTTP or TCP, the listener's InstanceProtocol must be HTTP or
	// TCP.
	InstanceProtocol *string

	// The Amazon Resource Name (ARN) of the server certificate.
	SSLCertificateId *string
}

// The policies enabled for a listener.
type ListenerDescription struct {

	// The listener.
	Listener *Listener

	// The policies. If there are no policies enabled, the list is empty.
	PolicyNames []*string
}

// The attributes for a load balancer.
type LoadBalancerAttributes struct {

	// If enabled, the load balancer captures detailed information of all requests and
	// delivers the information to the Amazon S3 bucket that you specify. For more
	// information, see Enable Access Logs
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html)
	// in the Classic Load Balancers Guide.
	AccessLog *AccessLog

	// Any additional attributes.
	AdditionalAttributes []*AdditionalAttribute

	// If enabled, the load balancer allows existing requests to complete before the
	// load balancer shifts traffic away from a deregistered or unhealthy instance. For
	// more information, see Configure Connection Draining
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-conn-drain.html)
	// in the Classic Load Balancers Guide.
	ConnectionDraining *ConnectionDraining

	// If enabled, the load balancer allows the connections to remain idle (no data is
	// sent over the connection) for the specified duration. By default, Elastic Load
	// Balancing maintains a 60-second idle connection timeout for both front-end and
	// back-end connections of your load balancer. For more information, see Configure
	// Idle Connection Timeout
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/config-idle-timeout.html)
	// in the Classic Load Balancers Guide.
	ConnectionSettings *ConnectionSettings

	// If enabled, the load balancer routes the request traffic evenly across all
	// instances regardless of the Availability Zones. For more information, see
	// Configure Cross-Zone Load Balancing
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-disable-crosszone-lb.html)
	// in the Classic Load Balancers Guide.
	CrossZoneLoadBalancing *CrossZoneLoadBalancing
}

// Information about a load balancer.
type LoadBalancerDescription struct {

	// The Availability Zones for the load balancer.
	AvailabilityZones []*string

	// Information about your EC2 instances.
	BackendServerDescriptions []*BackendServerDescription

	// The DNS name of the load balancer. For more information, see Configure a Custom
	// Domain Name
	// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/using-domain-names-with-elb.html)
	// in the Classic Load Balancers Guide.
	CanonicalHostedZoneName *string

	// The ID of the Amazon Route 53 hosted zone for the load balancer.
	CanonicalHostedZoneNameID *string

	// The date and time the load balancer was created.
	CreatedTime *time.Time

	// The DNS name of the load balancer.
	DNSName *string

	// Information about the health checks conducted on the load balancer.
	HealthCheck *HealthCheck

	// The IDs of the instances for the load balancer.
	Instances []*Instance

	// The listeners for the load balancer.
	ListenerDescriptions []*ListenerDescription

	// The name of the load balancer.
	LoadBalancerName *string

	// The policies defined for the load balancer.
	Policies *Policies

	// The type of load balancer. Valid only for load balancers in a VPC. If Scheme is
	// internet-facing, the load balancer has a public DNS name that resolves to a
	// public IP address. If Scheme is internal, the load balancer has a public DNS
	// name that resolves to a private IP address.
	Scheme *string

	// The security groups for the load balancer. Valid only for load balancers in a
	// VPC.
	SecurityGroups []*string

	// The security group for the load balancer, which you can use as part of your
	// inbound rules for your registered instances. To only allow traffic from load
	// balancers, add a security group rule that specifies this source security group
	// as the inbound source.
	SourceSecurityGroup *SourceSecurityGroup

	// The IDs of the subnets for the load balancer.
	Subnets []*string

	// The ID of the VPC for the load balancer.
	VPCId *string
}

// The policies for a load balancer.
type Policies struct {

	// The stickiness policies created using CreateAppCookieStickinessPolicy.
	AppCookieStickinessPolicies []*AppCookieStickinessPolicy

	// The stickiness policies created using CreateLBCookieStickinessPolicy.
	LBCookieStickinessPolicies []*LBCookieStickinessPolicy

	// The policies other than the stickiness policies.
	OtherPolicies []*string
}

// Information about a policy attribute.
type PolicyAttribute struct {

	// The name of the attribute.
	AttributeName *string

	// The value of the attribute.
	AttributeValue *string
}

// Information about a policy attribute.
type PolicyAttributeDescription struct {

	// The name of the attribute.
	AttributeName *string

	// The value of the attribute.
	AttributeValue *string
}

// Information about a policy attribute type.
type PolicyAttributeTypeDescription struct {

	// The name of the attribute.
	AttributeName *string

	// The type of the attribute. For example, Boolean or Integer.
	AttributeType *string

	// The cardinality of the attribute. Valid values:
	//
	// * ONE(1) : Single value
	// required
	//
	// * ZERO_OR_ONE(0..1) : Up to one value is allowed
	//
	// * ZERO_OR_MORE(0..*)
	// : Optional. Multiple values are allowed
	//
	// * ONE_OR_MORE(1..*0) : Required.
	// Multiple values are allowed
	Cardinality *string

	// The default value of the attribute, if applicable.
	DefaultValue *string

	// A description of the attribute.
	Description *string
}

// Information about a policy.
type PolicyDescription struct {

	// The policy attributes.
	PolicyAttributeDescriptions []*PolicyAttributeDescription

	// The name of the policy.
	PolicyName *string

	// The name of the policy type.
	PolicyTypeName *string
}

// Information about a policy type.
type PolicyTypeDescription struct {

	// A description of the policy type.
	Description *string

	// The description of the policy attributes associated with the policies defined by
	// Elastic Load Balancing.
	PolicyAttributeTypeDescriptions []*PolicyAttributeTypeDescription

	// The name of the policy type.
	PolicyTypeName *string
}

// Information about a source security group.
type SourceSecurityGroup struct {

	// The name of the security group.
	GroupName *string

	// The owner of the security group.
	OwnerAlias *string
}

// Information about a tag.
type Tag struct {

	// The key of the tag.
	//
	// This member is required.
	Key *string

	// The value of the tag.
	Value *string
}

// The tags associated with a load balancer.
type TagDescription struct {

	// The name of the load balancer.
	LoadBalancerName *string

	// The tags.
	Tags []*Tag
}

// The key of a tag.
type TagKeyOnly struct {

	// The name of the key.
	Key *string
}

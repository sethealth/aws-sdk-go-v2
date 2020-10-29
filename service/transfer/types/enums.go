// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EndpointType string

// Enum values for EndpointType
const (
	EndpointTypePublic      EndpointType = "PUBLIC"
	EndpointTypeVpc         EndpointType = "VPC"
	EndpointTypeVpcEndpoint EndpointType = "VPC_ENDPOINT"
)

// Values returns all known values for EndpointType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EndpointType) Values() []EndpointType {
	return []EndpointType{
		"PUBLIC",
		"VPC",
		"VPC_ENDPOINT",
	}
}

type HomeDirectoryType string

// Enum values for HomeDirectoryType
const (
	HomeDirectoryTypePath    HomeDirectoryType = "PATH"
	HomeDirectoryTypeLogical HomeDirectoryType = "LOGICAL"
)

// Values returns all known values for HomeDirectoryType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (HomeDirectoryType) Values() []HomeDirectoryType {
	return []HomeDirectoryType{
		"PATH",
		"LOGICAL",
	}
}

type IdentityProviderType string

// Enum values for IdentityProviderType
const (
	IdentityProviderTypeServiceManaged IdentityProviderType = "SERVICE_MANAGED"
	IdentityProviderTypeApiGateway     IdentityProviderType = "API_GATEWAY"
)

// Values returns all known values for IdentityProviderType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IdentityProviderType) Values() []IdentityProviderType {
	return []IdentityProviderType{
		"SERVICE_MANAGED",
		"API_GATEWAY",
	}
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolSftp Protocol = "SFTP"
	ProtocolFtp  Protocol = "FTP"
	ProtocolFtps Protocol = "FTPS"
)

// Values returns all known values for Protocol. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Protocol) Values() []Protocol {
	return []Protocol{
		"SFTP",
		"FTP",
		"FTPS",
	}
}

type State string

// Enum values for State
const (
	StateOffline     State = "OFFLINE"
	StateOnline      State = "ONLINE"
	StateStarting    State = "STARTING"
	StateStopping    State = "STOPPING"
	StateStartFailed State = "START_FAILED"
	StateStopFailed  State = "STOP_FAILED"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"OFFLINE",
		"ONLINE",
		"STARTING",
		"STOPPING",
		"START_FAILED",
		"STOP_FAILED",
	}
}

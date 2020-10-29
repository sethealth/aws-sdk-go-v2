// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type BrokerAZDistribution string

// Enum values for BrokerAZDistribution
const (
	BrokerAZDistributionDefault BrokerAZDistribution = "DEFAULT"
)

// Values returns all known values for BrokerAZDistribution. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (BrokerAZDistribution) Values() []BrokerAZDistribution {
	return []BrokerAZDistribution{
		"DEFAULT",
	}
}

type ClientBroker string

// Enum values for ClientBroker
const (
	ClientBrokerTls          ClientBroker = "TLS"
	ClientBrokerTlsPlaintext ClientBroker = "TLS_PLAINTEXT"
	ClientBrokerPlaintext    ClientBroker = "PLAINTEXT"
)

// Values returns all known values for ClientBroker. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ClientBroker) Values() []ClientBroker {
	return []ClientBroker{
		"TLS",
		"TLS_PLAINTEXT",
		"PLAINTEXT",
	}
}

type ClusterState string

// Enum values for ClusterState
const (
	ClusterStateActive   ClusterState = "ACTIVE"
	ClusterStateCreating ClusterState = "CREATING"
	ClusterStateUpdating ClusterState = "UPDATING"
	ClusterStateDeleting ClusterState = "DELETING"
	ClusterStateFailed   ClusterState = "FAILED"
)

// Values returns all known values for ClusterState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ClusterState) Values() []ClusterState {
	return []ClusterState{
		"ACTIVE",
		"CREATING",
		"UPDATING",
		"DELETING",
		"FAILED",
	}
}

type ConfigurationState string

// Enum values for ConfigurationState
const (
	ConfigurationStateActive       ConfigurationState = "ACTIVE"
	ConfigurationStateDeleting     ConfigurationState = "DELETING"
	ConfigurationStateDeleteFailed ConfigurationState = "DELETE_FAILED"
)

// Values returns all known values for ConfigurationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConfigurationState) Values() []ConfigurationState {
	return []ConfigurationState{
		"ACTIVE",
		"DELETING",
		"DELETE_FAILED",
	}
}

type EnhancedMonitoring string

// Enum values for EnhancedMonitoring
const (
	EnhancedMonitoringDefault           EnhancedMonitoring = "DEFAULT"
	EnhancedMonitoringPerBroker         EnhancedMonitoring = "PER_BROKER"
	EnhancedMonitoringPerTopicPerBroker EnhancedMonitoring = "PER_TOPIC_PER_BROKER"
)

// Values returns all known values for EnhancedMonitoring. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnhancedMonitoring) Values() []EnhancedMonitoring {
	return []EnhancedMonitoring{
		"DEFAULT",
		"PER_BROKER",
		"PER_TOPIC_PER_BROKER",
	}
}

type KafkaVersionStatus string

// Enum values for KafkaVersionStatus
const (
	KafkaVersionStatusActive     KafkaVersionStatus = "ACTIVE"
	KafkaVersionStatusDeprecated KafkaVersionStatus = "DEPRECATED"
)

// Values returns all known values for KafkaVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KafkaVersionStatus) Values() []KafkaVersionStatus {
	return []KafkaVersionStatus{
		"ACTIVE",
		"DEPRECATED",
	}
}

type NodeType string

// Enum values for NodeType
const (
	NodeTypeBroker NodeType = "BROKER"
)

// Values returns all known values for NodeType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (NodeType) Values() []NodeType {
	return []NodeType{
		"BROKER",
	}
}

// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeGroup         ResourceType = "GROUP"
	ResourceTypeUser          ResourceType = "USER"
	ResourceTypeIdentityStore ResourceType = "IDENTITY_STORE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"GROUP",
		"USER",
		"IDENTITY_STORE",
	}
}

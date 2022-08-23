// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	operatorv1 "github.com/openshift/api/operator/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// ClusterCSIDriverSpecApplyConfiguration represents an declarative configuration of the ClusterCSIDriverSpec type for use
// with apply.
type ClusterCSIDriverSpecApplyConfiguration struct {
	OperatorSpecApplyConfiguration `json:",inline"`
	StorageClassState              *operatorv1.StorageClassStateName `json:"storageClassState,omitempty"`
}

// ClusterCSIDriverSpecApplyConfiguration constructs an declarative configuration of the ClusterCSIDriverSpec type for use with
// apply.
func ClusterCSIDriverSpec() *ClusterCSIDriverSpecApplyConfiguration {
	return &ClusterCSIDriverSpecApplyConfiguration{}
}

// WithManagementState sets the ManagementState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagementState field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithManagementState(value operatorv1.ManagementState) *ClusterCSIDriverSpecApplyConfiguration {
	b.ManagementState = &value
	return b
}

// WithLogLevel sets the LogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LogLevel field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithLogLevel(value operatorv1.LogLevel) *ClusterCSIDriverSpecApplyConfiguration {
	b.LogLevel = &value
	return b
}

// WithOperatorLogLevel sets the OperatorLogLevel field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OperatorLogLevel field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithOperatorLogLevel(value operatorv1.LogLevel) *ClusterCSIDriverSpecApplyConfiguration {
	b.OperatorLogLevel = &value
	return b
}

// WithUnsupportedConfigOverrides sets the UnsupportedConfigOverrides field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UnsupportedConfigOverrides field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithUnsupportedConfigOverrides(value runtime.RawExtension) *ClusterCSIDriverSpecApplyConfiguration {
	b.UnsupportedConfigOverrides = &value
	return b
}

// WithObservedConfig sets the ObservedConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedConfig field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithObservedConfig(value runtime.RawExtension) *ClusterCSIDriverSpecApplyConfiguration {
	b.ObservedConfig = &value
	return b
}

// WithStorageClassState sets the StorageClassState field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageClassState field is set to the value of the last call.
func (b *ClusterCSIDriverSpecApplyConfiguration) WithStorageClassState(value operatorv1.StorageClassStateName) *ClusterCSIDriverSpecApplyConfiguration {
	b.StorageClassState = &value
	return b
}

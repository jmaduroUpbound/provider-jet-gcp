/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AddressObservation struct {
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type AddressParameters struct {

	// The static external IP address represented by this resource. Only
	// IPv4 is supported. An address may only be specified for INTERNAL
	// address types. The IP address must be inside the specified subnetwork,
	// if any.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The type of address to reserve. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"]
	// +kubebuilder:validation:Optional
	AddressType *string `json:"addressType,omitempty" tf:"address_type,omitempty"`

	// An optional description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The URL of the network in which to reserve the address. This field
	// can only be used with INTERNAL type with the VPC_PEERING and
	// IPSEC_INTERCONNECT purposes.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Optional
	NetworkRef *v1.Reference `json:"networkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	NetworkSelector *v1.Selector `json:"networkSelector,omitempty" tf:"-"`

	// The networking tier used for configuring this address. If this field is not
	// specified, it is assumed to be PREMIUM. Possible values: ["PREMIUM", "STANDARD"]
	// +kubebuilder:validation:Optional
	NetworkTier *string `json:"networkTier,omitempty" tf:"network_tier,omitempty"`

	// The prefix length if the resource represents an IP range.
	// +kubebuilder:validation:Optional
	PrefixLength *int64 `json:"prefixLength,omitempty" tf:"prefix_length,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The purpose of this resource, which can be one of the following values:
	//
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP
	// ranges, internal load balancers, and similar resources.
	//
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple
	// internal load balancers.
	//
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//
	// * IPSEC_INTERCONNECT for addresses created from a private IP range
	// that are reserved for a VLAN attachment in an IPsec-encrypted Cloud
	// Interconnect configuration. These addresses are regional resources.
	//
	// * PRIVATE_SERVICE_CONNECT for a private network address that is used
	// to configure Private Service Connect. Only global internal addresses
	// can use this purpose.
	//
	// This should only be set when using an Internal address.
	// +kubebuilder:validation:Optional
	Purpose *string `json:"purpose,omitempty" tf:"purpose,omitempty"`

	// The Region in which the created address should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The URL of the subnetwork in which to reserve the address. If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	// +crossplane:generate:reference:type=Subnetwork
	// +kubebuilder:validation:Optional
	Subnetwork *string `json:"subnetwork,omitempty" tf:"subnetwork,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetworkRef *v1.Reference `json:"subnetworkRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetworkSelector *v1.Selector `json:"subnetworkSelector,omitempty" tf:"-"`
}

// AddressSpec defines the desired state of Address
type AddressSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AddressParameters `json:"forProvider"`
}

// AddressStatus defines the observed state of Address.
type AddressStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AddressObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Address is the Schema for the Addresss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfgcp}
type Address struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AddressSpec   `json:"spec"`
	Status            AddressStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AddressList contains a list of Addresss
type AddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Address `json:"items"`
}

// Repository type metadata.
var (
	Address_Kind             = "Address"
	Address_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Address_Kind}.String()
	Address_KindAPIVersion   = Address_Kind + "." + CRDGroupVersion.String()
	Address_GroupVersionKind = CRDGroupVersion.WithKind(Address_Kind)
)

func init() {
	SchemeBuilder.Register(&Address{}, &AddressList{})
}

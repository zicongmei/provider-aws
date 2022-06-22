/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PartitionIndexObservation_2 struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	PartitionIndex []PartitionIndexPartitionIndexObservation `json:"partitionIndex,omitempty" tf:"partition_index,omitempty"`
}

type PartitionIndexParameters_2 struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`

	// +kubebuilder:validation:Required
	PartitionIndex []PartitionIndexPartitionIndexParameters `json:"partitionIndex" tf:"partition_index,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

type PartitionIndexPartitionIndexObservation struct {
	IndexStatus *string `json:"indexStatus,omitempty" tf:"index_status,omitempty"`
}

type PartitionIndexPartitionIndexParameters struct {

	// +kubebuilder:validation:Optional
	Keys []*string `json:"keys,omitempty" tf:"keys,omitempty"`
}

// PartitionIndexSpec defines the desired state of PartitionIndex
type PartitionIndexSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PartitionIndexParameters_2 `json:"forProvider"`
}

// PartitionIndexStatus defines the observed state of PartitionIndex.
type PartitionIndexStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PartitionIndexObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PartitionIndex is the Schema for the PartitionIndexs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PartitionIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartitionIndexSpec   `json:"spec"`
	Status            PartitionIndexStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PartitionIndexList contains a list of PartitionIndexs
type PartitionIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PartitionIndex `json:"items"`
}

// Repository type metadata.
var (
	PartitionIndex_Kind             = "PartitionIndex"
	PartitionIndex_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PartitionIndex_Kind}.String()
	PartitionIndex_KindAPIVersion   = PartitionIndex_Kind + "." + CRDGroupVersion.String()
	PartitionIndex_GroupVersionKind = CRDGroupVersion.WithKind(PartitionIndex_Kind)
)

func init() {
	SchemeBuilder.Register(&PartitionIndex{}, &PartitionIndexList{})
}

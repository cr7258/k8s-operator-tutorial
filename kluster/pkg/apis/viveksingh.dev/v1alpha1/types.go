package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type Kluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec KlusterSpec
}

type KlusterSpec struct {
	Name    string
	Region  string
	Version string

	NodePools []NodePool
}

type NodePool struct {
	Size  string
	Name  string
	count int
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KlusterList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []Kluster
}

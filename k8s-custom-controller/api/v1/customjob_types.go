package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// CustomJobSpec defines the desired state of CustomJob
type CustomJobSpec struct {
	Foo string `json:"foo,omitempty"`
}

// CustomJobStatus defines the observed state of CustomJob
type CustomJobStatus struct {
	Bar string `json:"bar,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomJob is the Schema for the customjobs API
type CustomJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomJobSpec   `json:"spec,omitempty"`
	Status            CustomJobStatus `json:"status,omitempty"`
}

func (c *CustomJob) DeepCopyObject() runtime.Object {
	return c.DeepCopy()
}

func (c *CustomJob) GetObjectKind() schema.ObjectKind {
	return &c.TypeMeta
}

//+kubebuilder:object:root=true

// CustomJobList contains a list of CustomJob
type CustomJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomJob `json:"items"`
}

func (c *CustomJobList) DeepCopyObject() runtime.Object {
	return c.DeepCopy()
}

func (c *CustomJobList) GetObjectKind() schema.ObjectKind {
	return &c.TypeMeta
}

func init() {
	SchemeBuilder.Register(&CustomJob{}, &CustomJobList{})
}

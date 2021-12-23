package common

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// TODO IMPLEMENT

func ToUnstructured(runtimeObj runtime.Object) (*unstructured.Unstructured, error) {
	unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(runtimeObj)
	return &unstructured.Unstructured{Object: unstructuredObj}, err
}

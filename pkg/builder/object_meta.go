/*
Copyright 2019 the Velero contributors.

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

package builder

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// ObjectMetaOpt is a functional option for ObjectMeta.
type ObjectMetaOpt func(metav1.Object)

// WithName is a functional option that applies the specified
// name to an object.
func WithName(val string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetName(val)
	}
}

// WithLabels is a functional option that applies the specified
// label keys/values to an object.
func WithLabels(vals ...string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetLabels(setMapEntries(obj.GetLabels(), vals...))
	}
}

// WithAnnotations is a functional option that applies the specified
// annotation keys/values to an object.
func WithAnnotations(vals ...string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetAnnotations(setMapEntries(obj.GetAnnotations(), vals...))
	}
}

func setMapEntries(m map[string]string, vals ...string) map[string]string {
	if m == nil {
		m = make(map[string]string)
	}

	// if we don't have a value for every key, add an empty
	// string at the end to serve as the value for the last
	// key.
	if len(vals)%2 != 0 {
		vals = append(vals, "")
	}

	for i := 0; i < len(vals); i += 2 {
		key := vals[i]
		val := vals[i+1]

		m[key] = val
	}

	return m
}

// WithClusterName is a functional option that applies the specified
// cluster name to an object.
func WithClusterName(val string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetClusterName(val)
	}
}

// WithFinalizers is a functional option that applies the specified
// finalizers to an object.
func WithFinalizers(vals ...string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetFinalizers(vals)
	}
}

// WithDeletionTimestamp is a functional option that applies the specified
// deletion timestamp to an object.
func WithDeletionTimestamp(val time.Time) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetDeletionTimestamp(&metav1.Time{Time: val})
	}
}

// WithUID is a functional option that applies the specified UID to an object.
func WithUID(val string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetUID(types.UID(val))
	}
}

// WithGenerateName is a functional option that applies the specified generate name to an object.
func WithGenerateName(val string) func(obj metav1.Object) {
	return func(obj metav1.Object) {
		obj.SetGenerateName(val)
	}
}

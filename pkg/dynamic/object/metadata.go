/*
Copyright 2018 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// TODO remove in favour of  controllerutil methods

package object

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

// HasFinalizer returns true if obj has the named finalizer.
// Deprecated: use controllerutil.ContainsFinalizer instead
func HasFinalizer(obj client.Object, name string) bool {
	return controllerutil.ContainsFinalizer(obj, name)
}

// AddFinalizer adds the named finalizer to obj, if it isn't already present.
// Deprecated: use controllerutil.AddFinalizer instead
func AddFinalizer(obj client.Object, name string) {
	controllerutil.AddFinalizer(obj, name)
}

// RemoveFinalizer removes the named finalizer from obj, if it's present.
// Deprecated: use controllerutil.RemoveFinalizer instead
func RemoveFinalizer(obj client.Object, name string) {
	controllerutil.RemoveFinalizer(obj, name)
}

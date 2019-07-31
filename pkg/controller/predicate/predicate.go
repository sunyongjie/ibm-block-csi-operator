/**
 * Copyright 2019 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package predicate

import (
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// CreateDeletePredicate implements a default predicate on resource creation or deletion events
type CreateDeletePredicate struct {
	predicate.Funcs
}

// no watch for update events
func (p CreateDeletePredicate) Update(e event.UpdateEvent) bool {
	return false
}

// no watch for generic events
func (p CreateDeletePredicate) Generic(e event.GenericEvent) bool {
	return false
}

// CreatePredicate implements a default predicate on resource creation events
type CreatePredicate struct {
	predicate.Funcs
}

// no watch for update events
func (p CreatePredicate) Update(e event.UpdateEvent) bool {
	return false
}

// no watch for update events
func (p CreatePredicate) Delete(e event.DeleteEvent) bool {
	return false
}

// no watch for generic events
func (p CreatePredicate) Generic(e event.GenericEvent) bool {
	return false
}

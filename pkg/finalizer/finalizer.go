/*
 * The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
 * (the "License"). You may not use this work except in compliance with the License, which is
 * available at www.apache.org/licenses/LICENSE-2.0
 *
 * This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied, as more fully set forth in the License.
 *
 * See the NOTICE file distributed with this work for information regarding copyright ownership.
 */

package finalizer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

const dummyFinalizer = "k8s-operator.alluxio.com/dummyFinalizer"

// Add this finalizer so that we can still access information of the deleted resources.
func AddDummyFinalizerIfNotExist(c client.Client, o client.Object, ctx context.Context) error {
	if !controllerutil.ContainsFinalizer(o, dummyFinalizer) {
		controllerutil.AddFinalizer(o, dummyFinalizer)
	}
	return c.Update(ctx, o)
}

func RemoveDummyFinalizerIfExist(c client.Client, o client.Object, ctx context.Context) error {
	if controllerutil.ContainsFinalizer(o, dummyFinalizer) {
		controllerutil.RemoveFinalizer(o, dummyFinalizer)
	}
	return c.Update(ctx, o)
}

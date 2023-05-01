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

package alluxiocluster

import (
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/alluxio/k8s-operator/pkg/logger"
)

const dummyFinalizer = "k8s-operator.alluxio.com/dummyFinalizer"

// We add this finalizer so that we can still access information of the deleted AlluxioCluster
func addDummyFinalizerIfNotExist(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) error {
	if !controllerutil.ContainsFinalizer(ctx.AlluxioCluster, dummyFinalizer) {
		controllerutil.AddFinalizer(ctx.AlluxioCluster, dummyFinalizer)
		err := r.Update(ctx, ctx.AlluxioCluster)
		if err != nil {
			logger.Errorf("Failed to add finalizer to alluxio cluster. %v", err)
			return err
		}
	}
	return nil
}

func removeDummyFinalizerIfExist(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) error {
	if controllerutil.ContainsFinalizer(ctx.AlluxioCluster, dummyFinalizer) {
		controllerutil.RemoveFinalizer(ctx.AlluxioCluster, dummyFinalizer)
		err := r.Update(ctx, ctx.AlluxioCluster)
		if err != nil {
			logger.Errorf("Failed to remove finalizer to alluxio cluster. Add it back. %v", err)
			if err := addDummyFinalizerIfNotExist(r, ctx); err != nil {
				return err
			}
		}
		return err
	}
	return nil
}

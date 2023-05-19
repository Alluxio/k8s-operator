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

package dataset

import (
	"reflect"

	ctrl "sigs.k8s.io/controller-runtime"

	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
)

func UpdateDatasetStatus(ctx DatasetReconcilerReqCtx) (ctrl.Result, error) {
	newestDataset := &alluxiov1alpha1.Dataset{}
	if err := ctx.Get(ctx, ctx.NamespacedName, newestDataset); err != nil {
		logger.Errorf("Error getting newest dataset status: %v", err)
		return ctrl.Result{}, err
	}
	if !reflect.DeepEqual(newestDataset.Status, ctx.Dataset.Status) {
		newestDataset.Status = ctx.Dataset.Status
		if err := ctx.Client.Status().Update(ctx.Context, newestDataset); err != nil {
			logger.Errorf("Error updating dataset %s status: %v", ctx.NamespacedName.String(), err)
			return ctrl.Result{}, err
		}
	}
	return ctrl.Result{}, nil
}

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
	"context"
	alluxiocomv1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// DatasetReconciler reconciles a Dataset object
type DatasetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type DatasetReconcilerReqCtx struct {
	*alluxiocomv1alpha1.Dataset
	client.Client
	context.Context
	types.NamespacedName
}

func (r *DatasetReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx := DatasetReconcilerReqCtx{
		Client:         r.Client,
		Context:        context,
		NamespacedName: req.NamespacedName,
	}

	dataset := &alluxiocomv1alpha1.Dataset{}
	ctx.Dataset = dataset

	if err := r.Get(context, req.NamespacedName, dataset); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("Data set %v in namespace %v not found. It is being deleted or already deleted.", req.Name, req.Namespace)
		} else {
			logger.Errorf("Failed to get data set %v in namespace %v: %v", req.Name, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}

	if dataset.ObjectMeta.UID == "" {
		return DeleteDatasetIfExist(req)
	}
	if dataset.Status.Phase == alluxiocomv1alpha1.DatasetPhaseNone {
		dataset.Status.Phase = alluxiocomv1alpha1.DatasetPhasePending
		return UpdateDatasetStatus(ctx)
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DatasetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alluxiocomv1alpha1.Dataset{}).
		Complete(r)
}

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
	"context"
	alluxiocomv1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

// AlluxioClusterReconciler reconciles a AlluxioCluster object
type AlluxioClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type AlluxioClusterReconcileReqCtx struct {
	*alluxiocomv1alpha1.AlluxioCluster
	client.Client
	context.Context
	*alluxiocomv1alpha1.Dataset
	types.NamespacedName
}

func (r *AlluxioClusterReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger.Infof("Reconciling AlluxioCluster. Name: %s. Namespace: %s", req.NamespacedName.Name, req.NamespacedName.Namespace)
	ctx := AlluxioClusterReconcileReqCtx{
		Client:         r.Client,
		Context:        context,
		NamespacedName: req.NamespacedName,
	}

	alluxioCluster := &alluxiocomv1alpha1.AlluxioCluster{}
	ctx.AlluxioCluster = alluxioCluster
	if err := r.Get(context, req.NamespacedName, alluxioCluster); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("Alluxio cluster %v in namespace %v not found. It is being deleted or already deleted.", req.Name, req.Namespace)
		} else {
			logger.Errorf("Failed to get Alluxio cluster %v in namespace %v: %v", req.Name, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}

	if alluxioCluster.ObjectMeta.UID == "" {
		if err := deleteConfYamlFile(ctx); err != nil {
			logger.Infof("Failed to delete configuration yaml file. You may need to manually delete it to avoid unexpected behavior. %v", err)
			return ctrl.Result{}, err
		}
		return DeleteAlluxioClusterIfExist(ctx)
	}

	datasetToBound := &alluxiocomv1alpha1.Dataset{}
	ctx.Dataset = datasetToBound
	if err := r.Get(context, req.NamespacedName, datasetToBound); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("Dataset %v in namespace %v not found. It is deleted or hasn't been created yet.", req.Name, req.Namespace)
			return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
		} else {
			logger.Errorf("Failed to get Dataset %v in namespace %s: %s", req.Name, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}
	if alluxioCluster.Status.Phase == alluxiocomv1alpha1.ClusterPhaseNone {
		if err := CreateAlluxioClusterIfNotExist(ctx); err != nil {
			return ctrl.Result{}, err
		}
		return UpdateStatus(r, ctx, datasetToBound)
	}

	return UpdateStatus(r, ctx, datasetToBound)
}

// SetupWithManager sets up the controller with the Manager.
func (r *AlluxioClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alluxiocomv1alpha1.AlluxioCluster{}).
		Complete(r)
}

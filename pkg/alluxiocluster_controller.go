/*
Copyright 2023.

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

package pkg

import (
	"context"
	alluxiocomv1alpha1 "github.com/Alluxio/k8s-operator/api/v1alpha1"
	"github.com/Alluxio/k8s-operator/pkg/logger"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AlluxioClusterReconciler reconciles a AlluxioCluster object
type AlluxioClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type ReconcileRequestContext struct {
	*alluxiocomv1alpha1.AlluxioCluster
	client.Client
	context.Context
	types.NamespacedName
}

func (r *AlluxioClusterReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger.Infof("Reconciling. Name: %v. Namespace: %v", req.Name, req.Namespace)
	ctx := ReconcileRequestContext{
		Client:         r.Client,
		Context:        context,
		NamespacedName: req.NamespacedName,
	}

	alluxioCluster := &alluxiocomv1alpha1.AlluxioCluster{}
	ctx.AlluxioCluster = alluxioCluster

	err := r.Get(ctx, req.NamespacedName, alluxioCluster)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("Alluxio cluster %v in namespace %v not found. It is being deleted or already deleted.", req.Name, req.Namespace)
		} else {
			logger.Errorf("Failed to get Alluxio cluster %v in namespace %v: %v", req.Name, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}

	if alluxioCluster.ObjectMeta.UID == "" {
		return DeleteAlluxioClusterIfExist(ctx)
	}

	if alluxioCluster.Status.Phase == alluxiocomv1alpha1.ClusterPhaseNone {
		err = CreateAlluxioClusterIfNotExist(ctx)
		if err != nil {
			return ctrl.Result{}, err
		}
		return ctx.UpdateClusterStatus(r)
	}

	return ctx.UpdateClusterStatus(r)
}

// SetupWithManager sets up the controller with the Manager.
func (r *AlluxioClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alluxiocomv1alpha1.AlluxioCluster{}).
		Complete(r)
}

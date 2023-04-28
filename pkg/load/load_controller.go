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

package load

import (
	"context"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
)

type LoadReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

type LoadReconcilerReqCtx struct {
	*alluxiov1alpha1.AlluxioCluster
	*alluxiov1alpha1.Load
	client.Client
	context.Context
	types.NamespacedName
}

func (r *LoadReconciler) Reconcile(context context.Context, req ctrl.Request) (ctrl.Result, error) {
	ctx := LoadReconcilerReqCtx{
		Client:         r.Client,
		Context:        context,
		NamespacedName: req.NamespacedName,
	}
	load := &alluxiov1alpha1.Load{}
	if err := r.Get(context, req.NamespacedName, load); err != nil {
		if errors.IsNotFound(err) {
			logger.Infof("Load object %v in namespace %v not found. It is being deleted or already deleted.", req.Name, req.Namespace)
		} else {
			logger.Errorf("Failed to get load job %v in namespace %v: %v", req.Name, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}
	ctx.Load = load

	if load.ObjectMeta.UID == "" {
		// TODO: shall we stop the load if still loading?
		return r.deleteJob(ctx)
	}

	alluxioCluster := &alluxiov1alpha1.AlluxioCluster{}
	alluxioNamespacedName := types.NamespacedName{
		Namespace: req.Namespace,
		Name:      load.Spec.Dataset,
	}
	if err := r.Get(context, alluxioNamespacedName, alluxioCluster); err != nil {
		if errors.IsNotFound(err) {
			logger.Errorf("Dataset %s in namespace %v is not found. Please double check your configuration.", load.Spec.Dataset, req.Namespace)
			load.Status.Phase = alluxiov1alpha1.LoadPhaseFailed
			return r.updateLoadStatus(ctx)
		} else {
			logger.Errorf("Error getting alluxio cluster %s in namespace %s: %v", load.Spec.Dataset, req.Namespace, err)
			return ctrl.Result{}, err
		}
	}
	ctx.AlluxioCluster = alluxioCluster

	if alluxioCluster.Status.Phase != alluxiov1alpha1.ClusterPhaseReady {
		load.Status.Phase = alluxiov1alpha1.LoadPhaseWaiting
		return r.updateLoadStatus(ctx)
	}

	switch load.Status.Phase {
	case alluxiov1alpha1.LoadPhaseNone, alluxiov1alpha1.LoadPhaseWaiting:
		return r.createLoadJob(ctx)
	case alluxiov1alpha1.LoadPhaseLoading:
		return r.waitLoadJobFinish(ctx)
	default:
		return ctrl.Result{}, nil
	}
}

func (r *LoadReconciler) waitLoadJobFinish(ctx LoadReconcilerReqCtx) (ctrl.Result, error) {
	loadJob, err := r.getLoadJob(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	if loadJob.Status.Succeeded == 1 {
		ctx.Load.Status.Phase = alluxiov1alpha1.LoadPhaseLoaded
		if _, err := r.updateLoadStatus(ctx); err != nil {
			logger.Errorf("Data is loaded but failed to update status. %v", err)
			return ctrl.Result{Requeue: true}, err
		}
		return ctrl.Result{}, nil
	} else if loadJob.Status.Failed == 1 {
		ctx.Load.Status.Phase = alluxiov1alpha1.LoadPhaseFailed
		if _, err := r.updateLoadStatus(ctx); err != nil {
			logger.Errorf("Failed to update status. %v", err)
			return ctrl.Result{Requeue: true}, err
		}
		logger.Errorf("Load data job failed. Please check the log of the pod for errors.")
		return ctrl.Result{}, nil
	} else {
		return ctrl.Result{RequeueAfter: 15 * time.Second}, nil
	}
}

func (r *LoadReconciler) getLoadJob(ctx LoadReconcilerReqCtx) (*batchv1.Job, error) {
	loadJob := &batchv1.Job{}
	loadJobNamespacedName := types.NamespacedName{
		Name:      utils.GetLoadJobName(ctx.Name),
		Namespace: ctx.Namespace,
	}
	if err := r.Get(ctx.Context, loadJobNamespacedName, loadJob); err != nil {
		logger.Errorf("Error getting load job %s: %v", ctx.NamespacedName.String(), err)
		return nil, err
	}
	return loadJob, nil
}

func (r *LoadReconciler) updateLoadStatus(ctx LoadReconcilerReqCtx) (ctrl.Result, error) {
	if err := r.Client.Status().Update(ctx.Context, ctx.Load); err != nil {
		logger.Errorf("Failed updating load job status: %v", err)
		return ctrl.Result{}, err
	}
	return ctrl.Result{RequeueAfter: 15 * time.Second}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *LoadReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&alluxiov1alpha1.Load{}).
		Complete(r)
}

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
	"reflect"
	"time"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/dataset"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
)

type componentStatus struct {
	Master *v1.StatefulSet `json:"master"`
	Worker *v1.Deployment  `json:"worker"`
	Fuse   *v1.DaemonSet   `json:"fuse,omitempty"`
	Proxy  *v1.DaemonSet   `json:"proxy,omitempty"`
}

func UpdateStatus(r *AlluxioClusterReconciler,
	alluxioClusterCtx AlluxioClusterReconcileReqCtx, datasetToUpdate *alluxiov1alpha1.Dataset) (ctrl.Result, error) {
	alluxioOriginalStatusCopy := alluxioClusterCtx.AlluxioCluster.Status.DeepCopy()
	datasetOriginalStatusCopy := datasetToUpdate.Status.DeepCopy()

	alluxioClusterNewPhase := alluxioOriginalStatusCopy.Phase
	datasetNewPhase := datasetToUpdate.Status.Phase

	if alluxioOriginalStatusCopy.Phase == alluxiov1alpha1.ClusterPhaseNone {
		alluxioClusterNewPhase = alluxiov1alpha1.ClusterPhaseCreatingOrUpdating
		datasetNewPhase = alluxiov1alpha1.DatasetPhaseBounding
	} else {
		if ClusterReady(r, alluxioClusterCtx) {
			alluxioClusterNewPhase = alluxiov1alpha1.ClusterPhaseReady
			datasetNewPhase = alluxiov1alpha1.DatasetPhaseReady
		} else {
			alluxioClusterNewPhase = alluxiov1alpha1.ClusterPhaseCreatingOrUpdating
			datasetNewPhase = alluxiov1alpha1.DatasetPhaseBounding
		}
	}
	alluxioClusterCtx.AlluxioCluster.Status.Phase = alluxioClusterNewPhase
	datasetToUpdate.Status.Phase = datasetNewPhase

	if !reflect.DeepEqual(alluxioOriginalStatusCopy, alluxioClusterCtx.AlluxioCluster.Status) {
		if err := r.Client.Status().Update(alluxioClusterCtx.Context, alluxioClusterCtx.AlluxioCluster); err != nil {
			logger.Errorf("Error updating cluster status: %v", err)
			return ctrl.Result{}, err
		}
	}
	if !reflect.DeepEqual(datasetOriginalStatusCopy, datasetToUpdate.Status) {
		if err := updateDatasetStatus(alluxioClusterCtx); err != nil {
			return ctrl.Result{}, err
		}
	}

	if alluxioClusterNewPhase != alluxiov1alpha1.ClusterPhaseReady {
		return ctrl.Result{RequeueAfter: 15 * time.Second}, nil
	}
	return ctrl.Result{RequeueAfter: 2 * time.Minute}, nil
}

func ClusterReady(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) bool {
	master, err := GetMasterStatus(r, ctx)
	if err != nil || master.Status.AvailableReplicas != master.Status.Replicas {
		return false
	}
	worker, err := GetWorkerStatus(r, ctx)
	if err != nil || worker.Status.AvailableReplicas != worker.Status.Replicas {
		return false
	}
	if ctx.AlluxioCluster.Spec.Fuse.Enabled {
		fuse, err := GetFuseStatus(r, ctx)
		if err != nil || fuse.Status.NumberAvailable != fuse.Status.DesiredNumberScheduled {
			return false
		}
	}
	if ctx.AlluxioCluster.Spec.Proxy.Enabled {
		proxy, err := GetProxyStatus(r, ctx)
		if err != nil || proxy.Status.NumberAvailable != proxy.Status.DesiredNumberScheduled {
			return false
		}
	}
	return true
}

func GetMasterStatus(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) (*v1.StatefulSet, error) {
	master := &v1.StatefulSet{}
	if err := r.Get(ctx.Context, utils.GetMasterStatefulSetNamespacedName(ctx.AlluxioCluster.Spec.NameOverride, ctx.NamespacedName), master); err != nil {
		logger.Errorf("Error getting Alluxio master StatefulSet from k8s api server: %v", err)
		return nil, err
	}
	return master, nil
}

func GetWorkerStatus(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) (*v1.Deployment, error) {
	worker := &v1.Deployment{}
	if err := r.Get(ctx.Context, utils.GetWorkerDeploymentNamespacedName(ctx.AlluxioCluster.Spec.NameOverride, ctx.NamespacedName), worker); err != nil {
		logger.Errorf("Error getting Alluxio worker Deployment from k8s api server: %v", err)
		return nil, err
	}
	return worker, nil
}

func GetFuseStatus(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) (*v1.DaemonSet, error) {
	fuse := &v1.DaemonSet{}
	if err := r.Get(ctx.Context, utils.GetFuseDaemonSetNamespacedName(ctx.AlluxioCluster.Spec.NameOverride, ctx.NamespacedName), fuse); err != nil {
		logger.Errorf("Error getting Alluxio fuse DaemonSet from k8s api server: %v", err)
		return nil, err
	}
	return fuse, nil
}

func GetProxyStatus(r *AlluxioClusterReconciler, ctx AlluxioClusterReconcileReqCtx) (*v1.DaemonSet, error) {
	proxy := &v1.DaemonSet{}
	if err := r.Get(ctx.Context, utils.GetProxyDaemonSetNamespacedName(ctx.AlluxioCluster.Spec.NameOverride, ctx.NamespacedName), proxy); err != nil {
		logger.Errorf("Error getting Alluxio proxy DaemonSet from k8s api server: %v", err)
		return nil, err
	}
	return proxy, nil
}

func updateDatasetStatus(alluxioClusterCtx AlluxioClusterReconcileReqCtx) error {
	datasetCtx := dataset.DatasetReconcilerReqCtx{
		Dataset: alluxioClusterCtx.Dataset,
		Client:  alluxioClusterCtx.Client,
		Context: alluxioClusterCtx.Context,
		NamespacedName: types.NamespacedName{
			Name:      alluxioClusterCtx.Dataset.Name,
			Namespace: alluxioClusterCtx.Namespace,
		},
	}
	if _, err := dataset.UpdateDatasetStatus(datasetCtx); err != nil {
		return err
	}
	return nil
}

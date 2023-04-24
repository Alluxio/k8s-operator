package pkg

import (
	alluxiocomv1alpha1 "github.com/Alluxio/k8s-operator/api/v1alpha1"
	"github.com/Alluxio/k8s-operator/pkg/logger"
	v1 "k8s.io/api/apps/v1"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"time"
)

func (ctx ReconcileRequestContext) UpdateClusterStatus(r *AlluxioClusterReconciler) (ctrl.Result, error) {
	originalStatusCopy := ctx.AlluxioCluster.DeepCopy().Status

	if ctx.Spec.Master.Enabled {
		ctx.GetMasterStatus(r)
	}
	ctx.GetWorkerStatus(r)
	if ctx.Spec.Fuse.Enabled {
		ctx.GetFuseStatus(r)
	}
	if ctx.Spec.Proxy.Enabled {
		ctx.GetProxyStatus(r)
	}

	newPhase := originalStatusCopy.Phase
	if originalStatusCopy.Phase == alluxiocomv1alpha1.ClusterPhaseNone {
		newPhase = alluxiocomv1alpha1.ClusterPhaseCreatingOrUpdating
	} else if originalStatusCopy.Phase == alluxiocomv1alpha1.ClusterPhaseCreatingOrUpdating {
		newPhase = alluxiocomv1alpha1.ClusterPhaseReady
	}
	if !ClusterReady(ctx.AlluxioCluster.Status) {
		newPhase = alluxiocomv1alpha1.ClusterPhaseCreatingOrUpdating
	}
	ctx.AlluxioCluster.Status.Phase = newPhase

	if !reflect.DeepEqual(originalStatusCopy, ctx.AlluxioCluster.Status) {
		if err := r.Client.Status().Update(ctx.Context, ctx.AlluxioCluster); err != nil {
			logger.Errorf("Error updating cluster status: %v", err)
			return ctrl.Result{}, err
		}
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}
	return ctrl.Result{RequeueAfter: 2 * time.Minute}, nil
}

func (ctx ReconcileRequestContext) GetMasterStatus(r *AlluxioClusterReconciler) error {
	master := &v1.StatefulSet{}
	if err := r.Get(ctx.Context, ctx.getMasterStatefulSetNamespacedName(), master); err != nil {
		logger.Errorf("Error getting Alluxio master StatefulSet from k8s api server: %v", err)
		return err
	}
	ctx.AlluxioCluster.Status.Master = master
	return nil
}

func (ctx ReconcileRequestContext) GetWorkerStatus(r *AlluxioClusterReconciler) error {
	worker := &v1.Deployment{}
	if err := r.Get(ctx.Context, ctx.getWorkerDeploymentNamespacedName(), worker); err != nil {
		logger.Errorf("Error getting Alluxio worker Deployment from k8s api server: %v", err)
		return err
	}
	ctx.AlluxioCluster.Status.Worker = worker
	return nil
}

func (ctx ReconcileRequestContext) GetFuseStatus(r *AlluxioClusterReconciler) error {
	fuse := &v1.DaemonSet{}
	if err := r.Get(ctx.Context, ctx.getMasterStatefulSetNamespacedName(), fuse); err != nil {
		logger.Errorf("Error getting Alluxio fuse DaemonSet from k8s api server: %v", err)
		return err
	}
	ctx.AlluxioCluster.Status.Fuse = fuse
	return nil
}

func (ctx ReconcileRequestContext) GetProxyStatus(r *AlluxioClusterReconciler) error {
	proxy := &v1.DaemonSet{}
	if err := r.Get(ctx.Context, ctx.getMasterStatefulSetNamespacedName(), proxy); err != nil {
		logger.Errorf("Error getting Alluxio proxy DaemonSet from k8s api server: %v", err)
		return err
	}
	ctx.AlluxioCluster.Status.Proxy = proxy
	return nil
}

func ClusterReady(status alluxiocomv1alpha1.AlluxioClusterStatus) bool {
	ready := true
	if status.Master != nil && status.Master.Status.AvailableReplicas != status.Master.Status.Replicas {
		ready = false
	}
	if status.Worker.Status.AvailableReplicas != status.Worker.Status.Replicas {
		ready = false
	}
	if status.Fuse != nil && status.Fuse.Status.NumberAvailable != status.Fuse.Status.DesiredNumberScheduled {
		ready = false
	}
	if status.Proxy != nil && status.Proxy.Status.NumberAvailable != status.Proxy.Status.DesiredNumberScheduled {
		ready = false
	}
	return ready
}

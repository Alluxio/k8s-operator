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
	"github.com/alluxio/k8s-operator/pkg/logger"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *LoadReconciler) deleteJob(ctx LoadReconcilerReqCtx) (ctrl.Result, error) {
	loadJob, err := r.getLoadJob(ctx)
	if err != nil {
		return ctrl.Result{}, err
	}
	propagationPolicy := metav1.DeletePropagationBackground // for deleting the pod along with the job
	if err := r.Delete(ctx.Context, loadJob, &client.DeleteOptions{PropagationPolicy: &propagationPolicy}); err != nil {
		logger.Errorf("Error deleting load job %s in namespace %s: %v", loadJob.Name, loadJob.Namespace, err)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

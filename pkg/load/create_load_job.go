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
	"fmt"
	"os"
	"time"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
)

func (r *LoadReconciler) createLoadJob(ctx LoadReconcilerReqCtx) (ctrl.Result, error) {
	// Update the status before job creation instead of after, because otherwise if the status update fails,
	// the reconciler will loop again and create another same job, leading to failure to create duplicated job which is confusing.
	ctx.Load.Status.Phase = alluxiov1alpha1.LoadPhaseLoading
	_, err := r.updateLoadStatus(ctx)
	if err != nil {
		logger.Infof("Job is pending because status was not updated successfully")
		return ctrl.Result{}, err
	}
	loadJob, err := getLoadJobFromYaml()
	if err != nil {
		return ctrl.Result{}, err
	}
	constructLoadJob(ctx.AlluxioCluster, ctx.Load, loadJob)
	if err := r.Create(ctx.Context, loadJob); err != nil {
		logger.Errorf("Failed to load data of dataset %s: %v", ctx.NamespacedName.String(), err)
		return ctrl.Result{}, err
	}
	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

func getLoadJobFromYaml() (*batchv1.Job, error) {
	loadJobYaml, err := os.ReadFile("/opt/alluxio-jobs/load.yaml")
	if err != nil {
		logger.Errorf("Failed to read load job yaml file at /opt/alluxio-jobs/load.yaml: %v", err)
		return nil, err
	}
	loadJob, _, err := scheme.Codecs.UniversalDeserializer().Decode(loadJobYaml, nil, nil)
	if err != nil {
		logger.Errorf("Failed to parse load job yaml file: %v", err)
	}
	return loadJob.(*batchv1.Job), nil
}

func constructLoadJob(alluxio *alluxiov1alpha1.AlluxioCluster, load *alluxiov1alpha1.Load, loadJob *batchv1.Job) {
	loadJob.Name = utils.GetLoadJobName(load.Name)
	loadJob.Namespace = alluxio.Namespace
	var imagePullSecrets []corev1.LocalObjectReference
	for _, secret := range alluxio.Spec.ImagePullSecrets {
		imagePullSecrets = append(imagePullSecrets, corev1.LocalObjectReference{Name: secret})

	}
	loadJob.Spec.Template.Spec.ImagePullSecrets = imagePullSecrets
	loadJob.Spec.Template.Spec.ServiceAccountName = alluxio.Spec.ServiceAccountName
	loadJob.Spec.Template.Spec.Containers[0].Image = fmt.Sprintf("%s:%s", alluxio.Spec.Image, alluxio.Spec.ImageTag)
	loadJob.Spec.Template.Spec.Containers[0].Command = []string{"go", "run", "/load.go", load.Spec.Path}
	alluxioConfigMapName := utils.GetAlluxioConfigMapName(alluxio.Spec.NameOverride, alluxio.Name)
	loadConfigMapName := utils.GetLoadConfigmapName(alluxio.Spec.NameOverride, alluxio.Name)
	loadJob.Spec.Template.Spec.Containers[0].VolumeMounts = []corev1.VolumeMount{
		{
			Name:      alluxioConfigMapName,
			MountPath: "/opt/alluxio/conf",
		},
		{
			Name:      loadConfigMapName,
			MountPath: "/load.sh",
		},
	}
	loadJob.Spec.Template.Spec.Volumes = []corev1.Volume{
		{
			Name: alluxioConfigMapName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: alluxioConfigMapName,
					},
				},
			},
		},
		{
			Name: loadConfigMapName,
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: loadConfigMapName,
					},
				},
			},
		},
	}
}

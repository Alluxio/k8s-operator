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
	"os"

	ctrl "sigs.k8s.io/controller-runtime"

	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
)

func DeleteAlluxioClusterIfExist(ctx AlluxioClusterReconcileReqCtx) (ctrl.Result, error) {
	logger.Infof("Uninstalling Alluxio cluster %s.", ctx.NamespacedName.String())

	helmCtx := utils.HelmContext{
		ReleaseName: ctx.Name,
		Namespace:   ctx.Namespace,
	}
	if err := utils.HelmDeleteIfExist(helmCtx); err != nil {
		logger.Errorf("failed to delete helm release %s: %v", ctx.NamespacedName.String(), err)
		return ctrl.Result{}, err
	}

	ctx.Dataset.Status.Phase = alluxiov1alpha1.DatasetPhasePending
	if err := updateDatasetStatus(ctx); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func deleteConfYamlFile(ctx AlluxioClusterReconcileReqCtx) error {
	confYamlFilePath := utils.GetConfYamlPath(ctx.NamespacedName)
	if err := os.Remove(confYamlFilePath); err != nil {
		return err
	}
	return nil
}

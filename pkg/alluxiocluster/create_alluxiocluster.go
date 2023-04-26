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
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
	"os"
	"sigs.k8s.io/yaml"
)

const chartPath = "/opt/alluxio-helm-chart"

func CreateAlluxioClusterIfNotExist(ctx AlluxioClusterReconcileReqCtx) error {
	// if the release has already been deployed, requeue without further actions
	helmCtx := utils.HelmContext{
		Namespace:   ctx.Namespace,
		ReleaseName: ctx.Name,
	}
	exists, err := utils.IfHelmReleaseExists(helmCtx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	logger.Infof("Creating Alluxio cluster %v in namespace %v.", ctx.Name, ctx.Namespace)
	// Construct config.yaml file
	clusterYaml, err := yaml.Marshal(ctx.AlluxioCluster.Spec)
	if err != nil {
		return err
	}
	confYamlFilePath := utils.GetConfYamlPath(ctx.NamespacedName)
	confYamlFile, err := os.OpenFile(confYamlFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	defer confYamlFile.Close()
	if err != nil {
		logger.Errorf("failed to create empty config file: %v", err)
		return err
	}
	if _, err := confYamlFile.Write(clusterYaml); err != nil {
		logger.Errorf("Error writing to config file: %v", err)
		return err
	}
	datasetYalm, err := yaml.Marshal(ctx.Dataset.Spec)
	if err != nil {
		return err
	}
	if _, err = confYamlFile.WriteString(string(datasetYalm)); err != nil {
		logger.Errorf("Error writing to config file: %v", err)
		return err
	}
	// helm install release with the constructed config.yaml
	helmCtx = utils.HelmContext{
		HelmChartPath:  chartPath,
		ConfigFilePath: confYamlFilePath,
		Namespace:      ctx.Namespace,
		ReleaseName:    ctx.Name,
	}
	if err := utils.HelmInstall(helmCtx); err != nil {
		logger.Errorf("error installing helm release. Uninstalling...")
		if _, err := DeleteAlluxioClusterIfExist(ctx); err != nil {
			logger.Errorf("failed to delete failed helm release %v in namespace %v: %v", ctx.Name, ctx.Namespace, err)
			return err
		}
	}
	return nil
}

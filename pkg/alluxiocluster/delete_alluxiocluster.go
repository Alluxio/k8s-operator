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

	"k8s.io/apimachinery/pkg/types"

	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/alluxio/k8s-operator/pkg/utils"
)

func DeleteAlluxioClusterIfExist(namespacedName types.NamespacedName) error {
	logger.Infof("Uninstalling Alluxio cluster %s.", namespacedName.String())

	helmCtx := utils.HelmContext{
		ReleaseName: namespacedName.Name,
		Namespace:   namespacedName.Namespace,
	}
	if err := utils.HelmDeleteIfExist(helmCtx); err != nil {
		logger.Errorf("failed to delete helm release %s: %v", namespacedName.String(), err)
		return err
	}
	return nil
}

func deleteConfYamlFileIfExist(namespacedName types.NamespacedName) error {
	confYamlFilePath := utils.GetConfYamlPath(namespacedName)
	if _, err := os.Stat(confYamlFilePath); err != nil && os.IsNotExist(err) {
		logger.Errorf("Error getting information of configuration yaml file. %v", err)
		return err
	}
	if err := os.Remove(confYamlFilePath); err != nil {
		logger.Infof("Failed to delete configuration yaml file. You may need to manually delete it to avoid unexpected behavior. %v", err)
		return err
	}
	return nil
}

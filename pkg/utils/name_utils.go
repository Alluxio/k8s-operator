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

package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/apimachinery/pkg/types"
	k8sStringsUtil "k8s.io/utils/strings"
)

func GetMasterStatefulSetNamespacedName(nameOverride string, clusterNamespacedName types.NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Namespace: clusterNamespacedName.Namespace,
		Name:      fmt.Sprintf("%s-master", getClusterFullName(nameOverride, clusterNamespacedName.Name)),
	}
}

func GetWorkerDeploymentNamespacedName(nameOverride string, clusterNamespacedName types.NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Namespace: clusterNamespacedName.Namespace,
		Name:      fmt.Sprintf("%s-worker", getClusterFullName(nameOverride, clusterNamespacedName.Name)),
	}
}

func GetFuseDaemonSetNamespacedName(nameOverride string, clusterNamespacedName types.NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Namespace: clusterNamespacedName.Namespace,
		Name:      fmt.Sprintf("%s-fuse", getClusterFullName(nameOverride, clusterNamespacedName.Name)),
	}
}

func GetProxyDaemonSetNamespacedName(nameOverride string, clusterNamespacedName types.NamespacedName) types.NamespacedName {
	return types.NamespacedName{
		Namespace: clusterNamespacedName.Namespace,
		Name:      fmt.Sprintf("%s-proxy", getClusterFullName(nameOverride, clusterNamespacedName.Name)),
	}
}

func GetAlluxioConfigMapName(nameOverride, clusterName string) string {
	return fmt.Sprintf("%s-alluxio-conf", getClusterFullName(nameOverride, clusterName))
}

func GetLoadJobName(loadName string) string {
	return fmt.Sprintf("%s-load-job", loadName)
}

// The same function that constructs alluxio.fullName in helm chart
func getClusterFullName(nameOverride, helmReleaseName string) string {
	if nameOverride == "" {
		nameOverride = "alluxio"
	} else {
		nameOverride = strings.TrimSuffix(k8sStringsUtil.ShortenString(nameOverride, 63), "-")
	}
	if strings.Contains(helmReleaseName, nameOverride) {
		return helmReleaseName
	} else {
		fullName := fmt.Sprintf("%v-%v", helmReleaseName, nameOverride)
		return strings.TrimSuffix(k8sStringsUtil.ShortenString(fullName, 32), "-")
	}
}

func GetConfYamlPath(namespacedName types.NamespacedName) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("%s-%s-config.yaml", namespacedName.Namespace, namespacedName.Name))
}

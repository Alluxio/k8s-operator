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
	"github.com/alluxio/k8s-operator/pkg/logger"
	"os/exec"
)

type HelmContext struct {
	HelmChartPath  string
	ConfigFilePath string
	Namespace      string
	ReleaseName    string
}

func HelmInstall(ctx HelmContext) error {
	args := []string{"install", "-f", ctx.ConfigFilePath, "--namespace", ctx.Namespace, ctx.ReleaseName, ctx.HelmChartPath}
	if _, err := executeHelmCommand(args); err != nil {
		logger.Errorf("Error installing Helm release with name %v in namespace %v.", ctx.ReleaseName, ctx.Namespace)
		return err
	}
	return nil
}

func HelmDeleteIfExist(ctx HelmContext) error {
	exists, err := IfHelmReleaseExists(ctx)
	if err != nil {
		logger.Errorf("Error checking if Helm release with name %v in namespace %v exists", ctx.ReleaseName, ctx.Namespace)
		return err
	}
	if exists {
		args := []string{"delete", ctx.ReleaseName, "-n", ctx.Namespace}
		if _, err := executeHelmCommand(args); err != nil {
			logger.Errorf("Error deleting Helm release with name %v in namespace %v.", ctx.ReleaseName, ctx.Namespace)
			return err
		}
	}
	return nil
}

func IfHelmReleaseExists(ctx HelmContext) (bool, error) {
	args := []string{"list", "--filter", ctx.ReleaseName, "-n", ctx.Namespace, "-q"}
	output, err := executeHelmCommand(args)
	if err != nil {
		logger.Errorf("Error listing Helm release with name %v in namespace %v", ctx.ReleaseName, ctx.Namespace)
		return false, err
	}
	if len(output) != 0 {
		return true, nil
	}
	return false, nil
}

func executeHelmCommand(args []string) (string, error) {
	cmd := exec.Command("helm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Error running command %v. Output: %v. Error: %v", cmd.String(), string(out), err)
		return "", err
	}
	return string(out), nil
}

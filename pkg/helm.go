package pkg

import (
	"github.com/Alluxio/k8s-operator/pkg/logger"
	"os/exec"
)

type HelmContext struct {
	HelmChartPath      string
	ConfigFilePath     string
	Namespace          string
	OverrideProperties map[string]string
	ReleaseName        string
}

func HelmInstall(ctx HelmContext) error {
	args := []string{"install", "-f", ctx.ConfigFilePath, "--namespace", ctx.Namespace, ctx.ReleaseName, ctx.HelmChartPath}
	if err := executeHelmCommand(args); err != nil {
		logger.Errorf("Error installing Helm release with name %v in namespace %v.", ctx.ReleaseName, ctx.Namespace)
		return err
	}
	return nil
}

func HelmDeleteIfExist(ctx HelmContext) error {
	if err := HelmStatus(ctx); err == nil {
		args := []string{"delete", ctx.ReleaseName, "-n", ctx.Namespace}
		if err := executeHelmCommand(args); err != nil {
			logger.Errorf("Error deleting Helm release with name %v in namespace %v", ctx.ReleaseName, ctx.Namespace)
			return err
		}
	}
	return nil
}

func HelmStatus(ctx HelmContext) error {
	args := []string{"status", ctx.ReleaseName, "-n", ctx.Namespace}
	if err := executeHelmCommand(args); err != nil {
		logger.Errorf("Error checking status of Helm release with name %v in namespace %v", ctx.ReleaseName, ctx.Namespace)
		return err
	}
	return nil
}

func executeHelmCommand(args []string) error {
	cmd := exec.Command("helm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Errorf("Error running command %v. Output: %v. Error: %v", cmd.String(), string(out), err)
		return err
	}
	return nil
}

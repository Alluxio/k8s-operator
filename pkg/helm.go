package pkg

import (
	"fmt"
	"github.com/go-logr/logr"
	"os/exec"
)

type HelmContext struct {
	HelmChartPath      string
	ConfigFilePath     string
	Namespace          string
	OverrideProperties map[string]string
	ReleaseName        string
}

func HelmInstall(ctx HelmContext, logger logr.Logger) error {
	args := []string{"install", "-f", ctx.ConfigFilePath, "--namespace", ctx.Namespace, ctx.ReleaseName, ctx.HelmChartPath}
	if err := executeHelmCommand(args, logger); err != nil {
		logger.Error(err, fmt.Sprintf("Error installing Helm release with name %v in namespace %v.", ctx.ReleaseName, ctx.Namespace))
		return err
	}
	return nil
}

func HelmDeleteIfExist(ctx HelmContext, logger logr.Logger) error {
	if err := HelmStatus(ctx, logger); err == nil {
		args := []string{"delete", ctx.ReleaseName, "-n", ctx.Namespace}
		if err := executeHelmCommand(args, logger); err != nil {
			logger.Error(err, fmt.Sprintf("Error deleting Helm release with name %v in namespace %v", ctx.ReleaseName, ctx.Namespace))
			return err
		}
	}
	return nil
}

func HelmStatus(ctx HelmContext, logger logr.Logger) error {
	args := []string{"status", ctx.ReleaseName, "-n", ctx.Namespace}
	if err := executeHelmCommand(args, logger); err != nil {
		logger.Error(err, fmt.Sprintf("Error checking status of Helm release with name %v in namespace %v", ctx.ReleaseName, ctx.Namespace))
		return err
	}
	return nil
}

func executeHelmCommand(args []string, logger logr.Logger) error {
	cmd := exec.Command("helm", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Error running command %v. Output: %v", cmd.String(), string(out)))
		return err
	}
	return nil
}

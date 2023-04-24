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
	if _, err := executeHelmCommand(args); err != nil {
		logger.Errorf("Error installing Helm release with name %v in namespace %v.", ctx.ReleaseName, ctx.Namespace)
		return err
	}
	return nil
}

func HelmDeleteIfExist(ctx HelmContext) error {
	exists, err := checkIfHelmReleaseExists(ctx)
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

func checkIfHelmReleaseExists(ctx HelmContext) (bool, error) {
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

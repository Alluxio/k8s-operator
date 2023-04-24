package pkg

import (
	"fmt"
	"github.com/Alluxio/k8s-operator/pkg/logger"
	"os"
	"sigs.k8s.io/yaml"
)

const chartPath = "/opt/alluxio-helm-chart"

func CreateAlluxioClusterIfNotExist(ctx ReconcileRequestContext) error {
	// if the release has already been deployed, requeue without further actions
	helmCtx := HelmContext{
		Namespace:   ctx.Namespace,
		ReleaseName: ctx.Name,
	}
	exists, err := checkIfHelmReleaseExists(helmCtx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	logger.Infof("Creating Alluxio cluster %v in namespace %v.", ctx.Name, ctx.Namespace)
	// Construct config.yaml file
	data, err := yaml.Marshal(ctx.AlluxioCluster.Spec)
	if err != nil {
		return err
	}
	configYamlFile, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("%s-%s-config.yaml", ctx.Namespace, ctx.Name))
	configYamlFilePath := configYamlFile.Name()
	if err != nil {
		logger.Errorf("failed to create empty config file: %v", err)
		return err
	}
	err = os.WriteFile(configYamlFilePath, data, 0400)
	if err != nil {
		logger.Errorf("failed saving config file: %v", err)
		return err
	}
	// helm install release with the constructed config.yaml
	helmCtx = HelmContext{
		HelmChartPath:      chartPath,
		ConfigFilePath:     configYamlFilePath,
		Namespace:          ctx.Namespace,
		OverrideProperties: nil,
		ReleaseName:        ctx.Name,
	}
	if err := HelmInstall(helmCtx); err != nil {
		logger.Errorf("error installing helm release. Uninstalling...")
		if err := HelmDeleteIfExist(helmCtx); err != nil {
			logger.Errorf("failed to delete failed helm release %v in namespace %v: %v", ctx.Name, ctx.Namespace, err)
			return err
		}
	}
	return nil
}

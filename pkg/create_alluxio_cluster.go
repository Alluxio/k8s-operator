package pkg

import (
	"fmt"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/yaml"
)

const chartPath = "/opt/alluxio-helm-chart"

func CreateAlluxioCluster(ctx ReconcileRequestContext) ReconcileResponse {
	logger := ctx.Logger
	logger.Info("Creating Alluxio cluster.", "cluster name", ctx.Name, "namespace", ctx.Namespace)

	// Construct config.yaml file
	data, err := yaml.Marshal(ctx.AlluxioCluster)
	if err != nil {
		return ReconcileResponse{
			Err:        err,
			NeedReturn: true,
			Result:     ctrl.Result{},
		}
	}
	configYamlFile, err := os.CreateTemp(os.TempDir(), fmt.Sprintf("%s-%s-config.yaml", ctx.Namespace, ctx.Name))
	configYamlFilePath := configYamlFile.Name()
	if err != nil {
		logger.Error(err, "failed to create empty config file", "config file", configYamlFile)
		return ReconcileResponse{
			Err:        err,
			NeedReturn: true,
			Result:     ctrl.Result{},
		}
	}
	err = os.WriteFile(configYamlFilePath, data, 0400)
	if err != nil {
		logger.Error(err, "failed saving config file.", "config file", configYamlFilePath)
		return ReconcileResponse{
			Err:        err,
			NeedReturn: true,
			Result:     ctrl.Result{},
		}
	}

	// helm install release with the constructed config.yaml
	helmCtx := HelmContext{
		HelmChartPath:      chartPath,
		ConfigFilePath:     configYamlFilePath,
		Namespace:          ctx.Namespace,
		OverrideProperties: nil,
		ReleaseName:        ctx.Name,
	}
	if err := HelmInstall(helmCtx, logger); err != nil {
		logger.Error(err, "error installing helm release. Uninstalling...")
		if err := HelmDeleteIfExist(helmCtx, logger); err != nil {
			logger.Error(err, "failed to delete failed helm release %v in namespace %v", ctx.Name, ctx.Namespace)
			return ReconcileResponse{
				Err:        err,
				NeedReturn: true,
				Result:     ctrl.Result{},
			}
		}
	}
	return ReconcileResponse{
		Err:        nil,
		NeedReturn: true,
		Result:     ctrl.Result{Requeue: true},
	}
}

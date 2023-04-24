package pkg

import (
	"github.com/Alluxio/k8s-operator/pkg/logger"
	ctrl "sigs.k8s.io/controller-runtime"
)

func DeleteAlluxioClusterIfExist(ctx ReconcileRequestContext) (ctrl.Result, error) {
	logger.Infof("Uninstalling Alluxio cluster %v in namespace %v.", ctx.Name, ctx.Namespace)

	helmCtx := HelmContext{
		ReleaseName: ctx.Name,
		Namespace:   ctx.Namespace,
	}
	if err := HelmDeleteIfExist(helmCtx); err != nil {
		logger.Errorf("failed to delete helm release %v in namespace %v: %v", ctx.Name, ctx.Namespace, err)
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

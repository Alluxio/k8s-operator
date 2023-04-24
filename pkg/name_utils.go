package pkg

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/types"
	k8sStringsUtil "k8s.io/utils/strings"
)

func (ctx ReconcileRequestContext) getMasterStatefulSetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: ctx.Namespace,
		Name:      fmt.Sprintf("%s-master", getClusterFullName(ctx.Spec.NameOverride, ctx.Name)),
	}
}

func (ctx ReconcileRequestContext) getWorkerDeploymentNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: ctx.Namespace,
		Name:      fmt.Sprintf("%s-worker", getClusterFullName(ctx.Spec.NameOverride, ctx.Name)),
	}
}

func (ctx ReconcileRequestContext) getFuseDaemonSetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: ctx.Namespace,
		Name:      fmt.Sprintf("%s-fuse", getClusterFullName(ctx.Spec.NameOverride, ctx.Name)),
	}
}

func (ctx ReconcileRequestContext) getProxyDaemonSetNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: ctx.Namespace,
		Name:      fmt.Sprintf("%s-proxy", getClusterFullName(ctx.Spec.NameOverride, ctx.Name)),
	}
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

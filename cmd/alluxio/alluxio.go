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

package alluxio

import (
	"github.com/Alluxio/k8s-operator/pkg"
	"os"

	alluxiov1alpha1 "github.com/Alluxio/k8s-operator/api/v1alpha1"
	"github.com/Alluxio/k8s-operator/pkg/logger"

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme = runtime.NewScheme()
)

func NewAlluxioManagerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "alluxio-manager start",
		Short: "Start the manager of alluxio",
		Run: func(cmd *cobra.Command, args []string) {
			startAlluxioManager()
		},
	}
	return cmd
}

func startAlluxioManager() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(alluxiov1alpha1.AddToScheme(scheme))

	manager, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Port:   9443,
	})
	if err != nil {
		logger.Fatalf("Unable to create Alluxio manager: %v", err)
		os.Exit(1)
	}

	if err = (&pkg.AlluxioClusterReconciler{
		Client: manager.GetClient(),
		Scheme: manager.GetScheme(),
	}).SetupWithManager(manager); err != nil {
		logger.Fatalf("unable to create Alluxio Reconciler: %v", err)
		os.Exit(1)
	}

	logger.Infof("starting manager")
	if err = manager.Start(ctrl.SetupSignalHandler()); err != nil {
		logger.Fatalf("Error starting Alluxio manager: %v", err)
		os.Exit(1)
	}
}

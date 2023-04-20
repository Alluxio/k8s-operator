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

	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup alluxio manager")
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

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&zap.Options{Development: true})))

	manager, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Port:   9443,
	})
	if err != nil {
		setupLog.Error(err, "Unable to create Alluxio manager.")
		os.Exit(1)
	}

	if err = (&pkg.AlluxioClusterReconciler{
		Client: manager.GetClient(),
		Scheme: manager.GetScheme(),
	}).SetupWithManager(manager); err != nil {
		setupLog.Error(err, "unable to create Alluxio Reconciler")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := manager.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "Error starting Alluxio manager.")
		os.Exit(1)
	}
}

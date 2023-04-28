/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dataset

import (
	alluxiov1alpha1 "github.com/alluxio/k8s-operator/api/v1alpha1"
	"github.com/alluxio/k8s-operator/pkg/dataset"
	"github.com/alluxio/k8s-operator/pkg/load"
	"github.com/alluxio/k8s-operator/pkg/logger"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	scheme = runtime.NewScheme()
)

func NewDatasetManagerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dataset-manager start",
		Short: "Start the manager of dataset",
		Run: func(cmd *cobra.Command, args []string) {
			startDatasetManager()
		},
	}
	return cmd
}

func startDatasetManager() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(alluxiov1alpha1.AddToScheme(scheme))

	manager, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: scheme,
		Port:   9443,
	})
	if err != nil {
		logger.Fatalf("Unable to create Dataset manager: %v", err)
		os.Exit(1)
	}

	if err = (&dataset.DatasetReconciler{
		Client: manager.GetClient(),
		Scheme: manager.GetScheme(),
	}).SetupWithManager(manager); err != nil {
		logger.Fatalf("unable to create Dataset controller: %v", err)
		os.Exit(1)
	}

	if err = (&load.LoadReconciler{
		Client: manager.GetClient(),
		Scheme: manager.GetScheme(),
	}).SetupWithManager(manager); err != nil {
		logger.Fatalf("unable to create Load controller: %v", err)
		os.Exit(1)
	}

	logger.Infof("starting manager")
	if err = manager.Start(ctrl.SetupSignalHandler()); err != nil {
		logger.Fatalf("Error starting Dataset manager: %v", err)
		os.Exit(1)
	}
}

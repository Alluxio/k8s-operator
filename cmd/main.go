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

package main

import (
	"github.com/Alluxio/k8s-operator/cmd/alluxio"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	setupLog = ctrl.Log.WithName("setup main")
)

func main() {
	command := alluxio.NewAlluxioManagerCommand()
	if err := command.Execute(); err != nil {
		setupLog.Error(err, "Failed to launch Alluxio controller.")
		os.Exit(1)
	}
	setupLog.Info("succeeded. exiting...")
	os.Exit(0)
}

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
	"os"

	"github.com/alluxio/k8s-operator/cmd/alluxio/alluxio"
	"github.com/alluxio/k8s-operator/pkg/logger"
)

func main() {
	command := alluxio.NewAlluxioManagerCommand()
	if err := command.Execute(); err != nil {
		logger.Fatalf("Failed to launch Alluxio controller: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

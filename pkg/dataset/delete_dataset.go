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

package dataset

import (
	"github.com/alluxio/k8s-operator/pkg/logger"
	ctrl "sigs.k8s.io/controller-runtime"
)

// TODO: shall we free worker space when called?
func DeleteDatasetIfExist(req ctrl.Request) (ctrl.Result, error) {
	logger.Infof("Uninstalling Dataset %v in namespace %v.", req.Name, req.Namespace)

	return ctrl.Result{}, nil
}

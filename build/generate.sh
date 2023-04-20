#
# The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
# (the "License"). You may not use this work except in compliance with the License, which is
# available at www.apache.org/licenses/LICENSE-2.0
#
# This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
# either express or implied, as more fully set forth in the License.
#
# See the NOTICE file distributed with this work for information regarding copyright ownership.
#

#!/usr/bin/env bash
# Use controller-gen to generate files of CRDs and deepcopy methods for AlluxioCluster.

GOBIN="$(go env GOBIN)"
CONTROLLER_GEN="$(which controller-gen)"
SCRIPT_DIR="$(dirname -- "$0")"

main() {
  if [[ -z "$CONTROLLER_GEN" ]]; then
    install_controller_gen
  fi
  generate_deepcopy_methods_file
  generate_crds
  copy_crds_to_chart
}

set_gobin() {
  GOBIN="$(go env GOPATH)/bin"
}

install_controller_gen() {
  go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.10.0
  if [[ -z $"$GOBIN" ]]; then
    set_gobin
  fi
	CONTROLLER_GEN="${GOBIN}/controller-gen"
}

generate_deepcopy_methods_file() {
  local cmd=("${CONTROLLER_GEN}" object:headerFile="${SCRIPT_DIR}"/license-header/go.txt paths="${SCRIPT_DIR}"/../...)
  "${cmd[@]}"
}

generate_crds() {
  local cmd=("${CONTROLLER_GEN}" crd paths="${SCRIPT_DIR}"/../... output:crd:artifacts:config="${SCRIPT_DIR}"/../resources/crds)
  "${cmd[@]}"
  for file in "${SCRIPT_DIR}"/../resources/crds/*; do
    insert_license_header $file
  done
}

insert_license_header() {
  if [[ $# != 1 ]]; then
    echo "insert_license_header function only takes one file"
    exit 1
  fi
  cp $1 $1.tmp
  cp "${SCRIPT_DIR}"/license-header/yaml.txt $1
  cat $1.tmp >> $1
  rm $1.tmp
}

copy_crds_to_chart() {
  cp -r "${SCRIPT_DIR}"/../resources/crds "${SCRIPT_DIR}"/../deploy/charts/alluxio-operator/
}

main

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

# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /opt/alluxio-k8s-operator

COPY .. .

RUN bash -c ./build/generate.sh

RUN go mod download

WORKDIR /opt/alluxio-k8s-operator/cmd/alluxio
RUN CGO_ENABLED=0 go build -a -o ../../bin/alluxio-manager

WORKDIR /opt/alluxio-k8s-operator/cmd/dataset
RUN CGO_ENABLED=0 go build -a -o ../../bin/dataset-manager

WORKDIR /opt/helm

ARG HELM_VERSION=v3.11.2
RUN curl -o helm-${HELM_VERSION}-linux-amd64.tar.gz https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar -xvf helm-${HELM_VERSION}-linux-amd64.tar.gz

FROM centos:7

RUN groupadd --gid 1001 k8soperator && \
    useradd -m --uid 1001 --gid 1001 k8soperator

WORKDIR /

COPY --chown=1001:1001 ../deploy/charts/alluxio /opt/alluxio-helm-chart
COPY --chown=1001:1001 ../deploy/jobs /opt/alluxio-jobs
COPY --chown=1001:1001 --from=builder /opt/alluxio-k8s-operator/bin/alluxio-manager .
COPY --chown=1001:1001 --from=builder /opt/alluxio-k8s-operator/bin/dataset-manager .
COPY --chown=1001:1001 --from=builder /opt/helm/linux-amd64/helm /usr/local/bin/helm

USER 1001

# Override in k8s yaml
CMD ["sleep", "infinity"]

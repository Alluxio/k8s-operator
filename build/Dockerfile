# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /opt/alluxio-k8s-operator

COPY .. .

RUN bash -c ./build/generate.sh

RUN go mod download

WORKDIR /opt/alluxio-k8s-operator/cmd

RUN go build -a -o ../alluxio-manager

WORKDIR /opt/helm

ARG HELM_VERSION=v3.11.2
RUN curl -o helm-${HELM_VERSION}-linux-amd64.tar.gz https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz && \
    tar -xvf helm-${HELM_VERSION}-linux-amd64.tar.gz

FROM centos:7

RUN groupadd --gid 1001 k8sOperator && \
    useradd -m --uid 1001 --gid 1001 k8sOperator

WORKDIR /

COPY --chown=1001:1001 ../deploy/charts/alluxio /opt/alluxio-helm-chart
COPY --chown=1001:1001 --from=builder /opt/alluxio-k8s-operator/alluxio-manager .
COPY --chown=1001:1001 --from=builder /opt/helm/linux-amd64/helm /usr/local/bin/helm

USER 1001

CMD ["/alluxio-manager", "start"]
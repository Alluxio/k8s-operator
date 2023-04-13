# Build the manager binary
FROM golang:1.19 as builder

WORKDIR /opt/alluxio-k8s-operator

COPY .. .

RUN go mod download

WORKDIR /opt/alluxio-k8s-operator/cmd

RUN go build -a -o ../alluxio-manager

FROM centos:7

RUN groupadd --gid 1001 k8sOperator && \
    useradd -m --uid 1001 --gid 1001 k8sOperator

WORKDIR /

COPY --chown=1001:1001 --from=builder /opt/alluxio-k8s-operator/alluxio-manager .

USER 1001

CMD ["/alluxio-manager", "start"]

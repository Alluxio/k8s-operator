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

FROM golang:1.19-alpine3.17 AS csi-dev

ENV GO111MODULE=on
RUN mkdir -p /alluxio-csi
COPY ../../csi /alluxio-csi
RUN cd /alluxio-csi && \
    CGO_ENABLED=0 go build -o /usr/local/bin/alluxio-csi

FROM alpine:3.17 AS final

RUN addgroup -g 1001 alluxio && \
    adduser -u 1001 -G alluxio alluxio -D

COPY --chown=1001:1001 --from=csi-dev /usr/local/bin/alluxio-csi /usr/local/bin/

USER 1001

ENTRYPOINT ["/usr/local/bin/alluxio-csi"]

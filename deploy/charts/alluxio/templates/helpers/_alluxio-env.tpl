{{/* The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
(the "License"). You may not use this work except in compliance with the License, which is
available at www.apache.org/licenses/LICENSE-2.0

This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied, as more fully set forth in the License.

See the NOTICE file distributed with this work for information regarding copyright ownership. */}}

{{/* vim: set filetype=mustache: */}}

{{- /* ===================================== */}}
{{- /*       ALLUXIO_MASTER_JAVA_OPTS        */}}
{{- /* ===================================== */}}
{{- define "alluxio.master.env" -}}
{{- $masterJavaOpts := list }}
{{- $masterJavaOpts = print "-Dalluxio.master.hostname=${ALLUXIO_MASTER_HOSTNAME}" | append $masterJavaOpts }}
{{- range $key, $val := .Values.master.properties }}
  {{- $masterJavaOpts = printf "-D%v=%v" $key $val | append $masterJavaOpts }}
{{- end }}
{{- if .Values.master.jvmOptions }}
  {{- $masterJavaOpts = concat $masterJavaOpts .Values.master.jvmOptions }}
{{- end }}
{{- range $opt := $masterJavaOpts }}{{ printf "%v " $opt }}{{ end }}
{{- end -}}

{{- /* ===================================== */}}
{{- /*       ALLUXIO_WORKER_JAVA_OPTS        */}}
{{- /* ===================================== */}}
{{- define "alluxio.worker.env" -}}
{{- $workerJavaOpts := list }}
{{- $workerJavaOpts = print "-Dalluxio.worker.hostname=${ALLUXIO_WORKER_HOSTNAME}" | append $workerJavaOpts }}
{{- range $key, $val := .Values.worker.properties }}
  {{- $workerJavaOpts = printf "-D%v=%v" $key $val | append $workerJavaOpts }}
{{- end }}
{{- if .Values.worker.jvmOptions }}
  {{- $workerJavaOpts = concat $workerJavaOpts .Values.worker.jvmOptions }}
{{- end }}
{{- range $opt := $workerJavaOpts }}{{ printf "%v " $opt }}{{ end }}
{{- end -}}

{{- /* ===================================== */}}
{{- /*       ALLUXIO_PROXY_JAVA_OPTS        */}}
{{- /* ===================================== */}}
{{- define "alluxio.proxy.env" -}}
{{- $proxyJavaOpts := list }}
{{- if .Values.proxy.enabled }}
  {{- $proxyJavaOpts = print "-Dalluxio.user.hostname=${ALLUXIO_CLIENT_HOSTNAME}" | append $proxyJavaOpts }}
  {{- range $key, $val := .Values.proxy.properties }}
    {{- $proxyJavaOpts = printf "-D%v=%v" $key $val | append $proxyJavaOpts }}
  {{- end }}
  {{- if .Values.proxy.jvmOptions }}
    {{- $proxyJavaOpts = concat $proxyJavaOpts .Values.proxy.jvmOptions }}
  {{- end }}
{{- end }}
{{- range $opt := $proxyJavaOpts }}{{ printf "%v " $opt }}{{ end }}
{{- end -}}

{{- /* ===================================== */}}
{{- /*        ALLUXIO_FUSE_JAVA_OPTS         */}}
{{- /* ===================================== */}}
{{- define "alluxio.fuse.env" -}}
{{- $fuseJavaOpts := list }}
{{- if .Values.fuse.enabled }}
  {{- $fuseJavaOpts = print "-Dalluxio.user.hostname=${ALLUXIO_CLIENT_HOSTNAME}" | append $fuseJavaOpts }}
  {{- range $key, $val := .Values.fuse.properties }}
    {{- $fuseJavaOpts = printf "-D%v=%v" $key $val | append $fuseJavaOpts }}
  {{- end }}
  {{- if .Values.fuse.jvmOptions }}
    {{- $fuseJavaOpts = concat $fuseJavaOpts .Values.fuse.jvmOptions }}
  {{- end }}
{{- end }}
{{- range $opt := $fuseJavaOpts }}{{ printf "%v " $opt }}{{ end }}
{{- end -}}

{{- define "alluxio.env" -}}
ALLUXIO_MASTER_JAVA_OPTS+="{{ include "alluxio.master.env" . }}"
ALLUXIO_WORKER_JAVA_OPTS+="{{ include "alluxio.worker.env" . }}"
ALLUXIO_PROXY_JAVA_OPTS+="{{ include "alluxio.proxy.env" . }}"
ALLUXIO_FUSE_JAVA_OPTS+="{{ include "alluxio.fuse.env" . }}"
{{- end -}}

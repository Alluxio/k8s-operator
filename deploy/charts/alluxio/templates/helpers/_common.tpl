{{/* The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
(the "License"). You may not use this work except in compliance with the License, which is
available at www.apache.org/licenses/LICENSE-2.0

This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied, as more fully set forth in the License.

See the NOTICE file distributed with this work for information regarding copyright ownership. */}}

{{/* vim: set filetype=mustache: */}}

{{/*
Expand the name of the chart.
*/}}
{{- define "alluxio.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 32 chars because some Kubernetes name fields are limited to 63 chars (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "alluxio.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 32 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 32 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 32 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "alluxio.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 32 | trimSuffix "-" -}}
{{- end -}}

{{- define "alluxio.mount.basePath" -}}
{{- printf "/mnt/alluxio%v" . }}
{{- end -}}

{{- define "alluxio.imagePullSecrets" -}}
imagePullSecrets:
{{- range $name := .Values.imagePullSecrets }}
  - name: {{ $name }}
{{- end -}}
{{- end -}}

{{- define "alluxio.hostAliases" -}}
hostAliases:
{{- range .Values.hostAliases }}
  - ip: {{ .ip }}
    hostnames:
    {{- range .hostnames }}
      - {{ . }}
    {{- end }}
{{- end }}
{{- end -}}

{{- define "alluxio.resources" -}}
resources:
  {{- if .limits }}
  limits:
    {{- if .limits.cpu  }}
    cpu: {{ .limits.cpu }}
    {{- end }}
    {{- if .limits.memory  }}
    memory: {{ .limits.memory }}
    {{- end }}
  {{- end }}
  {{- if .requests }}
  requests:
    {{- if .requests.cpu  }}
    cpu: {{ .requests.cpu }}
    {{- end }}
    {{- if .requests.memory  }}
    memory: {{ .requests.memory }}
    {{- end }}
  {{- end }}
{{- end -}}

{{- define "alluxio.volumeMounts" -}}
{{- $readOnly := .readOnly }}
{{- range $key, $val := .volumeMounts }}
- name: {{ $key }}-volume
  mountPath: {{ $val }}
  readOnly: {{ $readOnly }}
{{- end }}
{{- end -}}

{{- define "alluxio.secretVolumes" -}}
{{- range $key, $val := . }}
- name: {{ $key }}-volume
  secret:
    secretName: {{ $key }}
    defaultMode: 256
{{- end }}
{{- end -}}

{{- define "alluxio.configMapVolumes" -}}
{{- range $key, $val := . }}
- name: {{ $key }}-volume
  configMap:
    name: {{ $key }}
{{- end }}
{{- end -}}

{{- define "alluxio.persistentVolumeClaims" -}}
{{- range $key, $val := . }}
- name: {{ $key }}-volume
  persistentVolumeClaim:
    claimName: {{ $key }}
{{- end }}
{{- end -}}

{{- define "alluxio.getVolumeName" -}}
{{ printf "%v-%v-volume" .prefix .component }}
{{- end -}}

{{- define "alluxio.getPvcName" -}}
{{ printf "%v-%v-pvc" .prefix .component }}
{{- end -}}

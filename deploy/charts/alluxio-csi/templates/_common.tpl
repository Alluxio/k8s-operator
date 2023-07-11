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
{{- define "alluxio-csi.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 32 chars because some Kubernetes name fields are limited to 63 chars (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "alluxio-csi.fullname" -}}
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
{{- define "alluxio-csi.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 32 | trimSuffix "-" -}}
{{- end -}}

{{- define "alluxio-csi.serviceAccountName" -}}
alluxio-csi-sa
{{- end -}}

{{- define "alluxio-csi.imagePullSecrets" -}}
imagePullSecrets:
{{- range $name := .Values.imagePullSecrets }}
  - name: {{ $name }}
{{- end -}}
{{- end -}}

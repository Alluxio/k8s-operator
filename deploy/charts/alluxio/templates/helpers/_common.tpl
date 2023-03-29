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

{{- define "alluxio.mount.path" -}}
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

{{- define "alluxio.persistentVolumes" -}}
{{- range $key, $val := . }}
- name: {{ $key }}-volume
  persistentVolumeClaim:
    claimName: {{ $key }}
{{- end }}
{{- end -}}

{{- define "alluxio.worker.tieredstoreVolumeMounts" -}}
  {{- if .Values.tieredstore.levels }}
    {{- range .Values.tieredstore.levels }}
      {{- /* The mediumtype can have multiple parts like MEM,SSD */}}
      {{- if .mediumtype }}
        {{- /* Mount each part */}}
        {{- if contains "," .mediumtype }}
          {{- $type := .type }}
          {{- $path := .path }}
          {{- $parts := splitList "," .mediumtype }}
          {{- range $i, $val := $parts }}
            {{- /* Example: For path="/tmp/mem,/tmp/ssd", mountPath resolves to /tmp/mem and /tmp/ssd */}}
            - mountPath: {{ index ($path | splitList ",") $i }}
              name: {{ $val | lower }}-{{ $i }}
          {{- end}}
        {{- /* The mediumtype is a single value. */}}
        {{- else}}
            - mountPath: {{ .path }}
              name: {{ .mediumtype | replace "," "-" | lower }}
        {{- end}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end -}}

{{- define "alluxio.worker.tieredstoreVolumes" -}}
  {{- if .Values.tieredstore.levels }}
    {{- range .Values.tieredstore.levels }}
      {{- if .mediumtype }}
        {{- /* The mediumtype can have multiple parts like MEM,SSD */}}
        {{- if contains "," .mediumtype }}
          {{- $parts := splitList "," .mediumtype }}
          {{- $type := .type }}
          {{- $path := .path }}
          {{- $volumeName := .name }}
          {{- /* A volume will be generated for each part */}}
          {{- range $i, $val := $parts }}
            {{- /* Example: For mediumtype="MEM,SSD", mediumName resolves to mem-0 and ssd-1 */}}
            {{- $mediumName := printf "%v-%v" (lower $val) $i }}
            {{- if eq $type "hostPath"}}
        - hostPath:
            path: {{ index ($path | splitList ",") $i }}
            type: DirectoryOrCreate
          name: {{ $mediumName }}
            {{- else if eq $type "persistentVolumeClaim" }}
        - name: {{ $mediumName }}
          persistentVolumeClaim:
            {{- /* Example: For volumeName="/tmp/mem,/tmp/ssd", claimName resolves to /tmp/mem and /tmp/ssd */}}
            claimName: {{ index ($volumeName | splitList ",") $i }}
            {{- else }}
        - name: {{ $mediumName }}
          emptyDir:
            medium: "Memory"
              {{- if .quota }}
            sizeLimit: {{ .quota }}
              {{- end}}
            {{- end}}
          {{- end}}
        {{- /* The mediumtype is a single value like MEM. */}}
        {{- else}}
          {{- $mediumName := .mediumtype | lower }}
          {{- if eq .type "hostPath"}}
        - hostPath:
            path: {{ .path }}
            type: DirectoryOrCreate
          name: {{ $mediumName }}
          {{- else if eq .type "persistentVolumeClaim" }}
        - name: {{ $mediumName }}
          persistentVolumeClaim:
            claimName: {{ .name }}
          {{- else }}
        - name: {{ $mediumName }}
          emptyDir:
            medium: "Memory"
            {{- if .quota }}
            sizeLimit: {{ .quota }}
            {{- end}}
          {{- end}}
        {{- end}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end -}}

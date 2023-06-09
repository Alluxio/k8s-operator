{{/* The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
(the "License"). You may not use this work except in compliance with the License, which is
available at www.apache.org/licenses/LICENSE-2.0

This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied, as more fully set forth in the License.

See the NOTICE file distributed with this work for information regarding copyright ownership. */}}

{{/* vim: set filetype=mustache: */}}

{{- define "alluxio.csi.fuse.yaml" -}}
{{- $name := include "alluxio.name" . }}
{{- $fullName := include "alluxio.fullname" . }}
{{- $alluxioFuseMountPoint := include "alluxio.mount.basePath" "/fuse" }}
kind: Pod
apiVersion: v1
metadata:
  name: {{ $fullName }}-fuse
  labels:
    name: {{ $fullName }}-fuse
    app: {{ $name }}
    role: alluxio-fuse
spec:
  nodeName:
  securityContext:
    runAsUser: 0 # required for mounting to csi designated path
    runAsGroup: 0
    fsGroup: 0
  {{- if .Values.serviceAccountName }}
  serviceAccountName: {{ .Values.serviceAccountName }}
  {{- end }}
  {{- if .Values.imagePullSecrets }}
{{ include "alluxio.imagePullSecrets" . | indent 2 }}
  {{- end}}
  initContainers:
  - name: wait-master
    image: {{ .Values.image }}:{{ .Values.imageTag }}
    command: ["/bin/sh", "-c"]
    args:
      - until nslookup {{ $fullName }}-master-0;
        do sleep 2;
        done
    volumeMounts:
    - name: {{ $fullName }}-alluxio-conf
      mountPath: /opt/alluxio/conf
  containers:
    - name: alluxio-fuse
      image: {{ .Values.image }}:{{ .Values.imageTag }}
      imagePullPolicy: {{ .Values.imagePullPolicy }}
      {{- if .Values.fuse.resources }}
      resources:
        {{- if .Values.fuse.resources.limits }}
        limits:
          cpu: {{ .Values.fuse.resources.limits.cpu }}
          memory: {{ .Values.fuse.resources.limits.memory }}
        {{- end }}
        {{- if .Values.fuse.resources.requests }}
          cpu: {{ .Values.fuse.resources.requests.cpu }}
          memory: {{ .Values.fuse.resources.requests.memory }}
        {{- end }}
      {{- end }}
      command: [ "/entrypoint.sh" ]
      args:
        - fuse
        - {{ required "The path of the dataset must be set." .Values.dataset.path }}
        - {{ $alluxioFuseMountPoint }}
        {{- range .Values.fuse.mountOptions }}
        - -o {{ . }}
        {{- end }}
      env:
      {{- range $key, $value := .Values.fuse.env }}
      - name: "{{ $key }}"
        value: "{{ $value }}"
      {{- end }}
      securityContext:
        privileged: true # required by bidirectional mount
      volumeMounts:
        - name: {{ $fullName }}-alluxio-conf
          mountPath: /opt/alluxio/conf
        - name: pods-mount-dir
          mountPath: /var/lib/kubelet
          mountPropagation: "Bidirectional"
        {{- if .Values.secrets }}
{{- include "alluxio.volumeMounts" (dict "volumeMounts" .Values.secrets.fuse "readOnly" true) | indent 8 }}
        {{- end }}
        {{- if .Values.configMaps }}
{{- include "alluxio.volumeMounts" (dict "volumeMounts" .Values.configMaps.fuse "readOnly" true) | indent 8 }}
        {{- end }}
        {{- if .Values.pvcMounts }}
{{- include "alluxio.volumeMounts" (dict "volumeMounts" .Values.pvcMounts.fuse "readOnly" false) | indent 8 }}
        {{- end }}
  restartPolicy: Always
  volumes:
    - name: pods-mount-dir
      hostPath:
        path: /var/lib/kubelet
        type: Directory
    - name: {{ $fullName }}-alluxio-conf
      configMap:
        name: {{ $fullName }}-alluxio-conf
    {{- if .Values.secrets }}
{{- include "alluxio.secretVolumes" .Values.secrets.fuse | indent 4 }}
    {{- end }}
{{- if .Values.configMaps }}
    {{- include "alluxio.configMapVolumes" .Values.configMaps.fuse | indent 4 }}
{{- end }}
    {{- if .Values.pvcMounts }}
{{- include "alluxio.persistentVolumeClaims" .Values.pvcMounts.fuse | indent 4 }}
    {{- end }}
{{- end -}}

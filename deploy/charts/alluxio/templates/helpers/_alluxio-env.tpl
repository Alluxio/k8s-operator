{{- /* ===================================== */}}
{{- /*       ALLUXIO_MASTER_JAVA_OPTS        */}}
{{- /* ===================================== */}}
{{- define "alluxio.master.env" -}}
{{- $masterJavaOpts := list }}
{{- $masterJavaOpts = print "-Dalluxio.master.hostname=${ALLUXIO_MASTER_HOSTNAME}" | append $masterJavaOpts }}
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

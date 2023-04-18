{{/* The Alluxio Open Foundation licenses this work under the Apache License, version 2.0
(the "License"). You may not use this work except in compliance with the License, which is
available at www.apache.org/licenses/LICENSE-2.0

This software is distributed on an "AS IS" basis, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied, as more fully set forth in the License.

See the NOTICE file distributed with this work for information regarding copyright ownership. */}}

{{/* vim: set filetype=mustache: */}}

{{- define "alluxio.metrics.properties" -}}
{{- if .Values.metrics.consoleSink.enabled -}}
sink.console.class=alluxio.metrics.sink.consoleSink
sink.console.period={{ .Values.metrics.consoleSink.period }}
sink.console.unit={{ .Values.metrics.consoleSink.unit }}
{{- end -}}
{{- if .Values.metrics.csvSink.enabled -}}
sink.csv.class=alluxio.metrics.sink.csvSink
sink.csv.period={{ .Values.metrics.csvSink.period }}
sink.csv.unit={{ .Values.metrics.csvSink.unit }}
sink.csv.directory={{ .Values.metrics.csvSink.directory }}
{{- end -}}
{{- if .Values.metrics.jmxSink.enabled -}}
sink.jmx.class=alluxio.metrics.sink.jmxSink
sink.jmx.domain={{ .Values.metrics.jmxSink.domain }}
{{- end -}}
{{- if .Values.metrics.graphiteSink.enabled -}}
sink.graphite.class=alluxio.metrics.sink.graphiteSink
sink.graphite.host={{ .Values.metrics.graphiteSink.host }}
sink.graphite.port={{ .Values.metrics.graphiteSink.port }}
sink.graphite.period={{ .Values.metrics.graphiteSink.period }}
sink.graphite.unit={{ .Values.metrics.graphiteSink.unit }}
sink.graphite.prefix={{ .Values.metrics.graphiteSink.prefix }}
{{- end -}}
{{- if .Values.metrics.slf4jSink.enabled -}}
sink.slf4j.class=alluxio.metrics.sink.slf4jSink
sink.slf4j.period={{ .Values.metrics.slf4jSink.period }}
sink.slf4j.unit={{ .Values.metrics.slf4jSink.unit }}
sink.slf4j.filter-class={{ .Values.metrics.slf4jSink.filterClass }}
sink.slf4j.filter-regex={{ .Values.metrics.slf4jSink.filterRegex }}
{{- end -}}
{{- if .Values.metrics.prometheusMetricsServlet.enabled -}}
sink.prometheus.class=alluxio.metrics.sink.prometheusMetricsServlet
{{- end -}}
{{- end -}}
{{- define "alluxio.site.properties" -}}
# Common properties
{{- range $key, $val := .Values.properties }}
{{- printf "%v=%v" $key $val }}
{{- end }}

# Master address if single master
{{- if eq (int .Values.master.count) 1 }}
alluxio.master.hostname={{ include "alluxio.fullname" . }}-master-0
{{- end }}

# Journal properties
{{ printf "alluxio.master.journal.type=EMBEDDED" }}
{{ printf "alluxio.master.journal.folder=%v" (include "alluxio.mount.path" "/journal") }}
{{- if gt (int .Values.master.count) 1 }}
{{- $embeddedJournalAddresses := ""}}
{{- range $i := until (int .Values.master.count) }}
  {{ $embeddedJournalAddresses = printf "%v,%v-master-%v:19200" $embeddedJournalAddresses (include "alluxio.fullname" .) $i }}
{{- end }}
{{ printf "alluxio.master.embedded.journal.addresses=%v" $embeddedJournalAddresses }}
{{- end }}

# Tiered Storage
{{- if .Values.tieredstore }}
{{ printf "alluxio.worker.tieredstore.levels=%v" (len .Values.tieredstore.levels) }}
{{- range .Values.tieredstore.levels }}
{{- $tierName := printf "alluxio.worker.tieredstore.level%v" .level }}
{{- if .alias }}
{{ printf "%v.alias=%v" $tierName .alias -}}
{{- end }}
{{ printf "%v.dirs.mediumtype=%v" $tierName .mediumtype }}
{{- if .path }}
{{ printf "%v.dirs.path=%v" $tierName .path }}
{{- end}}
{{- if .quota }}
{{ printf "%v.dirs.quota=%v" $tierName .quota }}
{{- end }}
{{- end }}
{{- end }}
{{- end -}}

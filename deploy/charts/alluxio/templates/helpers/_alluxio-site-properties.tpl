{{- define "alluxio.site.properties" -}}
# Enable Dora
alluxio.dora.client.read.location.policy.enabled=true
alluxio.user.short.circuit.enabled=false
alluxio.master.worker.register.lease.enabled=false

# Common properties
{{- range $key, $val := .Values.properties }}
{{ printf "%v=%v" $key $val }}
{{- end }}

{{- if eq (int .Values.master.count) 1 }}
# Master address for single master
alluxio.master.hostname={{ include "alluxio.fullname" . }}-master-0
{{- end }}

# Journal properties
{{ printf "alluxio.master.journal.type=EMBEDDED" }}
{{ printf "alluxio.master.journal.folder=%v" (include "alluxio.mount.path" "/journal") }}
{{- if gt (int .Values.master.count) 1 }}
{{- $embeddedJournalAddresses := ""}}
{{- range $i := until (int .Values.master.count) }}
  {{- $embeddedJournalAddresses = printf "%v,%v-master-%v:19200" $embeddedJournalAddresses (include "alluxio.fullname" $) $i }}
{{- end }}
{{ printf "alluxio.master.embedded.journal.addresses=%v" $embeddedJournalAddresses }}
{{- end }}

# Page Storage
alluxio.worker.block.store.type=PAGE
alluxio.worker.page.store.type=LOCAL
alluxio.worker.page.store.dirs=/mnt/alluxio/pagestore
{{ printf "alluxio.worker.page.store.sizes=%v" .Values.pagestore.quota }}

{{- end -}}

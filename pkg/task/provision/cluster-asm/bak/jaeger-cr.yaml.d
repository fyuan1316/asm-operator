apiVersion: "jaegertracing.io/v1"
kind: Jaeger
metadata:
  name: jaeger-prod
  namespace: {{ .Values.istio.namespace }}
  labels:
    heritage: {{ $.Release.Service }}
    release: {{ $.Release.Name }}
spec:
  strategy: production
  collector:
    image: "{{ template "collector.image" . }}"
    replicas: {{ .Values.jaeger.collector.replicas }}
  ingress:
    enabled: {{ .Values.jaeger.ingress.enabled }}
  agent:
    image: "{{ template "agent.image" . }}"
  query:
    image: "{{ template "query.image" . }}"
    replicas: {{ .Values.jaeger.query.replicas }}
    options:
      query:
        base-path: {{ .Values.jaeger.query.basepath }}
  storage:
    type: elasticsearch
    options:
      es.index-prefix: "{{ .Values.jaeger.elasticsearch.indexprefix }}"
      es.server-urls: "{{ .Values.jaeger.elasticsearch.serverurl }}"
{{ if .Values.jaeger.elasticsearch.username }}
      es.username: {{ .Values.jaeger.elasticsearch.username }}
{{ end }}
{{ if .Values.jaeger.elasticsearch.password }}
      es.password: {{ .Values.jaeger.elasticsearch.password }}
{{ end }}
    esIndexCleaner:
      image: "{{ template "es_index_cleaner.image" .}}"
      numberOfDays: {{ .Values.jaeger.elasticsearch.indexRetainsDays }}
    dependencies:
      enabled: false
  labels:
    asm.{{ .Values.global.labelBaseDomain }}/healthcheck-component: "true"
    service_name: jaeger
    {{ .Values.global.labelBaseDomain }}/product: "Service-Mesh"
  resources:
{{- if .Values.jaeger.resources }}
{{ toYaml .Values.jaeger.resources | indent 4 }}
{{- end }}

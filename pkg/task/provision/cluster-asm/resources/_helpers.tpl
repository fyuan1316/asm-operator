{{/* vim: set filetype=mustache: */}}
{{/*
Expand the name of the chart.
*/}}
{{- define "istio.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "istio.fullname" -}}
{{- $name := default .Chart.Name .Values.nameOverride -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a fully qualified configmap name.
*/}}
{{- define "istio.configmap.fullname" -}}
{{- printf "%s-%s" .Release.Name "istio-mesh-config" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Configmap checksum.
*/}}
{{- define "istio.configmap.checksum" -}}
{{- print $.Template.BasePath "/configmap.yaml" | sha256sum -}}
{{- end -}}

{{- define "ingressHost" -}}
{{- $scheme :=  .Values.global.scheme -}}
{{- printf "%s://%s" $scheme .Values.global.host | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{- define "query.image" -}}
{{- $registryAddress :=  .Values.global.registry.address -}}
{{- $repositoryName := .Values.global.images.jaeger_query.repository -}}
{{- $tag := .Values.global.images.jaeger_query.tag -}}
{{- printf "%s/%s:%s" .Values.global.registry.address $repositoryName $tag -}}
{{- end -}}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "collector.image" -}}
{{- $registryAddress :=  .Values.global.registry.address -}}
{{- $repositoryName := .Values.global.images.jaeger_collector.repository -}}
{{- $tag := .Values.global.images.jaeger_collector.tag -}}
{{- printf "%s/%s:%s" .Values.global.registry.address $repositoryName $tag -}}
{{- end -}}

{{- define "es_index_cleaner.image" -}}
{{- $registryAddress :=  .Values.global.registry.address -}}
{{- $repositoryName := .Values.global.images.jaeger_es_index_cleaner.repository -}}
{{- $tag := .Values.global.images.jaeger_es_index_cleaner.tag -}}
{{- printf "%s/%s:%s" .Values.global.registry.address $repositoryName $tag -}}
{{- end -}}

{{- define "agent.image" -}}
{{- $registryAddress :=  .Values.global.registry.address -}}
{{- $repositoryName := .Values.global.images.jaeger_agent.repository -}}
{{- $tag := .Values.global.images.jaeger_agent.tag -}}
{{- printf "%s/%s:%s" .Values.global.registry.address $repositoryName $tag -}}
{{- end -}}



{{/* affinity - https://kubernetes.io/docs/concepts/configuration/assign-pod-node/ */}}


{{- define "podAntiAffinity" }}
{{- $arg1 := index . 0 }}
{{- $arg2 := index . 1 }}
  podAntiAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 100
        podAffinityTerm:
          labelSelector:
            matchLabels:
              {{ $arg1 }}: {{ $arg2 }}
          topologyKey: kubernetes.io/hostname
{{- end }}

apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: asm-controller-webhook
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: asm-controller-webhook
      namespace: {{ $.Release.Namespace }}
      path: /webhook/validate/virtualservices
  admissionReviewVersions:
  - v1beta1
  failurePolicy: Fail
  name: vvirtualservices.kb.io
  objectSelector:
    matchLabels:
      asm.{{ .Values.global.labelBaseDomain }}/vstype: root
  rules:
  - apiGroups:
    - networking.istio.io
    apiVersions:
    - v1alpha3
    operations:
    - CREATE
    - UPDATE
    resources:
    - virtualservices
  sideEffects: None
---
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  creationTimestamp: null
  name: serviceentry-v1alpha3-validation
webhooks:
- clientConfig:
    caBundle: Cg==
    service:
      name: asm-controller-webhook
      namespace: {{ $.Release.Namespace }}
      path: /webhook/validate/serviceentry
  admissionReviewVersions:
  - v1beta1
  failurePolicy: Fail
  name: vserviceentries.kb.io
  rules:
  - apiGroups:
    - networking.istio.io
    apiVersions:
    - v1alpha3
    operations:
    - CREATE
    resources:
    - serviceentries
  sideEffects: None
---
apiVersion: v1
kind: Service
metadata:
  name: asm-controller-webhook
  namespace: {{ $.Release.Namespace }}
spec:
  ports:
    - port: 443
      targetPort: 9443
  selector:
    app: asm-controller

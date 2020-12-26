k apply -f - <<EOF
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    control-plane: controller-manager
  name: controller-manager
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      control-plane: controller-manager
  template:
    metadata:
      labels:
        control-plane: controller-manager
    spec:
      containers:
      - args:
        - --enable-leader-election
        command:
        - /manager
        image: harbor-b.alauda.cn/aml/asm-operator:v0.1.0
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 30Mi
          requests:
            cpu: 100m
            memory: 20Mi
      serviceAccount: asm-operator
      terminationGracePeriodSeconds: 10
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: asm-operator
rules:
- apiGroups:
  - "operator.alauda.io"
  resources:
  - configmaps
  - asms
  verbs:
  - '*'
- apiGroups:
  - "apiextensions.k8s.io"
  resources:
  - customresourcedefinitions
  verbs:
  - '*'
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: asm-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: admin
subjects:
- kind: ServiceAccount
  name: asm-operator
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: asm-operator1
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: asm-operator
subjects:
- kind: ServiceAccount
  name: asm-operator
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: asm-operator
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: asm-operator
subjects:
- kind: ServiceAccount
  name: asm-operator
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: asm-operator
  namespace: default
rules:
- apiGroups:
  - "operator.alauda.io"
  resources:
  - configmaps
  - asms
  verbs:
  - '*'
- apiGroups:
  - ""
  resources:
  - events
  verbs:
  - '*'

EOF

k create sa asm-operator



k set image deploy/controller-manager manager=harbor-b.alauda.cn/aml/asm-operator:v0.2.0
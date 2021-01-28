module gitlab-ce.alauda.cn/micro-service/asm-operator

go 1.13

replace (
	//github.com/fyuan1316/klient v0.0.0-20210106101119-962edf8f6e33 => ./local/klient // indirect
	//github.com/fyuan1316/operatorlib v0.0.0-20210127105047-5c3456252ba9 => ./local/operatorlib
)

require (
	github.com/fyuan1316/klient v0.0.0-20210106235436-d9ae9fa37eec // indirect
	github.com/fyuan1316/operatorlib v0.0.0-20210127105047-5c3456252ba9
	github.com/go-logr/logr v0.2.1
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.2
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.7.0
	go.uber.org/zap v1.16.0
	k8s.io/api v0.19.4
	k8s.io/apiextensions-apiserver v0.19.4
	k8s.io/apimachinery v0.19.4
	k8s.io/client-go v0.19.4
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

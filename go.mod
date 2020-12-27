module github.com/fyuan1316/asm-operator

go 1.13

//replace bitbucket.org/mathildetech/asm-controller v1.8.5-0.20200928070659-0203e205a584 => gomod.alauda.cn/asm-controller v1.8.5-0.20200928070659-0203e205a584

require (
	//bitbucket.org/mathildetech/asm-controller v1.8.5-0.20200928070659-0203e205a584
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/coreos/prometheus-operator v0.41.0
	github.com/go-logr/logr v0.1.0
	github.com/gogo/protobuf v1.3.1
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/mattbaird/jsonpatch v0.0.0-20200820163806-098863c1fc24
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	k8s.io/api v0.18.6
	k8s.io/apiextensions-apiserver v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v0.18.6
	sigs.k8s.io/controller-runtime v0.6.3
	sigs.k8s.io/yaml v1.2.0
)

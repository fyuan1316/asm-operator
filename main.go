/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/crd"

	//"gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/crd"
	//promv1 "gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/monitoring/v1"
	//depv1alphba1 "gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/v1alpha1"
	//depv1beta1 "gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/v1beta1"
	//depv1beta2 "gitlab-ce.alauda.cn/micro-service/asm-operator/api/dep/v1beta2"
	"os"

	"go.uber.org/zap/zapcore"

	"gitlab-ce.alauda.cn/micro-service/asm-operator/pkg/task/entry"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	operatorv1alpha1 "gitlab-ce.alauda.cn/micro-service/asm-operator/api/v1alpha1"
	"gitlab-ce.alauda.cn/micro-service/asm-operator/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(crd.AddToScheme(scheme))
	//utilruntime.Must(depv1beta1.AddToScheme(scheme))
	//utilruntime.Must(depv1beta2.AddToScheme(scheme))
	//utilruntime.Must(depv1alphba1.AddToScheme(scheme))
	//utilruntime.Must(promv1.AddToScheme(scheme))

	utilruntime.Must(operatorv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true), zap.Level(zapcore.DebugLevel)))
	//watchNamespace, err := k8sutil.GetWatchNamespace()
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "efff73cd.alauda.io",
		Namespace:          "",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}
	//dynamicClient, err := dynamic.NewForConfig(mgr.GetConfig())
	//if err != nil {
	//	panic(err)
	//}

	if err = (&controllers.AsmReconciler{
		Client: mgr.GetClient(),
		//DynamicClient: dynamicClient,
		//Config:        mgr.GetConfig(),
		Log:      ctrl.Log.WithName("asm-operator").WithName("Asm"),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor("asm-operator"),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Asm")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	//setup tasks
	if err := entry.SetUp(); err != nil {
		panic(err)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

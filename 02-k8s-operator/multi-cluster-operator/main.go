/*
Copyright 2022.

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
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	jobv1 "github.com/mohuishou/blog-code/02-k8s-operator/multi-cluster-operator/api/v1"
	"github.com/mohuishou/blog-code/02-k8s-operator/multi-cluster-operator/controllers"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(jobv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var probeAddr string
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8090", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8091", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	conf := NewMainConfig("kind-main")

	mgr, err := ctrl.NewManager(conf, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "6837ed1d.lailin.xyz",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	clusters := NewSubClusters(mgr, "kind-sub")
	_, err = controllers.NewTestReconciler(mgr, clusters)
	checkErr(err, "NewTestReconciler")

	_, err = controllers.NewTestJobReconciler(mgr, clusters)
	checkErr(err, "NewTestJobReconciler")

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// NewMainConfig 初始化主集群的 config
func NewMainConfig(context string) *rest.Config {
	conf, err := config.GetConfigWithContext(context)
	checkErr(err, "get client config fail", "context", context)
	return conf
}

// NewSubClusters 初始化子集群
// 在 ~/.kube/config 文件中需要有这两个 context 集群
func NewSubClusters(mgr ctrl.Manager, clientContexts ...string) map[string]cluster.Cluster {
	clusters := map[string]cluster.Cluster{}

	for _, v := range clientContexts {
		conf, err := config.GetConfigWithContext(v)
		checkErr(err, "get client config fail", "context", v)

		c, err := cluster.New(conf)
		checkErr(err, "new cluster fail", "context", v)

		err = mgr.Add(c)
		checkErr(err, "add cluster in manager", "context", v)

		clusters[v] = c
	}
	return clusters
}

func checkErr(err error, msg string, keysAndValues ...interface{}) {
	if err == nil {
		return
	}
	fatal(err, msg, keysAndValues...)
}

func fatal(err error, msg string, keysAndValues ...interface{}) {
	setupLog.Error(err, msg, keysAndValues...)
	os.Exit(1)
}

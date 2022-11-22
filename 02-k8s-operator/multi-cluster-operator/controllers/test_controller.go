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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"

	jobv1 "github.com/mohuishou/blog-code/02-k8s-operator/multi-cluster-operator/api/v1"
)

// TestReconciler reconciles a Test object
type TestReconciler struct {
	// 主集群 client
	client.Client

	// 所有集群的客户端列表
	Clients map[string]client.Client

	Scheme *runtime.Scheme
}

// NewTestReconciler ...
func NewTestReconciler(mgr ctrl.Manager, clusters map[string]cluster.Cluster) (*TestReconciler, error) {
	r := TestReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Clients: map[string]client.Client{
			"main": mgr.GetClient(),
		},
	}
	for name, cluster := range clusters {
		r.Clients[name] = cluster.GetClient()
	}

	err := r.SetupWithManager(mgr)
	return &r, err
}

//+kubebuilder:rbac:groups=job.lailin.xyz,resources=tests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=job.lailin.xyz,resources=tests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=job.lailin.xyz,resources=tests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *TestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var test jobv1.Test
	var res ctrl.Result

	err := r.Get(ctx, req.NamespacedName, &test)
	if err != nil {
		return res, client.IgnoreNotFound(err)
	}

	job := test.Job()

	for _, c := range r.Clients {
		j := job.DeepCopy()
		err := c.Get(ctx, client.ObjectKeyFromObject(&job), j)
		if client.IgnoreNotFound(err) != nil {
			return res, err
		}

		// 已存在
		if err == nil {
			continue
		}

		err = c.Create(ctx, job.DeepCopy())
		if err != nil {
			return res, err
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	builder := ctrl.NewControllerManagedBy(mgr).
		For(&jobv1.Test{})
	return builder.Complete(r)
}

/*
Copyright 2021.

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

	"k8s.io/client-go/tools/record"

	"github.com/go-logr/logr"
	nodesv1 "github.com/mohuishou/blog-code/k8s-operator/node-pool-operator/api/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/node/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NodePoolReconciler reconciles a NodePool object
type NodePoolReconciler struct {
	client.Client
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

const nodeFinalizer = "node.finalizers.node-pool.lailin.xyz"

//+kubebuilder:rbac:groups=nodes.lailin.xyz,resources=nodepools,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=nodes.lailin.xyz,resources=nodepools/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=nodes.lailin.xyz,resources=nodepools/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// Modify the Reconcile function to compare the state specified by
// the NodePool object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.7.2/pkg/reconcile
func (r *NodePoolReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = r.Log.WithValues("nodepool", req.NamespacedName)
	// 获取对象
	pool := &nodesv1.NodePool{}
	err := r.Get(ctx, req.NamespacedName, pool)
	if err != nil {
		r.Log.Error(err, "can not get node pool")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var nodes corev1.NodeList

	// 查看是否存在对应的节点，如果存在那么就给这些节点加上数据
	err = r.List(ctx, &nodes, &client.ListOptions{LabelSelector: pool.NodeLabelSelector()})
	if client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	// 进入预删除流程
	if !pool.DeletionTimestamp.IsZero() {
		return ctrl.Result{}, r.nodeFinalizer(ctx, pool, nodes.Items)
	}

	// 如果删除时间戳为空说明现在不需要删除该数据，我们将 nodeFinalizer 加入到资源中
	if !containsString(pool.Finalizers, nodeFinalizer) {
		pool.Finalizers = append(pool.Finalizers, nodeFinalizer)
		if err := r.Client.Update(ctx, pool); err != nil {
			return ctrl.Result{}, err
		}
	}

	if len(nodes.Items) > 0 {
		r.Log.Info("find nodes, will merge data", "nodes", len(nodes.Items))
		pool.Status.Allocatable = corev1.ResourceList{}
		pool.Status.NodeCount = len(nodes.Items)
		for _, n := range nodes.Items {
			n := n

			// 更新节点的标签和污点信息
			err := r.Update(ctx, pool.Spec.ApplyNode(n))
			if err != nil {
				return ctrl.Result{}, err
			}

			for name, quantity := range n.Status.Allocatable {
				q, ok := pool.Status.Allocatable[name]
				if ok {
					q.Add(quantity)
					pool.Status.Allocatable[name] = q
					continue
				}
				pool.Status.Allocatable[name] = quantity
			}
		}
	}

	runtimeClass := &v1beta1.RuntimeClass{}
	err = r.Get(ctx, client.ObjectKeyFromObject(pool.RuntimeClass()), runtimeClass)
	if client.IgnoreNotFound(err) != nil {
		return ctrl.Result{}, err
	}

	// 如果不存在创建一个新的
	if runtimeClass.Name == "" {
		runtimeClass = pool.RuntimeClass()
		err = ctrl.SetControllerReference(pool, runtimeClass, r.Scheme)
		if err != nil {
			return ctrl.Result{}, err
		}
		err = r.Create(ctx, runtimeClass)
		return ctrl.Result{}, err
	}

	// 如果存在则更新
	runtimeClass.Scheduling = pool.RuntimeClass().Scheduling
	runtimeClass.Handler = pool.RuntimeClass().Handler
	err = r.Client.Update(ctx, runtimeClass)
	if err != nil {
		return ctrl.Result{}, err
	}

	pool.Status.Status = 200
	err = r.Status().Update(ctx, pool)
	return ctrl.Result{}, err
}

// 节点预删除逻辑
func (r *NodePoolReconciler) nodeFinalizer(ctx context.Context, pool *nodesv1.NodePool, nodes []corev1.Node) error {
	// 不为空就说明进入到预删除流程
	for _, n := range nodes {
		n := n

		// 更新节点的标签和污点信息
		err := r.Update(ctx, pool.Spec.CleanNode(n))
		if err != nil {
			return err
		}
	}

	// 预删除执行完毕，移除 nodeFinalizer
	pool.Finalizers = removeString(pool.Finalizers, nodeFinalizer)
	return r.Client.Update(ctx, pool)
}

// SetupWithManager sets up the controller with the Manager.
func (r *NodePoolReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&nodesv1.NodePool{}).
		Complete(r)
}

func nodeReady(status corev1.NodeStatus) bool {
	for _, condition := range status.Conditions {
		if condition.Status == "True" && condition.Type == "Ready" {
			return true
		}
	}
	return false
}

// 辅助函数用于检查并从字符串切片中删除字符串。
func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

package controllers

import (
	"context"
	"strings"

	jobv1 "github.com/mohuishou/blog-code/02-k8s-operator/multi-cluster-operator/api/v1"
	"github.com/mohuishou/blog-code/02-k8s-operator/multi-cluster-operator/pkg/clusterx"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// TestJobReconciler reconciles a Test object
type TestJobReconciler struct {
	// 主集群 client
	client.Client

	// 所有集群的客户端列表
	Clients map[string]client.Client

	Scheme *runtime.Scheme
}

// NewTestReconciler ...
func NewTestJobReconciler(mgr ctrl.Manager, clusters map[string]cluster.Cluster) (*TestJobReconciler, error) {
	r := TestJobReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		Clients: map[string]client.Client{
			"main": mgr.GetClient(),
		},
	}
	for name, cluster := range clusters {
		r.Clients[name] = cluster.GetClient()
	}

	err := r.SetupWithManager(mgr, clusters)
	return &r, err
}

func (r *TestJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var res ctrl.Result

	logger := log.FromContext(ctx)

	var job batchv1.Job
	cluster, ns := GetClusterNameNs(req.Namespace)
	req.Namespace = ns

	logger.Info("get job", "cluster", cluster)

	err := r.GetClient(cluster).Get(ctx, req.NamespacedName, &job)
	if err != nil {
		return res, client.IgnoreNotFound(err)
	}

	if job.Status.CompletionTime.IsZero() {
		return res, nil
	}
	logger.Info("job complete", "cluster", cluster)

	var test jobv1.Test
	err = r.Get(ctx, clusterx.GetOwnerNameNs(&job), &test)
	if err != nil {
		return res, client.IgnoreNotFound(err)
	}

	test.Status.Phase = "finished"
	err = r.Client.Status().Update(ctx, &test)
	return ctrl.Result{}, err
}

// MuiltClustersEnqueue 多集群入队器
// 将集群名称附加在 Namespace 上
func MuiltClustersEnqueue(clusterName string) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(o client.Object) []reconcile.Request {
		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Name:      o.GetName(),
					Namespace: clusterName + "/" + o.GetNamespace(),
				},
			},
		}
	})
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestJobReconciler) SetupWithManager(mgr ctrl.Manager, cs map[string]cluster.Cluster) error {
	build := ctrl.NewControllerManagedBy(mgr).
		For(&batchv1.Job{})

		// 监听多个集群
	for name, cluster := range cs {
		build = build.Watches(
			source.NewKindWithCache(&batchv1.Job{}, cluster.GetCache()),
			MuiltClustersEnqueue(name),
		)
	}
	return build.Complete(r)
}

// GetClient by cluster name
func (r *TestJobReconciler) GetClient(name string) client.Client {
	if c, ok := r.Clients[name]; ok {
		return c
	}

	return r.Client
}

// GetClientByNs 通过命名空间获取 client
func (r *TestJobReconciler) GetClientByNs(ns string) client.Client {
	name, _ := GetClusterNameNs(ns)
	return r.GetClient(name)
}

// GetClusterNameNs 获取集群名称和 NS
func GetClusterNameNs(ns string) (clusterName, namespace string) {
	ss := strings.Split(ns, "/")
	if len(ss) < 2 {
		return "main", ns
	}

	return ss[0], ss[1]
}

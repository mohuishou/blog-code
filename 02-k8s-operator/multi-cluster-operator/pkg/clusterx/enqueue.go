package clusterx

import (
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MuiltClustersEnqueue 多集群入队器
// 将集群名称附加在 Namespace 上
func MuiltClustersEnqueue(clusterName string) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(o client.Object) []reconcile.Request {
		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Name:      o.GetName(),
					Namespace: clusterName + ClusterSeparator + o.GetNamespace(),
				},
			},
		}
	})
}

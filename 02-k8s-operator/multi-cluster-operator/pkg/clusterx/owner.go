package clusterx

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// owner 资源关联
const (
	OwnerNameKey      = "owner.name"
	OwnerNamespaceKey = "owner.namespace"
	OwnerClusterKey   = "owner.cluster"
)

// 创建时给资源添加一个注解
const ClusterNameKey = "cluster.name"

// SetOwner 注入 owner 相关标签
func SetOwner(owner, sub metav1.Object) {
	labels := sub.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	labels[OwnerNamespaceKey] = owner.GetNamespace()
	labels[OwnerNameKey] = owner.GetName()
	labels[OwnerClusterKey] = GetClusterName(owner)
	sub.SetLabels(labels)
}

// GetOwnerNameNs 获取
func GetOwnerNameNs(sub metav1.Object) types.NamespacedName {
	labels := sub.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	return types.NamespacedName{
		Namespace: labels[OwnerNamespaceKey],
		Name:      labels[OwnerNameKey],
	}
}

// GetOwnerClusterName 获取当前资源 owner 所在集群
func GetOwnerClusterName(sub metav1.Object) string {
	return sub.GetLabels()[OwnerClusterKey]
}

// GetClusterName 获取资源当前的集群名称
func GetClusterName(obj metav1.Object) string {
	return obj.GetAnnotations()[ClusterNameKey]
}

// SetClusterName 设置资源当前的集群名称
func SetClusterName(obj metav1.Object, cluster string) {
	data := obj.GetAnnotations()
	data[ClusterNameKey] = cluster
}

package clusterx

import (
	"strings"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// MulitClusterClient 多集群 client
type MulitClusterClient struct {
	// 默认的主集群
	client client.Client
	// 所有集群的 client 列表
	clients map[string]client.Client
}

const (
	// ClusterSeparator 集群分割
	ClusterSeparator = "/"
	// DefaultClusterName 默认集群名字
	DefaultClusterName = "default"
)

// GetClient by cluster name
func (mc MulitClusterClient) GetClient(name string) client.Client {
	if c, ok := mc.clients[name]; ok {
		return c
	}

	return mc.client
}

// GetClientByNs 通过命名空间获取 client
func (mc MulitClusterClient) GetClientByNs(ns string) client.Client {
	name, _ := GetClusterNameNs(ns)
	return mc.GetClient(name)
}

// GetOwnerClient ...
func (mc MulitClusterClient) GetOwnerClient(obj client.Object) client.Client {
	return mc.GetClient(GetOwnerClusterName(obj))
}

// GetClientByObj ...
func (mc MulitClusterClient) GetClientByObj(obj client.Object) client.Client {
	return mc.GetClient(GetClusterName(obj))
}

// GetClusterNameNs 获取集群名称和 NS
func GetClusterNameNs(ns string) (clusterName, namespace string) {
	ss := strings.Split(ns, ClusterSeparator)
	if len(ss) < 2 {
		return DefaultClusterName, ns
	}

	return ss[0], ss[1]
}

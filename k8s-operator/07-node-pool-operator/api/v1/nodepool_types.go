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

package v1

import (
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/api/node/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodePoolSpec 节点池，MVP 版本只允许一个节点属于一个节点池
// TODO: 后续需要支持一个节点属于多个资源池
type NodePoolSpec struct {
	// Taints 污点
	Taints []corev1.Taint `json:"taints,omitempty"`

	// Labels 标签
	Labels map[string]string `json:"labels,omitempty"`

	// Handler 对应 Runtime Class 的 Handler
	Handler string `json:"handler,omitempty"`
}

// CleanNode 清理节点标签污点信息，仅保留系统标签
func (s *NodePoolSpec) CleanNode(node corev1.Node) *corev1.Node {
	// 除了节点池的标签之外，我们只保留 k8s 的相关标签
	// 注意：这里的逻辑如果一个节点只能属于一个节点池
	nodeLabels := map[string]string{}
	for k, v := range node.Labels {
		if strings.Contains(k, "kubernetes") {
			nodeLabels[k] = v
		}
	}
	node.Labels = nodeLabels

	// 污点同理
	var taints []corev1.Taint
	for _, taint := range node.Spec.Taints {
		if strings.Contains(taint.Key, "kubernetes") {
			taints = append(taints, taint)
		}
	}
	node.Spec.Taints = taints
	return &node
}

// ApplyNode 生成 Node 结构，可以用于 Patch 数据
func (s *NodePoolSpec) ApplyNode(node corev1.Node) *corev1.Node {
	n := s.CleanNode(node)

	for k, v := range s.Labels {
		n.Labels[k] = v
	}

	n.Spec.Taints = append(n.Spec.Taints, s.Taints...)
	return n
}

// NodePoolStatus defines the observed state of NodePool
type NodePoolStatus struct {

	// status=200 说明正常，其他情况为异常情况
	Status int `json:"status"`

	// 节点的数量
	NodeCount int `json:"nodeCount"`

	// 允许被调度的容量
	Allocatable corev1.ResourceList `json:"allocatable,omitempty" protobuf:"bytes,2,rep,name=allocatable,casttype=ResourceList,castkey=ResourceName"`
}

//+kubebuilder:printcolumn:JSONPath=".status.status",name=Status,type=integer
//+kubebuilder:printcolumn:JSONPath=".status.nodeCount",name=NodeCount,type=integer
//+kubebuilder:object:root=true
//+kubebuilder:resource:scope=Cluster
//+kubebuilder:subresource:status

// NodePool is the Schema for the nodepools API
type NodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodePoolSpec   `json:"spec,omitempty"`
	Status NodePoolStatus `json:"status,omitempty"`
}

// NodeRole 返回节点对应的 role 标签名
func (n *NodePool) NodeRole() string {
	return "node-role.kubernetes.io/" + n.Name
}

// NodeLabelSelector 返回节点 label 选择器
func (n *NodePool) NodeLabelSelector() labels.Selector {
	return labels.SelectorFromSet(map[string]string{
		n.NodeRole(): "",
	})
}

// RuntimeClass 生成对应的 runtime class 对象
func (n *NodePool) RuntimeClass() *v1beta1.RuntimeClass {
	s := n.Spec
	tolerations := make([]corev1.Toleration, len(s.Taints))
	for i, t := range s.Taints {
		tolerations[i] = corev1.Toleration{
			Key:      t.Key,
			Value:    t.Value,
			Effect:   t.Effect,
			Operator: corev1.TolerationOpEqual,
		}
	}

	return &v1beta1.RuntimeClass{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-pool-" + n.Name,
		},
		Handler: "runc",
		Scheduling: &v1beta1.Scheduling{
			NodeSelector: s.Labels,
			Tolerations:  tolerations,
		},
	}
}

//+kubebuilder:object:root=true

// NodePoolList contains a list of NodePool
type NodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodePool{}, &NodePoolList{})
}

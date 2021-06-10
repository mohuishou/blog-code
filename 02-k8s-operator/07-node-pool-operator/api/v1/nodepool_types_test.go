package v1

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
)

func TestNodePoolSpec_ApplyNode(t *testing.T) {
	type fields struct {
		Taints  []corev1.Taint
		Labels  map[string]string
		Handler string
	}
	type args struct {
		node v1.Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *corev1.Node
	}{
		{
			name: "label",
			fields: fields{
				Labels: map[string]string{
					"node-pool.lailin.xyz/test": "",
				},
			},
			args: args{
				node: v1.Node{
					ObjectMeta: metav1.ObjectMeta{
						Name: "worker",
						Labels: map[string]string{
							"kubernetes.io/arch": "amd64",
							"a":                  "b",
						},
					},
				},
			},
			want: &v1.Node{
				ObjectMeta: metav1.ObjectMeta{
					Name: "worker",
					Labels: map[string]string{
						"kubernetes.io/arch":        "amd64",
						"node-pool.lailin.xyz/test": "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &NodePoolSpec{
				Taints:  tt.fields.Taints,
				Labels:  tt.fields.Labels,
				Handler: tt.fields.Handler,
			}
			assert.Equal(t, tt.want, s.ApplyNode(tt.args.node))
		})
	}
}

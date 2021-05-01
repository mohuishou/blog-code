package v1

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
)

func TestNodePool_validate(t *testing.T) {
	type fields struct {
		Labels map[string]string
		Taints []corev1.Taint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr string
	}{
		{
			name: "one label key err",
			fields: fields{
				Labels: map[string]string{
					"xxx": "",
				},
				Taints: []corev1.Taint{
					{
						Key:    "",
						Value:  "",
						Effect: "",
					},
				},
			},
			wantErr: "xxx",
		},
		{
			name: "taint key empty",
			fields: fields{
				Labels: map[string]string{
					"node-pool.lailin.xyz":     "",
					"node-pool.lailin.xyz/xxx": "",
				},
				Taints: []corev1.Taint{
					{
						Key:    "",
						Value:  "",
						Effect: "",
					},
				},
			},
			wantErr: "taint key: ",
		},
		{
			name: "success",
			fields: fields{
				Labels: map[string]string{
					"node-pool.lailin.xyz":     "",
					"node-pool.lailin.xyz/xxx": "",
				},
				Taints: []corev1.Taint{
					{
						Key:    "node-pool.lailin.xyz",
						Value:  "",
						Effect: "",
					},
					{
						Key:    "node-pool.lailin.xyz/xxx",
						Value:  "",
						Effect: "",
					},
				},
			},
			wantErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &NodePool{
				Spec: NodePoolSpec{
					Taints: tt.fields.Taints,
					Labels: tt.fields.Labels,
				},
			}
			err := r.validate()
			if tt.wantErr == "" {
				require.Nil(t, err)
				return
			}
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

package util

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/require"
)

func TestGetNodeInternalIP(t *testing.T) {
	tests := []struct {
		name string
		node v1.Node
		exp4 string
		exp6 string
	}{
		{
			name: "correct",
			node: v1.Node{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       v1.NodeSpec{},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "192.168.0.2",
						},
						{
							Type:    "ExternalIP",
							Address: "192.188.0.4",
						},
						{
							Type:    "InternalIP",
							Address: "ffff:ffff:ffff:ffff:ffff::23",
						},
					},
				},
			},
			exp4: "192.168.0.2",
			exp6: "ffff:ffff:ffff:ffff:ffff::23",
		},
		{
			name: "correctWithDiff",
			node: v1.Node{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec:       v1.NodeSpec{},
				Status: v1.NodeStatus{
					Addresses: []v1.NodeAddress{
						{
							Type:    "InternalIP",
							Address: "ffff:ffff:ffff:ffff:ffff::23",
						},
						{
							Type:    "ExternalIP",
							Address: "192.188.0.4",
						},
						{
							Type:    "InternalIP",
							Address: "192.188.0.43",
						},
					},
				},
			},
			exp4: "192.188.0.43",
			exp6: "ffff:ffff:ffff:ffff:ffff::23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ret4, ret6 := GetNodeInternalIP(tt.node); ret4 != tt.exp4 || ret6 != tt.exp6 {
				t.Errorf("got %v, %v, want %v, %v", ret4, ret6, tt.exp4, tt.exp6)
			}
		})
	}
}

func TestServiceClusterIPs(t *testing.T) {
	tests := []struct {
		name string
		svc  v1.Service
		exp  []string
	}{
		{
			name: "service_with_one_cluster_ip",
			svc: v1.Service{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1.ServiceSpec{
					ClusterIP:  "10.96.0.1",
					ClusterIPs: []string{"10.96.0.1"},
				},
			},
			exp: []string{"10.96.0.1"},
		},
		{
			name: "service_with_two_cluster_ip",
			svc: v1.Service{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1.ServiceSpec{
					ClusterIP:  "10.96.0.1",
					ClusterIPs: []string{"10.96.0.1", "fd00:10:16::1"},
				},
			},
			exp: []string{"10.96.0.1", "fd00:10:16::1"},
		},
		{
			name: "service_with_no_cluster_ip",
			svc: v1.Service{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1.ServiceSpec{
					ClusterIP:  "",
					ClusterIPs: []string{},
				},
			},
			exp: []string{},
		},
		{
			name: "service_with_no_clusterips",
			svc: v1.Service{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1.ServiceSpec{
					ClusterIP:  "10.96.0.1",
					ClusterIPs: []string{},
				},
			},
			exp: []string{"10.96.0.1"},
		},
		{
			name: "service_with_invalid_cluster_ip",
			svc: v1.Service{
				TypeMeta:   metav1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{},
				Spec: v1.ServiceSpec{
					ClusterIP:  "",
					ClusterIPs: []string{"10.96.0.1", "invalid ip"},
				},
			},
			exp: []string{"10.96.0.1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if ret := ServiceClusterIPs(tt.svc); len(ret) != len(tt.exp) {
				t.Errorf("got %v, want %v", ret, tt.exp)
			}
		})
	}
}

func TestLabelSelectorNotEquals(t *testing.T) {
	selector, err := LabelSelectorNotEquals("key", "value")
	require.NoError(t, err)
	require.Equal(t, "key!=value", selector.String())
	// Test error case
	selector, err = LabelSelectorNotEquals("", "")
	require.Error(t, err)
	require.Nil(t, selector)
}

func TestLabelSelectorNotEmpty(t *testing.T) {
	selector, err := LabelSelectorNotEmpty("key")
	require.NoError(t, err)
	require.Equal(t, "key!=", selector.String())
	// Test error case
	selector, err = LabelSelectorNotEmpty("")
	require.Error(t, err)
	require.Nil(t, selector)
}

func TestGetTruncatedUID(t *testing.T) {
	uid := "12345678-1234-1234-1234-123456789012"
	require.Equal(t, "123456789012", GetTruncatedUID(uid))
}

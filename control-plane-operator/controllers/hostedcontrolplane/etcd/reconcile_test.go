package etcd

import (
	"testing"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"

	hyperv1 "github.com/openshift/hypershift/api/v1beta1"
	"github.com/openshift/hypershift/support/config"
)

func TestBuildEtcdInitContainer(t *testing.T) {
	tests := []struct {
		name        string
		params      EtcdParams
		expectedEnv []corev1.EnvVar
	}{
		{
			name: "single replica container env as expected",
			params: EtcdParams{
				EtcdImage: "animage",
				DeploymentConfig: config.DeploymentConfig{
					Replicas: 1,
				},
				StorageSpec: hyperv1.ManagedEtcdStorageSpec{
					RestoreSnapshotURL: []string{"arestoreurl"},
				},
			},
			expectedEnv: []corev1.EnvVar{
				{
					Name:  "RESTORE_URL_ETCD",
					Value: "arestoreurl",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGomegaWithT(t)
			c := corev1.Container{}
			buildEtcdInitContainer(&tt.params)(&c)
			g.Expect(c.Env).Should(ConsistOf(tt.expectedEnv))
		})
	}
}

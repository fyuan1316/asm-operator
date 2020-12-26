package resource

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/sync"
	appsv1 "k8s.io/api/apps/v1"
	"testing"
)

func TestSyncManager_LoadFile(t *testing.T) {
	type fields struct {
		K8sResource SyncResource
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test-load-file",
			fields:  fields{},
			args:    args{filePath: "./pod.yaml"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SyncManager{
				K8sResource: map[string]SyncResource{
					"test1": tt.fields.K8sResource,
				},
			}
			res := SyncResource{
				Object: &appsv1.Deployment{},
				Sync:   sync.FnDeployment,
			}
			if err := m.LoadFile(tt.args.filePath, res); (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println()
		})
	}
}

package resource

import (
	"reflect"
	"testing"
)

/*
import (
	"reflect"
	"testing"
)

func TestGetResourceFiles(t *testing.T) {
	type args struct {
		folderPath string
		opts       []Option
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test-1",
			args: args{
				folderPath: "/Users/yuan/Dev/work/GolangProjects/asm-operator/pkg/provision/cluster-asm/crds",
				opts:       []Option{Suffix(".yaml")},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFilesInFolder(tt.args.folderPath, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFilesInFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFilesInFolder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

*/

func TestGetChartResources(t *testing.T) {
	type args struct {
		folderPath string
		userValues map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "test-merge-values",
			args: args{
				folderPath: "/Users/yuan/Dev/work/GolangProjects/asm-operator/files/provision/cluster-asm",
				userValues: map[string]interface{}{
					"asm_controller": map[string]interface{}{
						"replicaCount": 2,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChartResources(tt.args.folderPath, tt.args.userValues)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChartResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChartResources() got = %v, want %v", got, tt.want)
			}
		})
	}
}

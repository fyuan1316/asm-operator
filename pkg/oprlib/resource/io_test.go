package resource

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
